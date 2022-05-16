package main

import (
	"errors"
	"net/http"
	"os"

	_ "github.com/go-mojito/defaults"
	"github.com/go-mojito/mojito"
	"github.com/go-mojito/mojito/handler"
	"github.com/go-mojito/mojito/pkg/router"
	"github.com/infinytum/starlink-dishboard/handlers"
)

func main() {
	mojito.GET("/*path", Nuxt)
	mojito.WithGroup("/ws", func(group router.Group) {
		group.GET("/omnibus/:timeframe", handlers.Omnibus)
	})
	handler.HandleAssets()
	mojito.ListenAndServe(":8123")
}

func Nuxt(ctx mojito.Context) error {
	fsPath := mojito.ResourcesDir() + "../web/app/.output/public"
	fsInfo, err := os.Stat(fsPath)
	if err != nil {
		mojito.DefaultLogger().Errorf("Error while accessing assets path %s: %s", fsPath, err.Error())
		return err
	}

	if !fsInfo.IsDir() {
		return errors.New("assets path is not a directory, cannot serve assets from path %s")
	}

	fileHandlerRequest := ctx.Request().GetRequest().Clone(ctx.Request().GetRequest().Context())
	fileHandlerRequest.URL.Path = "/" + ctx.Request().Param("path")
	http.FileServer(http.Dir(fsPath)).ServeHTTP(ctx.Response(), fileHandlerRequest)
	return nil
}
