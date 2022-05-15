package handlers

import (
	"strings"
	"time"

	"github.com/go-mojito/mojito"
	"github.com/infinytum/starlink-dishboard/pkg/starlink"
)

type TrafficPacket struct {
	Error     string  `json:"error,omitempty"`
	Down      float64 `json:"down,omitempty"`
	Up        float64 `json:"up,omitempty"`
	Timestamp int64   `json:"timestamp,omitempty"`
}

func Traffic(ctx mojito.WebSocketContext, sl *starlink.Service) error {
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
	if down, up, err := sl.TafficHistory(start, end); err == nil {
		for timestamp, down := range down {
			up := up[timestamp]
			ctx.Send(TrafficPacket{
				Down:      down,
				Up:        up,
				Timestamp: timestamp,
			})
		}
	}

	// Keep sending the latest ping to update graph every second
	for !ctx.Closed() {
		<-time.After(time.Second)
		down, up, err := sl.Traffic()
		if err != nil {
			return ctx.Send(TrafficPacket{Error: err.Error()})
		}
		ctx.Send(TrafficPacket{Down: down, Up: up, Timestamp: time.Now().Unix()})
	}
	return nil
}
