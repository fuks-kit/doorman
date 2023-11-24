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
	if ok := challenge.Validate(req.Challenge); !ok {
		return nil, fmt.Errorf("invalid challenge")
	}

	user, err := verifyUser(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return &pb.AccessCheckResponse{
		HasAccess:    user.HasAccess,
		IsFuks:       user.IsFuksMember,
		IsActiveFuks: user.IsActiveFuks,
	}, nil
}

func (server *DoormanServer) OpenDoor(ctx context.Context, req *pb.DoorOpenRequest) (*pb.DoorOpenResponse, error) {
	if ok := challenge.Validate(req.Challenge); !ok {
		return nil, fmt.Errorf("invalid challenge")
	}

	user, err := verifyUser(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	if !user.HasAccess {
		return nil, fmt.Errorf("permission denied")
	}

	log.Printf("Opening door for uid=%s name=%s email=%s", user.UID, user.Name, user.EMail)

	accessDuration := time.Second * 4

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
