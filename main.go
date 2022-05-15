package main

import (
	_ "github.com/go-mojito/defaults"
	"github.com/go-mojito/mojito"
	"github.com/go-mojito/mojito/handler"
	"github.com/go-mojito/mojito/pkg/router"
	fasthttp "github.com/go-mojito/router-fasthttp"
	"github.com/infinytum/starlink-dishboard/handlers"
)

func main() {
	fasthttp.AsDefault()
	mojito.GET("/", handler.View("dishboard"))
	mojito.WithGroup("/api", func(group router.Group) {
		group.GET("/latency", handlers.Latency)
		group.GET("/traffic", handlers.Traffic)
	})
	handler.HandleAssets()
	mojito.ListenAndServe(":8123")
}
