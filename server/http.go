package server

import (
	"explore-go/database"
	"explore-go/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func setupSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = App.Shutdown()

		sql, err := database.DB.DB()
		if err != nil {
			panic(err)
		}

		_ = sql.Close()
	}()
}

func httpListen() {
	setupSignal()
	var (
		host string = utils.GetEnv("HOST", "0.0.0.0")
		port uint64 = utils.ParseUint(utils.GetEnv("PORT", "3000"), 64)
	)
	if err := App.Listen(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Panic(err.Error())
	}
}
