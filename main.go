package main

import (
	_ "github.com/go-mojito/defaults"
	"github.com/go-mojito/mojito"
	"github.com/go-mojito/mojito/handler"
)

func main() {
	mojito.GET("/", handler.View("dishboard"))
	handler.HandleAssets()
	mojito.ListenAndServe(":8123")
}
