package handlers

import (
	"github.com/go-mojito/mojito"
	"github.com/infinytum/starlink-dishboard/pkg/starlink"
)

func Latency(ctx mojito.Context, sl *starlink.Service) {
	ping, _ := sl.GetPingHistory()
	ctx.JSON(ping)
}

func Traffic(ctx mojito.Context, sl *starlink.Service) {
	down, up, _ := sl.GetTafficHistory()
	ctx.JSON([][]float32{down, up})
}
