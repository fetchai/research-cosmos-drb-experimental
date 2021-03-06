// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/types/evidence.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	keys "github.com/tendermint/tendermint/proto/crypto/keys"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// DuplicateVoteEvidence contains evidence a validator signed two conflicting
// votes.
type DuplicateVoteEvidence struct {
	PubKey               *keys.PublicKey `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VoteA                *Vote           `protobuf:"bytes,2,opt,name=vote_a,json=voteA,proto3" json:"vote_a,omitempty"`
	VoteB                *Vote           `protobuf:"bytes,3,opt,name=vote_b,json=voteB,proto3" json:"vote_b,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DuplicateVoteEvidence) Reset()         { *m = DuplicateVoteEvidence{} }
func (m *DuplicateVoteEvidence) String() string { return proto.CompactTextString(m) }
func (*DuplicateVoteEvidence) ProtoMessage()    {}
func (*DuplicateVoteEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{0}
}
func (m *DuplicateVoteEvidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DuplicateVoteEvidence.Unmarshal(m, b)
}
func (m *DuplicateVoteEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DuplicateVoteEvidence.Marshal(b, m, deterministic)
}
func (m *DuplicateVoteEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DuplicateVoteEvidence.Merge(m, src)
}
func (m *DuplicateVoteEvidence) XXX_Size() int {
	return xxx_messageInfo_DuplicateVoteEvidence.Size(m)
}
func (m *DuplicateVoteEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_DuplicateVoteEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_DuplicateVoteEvidence proto.InternalMessageInfo

func (m *DuplicateVoteEvidence) GetPubKey() *keys.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *DuplicateVoteEvidence) GetVoteA() *Vote {
	if m != nil {
		return m.VoteA
	}
	return nil
}

func (m *DuplicateVoteEvidence) GetVoteB() *Vote {
	if m != nil {
		return m.VoteB
	}
	return nil
}

type PotentialAmnesiaEvidence struct {
	VoteA                *Vote    `protobuf:"bytes,1,opt,name=vote_a,json=voteA,proto3" json:"vote_a,omitempty"`
	VoteB                *Vote    `protobuf:"bytes,2,opt,name=vote_b,json=voteB,proto3" json:"vote_b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PotentialAmnesiaEvidence) Reset()         { *m = PotentialAmnesiaEvidence{} }
func (m *PotentialAmnesiaEvidence) String() string { return proto.CompactTextString(m) }
func (*PotentialAmnesiaEvidence) ProtoMessage()    {}
func (*PotentialAmnesiaEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{1}
}
func (m *PotentialAmnesiaEvidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PotentialAmnesiaEvidence.Unmarshal(m, b)
}
func (m *PotentialAmnesiaEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PotentialAmnesiaEvidence.Marshal(b, m, deterministic)
}
func (m *PotentialAmnesiaEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PotentialAmnesiaEvidence.Merge(m, src)
}
func (m *PotentialAmnesiaEvidence) XXX_Size() int {
	return xxx_messageInfo_PotentialAmnesiaEvidence.Size(m)
}
func (m *PotentialAmnesiaEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_PotentialAmnesiaEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_PotentialAmnesiaEvidence proto.InternalMessageInfo

func (m *PotentialAmnesiaEvidence) GetVoteA() *Vote {
	if m != nil {
		return m.VoteA
	}
	return nil
}

func (m *PotentialAmnesiaEvidence) GetVoteB() *Vote {
	if m != nil {
		return m.VoteB
	}
	return nil
}

// MockEvidence is used for testing pruposes
type MockEvidence struct {
	EvidenceHeight       int64     `protobuf:"varint,1,opt,name=evidence_height,json=evidenceHeight,proto3" json:"evidence_height,omitempty"`
	EvidenceTime         time.Time `protobuf:"bytes,2,opt,name=evidence_time,json=evidenceTime,proto3,stdtime" json:"evidence_time"`
	EvidenceAddress      []byte    `protobuf:"bytes,3,opt,name=evidence_address,json=evidenceAddress,proto3" json:"evidence_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MockEvidence) Reset()         { *m = MockEvidence{} }
func (m *MockEvidence) String() string { return proto.CompactTextString(m) }
func (*MockEvidence) ProtoMessage()    {}
func (*MockEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{2}
}
func (m *MockEvidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MockEvidence.Unmarshal(m, b)
}
func (m *MockEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MockEvidence.Marshal(b, m, deterministic)
}
func (m *MockEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockEvidence.Merge(m, src)
}
func (m *MockEvidence) XXX_Size() int {
	return xxx_messageInfo_MockEvidence.Size(m)
}
func (m *MockEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_MockEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_MockEvidence proto.InternalMessageInfo

func (m *MockEvidence) GetEvidenceHeight() int64 {
	if m != nil {
		return m.EvidenceHeight
	}
	return 0
}

func (m *MockEvidence) GetEvidenceTime() time.Time {
	if m != nil {
		return m.EvidenceTime
	}
	return time.Time{}
}

func (m *MockEvidence) GetEvidenceAddress() []byte {
	if m != nil {
		return m.EvidenceAddress
	}
	return nil
}

type MockRandomEvidence struct {
	EvidenceHeight       int64     `protobuf:"varint,1,opt,name=evidence_height,json=evidenceHeight,proto3" json:"evidence_height,omitempty"`
	EvidenceTime         time.Time `protobuf:"bytes,2,opt,name=evidence_time,json=evidenceTime,proto3,stdtime" json:"evidence_time"`
	EvidenceAddress      []byte    `protobuf:"bytes,3,opt,name=evidence_address,json=evidenceAddress,proto3" json:"evidence_address,omitempty"`
	RandBytes            []byte    `protobuf:"bytes,4,opt,name=rand_bytes,json=randBytes,proto3" json:"rand_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MockRandomEvidence) Reset()         { *m = MockRandomEvidence{} }
func (m *MockRandomEvidence) String() string { return proto.CompactTextString(m) }
func (*MockRandomEvidence) ProtoMessage()    {}
func (*MockRandomEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{3}
}
func (m *MockRandomEvidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MockRandomEvidence.Unmarshal(m, b)
}
func (m *MockRandomEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MockRandomEvidence.Marshal(b, m, deterministic)
}
func (m *MockRandomEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockRandomEvidence.Merge(m, src)
}
func (m *MockRandomEvidence) XXX_Size() int {
	return xxx_messageInfo_MockRandomEvidence.Size(m)
}
func (m *MockRandomEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_MockRandomEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_MockRandomEvidence proto.InternalMessageInfo

func (m *MockRandomEvidence) GetEvidenceHeight() int64 {
	if m != nil {
		return m.EvidenceHeight
	}
	return 0
}

func (m *MockRandomEvidence) GetEvidenceTime() time.Time {
	if m != nil {
		return m.EvidenceTime
	}
	return time.Time{}
}

func (m *MockRandomEvidence) GetEvidenceAddress() []byte {
	if m != nil {
		return m.EvidenceAddress
	}
	return nil
}

func (m *MockRandomEvidence) GetRandBytes() []byte {
	if m != nil {
		return m.RandBytes
	}
	return nil
}

type ConflictingHeadersEvidence struct {
	H1                   *SignedHeader `protobuf:"bytes,1,opt,name=h1,proto3" json:"h1,omitempty"`
	H2                   *SignedHeader `protobuf:"bytes,2,opt,name=h2,proto3" json:"h2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ConflictingHeadersEvidence) Reset()         { *m = ConflictingHeadersEvidence{} }
func (m *ConflictingHeadersEvidence) String() string { return proto.CompactTextString(m) }
func (*ConflictingHeadersEvidence) ProtoMessage()    {}
func (*ConflictingHeadersEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{4}
}
func (m *ConflictingHeadersEvidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConflictingHeadersEvidence.Unmarshal(m, b)
}
func (m *ConflictingHeadersEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConflictingHeadersEvidence.Marshal(b, m, deterministic)
}
func (m *ConflictingHeadersEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConflictingHeadersEvidence.Merge(m, src)
}
func (m *ConflictingHeadersEvidence) XXX_Size() int {
	return xxx_messageInfo_ConflictingHeadersEvidence.Size(m)
}
func (m *ConflictingHeadersEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_ConflictingHeadersEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_ConflictingHeadersEvidence proto.InternalMessageInfo

func (m *ConflictingHeadersEvidence) GetH1() *SignedHeader {
	if m != nil {
		return m.H1
	}
	return nil
}

func (m *ConflictingHeadersEvidence) GetH2() *SignedHeader {
	if m != nil {
		return m.H2
	}
	return nil
}

type LunaticValidatorEvidence struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Vote                 *Vote    `protobuf:"bytes,2,opt,name=vote,proto3" json:"vote,omitempty"`
	InvalidHeaderField   string   `protobuf:"bytes,3,opt,name=invalid_header_field,json=invalidHeaderField,proto3" json:"invalid_header_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LunaticValidatorEvidence) Reset()         { *m = LunaticValidatorEvidence{} }
func (m *LunaticValidatorEvidence) String() string { return proto.CompactTextString(m) }
func (*LunaticValidatorEvidence) ProtoMessage()    {}
func (*LunaticValidatorEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{5}
}
func (m *LunaticValidatorEvidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LunaticValidatorEvidence.Unmarshal(m, b)
}
func (m *LunaticValidatorEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LunaticValidatorEvidence.Marshal(b, m, deterministic)
}
func (m *LunaticValidatorEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LunaticValidatorEvidence.Merge(m, src)
}
func (m *LunaticValidatorEvidence) XXX_Size() int {
	return xxx_messageInfo_LunaticValidatorEvidence.Size(m)
}
func (m *LunaticValidatorEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_LunaticValidatorEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_LunaticValidatorEvidence proto.InternalMessageInfo

func (m *LunaticValidatorEvidence) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LunaticValidatorEvidence) GetVote() *Vote {
	if m != nil {
		return m.Vote
	}
	return nil
}

func (m *LunaticValidatorEvidence) GetInvalidHeaderField() string {
	if m != nil {
		return m.InvalidHeaderField
	}
	return ""
}

type BeaconInactivityEvidence struct {
	EvidenceHeight       int64     `protobuf:"varint,1,opt,name=evidence_height,json=evidenceHeight,proto3" json:"evidence_height,omitempty"`
	EvidenceTime         time.Time `protobuf:"bytes,2,opt,name=evidence_time,json=evidenceTime,proto3,stdtime" json:"evidence_time"`
	DefendantAddress     []byte    `protobuf:"bytes,3,opt,name=defendant_address,json=defendantAddress,proto3" json:"defendant_address,omitempty"`
	ComplainantAddress   []byte    `protobuf:"bytes,4,opt,name=complainant_address,json=complainantAddress,proto3" json:"complainant_address,omitempty"`
	AeonStart            int64     `protobuf:"varint,5,opt,name=aeon_start,json=aeonStart,proto3" json:"aeon_start,omitempty"`
	ComplainantSignature []byte    `protobuf:"bytes,6,opt,name=complainant_signature,json=complainantSignature,proto3" json:"complainant_signature,omitempty"`
	Threshold            int64     `protobuf:"varint,7,opt,name=threshold,proto3" json:"threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BeaconInactivityEvidence) Reset()         { *m = BeaconInactivityEvidence{} }
func (m *BeaconInactivityEvidence) String() string { return proto.CompactTextString(m) }
func (*BeaconInactivityEvidence) ProtoMessage()    {}
func (*BeaconInactivityEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{6}
}
func (m *BeaconInactivityEvidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconInactivityEvidence.Unmarshal(m, b)
}
func (m *BeaconInactivityEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconInactivityEvidence.Marshal(b, m, deterministic)
}
func (m *BeaconInactivityEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconInactivityEvidence.Merge(m, src)
}
func (m *BeaconInactivityEvidence) XXX_Size() int {
	return xxx_messageInfo_BeaconInactivityEvidence.Size(m)
}
func (m *BeaconInactivityEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconInactivityEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconInactivityEvidence proto.InternalMessageInfo

func (m *BeaconInactivityEvidence) GetEvidenceHeight() int64 {
	if m != nil {
		return m.EvidenceHeight
	}
	return 0
}

func (m *BeaconInactivityEvidence) GetEvidenceTime() time.Time {
	if m != nil {
		return m.EvidenceTime
	}
	return time.Time{}
}

func (m *BeaconInactivityEvidence) GetDefendantAddress() []byte {
	if m != nil {
		return m.DefendantAddress
	}
	return nil
}

func (m *BeaconInactivityEvidence) GetComplainantAddress() []byte {
	if m != nil {
		return m.ComplainantAddress
	}
	return nil
}

func (m *BeaconInactivityEvidence) GetAeonStart() int64 {
	if m != nil {
		return m.AeonStart
	}
	return 0
}

func (m *BeaconInactivityEvidence) GetComplainantSignature() []byte {
	if m != nil {
		return m.ComplainantSignature
	}
	return nil
}

func (m *BeaconInactivityEvidence) GetThreshold() int64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

type DKGEvidence struct {
	EvidenceHeight       int64     `protobuf:"varint,1,opt,name=evidence_height,json=evidenceHeight,proto3" json:"evidence_height,omitempty"`
	EvidenceTime         time.Time `protobuf:"bytes,2,opt,name=evidence_time,json=evidenceTime,proto3,stdtime" json:"evidence_time"`
	DefendantAddress     []byte    `protobuf:"bytes,3,opt,name=defendant_address,json=defendantAddress,proto3" json:"defendant_address,omitempty"`
	ComplainantAddress   []byte    `protobuf:"bytes,4,opt,name=complainant_address,json=complainantAddress,proto3" json:"complainant_address,omitempty"`
	ValidatorHeight      int64     `protobuf:"varint,5,opt,name=validator_height,json=validatorHeight,proto3" json:"validator_height,omitempty"`
	DkgId                int64     `protobuf:"varint,6,opt,name=dkg_id,json=dkgId,proto3" json:"dkg_id,omitempty"`
	ComplainantSignature []byte    `protobuf:"bytes,7,opt,name=complainant_signature,json=complainantSignature,proto3" json:"complainant_signature,omitempty"`
	Threshold            int64     `protobuf:"varint,8,opt,name=threshold,proto3" json:"threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DKGEvidence) Reset()         { *m = DKGEvidence{} }
func (m *DKGEvidence) String() string { return proto.CompactTextString(m) }
func (*DKGEvidence) ProtoMessage()    {}
func (*DKGEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{7}
}
func (m *DKGEvidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DKGEvidence.Unmarshal(m, b)
}
func (m *DKGEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DKGEvidence.Marshal(b, m, deterministic)
}
func (m *DKGEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DKGEvidence.Merge(m, src)
}
func (m *DKGEvidence) XXX_Size() int {
	return xxx_messageInfo_DKGEvidence.Size(m)
}
func (m *DKGEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_DKGEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_DKGEvidence proto.InternalMessageInfo

func (m *DKGEvidence) GetEvidenceHeight() int64 {
	if m != nil {
		return m.EvidenceHeight
	}
	return 0
}

func (m *DKGEvidence) GetEvidenceTime() time.Time {
	if m != nil {
		return m.EvidenceTime
	}
	return time.Time{}
}

func (m *DKGEvidence) GetDefendantAddress() []byte {
	if m != nil {
		return m.DefendantAddress
	}
	return nil
}

func (m *DKGEvidence) GetComplainantAddress() []byte {
	if m != nil {
		return m.ComplainantAddress
	}
	return nil
}

func (m *DKGEvidence) GetValidatorHeight() int64 {
	if m != nil {
		return m.ValidatorHeight
	}
	return 0
}

func (m *DKGEvidence) GetDkgId() int64 {
	if m != nil {
		return m.DkgId
	}
	return 0
}

func (m *DKGEvidence) GetComplainantSignature() []byte {
	if m != nil {
		return m.ComplainantSignature
	}
	return nil
}

func (m *DKGEvidence) GetThreshold() int64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

type Evidence struct {
	// Types that are valid to be assigned to Sum:
	//	*Evidence_DuplicateVoteEvidence
	//	*Evidence_ConflictingHeadersEvidence
	//	*Evidence_LunaticValidatorEvidence
	//	*Evidence_PotentialAmnesiaEvidence
	//	*Evidence_MockEvidence
	//	*Evidence_MockRandomEvidence
	//	*Evidence_BeaconInactivityEvidence
	//	*Evidence_DkgEvidence
	Sum                  isEvidence_Sum `protobuf_oneof:"sum"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{8}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Evidence.Unmarshal(m, b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
}
func (m *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(m, src)
}
func (m *Evidence) XXX_Size() int {
	return xxx_messageInfo_Evidence.Size(m)
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

type isEvidence_Sum interface {
	isEvidence_Sum()
}

type Evidence_DuplicateVoteEvidence struct {
	DuplicateVoteEvidence *DuplicateVoteEvidence `protobuf:"bytes,1,opt,name=duplicate_vote_evidence,json=duplicateVoteEvidence,proto3,oneof" json:"duplicate_vote_evidence,omitempty"`
}
type Evidence_ConflictingHeadersEvidence struct {
	ConflictingHeadersEvidence *ConflictingHeadersEvidence `protobuf:"bytes,2,opt,name=conflicting_headers_evidence,json=conflictingHeadersEvidence,proto3,oneof" json:"conflicting_headers_evidence,omitempty"`
}
type Evidence_LunaticValidatorEvidence struct {
	LunaticValidatorEvidence *LunaticValidatorEvidence `protobuf:"bytes,3,opt,name=lunatic_validator_evidence,json=lunaticValidatorEvidence,proto3,oneof" json:"lunatic_validator_evidence,omitempty"`
}
type Evidence_PotentialAmnesiaEvidence struct {
	PotentialAmnesiaEvidence *PotentialAmnesiaEvidence `protobuf:"bytes,4,opt,name=potential_amnesia_evidence,json=potentialAmnesiaEvidence,proto3,oneof" json:"potential_amnesia_evidence,omitempty"`
}
type Evidence_MockEvidence struct {
	MockEvidence *MockEvidence `protobuf:"bytes,5,opt,name=mock_evidence,json=mockEvidence,proto3,oneof" json:"mock_evidence,omitempty"`
}
type Evidence_MockRandomEvidence struct {
	MockRandomEvidence *MockRandomEvidence `protobuf:"bytes,6,opt,name=mock_random_evidence,json=mockRandomEvidence,proto3,oneof" json:"mock_random_evidence,omitempty"`
}
type Evidence_BeaconInactivityEvidence struct {
	BeaconInactivityEvidence *BeaconInactivityEvidence `protobuf:"bytes,7,opt,name=beacon_inactivity_evidence,json=beaconInactivityEvidence,proto3,oneof" json:"beacon_inactivity_evidence,omitempty"`
}
type Evidence_DkgEvidence struct {
	DkgEvidence *DKGEvidence `protobuf:"bytes,8,opt,name=dkg_evidence,json=dkgEvidence,proto3,oneof" json:"dkg_evidence,omitempty"`
}

func (*Evidence_DuplicateVoteEvidence) isEvidence_Sum()      {}
func (*Evidence_ConflictingHeadersEvidence) isEvidence_Sum() {}
func (*Evidence_LunaticValidatorEvidence) isEvidence_Sum()   {}
func (*Evidence_PotentialAmnesiaEvidence) isEvidence_Sum()   {}
func (*Evidence_MockEvidence) isEvidence_Sum()               {}
func (*Evidence_MockRandomEvidence) isEvidence_Sum()         {}
func (*Evidence_BeaconInactivityEvidence) isEvidence_Sum()   {}
func (*Evidence_DkgEvidence) isEvidence_Sum()                {}

func (m *Evidence) GetSum() isEvidence_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Evidence) GetDuplicateVoteEvidence() *DuplicateVoteEvidence {
	if x, ok := m.GetSum().(*Evidence_DuplicateVoteEvidence); ok {
		return x.DuplicateVoteEvidence
	}
	return nil
}

func (m *Evidence) GetConflictingHeadersEvidence() *ConflictingHeadersEvidence {
	if x, ok := m.GetSum().(*Evidence_ConflictingHeadersEvidence); ok {
		return x.ConflictingHeadersEvidence
	}
	return nil
}

func (m *Evidence) GetLunaticValidatorEvidence() *LunaticValidatorEvidence {
	if x, ok := m.GetSum().(*Evidence_LunaticValidatorEvidence); ok {
		return x.LunaticValidatorEvidence
	}
	return nil
}

func (m *Evidence) GetPotentialAmnesiaEvidence() *PotentialAmnesiaEvidence {
	if x, ok := m.GetSum().(*Evidence_PotentialAmnesiaEvidence); ok {
		return x.PotentialAmnesiaEvidence
	}
	return nil
}

func (m *Evidence) GetMockEvidence() *MockEvidence {
	if x, ok := m.GetSum().(*Evidence_MockEvidence); ok {
		return x.MockEvidence
	}
	return nil
}

func (m *Evidence) GetMockRandomEvidence() *MockRandomEvidence {
	if x, ok := m.GetSum().(*Evidence_MockRandomEvidence); ok {
		return x.MockRandomEvidence
	}
	return nil
}

func (m *Evidence) GetBeaconInactivityEvidence() *BeaconInactivityEvidence {
	if x, ok := m.GetSum().(*Evidence_BeaconInactivityEvidence); ok {
		return x.BeaconInactivityEvidence
	}
	return nil
}

func (m *Evidence) GetDkgEvidence() *DKGEvidence {
	if x, ok := m.GetSum().(*Evidence_DkgEvidence); ok {
		return x.DkgEvidence
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Evidence) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Evidence_DuplicateVoteEvidence)(nil),
		(*Evidence_ConflictingHeadersEvidence)(nil),
		(*Evidence_LunaticValidatorEvidence)(nil),
		(*Evidence_PotentialAmnesiaEvidence)(nil),
		(*Evidence_MockEvidence)(nil),
		(*Evidence_MockRandomEvidence)(nil),
		(*Evidence_BeaconInactivityEvidence)(nil),
		(*Evidence_DkgEvidence)(nil),
	}
}

// EvidenceData contains any evidence of malicious wrong-doing by validators
type EvidenceData struct {
	Evidence             []Evidence `protobuf:"bytes,1,rep,name=evidence,proto3" json:"evidence"`
	Hash                 []byte     `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EvidenceData) Reset()         { *m = EvidenceData{} }
func (m *EvidenceData) String() string { return proto.CompactTextString(m) }
func (*EvidenceData) ProtoMessage()    {}
func (*EvidenceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{9}
}
func (m *EvidenceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvidenceData.Unmarshal(m, b)
}
func (m *EvidenceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvidenceData.Marshal(b, m, deterministic)
}
func (m *EvidenceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvidenceData.Merge(m, src)
}
func (m *EvidenceData) XXX_Size() int {
	return xxx_messageInfo_EvidenceData.Size(m)
}
func (m *EvidenceData) XXX_DiscardUnknown() {
	xxx_messageInfo_EvidenceData.DiscardUnknown(m)
}

var xxx_messageInfo_EvidenceData proto.InternalMessageInfo

func (m *EvidenceData) GetEvidence() []Evidence {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *EvidenceData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type ProofOfLockChange struct {
	Votes                []Vote         `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes"`
	PubKey               keys.PublicKey `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProofOfLockChange) Reset()         { *m = ProofOfLockChange{} }
func (m *ProofOfLockChange) String() string { return proto.CompactTextString(m) }
func (*ProofOfLockChange) ProtoMessage()    {}
func (*ProofOfLockChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_86495eef24aeacc0, []int{10}
}
func (m *ProofOfLockChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProofOfLockChange.Unmarshal(m, b)
}
func (m *ProofOfLockChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProofOfLockChange.Marshal(b, m, deterministic)
}
func (m *ProofOfLockChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofOfLockChange.Merge(m, src)
}
func (m *ProofOfLockChange) XXX_Size() int {
	return xxx_messageInfo_ProofOfLockChange.Size(m)
}
func (m *ProofOfLockChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofOfLockChange.DiscardUnknown(m)
}

var xxx_messageInfo_ProofOfLockChange proto.InternalMessageInfo

func (m *ProofOfLockChange) GetVotes() []Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *ProofOfLockChange) GetPubKey() keys.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return keys.PublicKey{}
}

func init() {
	proto.RegisterType((*DuplicateVoteEvidence)(nil), "tendermint.proto.types.DuplicateVoteEvidence")
	proto.RegisterType((*PotentialAmnesiaEvidence)(nil), "tendermint.proto.types.PotentialAmnesiaEvidence")
	proto.RegisterType((*MockEvidence)(nil), "tendermint.proto.types.MockEvidence")
	proto.RegisterType((*MockRandomEvidence)(nil), "tendermint.proto.types.MockRandomEvidence")
	proto.RegisterType((*ConflictingHeadersEvidence)(nil), "tendermint.proto.types.ConflictingHeadersEvidence")
	proto.RegisterType((*LunaticValidatorEvidence)(nil), "tendermint.proto.types.LunaticValidatorEvidence")
	proto.RegisterType((*BeaconInactivityEvidence)(nil), "tendermint.proto.types.BeaconInactivityEvidence")
	proto.RegisterType((*DKGEvidence)(nil), "tendermint.proto.types.DKGEvidence")
	proto.RegisterType((*Evidence)(nil), "tendermint.proto.types.Evidence")
	proto.RegisterType((*EvidenceData)(nil), "tendermint.proto.types.EvidenceData")
	proto.RegisterType((*ProofOfLockChange)(nil), "tendermint.proto.types.ProofOfLockChange")
}

func init() { proto.RegisterFile("proto/types/evidence.proto", fileDescriptor_86495eef24aeacc0) }

var fileDescriptor_86495eef24aeacc0 = []byte{
	// 992 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xf7, 0x3a, 0xb1, 0xe3, 0xbc, 0xb8, 0x34, 0x1d, 0x12, 0xba, 0xb2, 0x52, 0x52, 0x19, 0x44,
	0x5b, 0xfe, 0xac, 0x53, 0x07, 0x21, 0x8e, 0xc4, 0x0d, 0xc5, 0x51, 0x8a, 0x88, 0x36, 0xa8, 0x07,
	0x0e, 0xac, 0x66, 0x77, 0xc6, 0xbb, 0x23, 0xef, 0xce, 0xac, 0x76, 0xc7, 0x96, 0x7c, 0xe7, 0xc0,
	0x91, 0x0b, 0x1f, 0x83, 0x2b, 0x57, 0x0e, 0x5c, 0x38, 0xf3, 0x01, 0xca, 0x97, 0xe0, 0x03, 0xa0,
	0x9d, 0xd9, 0x7f, 0x6d, 0xb3, 0x96, 0xe1, 0x82, 0xe0, 0x62, 0xad, 0xdf, 0x9f, 0xdf, 0xef, 0xcd,
	0x7b, 0x6f, 0xde, 0x1b, 0x18, 0xc4, 0x89, 0x90, 0x62, 0x24, 0x57, 0x31, 0x4d, 0x47, 0x74, 0xc9,
	0x08, 0xe5, 0x1e, 0xb5, 0x94, 0x10, 0xbd, 0x25, 0x29, 0x27, 0x34, 0x89, 0x18, 0x97, 0x5a, 0x62,
	0x29, 0xb3, 0xc1, 0x7b, 0x32, 0x60, 0x09, 0x71, 0x62, 0x9c, 0xc8, 0xd5, 0x48, 0xfb, 0xfb, 0xc2,
	0x17, 0xd5, 0x97, 0xb6, 0x1e, 0xdc, 0xad, 0x63, 0xab, 0xdf, 0x5c, 0x71, 0xec, 0x0b, 0xe1, 0x87,
	0x54, 0xfb, 0xba, 0x8b, 0xd9, 0x48, 0xb2, 0x88, 0xa6, 0x12, 0x47, 0x71, 0x6e, 0x70, 0x4f, 0x7b,
	0x7a, 0xc9, 0x2a, 0x96, 0x62, 0x34, 0xa7, 0xab, 0x97, 0xfc, 0x87, 0xbf, 0x1a, 0x70, 0x78, 0xbe,
	0x88, 0x43, 0xe6, 0x61, 0x49, 0x9f, 0x0b, 0x49, 0x3f, 0xcf, 0x03, 0x47, 0x9f, 0xc1, 0x4e, 0xbc,
	0x70, 0x9d, 0x39, 0x5d, 0x99, 0xc6, 0x7d, 0xe3, 0xe1, 0xde, 0xf8, 0x81, 0xf5, 0xda, 0x21, 0x34,
	0xaa, 0x95, 0xa1, 0x5a, 0x57, 0x0b, 0x37, 0x64, 0xde, 0x25, 0x5d, 0xd9, 0xdd, 0x78, 0xe1, 0x5e,
	0xd2, 0x15, 0x3a, 0x85, 0xee, 0x52, 0x48, 0xea, 0x60, 0xb3, 0xad, 0x00, 0x8e, 0xac, 0x9b, 0xb3,
	0x60, 0x65, 0xbc, 0x76, 0x27, 0xb3, 0x3d, 0x2b, 0x9d, 0x5c, 0x73, 0x6b, 0x53, 0xa7, 0xc9, 0xf0,
	0x3b, 0x03, 0xcc, 0x2b, 0x21, 0x29, 0x97, 0x0c, 0x87, 0x67, 0x11, 0xa7, 0x29, 0xc3, 0xe5, 0x41,
	0xaa, 0x30, 0x8c, 0x7f, 0x12, 0x46, 0x7b, 0xf3, 0x30, 0x7e, 0x32, 0xa0, 0xff, 0xa5, 0xf0, 0xe6,
	0x25, 0xf5, 0x03, 0xb8, 0x5d, 0x34, 0x82, 0x13, 0x50, 0xe6, 0x07, 0x52, 0xc5, 0xb0, 0x65, 0xbf,
	0x51, 0x88, 0xa7, 0x4a, 0x8a, 0x2e, 0xe0, 0x56, 0x69, 0x98, 0x55, 0x30, 0x67, 0x1d, 0x58, 0xba,
	0xbc, 0x56, 0x51, 0x5e, 0xeb, 0xeb, 0xa2, 0xbc, 0x93, 0xde, 0x6f, 0x2f, 0x8e, 0x5b, 0x3f, 0xfc,
	0x71, 0x6c, 0xd8, 0xfd, 0xc2, 0x35, 0x53, 0xa2, 0x47, 0xb0, 0x5f, 0x42, 0x61, 0x42, 0x12, 0x9a,
	0xa6, 0x2a, 0x95, 0x7d, 0xbb, 0x8c, 0xe5, 0x4c, 0x8b, 0x87, 0xbf, 0x1b, 0x80, 0xb2, 0x78, 0x6d,
	0xcc, 0x89, 0x88, 0xfe, 0x23, 0x51, 0xa3, 0x7b, 0x00, 0x09, 0xe6, 0xc4, 0x71, 0x57, 0x92, 0xa6,
	0xe6, 0xb6, 0x32, 0xda, 0xcd, 0x24, 0x93, 0x4c, 0x30, 0xfc, 0xde, 0x80, 0xc1, 0x13, 0xc1, 0x67,
	0x21, 0xf3, 0x24, 0xe3, 0xfe, 0x94, 0x62, 0x42, 0x93, 0xb4, 0x3c, 0xdc, 0xc7, 0xd0, 0x0e, 0x1e,
	0xe7, 0x9d, 0xf0, 0x6e, 0x53, 0x51, 0xaf, 0x99, 0xcf, 0x29, 0xd1, 0xae, 0x76, 0x3b, 0x78, 0xac,
	0xbc, 0xc6, 0xf9, 0xf1, 0x36, 0xf5, 0x1a, 0x0f, 0x7f, 0x36, 0xc0, 0x7c, 0xb6, 0xe0, 0x58, 0x32,
	0xef, 0x39, 0x0e, 0x19, 0xc1, 0x52, 0x24, 0x65, 0x20, 0x9f, 0x40, 0x37, 0x50, 0xa6, 0x79, 0x30,
	0x6f, 0x37, 0xc1, 0xe6, 0x80, 0xb9, 0x35, 0x3a, 0x81, 0xed, 0xac, 0xdb, 0x36, 0xea, 0x4b, 0x65,
	0x89, 0x4e, 0xe0, 0x80, 0xf1, 0x65, 0x16, 0x80, 0xa3, 0x31, 0x9c, 0x19, 0xa3, 0x21, 0x51, 0xf9,
	0xdd, 0xb5, 0x51, 0xae, 0xd3, 0x34, 0x4f, 0x33, 0xcd, 0xf0, 0x45, 0x1b, 0xcc, 0x09, 0xc5, 0x9e,
	0xe0, 0x17, 0x1c, 0x7b, 0x92, 0x2d, 0x99, 0x5c, 0xfd, 0xab, 0xed, 0xf1, 0x01, 0xdc, 0x21, 0x74,
	0x46, 0x39, 0xc1, 0x5c, 0xbe, 0xd2, 0x1f, 0xfb, 0xa5, 0xa2, 0x68, 0x90, 0x11, 0xbc, 0xe9, 0x89,
	0x28, 0x0e, 0x31, 0xe3, 0x75, 0x73, 0xdd, 0x29, 0xa8, 0xa6, 0xaa, 0x75, 0x14, 0xa6, 0x82, 0x3b,
	0xa9, 0xc4, 0x89, 0x34, 0x3b, 0xea, 0x30, 0xbb, 0x99, 0xe4, 0x3a, 0x13, 0xa0, 0x53, 0x38, 0xac,
	0xe3, 0xa5, 0xcc, 0xe7, 0x58, 0x2e, 0x12, 0x6a, 0x76, 0x15, 0xe2, 0x41, 0x4d, 0x79, 0x5d, 0xe8,
	0xd0, 0x11, 0xec, 0xca, 0x20, 0xa1, 0x69, 0x20, 0x42, 0x62, 0xee, 0x68, 0xc8, 0x52, 0x30, 0xfc,
	0xb3, 0x0d, 0x7b, 0xe7, 0x97, 0x5f, 0xfc, 0x8f, 0x73, 0xfa, 0x08, 0xf6, 0x97, 0x45, 0xcf, 0x17,
	0x47, 0xd2, 0x99, 0xbd, 0x5d, 0xca, 0xf3, 0x33, 0x1d, 0x42, 0x97, 0xcc, 0x7d, 0x87, 0x11, 0x95,
	0xd0, 0x2d, 0xbb, 0x43, 0xe6, 0xfe, 0x05, 0x69, 0x4e, 0xfb, 0xce, 0xa6, 0x69, 0xef, 0xbd, 0x9a,
	0xf6, 0x5f, 0xba, 0xd0, 0x2b, 0x73, 0xee, 0xc3, 0x5d, 0x52, 0x6c, 0x3e, 0x47, 0x0d, 0xfb, 0x22,
	0x3d, 0xf9, 0x8d, 0xfc, 0xa8, 0xe9, 0x6e, 0xdd, 0xb8, 0x30, 0xa7, 0x2d, 0xfb, 0x90, 0xdc, 0xb8,
	0x49, 0x97, 0x70, 0xe4, 0x55, 0x03, 0x29, 0xbf, 0x83, 0x69, 0xc5, 0xa6, 0x4b, 0x38, 0x6e, 0x62,
	0x6b, 0x1e, 0x66, 0xd3, 0x96, 0x3d, 0xf0, 0x9a, 0x47, 0x5d, 0x0c, 0x83, 0x50, 0x4f, 0x1f, 0xa7,
	0x2a, 0x45, 0xc9, 0xaa, 0xd7, 0xeb, 0x49, 0x13, 0x6b, 0xd3, 0xdc, 0x9a, 0xb6, 0x6c, 0x33, 0x6c,
	0x9a, 0x69, 0x31, 0x0c, 0xe2, 0x62, 0x0d, 0x3b, 0x58, 0xef, 0xe1, 0x8a, 0x71, 0x7b, 0x3d, 0x63,
	0xd3, 0x02, 0xcf, 0x18, 0xe3, 0xa6, 0xe5, 0x7e, 0x09, 0xb7, 0x22, 0xe1, 0xcd, 0x2b, 0x92, 0xce,
	0xfa, 0x19, 0x5d, 0x5f, 0xcf, 0xd3, 0x96, 0xdd, 0x8f, 0xea, 0xeb, 0xfa, 0x5b, 0x38, 0x50, 0x60,
	0x89, 0xda, 0x87, 0x15, 0x66, 0x57, 0x61, 0xbe, 0xbf, 0x0e, 0xf3, 0xe5, 0x15, 0x3a, 0x6d, 0xd9,
	0x28, 0x7a, 0x7d, 0xb1, 0xc6, 0x30, 0x70, 0xd5, 0x54, 0x75, 0x58, 0x39, 0x56, 0x2b, 0x96, 0x9d,
	0xf5, 0xe9, 0x69, 0x9a, 0xc7, 0x59, 0x7a, 0xdc, 0xa6, 0x59, 0x3d, 0x85, 0x7e, 0x76, 0xb5, 0x4a,
	0x8e, 0x9e, 0xe2, 0x78, 0xa7, 0xb1, 0xb1, 0xab, 0x91, 0x34, 0x6d, 0xd9, 0x7b, 0x64, 0xee, 0x17,
	0x7f, 0x27, 0x1d, 0xd8, 0x4a, 0x17, 0xd1, 0x70, 0x06, 0xfd, 0x42, 0x74, 0x8e, 0x25, 0x46, 0x13,
	0xe8, 0xd5, 0x6e, 0xcd, 0xd6, 0xc3, 0xbd, 0xf1, 0xfd, 0x26, 0xf0, 0x12, 0x6a, 0x3b, 0x1b, 0x48,
	0x76, 0xe9, 0x87, 0x10, 0x6c, 0x07, 0x38, 0x0d, 0xd4, 0x3d, 0xe8, 0xdb, 0xea, 0x7b, 0xf8, 0xa3,
	0x01, 0x77, 0xae, 0x12, 0x21, 0x66, 0x5f, 0xcd, 0x9e, 0x09, 0x6f, 0xfe, 0x24, 0xc0, 0xdc, 0xa7,
	0xe8, 0x53, 0x50, 0x2f, 0xad, 0x34, 0xa7, 0x5a, 0xbb, 0xfc, 0x72, 0x1a, 0xed, 0x80, 0x9e, 0x56,
	0xaf, 0xd9, 0xf6, 0xdf, 0x7a, 0xcd, 0xe6, 0x30, 0xf9, 0x9b, 0x76, 0x62, 0x7d, 0xf3, 0xa1, 0xcf,
	0x64, 0xb0, 0x70, 0x2d, 0x4f, 0x44, 0xa3, 0x0a, 0xa2, 0xfe, 0x59, 0x7b, 0xab, 0xbb, 0x5d, 0xf5,
	0xe7, 0xf4, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x55, 0x8f, 0x21, 0x1d, 0x0c, 0x00, 0x00,
}
