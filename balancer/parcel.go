package balancer

import (
	"encoding/json"
	"time"
)

const (
	RequestIdentityParcel int = iota
	ResponseIdentityParcel

	AssignmentParcel

	// HeartbeatParcel just keeps the bee in sync with balancer
	HeartbeatParcel

	// ChangeUserParcel is used to add/remove, active/deactivate a user
	ChangeUserParcel

	// AuditMessageParcel tells a slave to conduct an audit
	AuditMessageParcel

	// RebalanceUserParcel is sent from a bee, telling the balancer to
	// give another bee the user
	RebalanceUserParcel
)

type Parcel struct {
	// Header
	ID   string // Serves as 'To:' and 'From:' ID. It always refers to a Bee
	Type int

	// Body
	Message []byte
}

func NewRequestIDParcel(publicKey []byte) *Parcel {
	p := newParcel("", RequestIdentityParcel)
	p.Message = publicKey
	return p
}

type IDResponse struct {
	ID        string
	Users     []*User
	PublicKey []byte
}

func NewResponseIDParcel(id string, users []*User, public []byte) *Parcel {
	p := newParcel(id, ResponseIdentityParcel)

	m := new(IDResponse)
	m.ID = id
	m.Users = users
	m.PublicKey = public

	msg, _ := json.Marshal(m)
	p.Message = msg

	return p
}

type Assignment struct {
	Users []*User
}

func NewAssignment(id string, a Assignment) *Parcel {
	p := newParcel(id, AssignmentParcel)
	msg, _ := json.Marshal(a)
	p.Message = msg

	return p
}

type NewChangeUser struct {
	U      User
	Add    bool
	Active bool
}

func NewChangeUserParcelFromStruct(id string, cus NewChangeUser) *Parcel {
	p := newParcel(id, ChangeUserParcel)
	msg, _ := json.Marshal(cus)
	p.Message = msg

	return p
}

func NewChangeUserParcel(id string, u User, add, active bool) *Parcel {
	m := new(NewChangeUser)
	m.U = u
	m.Active = active
	m.Add = add

	return NewChangeUserParcelFromStruct(id, *m)
}

type Heartbeat struct {
	SentTime    time.Time
	Users       []*User
	ApiRate     float64
	LoanJobRate float64
}

func NewHeartbeat(id string, h Heartbeat) *Parcel {
	p := newParcel(id, HeartbeatParcel)
	msg, _ := json.Marshal(&h)
	p.Message = msg

	return p
}

type AuditParcel struct {
	Users []*User
}

func NewAuditParcel(id string, a AuditParcel) *Parcel {
	p := newParcel(id, AuditMessageParcel)
	msg, _ := json.Marshal(&a)
	p.Message = msg

	return p
}

type RebalanceUser struct {
	U User
}

func NewRebalanceUserParcel(id string, u User) *Parcel {
	p := newParcel(id, RebalanceUserParcel)

	ru := RebalanceUser{U: u}
	msg, _ := json.Marshal(&ru)
	p.Message = msg

	return p
}

func newParcel(id string, t int) *Parcel {
	p := new(Parcel)
	p.Type = t
	p.ID = id

	return p
}
