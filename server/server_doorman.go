package server

import (
	"context"
	"fmt"
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

	permission, err := verifyPermission(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.OfficePermission{
		HasAccess:    permission.HasAccess,
		IsFuksMember: permission.IsFuksMember,
		IsActiveFuks: permission.IsFuksMember,
	}, nil
}

func (server *DoormanServer) OpenDoor(ctx context.Context, _ *emptypb.Empty) (*pb.DoorState, error) {
	log.Printf("OpenDoor:")

	permission, err := verifyPermission(ctx)
	if err != nil {
		return nil, err
	}

	if !permission.HasAccess {
		return nil, fmt.Errorf("access denied")
	}

	accessDuration := time.Second * 5
	// TODO: uncomment when door is connected
	// go door.Open(accessDuration)

	return &pb.DoorState{
		Open:         true,
		OpenDuration: durationpb.New(accessDuration),
	}, nil
}
