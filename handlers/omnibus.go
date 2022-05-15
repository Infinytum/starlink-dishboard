package handlers

import (
	"sort"
	"strings"
	"time"

	"github.com/go-mojito/mojito"
	"github.com/infinytum/starlink-dishboard/pkg/starlink"
)

type OmnibusPacket struct {
	Error     string        `json:"error,omitempty"`
	Devices   []interface{} `json:"devices,omitempty"`
	Down      float64       `json:"down,omitempty"`
	Latency   float64       `json:"latency,omitempty"`
	Outages   []interface{} `json:"outages,omitempty"`
	Up        float64       `json:"up,omitempty"`
	Timestamp int64         `json:"timestamp,omitempty"`
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
	if pings, err := sl.PingHistory(start, end); err == nil {
		keys := make([]int, 0)
		for i := range pings {
			keys = append(keys, int(i))
		}
		sort.Ints(keys)

		for _, time := range keys {
			ctx.Send(OmnibusPacket{
				Latency:   pings[int64(time)],
				Timestamp: int64(time),
			})
		}
	}

	// Send Traffic History to pre-fill graph
	if down, up, err := sl.TafficHistory(start, end); err == nil {
		keys := make([]int, 0)
		for i := range down {
			keys = append(keys, int(i))
		}
		sort.Ints(keys)

		for _, time := range keys {
			ctx.Send(OmnibusPacket{
				Down:      down[int64(time)],
				Up:        up[int64(time)],
				Timestamp: int64(time),
			})
		}
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
		ctx.Send(OmnibusPacket{
			Latency:   ping,
			Down:      down,
			Up:        up,
			Timestamp: time.Now().Unix(),
		})
	}
	return nil
}
