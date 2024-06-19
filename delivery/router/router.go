package router

import (
	"io"
	"os"
	"server-ws-dummy/entities/ws"

	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
	"github.com/randyardiansyah25/wsbase-handler"
)

func Start() error {

	gin.SetMode(gin.ReleaseMode)

	//Discard semua output yang dicatat oleh gin karena print out akan dicetak sesuai kebutuhan programmer
	gin.DefaultWriter = io.Discard

	router := gin.Default() //create router engine by default
	router.Use(gin.Recovery())

	RegisterHandler(router)
	listenerPort := os.Getenv("app.listener_port")
	_ = glg.Logf("[HTTP] Listening at : %s", listenerPort)

	ws.Hub = wsbase.NewHub()
	ws.Hub.SetLogHandler(func(logType int, val string) {
		switch logType {
		case wsbase.LOG:
			_ = glg.Log(val)
		case wsbase.ERR:
			_ = glg.Error(val)
		default:
			_ = glg.Info(val)
		}
	})
	ws.Hub.SetOnReadMessageFunc(WSMessageListener)
	go ws.Hub.Run()

	return router.Run(":" + listenerPort)
}
