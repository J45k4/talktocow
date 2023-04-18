package call

import "github.com/google/uuid"

type Role int

const (
	RoleGuest Role = iota
	RoleHost
)

type MemberState int

const (
	MemberStateWaiting MemberState = iota
	MemberStateConnecting
	MemberStateConnected
)

type CallDevice struct {
	id string
}

type CallMember struct {
	UserId  int
	Role    Role
	State   MemberState
	Devices []CallDevice
}

type Call struct {
	id      string
	members map[int]*CallMember
}

func (c *Call) Join(userId int) {
	c.members[userId] = &CallMember{
		UserId: userId,
		Role:   RoleGuest,
	}
}

func (c *Call) Leave(userId int) {
	delete(c.members, userId)
}

func (c *Call) Invite(invitor int, invitee int) {
	member := c.members[invitee]

	member.Role = RoleGuest
}

func (c *Call) Accept(acceptor int, userId int) {
	member := c.members[userId]
	member.State = MemberStateConnecting
}

func (c *Call) AddDevice(userId int, deviceId string) {
	member := c.members[userId]

	member.Devices = append(member.Devices, CallDevice{
		id: deviceId,
	})
}

type CallManager struct {
	calls map[string]*Call
}

func (cm *CallManager) NewCall() *Call {
	newCall := &Call{
		id: uuid.New().String(),
	}

	cm.calls[newCall.id] = newCall

	return newCall
}

func (cm *CallManager) GetCall(callId string) *Call {
	return cm.calls[callId]
}

func (cm *CallManager) EndCall(callId string) {
	delete(cm.calls, callId)
}
