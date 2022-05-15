package main

import (
	_ "github.com/go-mojito/defaults"
	"github.com/go-mojito/mojito"
	"github.com/go-mojito/mojito/handler"
	"github.com/go-mojito/mojito/pkg/router"
	"github.com/infinytum/starlink-dishboard/handlers"
)

func main() {
	mojito.GET("/", handler.View("dishboard"))
	mojito.WithGroup("/ws", func(group router.Group) {
		group.GET("/latency/:timeframe", handlers.Latency)
		group.GET("/traffic/:timeframe", handlers.Traffic)
	})
	handler.HandleAssets()
	mojito.ListenAndServe(":8123")
}
