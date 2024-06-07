package router

import (
	"fmt"
	"math/rand"
	"net/http"
	"server-ws-dummy/entities"
	"server-ws-dummy/entities/ws"
	"server-ws-dummy/repository"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/randyardiansyah25/wsbase-handler"
)

func RegisterHandler(router *gin.Engine) {
	router.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "Router Template V0.0.0")
	})

	// router.GET("/list", handler.GetEmployeListHandler)
	// router.POST("/employee", handler.GetEmployee)

	router.GET("/connect/:id", func(ctx *gin.Context) {
		fmt.Println("register")
		ws.Hub.RegisterClient(ctx.Param("id"), ctx.Writer, ctx.Request)
	})

	router.GET("/browsenasabah", func(ctx *gin.Context) {
		reqId := GenerateId(entities.INSTITUTION_ID)
		//request.Table.Add(reqId, ctx)
		fmt.Println("receive Test")
		msg := wsbase.Message{
			Type:     wsbase.TypePrivate,
			SenderId: reqId,
			To:       []string{entities.INSTITUTION_ID},
			Action:   "req_browse_nasabah",
			Title:    "Hanya Test",
			Body:     "Hello world",
		}
		ws.Hub.PushMessage(msg)
		resp, er := repository.ListenResponse(reqId)
		if er != nil {
			ctx.String(200, er.Error())
		} else {
			ctx.JSON(200, resp)
		}
	})

	router.GET("/balance/check", func(ctx *gin.Context) {
		reqId := GenerateId(entities.INSTITUTION_ID)
		//request.Table.Add(reqId, ctx)
		fmt.Println("receive balance check")
		msg := wsbase.Message{
			Type:     wsbase.TypePrivate,
			SenderId: reqId,
			To:       []string{entities.INSTITUTION_ID},
			Action:   "req_cek_saldo",
			Title:    "Hanya Test",
			Body:     "Cek saldo saya berapa",
		}
		ws.Hub.PushMessage(msg)
		resp, er := repository.ListenResponse(reqId)
		if er != nil {
			ctx.String(200, er.Error())
		} else {
			ctx.JSON(200, resp)
		}
	})
}

func GenerateId(institutionId string) string {
	t := time.Now()
	dt := t.Format("20060102150405")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	num := make([]string, 0)

	for i := 0; i < 6; i++ {
		n := r1.Intn(9)
		num = append(num, strconv.Itoa(n))
	}

	seq := strings.Join(num, "")

	return fmt.Sprint(institutionId, dt, seq)
}
