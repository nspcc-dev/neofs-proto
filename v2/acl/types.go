package acl

import "github.com/nspcc-dev/neofs-api-go/v2/refs"

// HeaderFilter is a unified structure of FilterInfo
// message from proto definition.
type HeaderFilter struct {
	hdrType HeaderType

	matchType MatchType

	name, value string
}

// TargetInfo is a unified structure of TargetInfo
// message from proto definition.
type TargetInfo struct {
	target Target

	keys [][]byte
}

// Record is a unified structure of EACLRecord
// message from proto definition.
type Record struct {
	op Operation

	action Action

	filters []*HeaderFilter

	targets []*TargetInfo
}

// Table is a unified structure of EACLTable
// message from proto definition.
type Table struct {
	cid *refs.ContainerID

	records []*Record
}

type TokenLifetime struct {
	exp, nbf, iat uint64
}

type BearerTokenBody struct {
	eacl *Table

	ownerID *refs.OwnerID

	lifetime *TokenLifetime
}

type BearerToken struct {
	body *BearerTokenBody

	sig *refs.Signature
}

// TargetInfo is a unified enum of MatchType enum from proto definition.
type MatchType uint32

// HeaderType is a unified enum of HeaderType enum from proto definition.
type HeaderType uint32

// Action is a unified enum of Action enum from proto definition.
type Action uint32

// Operation is a unified enum of Operation enum from proto definition.
type Operation uint32

// Target is a unified enum of Target enum from proto definition.
type Target uint32

const (
	MatchTypeUnknown MatchType = iota
	MatchTypeStringEqual
	MatchTypeStringNotEqual
)

const (
	HeaderTypeUnknown HeaderType = iota
	HeaderTypeRequest
	HeaderTypeObject
)

const (
	ActionUnknown Action = iota
	ActionAllow
	ActionDeny
)

const (
	OperationUnknown Operation = iota
	OperationGet
	OperationHead
	OperationPut
	OperationDelete
	OperationSearch
	OperationRange
	OperationRangeHash
)

const (
	TargetUnknown Target = iota
	TargetUser
	TargetSystem
	TargetOthers
)

func (f *HeaderFilter) GetHeaderType() HeaderType {
	if f != nil {
		return f.hdrType
	}

	return HeaderTypeUnknown
}

func (f *HeaderFilter) SetHeaderType(v HeaderType) {
	if f != nil {
		f.hdrType = v
	}
}

func (f *HeaderFilter) GetMatchType() MatchType {
	if f != nil {
		return f.matchType
	}

	return MatchTypeUnknown
}

func (f *HeaderFilter) SetMatchType(v MatchType) {
	if f != nil {
		f.matchType = v
	}
}

func (f *HeaderFilter) GetName() string {
	if f != nil {
		return f.name
	}

	return ""
}

func (f *HeaderFilter) SetName(v string) {
	if f != nil {
		f.name = v
	}
}

func (f *HeaderFilter) GetValue() string {
	if f != nil {
		return f.value
	}

	return ""
}

func (f *HeaderFilter) SetValue(v string) {
	if f != nil {
		f.value = v
	}
}

func (t *TargetInfo) GetTarget() Target {
	if t != nil {
		return t.target
	}

	return TargetUnknown
}

func (t *TargetInfo) SetTarget(v Target) {
	if t != nil {
		t.target = v
	}
}

func (t *TargetInfo) GetKeyList() [][]byte {
	if t != nil {
		return t.keys
	}

	return nil
}

func (t *TargetInfo) SetKeyList(v [][]byte) {
	if t != nil {
		t.keys = v
	}
}

func (r *Record) GetOperation() Operation {
	if r != nil {
		return r.op
	}

	return OperationUnknown
}

func (r *Record) SetOperation(v Operation) {
	if r != nil {
		r.op = v
	}
}

func (r *Record) GetAction() Action {
	if r != nil {
		return r.action
	}

	return ActionUnknown
}

func (r *Record) SetAction(v Action) {
	if r != nil {
		r.action = v
	}
}

func (r *Record) GetFilters() []*HeaderFilter {
	if r != nil {
		return r.filters
	}

	return nil
}

func (r *Record) SetFilters(v []*HeaderFilter) {
	if r != nil {
		r.filters = v
	}
}

func (r *Record) GetTargets() []*TargetInfo {
	if r != nil {
		return r.targets
	}

	return nil
}

func (r *Record) SetTargets(v []*TargetInfo) {
	if r != nil {
		r.targets = v
	}
}

func (t *Table) GetContainerID() *refs.ContainerID {
	if t != nil {
		return t.cid
	}

	return nil
}

func (t *Table) SetContainerID(v *refs.ContainerID) {
	if t != nil {
		t.cid = v
	}
}

func (t *Table) GetRecords() []*Record {
	if t != nil {
		return t.records
	}

	return nil
}

func (t *Table) SetRecords(v []*Record) {
	if t != nil {
		t.records = v
	}
}

func (l *TokenLifetime) GetExp() uint64 {
	if l != nil {
		return l.exp
	}

	return 0
}

func (l *TokenLifetime) SetExp(v uint64) {
	if l != nil {
		l.exp = v
	}
}

func (l *TokenLifetime) GetNbf() uint64 {
	if l != nil {
		return l.nbf
	}

	return 0
}

func (l *TokenLifetime) SetNbf(v uint64) {
	if l != nil {
		l.nbf = v
	}
}

func (l *TokenLifetime) GetIat() uint64 {
	if l != nil {
		return l.iat
	}

	return 0
}

func (l *TokenLifetime) SetIat(v uint64) {
	if l != nil {
		l.iat = v
	}
}

func (bt *BearerTokenBody) GetEACL() *Table {
	if bt != nil {
		return bt.eacl
	}

	return nil
}

func (bt *BearerTokenBody) SetEACL(v *Table) {
	if bt != nil {
		bt.eacl = v
	}
}

func (bt *BearerTokenBody) GetOwnerID() *refs.OwnerID {
	if bt != nil {
		return bt.ownerID
	}

	return nil
}

func (bt *BearerTokenBody) SetOwnerID(v *refs.OwnerID) {
	if bt != nil {
		bt.ownerID = v
	}
}

func (bt *BearerTokenBody) GetLifetime() *TokenLifetime {
	if bt != nil {
		return bt.lifetime
	}

	return nil
}

func (bt *BearerTokenBody) SetLifetime(v *TokenLifetime) {
	if bt != nil {
		bt.lifetime = v
	}
}

func (bt *BearerToken) GetBody() *BearerTokenBody {
	if bt != nil {
		return bt.body
	}

	return nil
}

func (bt *BearerToken) SetBody(v *BearerTokenBody) {
	if bt != nil {
		bt.body = v
	}
}

func (bt *BearerToken) GetSignature() *refs.Signature {
	if bt != nil {
		return bt.sig
	}

	return nil
}

func (bt *BearerToken) SetSignature(v *refs.Signature) {
	if bt != nil {
		bt.sig = v
	}
}
