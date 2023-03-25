package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goombaio/namegenerator"
	"github.com/gorilla/websocket"
)

var serveName string
var count int

func main() {
	port := ":18086"
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	count = 0
	serveName = nameGenerator.Generate()

	router := gin.Default()
	router.GET("/api/welcome", welcome)
	router.GET("/api/wait", wait_welcome)

	router.GET("/websocket", handleWebSocket)

	log.Println("server start at :18086")
	log.Fatal(http.ListenAndServe(port, router))
}

func handleWebSocket(ctx *gin.Context) {
	upgrader := &websocket.Upgrader{
		//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer func() {
		log.Println("disconnect !!")
		conn.Close()
	}()

	for {
		mtype, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		current := time.Now().Format("2006-01-02 15:04:05")
		send_msg := fmt.Sprintf("@%s server @%v receive your message: %s \n", serveName, current, msg)
		send_msg_byte := []byte(send_msg)
		err = conn.WriteMessage(mtype, send_msg_byte)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func welcome(ctx *gin.Context) {
	res := fmt.Sprintf("Hello, I'am %s! call api %v times", serveName, count)
	count++
	ctx.JSON(http.StatusOK, res)
}

func wait_welcome(ctx *gin.Context) {
	start_time := time.Now().Format("2006-01-02 15:04:05")
	time.Sleep(time.Second * 30)
	end_time := time.Now().Format("2006-01-02 15:04:05")
	res := fmt.Sprintf("Hello, I'am %s! you call api @%v and get return @%v", serveName, start_time, end_time)
	count++
	ctx.JSON(http.StatusOK, res)
}
