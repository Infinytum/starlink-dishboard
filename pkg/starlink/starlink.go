package starlink

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-mojito/mojito"
	"github.com/infinytum/injector"
	"github.com/nakabonne/tstorage"
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
		s.startPollingLoop()
		if err != nil {
			panic(err)
		}
		return *s
	})
}

type Service struct {
	conn    *grpc.ClientConn
	client  pb.DeviceClient
	stopped bool
	storage tstorage.Storage
}

func (s *Service) startPollingLoop() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		for {
			select {
			case <-time.After(time.Second):
				s.pollDishy()
			case <-c:
				mojito.DefaultLogger().Info("Gracefully shutting down timeseries DB...")
				s.stopped = true
				s.storage.Close()
				mojito.DefaultLogger().Info("Timeseries DB saved!")
				os.Exit(0)
			}
		}
	}()
}

func (s *Service) Ping() (float64, error) {
	datapoints, err := s.storage.Select("latency", nil, time.Now().Add(-2*time.Second).Unix(), time.Now().Unix())
	if err != nil {
		return 0, err
	}
	return datapoints[0].Value, nil
}

func (s *Service) PingHistory(start, end time.Time) (map[int64]float64, error) {
	datapoints, err := s.storage.Select("latency", nil, start.Unix(), end.Unix())
	if err != nil {
		return nil, err
	}
	latencyMap := make(map[int64]float64)
	for _, datapoint := range datapoints {
		latencyMap[datapoint.Timestamp] = datapoint.Value
	}
	return latencyMap, nil
}

func (s *Service) GetPingHistory() ([]float32, error) {
	in := new(pb.Request)
	in.Request = &pb.Request_GetHistory{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// r is the gprc response
	r, err := s.client.Handle(ctx, in)
	if err != nil {
		return nil, err
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
		return nil, nil, err
	}

	pingResponse := r.Response.(*pb.Response_DishGetHistory)
	return pingResponse.DishGetHistory.DownlinkThroughputBps, pingResponse.DishGetHistory.UplinkThroughputBps, nil
}

func (s *Service) pollDishy() error {
	if s.stopped {
		return nil
	}
	in := new(pb.Request)
	in.Request = &pb.Request_GetStatus{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.client.Handle(ctx, in)
	if err != nil {
		return err
	}

	statusResponse := r.Response.(*pb.Response_DishGetStatus)
	timestamp := time.Now().Unix()
	return s.storage.InsertRows([]tstorage.Row{
		{
			Metric:    "latency",
			DataPoint: tstorage.DataPoint{Timestamp: timestamp, Value: float64(statusResponse.DishGetStatus.PopPingLatencyMs)},
		},
		{
			Metric:    "down",
			DataPoint: tstorage.DataPoint{Timestamp: timestamp, Value: float64(statusResponse.DishGetStatus.DownlinkThroughputBps)},
		},
		{
			Metric:    "up",
			DataPoint: tstorage.DataPoint{Timestamp: timestamp, Value: float64(statusResponse.DishGetStatus.UplinkThroughputBps)},
		},
	})
}

func NewService() (*Service, error) {
	conn, err := grpc.Dial(dishyAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	storage, err := tstorage.NewStorage(
		tstorage.WithTimestampPrecision(tstorage.Seconds),
		tstorage.WithDataPath("./data"),
	)
	if err != nil {
		return nil, err
	}
	return &Service{
		conn:    conn,
		client:  pb.NewDeviceClient(conn),
		storage: storage,
	}, nil
}
