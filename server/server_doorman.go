package server

import (
	"context"
	"fmt"
	"github.com/fuks-kit/doorman/challenge"
	"github.com/fuks-kit/doorman/door"
	pb "github.com/fuks-kit/doorman/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"log"
	"time"
)

type DoormanServer struct {
	pb.DoormanServer
}

func NewDoormanServer() *DoormanServer {
	return &DoormanServer{}
}

func (server *DoormanServer) CheckAccess(ctx context.Context, req *pb.AccessCheckRequest) (*pb.AccessCheckResponse, error) {
	log.Printf("CheckAccess: challenge=%v", req.Challenge)

	if ok := challenge.Validate(req.Challenge); !ok {
		return nil, fmt.Errorf("invalid challenge")
	}

	permission, err := verifyPermission(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return &pb.AccessCheckResponse{
		HasAccess:    permission.HasAccess,
		IsFuks:       permission.IsFuksMember,
		IsActiveFuks: permission.IsActiveFuks,
	}, nil
}

func (server *DoormanServer) OpenDoor(ctx context.Context, req *pb.DoorOpenRequest) (*pb.DoorOpenResponse, error) {
	log.Printf("OpenDoor: challenge=%v", req.Challenge)

	if ok := challenge.Validate(req.Challenge); !ok {
		return nil, fmt.Errorf("invalid challenge")
	}

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

	return &pb.DoorOpenResponse{
		Open:         true,
		OpenDuration: durationpb.New(accessDuration),
	}, nil
}
