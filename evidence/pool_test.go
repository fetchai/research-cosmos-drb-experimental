package evidence

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	dbm "github.com/tendermint/tm-db"

	sm "github.com/tendermint/tendermint/state"
	"github.com/tendermint/tendermint/store"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidences(cdc)

	code := m.Run()
	os.Exit(code)
}

func initializeValidatorState(valAddr []byte, height int64) dbm.DB {
	stateDB := dbm.NewMemDB()

	// create validator set and state
	valSet := &types.ValidatorSet{
		Validators: []*types.Validator{
			{Address: valAddr},
		},
	}
	state := sm.State{
		LastBlockHeight:             0,
		LastBlockTime:               tmtime.Now(),
		Validators:                  valSet,
		NextValidators:              valSet.CopyIncrementProposerPriority(1),
		LastHeightValidatorsChanged: 1,
		ConsensusParams: types.ConsensusParams{
			Evidence: types.EvidenceParams{
				MaxAgeNumBlocks: 10000,
				MaxAgeDuration:  48 * time.Hour,
			},
		},
	}

	// save all states up to height
	for i := int64(0); i < height; i++ {
		state.LastBlockHeight = i
		sm.SaveState(stateDB, state)
	}

	return stateDB
}

func TestEvidencePool(t *testing.T) {

	var (
		valAddr      = []byte("val1")
		height       = int64(100002)
		stateDB      = initializeValidatorState(valAddr, height)
		evidenceDB   = dbm.NewMemDB()
		blockStore   = store.NewBlockStore(dbm.NewMemDB())
		pool         = NewPool(stateDB, evidenceDB, blockStore)
		evidenceTime = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	)

	goodEvidence := types.NewMockEvidence(height, time.Now(), 0, valAddr)
	badEvidence := types.NewMockEvidence(1, evidenceTime, 0, valAddr)

	// bad evidence
	err := pool.AddEvidence(badEvidence)
	assert.Error(t, err)
	// err: evidence created at 2019-01-01 00:00:00 +0000 UTC has expired. Evidence can not be older than: ...

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		<-pool.EvidenceWaitChan()
		wg.Done()
	}()

	err = pool.AddEvidence(goodEvidence)
	assert.NoError(t, err)
	wg.Wait()

	assert.Equal(t, 1, pool.evidenceList.Len())

	// if we send it again, it shouldnt add and return an error
	err = pool.AddEvidence(goodEvidence)
	assert.Error(t, err)
	assert.Equal(t, 1, pool.evidenceList.Len())
}

func TestEvidencePoolIsCommitted(t *testing.T) {
	// Initialization:
	var (
		valAddr       = []byte("validator_address")
		height        = int64(42)
		lastBlockTime = time.Now()
		stateDB       = initializeValidatorState(valAddr, height)
		evidenceDB    = dbm.NewMemDB()
		blockStore    = store.NewBlockStore(dbm.NewMemDB())
		pool          = NewPool(stateDB, evidenceDB, blockStore)
	)

	// evidence not seen yet:
	evidence := types.NewMockEvidence(height, time.Now(), 0, valAddr)
	assert.False(t, pool.IsCommitted(evidence))

	// evidence seen but not yet committed:
	assert.NoError(t, pool.AddEvidence(evidence))
	assert.False(t, pool.IsCommitted(evidence))

	// evidence seen and committed:
	pool.MarkEvidenceAsCommitted(height, lastBlockTime, []types.Evidence{evidence})
	assert.True(t, pool.IsCommitted(evidence))
}

func TestAddEvidence(t *testing.T) {

	var (
		valAddr      = []byte("val1")
		height       = int64(100002)
		stateDB      = initializeValidatorState(valAddr, height)
		evidenceDB   = dbm.NewMemDB()
		blockStore   = store.NewBlockStore(dbm.NewMemDB())
		pool         = NewPool(stateDB, evidenceDB, blockStore)
		evidenceTime = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	)

	testCases := []struct {
		evHeight      int64
		evTime        time.Time
		expErr        bool
		evDescription string
	}{
		{height, time.Now(), false, "valid evidence"},
		{height, evidenceTime, false, "valid evidence (despite old time)"},
		{int64(1), time.Now(), false, "valid evidence (despite old height)"},
		{int64(1), evidenceTime, true,
			"evidence from height 1 (created at: 2019-01-01 00:00:00 +0000 UTC) is too old"},
	}

	for _, tc := range testCases {
		tc := tc
		ev := types.NewMockEvidence(tc.evHeight, tc.evTime, 0, valAddr)
		err := pool.AddEvidence(ev)
		if tc.expErr {
			assert.Error(t, err)
		}
	}
}

func TestAddNonUniqueBeaconEvidence(t *testing.T) {
	var (
		evidenceDB    = dbm.NewMemDB()
		height        = int64(10)
		stateDB       = dbm.NewMemDB()
		blockStore    = store.NewBlockStore(dbm.NewMemDB())
		pool          = NewPool(stateDB, evidenceDB, blockStore)
		valAddr       = []byte("val1")
		valAddr2      = []byte("val2")
		complainantPV = types.NewMockPV()
	)
	complainantPubKey, _ := complainantPV.GetPubKey()

	// Add unique evidence
	ev := types.NewBeaconInactivityEvidence(height, valAddr, complainantPubKey.Address(), 1, 1)
	hasEvidence := pool.store.Has(ev)
	assert.Equal(t, hasEvidence, false)
	_, err := pool.store.AddNewEvidence(ev, 10)
	assert.Nil(t, err)

	testCases := []struct {
		name           string
		evHeight       int64
		defAddr        []byte
		evidenceUnique bool
	}{
		{"Different height", height + 1, valAddr, false},
		{"Different defendant address", height, valAddr2, true},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ev := types.NewBeaconInactivityEvidence(tc.evHeight, tc.defAddr, complainantPubKey.Address(), 1, 1)
			hasEvidence := pool.store.Has(ev)
			assert.Equal(t, tc.evidenceUnique, !hasEvidence)
		})
	}
}
