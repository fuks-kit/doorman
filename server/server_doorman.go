package server

import (
	"context"
	"fmt"
	"github.com/fuks-kit/doorman/door"
	pb "github.com/fuks-kit/doorman/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

type DoormanServer struct {
	pb.DoormanServer
}

func NewDoormanServer() *DoormanServer {
	return &DoormanServer{}
}

func (server *DoormanServer) CheckPermissions(ctx context.Context, _ *emptypb.Empty) (*pb.OfficePermission, error) {
	log.Printf("CheckPermissions:")

	start := time.Now()
	permission, err := verifyPermission(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	log.Printf("CheckPermissions: took %s", time.Since(start))

	return &pb.OfficePermission{
		HasAccess:    permission.HasAccess,
		IsFuksMember: permission.IsFuksMember,
		IsActiveFuks: permission.IsActiveFuks,
	}, nil
}

func (server *DoormanServer) OpenDoor(ctx context.Context, _ *emptypb.Empty) (*pb.DoorState, error) {
	log.Printf("OpenDoor:")

	permission, err := verifyPermission(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	if !permission.HasAccess {
		return nil, fmt.Errorf("permission denied")
	}

	accessDuration := time.Second * 5

	err = door.Open(accessDuration)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return &pb.DoorState{
		Open:         true,
		OpenDuration: durationpb.New(accessDuration),
	}, nil
}
