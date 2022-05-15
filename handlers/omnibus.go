package handlers

import (
	"sort"
	"strings"
	"time"

	"github.com/go-mojito/mojito"
	"github.com/infinytum/starlink-dishboard/pkg/starlink"
)

type ChartDataPoint struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type OmnibusType string

const (
	OmnibusInitial     OmnibusType = "INIT"
	OmnibusChartUpdate OmnibusType = "CHART_UPDATE"
)

type OmnibusPacket struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Type  OmnibusType `json:"type,omitempty"`
}

type OmnibusDataInit struct {
	Down    []ChartDataPoint `json:"down,omitempty"`
	Latency []ChartDataPoint `json:"latency,omitempty"`
	Up      []ChartDataPoint `json:"up,omitempty"`
}

type OmnibusDataChartUpdate struct {
	Down    ChartDataPoint `json:"down,omitempty"`
	Latency ChartDataPoint `json:"latency,omitempty"`
	Up      ChartDataPoint `json:"up,omitempty"`
}

func Omnibus(ctx mojito.WebSocketContext, sl *starlink.Service) error {
	timeframe := ctx.Request().ParamOrDefault("timeframe", "live")
	var start, end time.Time
	var refreshDuration time.Duration

	switch strings.ToLower(timeframe) {
	case "today":
		start = time.Now().Add(-time.Hour * 24)
		end = time.Now()
		refreshDuration = time.Minute
	case "week":
		start = time.Now().Add(-time.Hour * 24 * 7)
		end = time.Now()
		refreshDuration = time.Minute * 30
	case "month":
		start = time.Now().Add(-time.Hour * 24 * 30)
		end = time.Now()
		refreshDuration = time.Hour
	default:
		start = time.Now().Add(-time.Minute * 5)
		end = time.Now()
		refreshDuration = time.Second
	}

	ctx.EnableReadCheck()

	// Send Ping History to pre-fill graph
	pings, err := sl.PingHistory(start, end)
	if err != nil {
		return ctx.Send(OmnibusPacket{Error: err.Error()})
	}

	// Send Traffic History to pre-fill graph
	down, up, err := sl.TafficHistory(start, end)
	if err != nil {
		return ctx.Send(OmnibusPacket{Error: err.Error()})
	}

	keys := make([]int, 0)
	for i := range pings {
		keys = append(keys, int(i))
	}
	sort.Ints(keys)

	downPoints, latencyPoints, upPoints := make([]ChartDataPoint, len(pings)), make([]ChartDataPoint, len(pings)), make([]ChartDataPoint, len(pings))
	for i, time := range keys {
		downPoints[i] = ChartDataPoint{Timestamp: int64(time), Value: down[int64(time)]}
		latencyPoints[i] = ChartDataPoint{Timestamp: int64(time), Value: pings[int64(time)]}
		upPoints[i] = ChartDataPoint{Timestamp: int64(time), Value: up[int64(time)]}
		ctx.Send(OmnibusPacket{
			Type: OmnibusInitial,
			Data: OmnibusDataInit{
				Down:    downPoints,
				Latency: latencyPoints,
				Up:      upPoints,
			},
		})
	}

	// Keep sending the latest ping to update graph every second
	for !ctx.Closed() {
		<-time.After(refreshDuration)
		ping, err := sl.Ping()
		if err != nil {
			return ctx.Send(OmnibusPacket{Error: err.Error()})
		}
		down, up, err := sl.Traffic()
		if err != nil {
			return ctx.Send(OmnibusPacket{Error: err.Error()})
		}
		timestamp := time.Now().Unix()
		ctx.Send(OmnibusPacket{
			Type: OmnibusChartUpdate,
			Data: OmnibusDataChartUpdate{
				Down:    ChartDataPoint{Timestamp: timestamp, Value: down},
				Latency: ChartDataPoint{Timestamp: timestamp, Value: ping},
				Up:      ChartDataPoint{Timestamp: timestamp, Value: up},
			},
		})
	}
	return nil
}
