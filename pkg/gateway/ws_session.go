package gateway

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jasonsoft/log"
	"github.com/jasonsoft/wakanda/internal/types"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	readWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = 5

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type WSSession struct {
	ID      uint64
	member  *types.Member
	socket  *websocket.Conn
	inChan  chan *WSMessage
	outChan chan *WSMessage
}

func NewWSSession(id uint64, member *types.Member, conn *websocket.Conn) *WSSession {
	return &WSSession{
		ID:      id,
		member:  member,
		socket:  conn,
		inChan:  make(chan *WSMessage, 1024),
		outChan: make(chan *WSMessage, 1024),
	}
}

func (s *WSSession) readLoop() {
	defer func() {
		s.socket.Close()
	}()

	s.socket.SetReadLimit(maxMessageSize)
	s.socket.SetPongHandler(func(string) error { s.socket.SetReadDeadline(time.Now().Add(readWait)); return nil })

	var (
		msgType int
		msgData []byte
		message *WSMessage
		err     error
	)

	for {
		s.socket.SetReadDeadline(time.Now().Add(readWait))
		msgType, msgData, err = s.socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorf("gateway: websocket message error: %v", err)
			}
			break
		}

		message = &WSMessage{
			MsgType: msgType,
			MsgData: msgData,
		}

		select {
		case s.inChan <- message:
		}
	}
}

func (s *WSSession) writeLoop() {
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		s.socket.Close()
	}()

	var (
		message *WSMessage
		err     error
	)
	for {
		select {
		case message = <-s.outChan:
			s.socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err = s.socket.WriteMessage(message.MsgType, message.MsgData); err != nil {
				return
			}
		case <-pingTicker.C:
			s.socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := s.socket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (s *WSSession) ReadMessage() *WSMessage {
	select {
	case message := <-s.inChan:
		return message
	}
}

func (s *WSSession) SendMessage(msg *WSMessage) {
	select {
	case s.outChan <- msg:
	}
}

func (s *WSSession) StartTasks() {
	go s.readLoop()
	go s.writeLoop()

	_manager.AddSession(s)

	var (
		message     *WSMessage
		commandReq  *Command
		commandResp *Command
		err         error
		buf         []byte
	)

	for {
		message = s.ReadMessage()

		if message.MsgType != websocket.TextMessage {
			continue
		}

		commandReq, err = CreateCommand(message.MsgData)
		commandResp = nil

		// handles all commands here
		switch commandReq.OP {
		case "JOIN":
			commandResp, err = s.handleJoin(commandReq)
			if err != nil {
				log.Errorf("gateway: handle JOIN command error: %v", err)
				continue
			}
		case "LEAVE":
			commandResp, err = s.handleLeave(commandReq)
			if err != nil {
				log.Errorf("gateway: handle LEAVE command error: %v", err)
				continue
			}
		}

		if commandResp != nil {
			buf, err = json.Marshal(*commandResp)
			if err != nil {
				continue
			}

			message = &WSMessage{websocket.TextMessage, buf}
			s.SendMessage(message)
		}
	}
}
