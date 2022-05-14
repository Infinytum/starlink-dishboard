package starlink

import (
	"context"
	"log"
	"time"

	"github.com/infinytum/injector"
	pb "github.com/starlink-community/starlink-grpc-go/pkg/spacex.com/api/device"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	dishyAddress = "192.168.100.1:9200"
)

func init() {
	injector.Singleton(func() Service {
		s, err := NewService()
		if err != nil {
			panic(err)
		}
		return *s
	})
}

type Service struct {
	conn   *grpc.ClientConn
	client pb.DeviceClient
}

func (s *Service) Ping() (float32, error) {
	in := new(pb.Request)
	in.Request = &pb.Request_GetStatus{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// r is the gprc response
	r, err := s.client.Handle(ctx, in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	pingResponse := r.Response.(*pb.Response_DishGetStatus)
	return pingResponse.DishGetStatus.PopPingLatencyMs, nil
}

func (s *Service) GetPingHistory() ([]float32, error) {
	in := new(pb.Request)
	in.Request = &pb.Request_GetHistory{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// r is the gprc response
	r, err := s.client.Handle(ctx, in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	pingResponse := r.Response.(*pb.Response_DishGetHistory)
	return pingResponse.DishGetHistory.PopPingLatencyMs, nil
}

func (s *Service) GetTafficHistory() ([]float32, []float32, error) {
	in := new(pb.Request)
	in.Request = &pb.Request_GetHistory{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// r is the gprc response
	r, err := s.client.Handle(ctx, in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	pingResponse := r.Response.(*pb.Response_DishGetHistory)
	return pingResponse.DishGetHistory.DownlinkThroughputBps, pingResponse.DishGetHistory.UplinkThroughputBps, nil
}

func NewService() (*Service, error) {
	conn, err := grpc.Dial(dishyAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		return nil, err
	}
	return &Service{
		conn:   conn,
		client: pb.NewDeviceClient(conn),
	}, nil
}
