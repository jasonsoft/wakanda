package http

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jasonsoft/log"
	"github.com/jasonsoft/napnap"
	"github.com/jasonsoft/wakanda/pkg/dispatcher/proto"
	dispatcherProto "github.com/jasonsoft/wakanda/pkg/dispatcher/proto"
	"github.com/jasonsoft/wakanda/pkg/gateway"
	"github.com/jasonsoft/wakanda/pkg/identity"
	routerProto "github.com/jasonsoft/wakanda/pkg/router/proto"
	"github.com/satori/go.uuid"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func NewGatewayRouter(h *GatewayHttpHandler) *napnap.Router {
	router := napnap.NewRouter()
	router.Get("/ws", h.wsEndpoint)
	router.Get("/rooms/:room_id", h.roomEndpoint)
	return router
}

type GatewayHttpHandler struct {
	manager          *gateway.Manager
	dispatcherClient proto.DispatcherServiceClient
	routerClient     routerProto.RouterServiceClient
}

func NewGatewayHttpHandler(manager *gateway.Manager, dispatcherClient proto.DispatcherServiceClient, routerClient routerProto.RouterServiceClient) *GatewayHttpHandler {
	return &GatewayHttpHandler{
		manager:          manager,
		dispatcherClient: dispatcherClient,
		routerClient:     routerClient,
	}
}

func (h *GatewayHttpHandler) wsEndpoint(c *napnap.Context) {
	ctx := c.StdContext()

	defer func() {
		log.Debug("gateway: ws socket endpoint end")
	}()

	claims, found := identity.FromContext(ctx)
	if found == false {
		c.SetStatus(403)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	sessionID := uuid.NewV4().String()
	wsSession := gateway.NewWSSession(sessionID, *claims, conn, h.manager, h.dispatcherClient, h.routerClient, "")
	wsSession.StartTasks()
}

type GatewayStatus struct {
	RoomStatus map[string]int
}

func (h *GatewayHttpHandler) roomEndpoint(c *napnap.Context) {
	ctx := c.StdContext()
	roomID := c.Param("room_id")

	defer func() {
		log.Debug("gateway: ws socket endpoint end")
	}()

	claims, found := identity.FromContext(ctx)
	if found == false {
		c.SetStatus(403)
		return
	}

	// ensure the member can join the room
	in := &dispatcherProto.DispatcherCommandRequest{
		OP:              "JOINROOM",
		Data:            []byte(roomID),
		SenderID:        claims["account_id"].(string),
		SenderFirstName: claims["first_name"].(string),
		SenderLastName:  claims["last_name"].(string),
	}

	handleCommandReply, err := h.dispatcherClient.HandleCommand(ctx, in)
	if err != nil {
		log.Errorf("gateway: command error from dispatcher server: %v", err)
		panic(err)
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	sessionID := uuid.NewV4().String()
	wsSession := gateway.NewWSSession(sessionID, *claims, conn, h.manager, h.dispatcherClient, h.routerClient, roomID)

	h.manager.JoinRoom(roomID, wsSession)
	wsSession.StartTasks()
}
