package main

import (
	"embed"
	"io/fs"
	"net/http"

	_ "github.com/Hilson-Alex/url_shortener/connection"
	"github.com/Hilson-Alex/url_shortener/shortener"
	"github.com/gin-gonic/gin"
)

//go:embed public
var staticFiles embed.FS

func main() {
	var router = gin.Default()
	var publicFolder, _ = fs.Sub(staticFiles, "public")
	shortener.SetupRoutes(router.Group("/short"))
	router.GET("/to/:key", shortener.GetEntry, shortener.Redirect)
	router.Any("/app/*fsPath", func(ctx *gin.Context) {
		ctx.FileFromFS(ctx.Param("fsPath"), http.FS(publicFolder))
	})
	router.Run()
}
