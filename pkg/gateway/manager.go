package gateway

import (
	"github.com/jasonsoft/wakanda/internal/hash"
)

type Manager struct {
	gatewayAddr string
	buckets     []*Bucket
	status      *Status
}

func NewManager() *Manager {
	m := &Manager{
		buckets: make([]*Bucket, 1024),
		status:  &Status{},
	}

	// inital bucket setting
	for idx, _ := range m.buckets {
		m.buckets[idx] = NewBucket(idx, 32)
	}
	return m
}

func (m *Manager) Status() *Status {
	return m.status
}

func (m *Manager) SetGatewayAddr(gatewayAddr string) {
	m.gatewayAddr = gatewayAddr
}

func (m *Manager) BucketBySessionID(sessionID string) *Bucket {
	hashNumber := hash.FNV32a(sessionID)
	return m.buckets[hashNumber%uint32(len(m.buckets))]
}

func (m *Manager) AddSession(session *WSSession) {
	bucket := m.BucketBySessionID(session.ID)
	bucket.addSession(session)
	m.status.IncreaseOnlinePeople()
}

func (m *Manager) DeleteSession(session *WSSession) {
	bucket := m.BucketBySessionID(session.ID)
	bucket.deleteSession(session)
	m.status.DecreaseOnlinePeople()

	// leave room
	session.rooms.Range(func(key, _ interface{}) bool {
		roomID, ok := key.(string)
		if ok {
			bucket.leaveRoom(roomID, session)
		}
		return true
	})
}

func (m *Manager) JoinRoom(roomID string, session *WSSession) {
	bucket := m.BucketBySessionID(session.ID)
	bucket.joinRoom(roomID, session)
}

func (m *Manager) LeaveRoom(roomID string, session *WSSession) {
	bucket := m.BucketBySessionID(session.ID)
	bucket.leaveRoom(roomID, session)
}

func (m *Manager) Push(sessionID string, command *Command) {
	b := m.BucketBySessionID(sessionID)
	b.push(sessionID, command)
}

func (m *Manager) PushAll(command *Command) {
	job := Job{
		OP:      OP_PUSH_ALL,
		Command: command,
	}
	for _, bucket := range m.buckets {
		bucket.jobChan <- job
	}
}

func (m *Manager) PushRoom(roomID string, command *Command) {
	job := Job{
		RoomID:  roomID,
		OP:      OP_PUSH_ROOM,
		Command: command,
	}
	for _, bucket := range m.buckets {
		bucket.jobChan <- job
	}
}
