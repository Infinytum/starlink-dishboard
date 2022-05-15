package handlers

import (
	"sort"
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
		start = time.Now().Add(-time.Second * 120)
		end = time.Now()
		refreshDuration = time.Second
	}

	ctx.EnableReadCheck()
	// Send Traffic History to pre-fill graph
	if down, up, err := sl.TafficHistory(start, end); err == nil {
		keys := make([]int, 0)
		for i := range down {
			keys = append(keys, int(i))
		}
		sort.Ints(keys)

		for _, time := range keys {
			ctx.Send(TrafficPacket{
				Down:      down[int64(time)],
				Up:        up[int64(time)],
				Timestamp: int64(time),
			})
		}
	}

	// Keep sending the latest ping to update graph every second
	for !ctx.Closed() {
		<-time.After(refreshDuration)
		down, up, err := sl.Traffic()
		if err != nil {
			return ctx.Send(TrafficPacket{Error: err.Error()})
		}
		ctx.Send(TrafficPacket{Down: down, Up: up, Timestamp: time.Now().Unix()})
	}
	return nil
}
