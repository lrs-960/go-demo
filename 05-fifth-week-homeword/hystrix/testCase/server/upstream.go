package server

import (
	"github.com/gin-gonic/gin"
	"go-demo/05-fifth-week-homeword/hystrix/internal"
	"io/ioutil"
	"net/http"
	"time"
)

func NewUpStreamServer(
	size,
	reqThreshold int,
	failedThreshold float64,
	duration time.Duration,
) *gin.Engine {
	app := gin.Default()
	app.GET("/api/up/v1", internal.Wrapper(
		size,
		reqThreshold,
		failedThreshold,
		duration,
	), upHandler)
	return app
}

func upHandler(c *gin.Context) {
	res, err := http.Get("http://localhost:8000/api/down/v1")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		c.String(res.StatusCode, string(data))
		return
	}
	c.String(res.StatusCode, string(data))
}
