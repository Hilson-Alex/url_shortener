package shortener

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(group *gin.RouterGroup) {
	group.GET("/:key", GetEntry, returnUrl)
	group.GET("/list", listAll)
	group.POST("/create", createShortUrl)
}

func GetEntry(ctx *gin.Context) {
	result, err := URLRepository().findByKey(ctx.Param("key"))
	if err != nil {
		var status = http.StatusInternalServerError
		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusNotFound
		}
		ctx.AbortWithError(status, err)
		return
	}
	result.GenShortUrl(getHost(ctx))
	ctx.Set("entry", result)
}

func Redirect(ctx *gin.Context) {
	var result = ctx.MustGet("entry").(*ShortURL)
	ctx.Redirect(http.StatusMovedPermanently, result.OriginalUrl)
}

func returnUrl(ctx *gin.Context) {
	var result = ctx.MustGet("entry").(*ShortURL)
	ctx.JSON(http.StatusOK, result)
}

func listAll(ctx *gin.Context) {
	var list, err = URLRepository().listUrls()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, item := range list {
		item.GenShortUrl(getHost(ctx))
	}
	ctx.JSON(http.StatusOK, gin.H{"data": list})
}

func createShortUrl(ctx *gin.Context) {
	var result = &ShortURL{ExpireDate: 10}
	if err := ctx.ShouldBindBodyWithJSON(result); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if expire := result.ExpireDate; expire < 1 || expire > 30 {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "The exipire date is outside the range! Should be between 1 and 30 days"},
		)
		return
	}
	if err := URLRepository().saveURL(result); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result.GenShortUrl(getHost(ctx))
	ctx.JSON(http.StatusOK, result)
}

func getHost(ctx *gin.Context) string {
	var scheme = "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + ctx.Request.Host + "/to/"
}
