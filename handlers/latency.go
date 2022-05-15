package handlers

import (
	"strings"
	"time"

	"github.com/go-mojito/mojito"
	"github.com/infinytum/starlink-dishboard/pkg/starlink"
)

func Traffic(ctx mojito.Context, sl *starlink.Service) {
	down, up, _ := sl.GetTafficHistory()
	ctx.JSON([][]float32{down, up})
}

type LatencyPacket struct {
	Error     string  `json:"error,omitempty"`
	Latency   float64 `json:"latency,omitempty"`
	Timestamp int64   `json:"timestamp,omitempty"`
}

// func LatencyWS(ctx mojito.WebSocketContext, sl *starlink.Service) error {
// 	ctx.EnableReadCheck()
// 	// Keep sending the latest ping to update graph every second
// 	for !ctx.Closed() {
// 		<-time.After(time.Second)
// 		ping, err := sl.Ping()
// 		if err != nil {
// 			return ctx.Send(LatencyPacket{Error: err.Error()})
// 		}
// 		ctx.Send(LatencyPacket{Latency: ping})
// 	}
// 	return nil
// }

func Latency(ctx mojito.WebSocketContext, sl *starlink.Service) error {
	timeframe := ctx.Request().ParamOrDefault("timeframe", "live")
	var start, end time.Time

	switch strings.ToLower(timeframe) {
	case "today":
		start = time.Now().Add(-time.Hour * 24)
		end = time.Now()
	case "week":
		start = time.Now().Add(-time.Hour * 24 * 7)
		end = time.Now()
	case "month":
		start = time.Now().Add(-time.Hour * 24 * 30)
		end = time.Now()
	default:
		start = time.Now().Add(-time.Second * 60)
		end = time.Now()
	}

	ctx.EnableReadCheck()
	// Send Ping History to pre-fill graph
	if pings, err := sl.PingHistory(start, end); err == nil {
		for time, ping := range pings {
			ctx.Send(LatencyPacket{
				Latency:   ping,
				Timestamp: time,
			})
		}
	}

	// Keep sending the latest ping to update graph every second
	for !ctx.Closed() {
		<-time.After(time.Second)
		ping, err := sl.Ping()
		if err != nil {
			return ctx.Send(LatencyPacket{Error: err.Error()})
		}
		ctx.Send(LatencyPacket{Latency: ping, Timestamp: time.Now().Unix()})
	}
	return nil
}
