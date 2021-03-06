package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	messengerProto "github.com/jasonsoft/wakanda/pkg/messenger/proto"
	"google.golang.org/grpc"

	"github.com/jasonsoft/log"
	"github.com/jasonsoft/napnap"
	"github.com/jasonsoft/wakanda/internal/config"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			// unknown error
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("unknown error: %v", err)
			}
			log.StackTrace().Error(err)
		}
	}()

	config := config.New("app.yml")
	err := initialize(config)
	if err != nil {
		log.Fatalf("messageer: main initialize failed: %v", err)
	}

	// start grpc server
	lis, err := net.Listen("tcp", config.Messenger.GRPCBind)
	if err != nil {
		log.Fatalf("messageer: bind grpc failed: %v", err)
	}
	s := grpc.NewServer()

	messengerProto.RegisterMessageServiceServer(s, _messageServer)
	go func() {
		log.Info("messenger: grpc service started")
		if err = s.Serve(lis); err != nil {
			log.Fatalf("messenger: failed to start grpc server: %v", err)
		}
	}()

	// start http service
	nap := napWithMiddlewares()
	httpEngine := napnap.NewHttpEngine(config.Messenger.HTTPBind)
	go func() {
		log.Info("messenger: http service started")
		err := nap.Run(httpEngine)
		if err != nil {
			log.Error(err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)
	<-stopChan
	log.Info("messenger: http shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpEngine.Shutdown(ctx); err != nil {
		log.Errorf("messenger: http hanlder shutdown error: %v", err)
	} else {
		log.Info("messenger: http gracefully stopped")
	}
}
