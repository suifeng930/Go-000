package main

import (
	"github.com/Go-000/Week02/global"
	"github.com/Go-000/Week02/internal/model"
	routers "github.com/Go-000/Week02/internal/router"
	"github.com/Go-000/Week02/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)


func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err:%v", err)

	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

}

func main() {

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Println("s.ListenAndServe()", err)

	}
	
}



func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DataBaseSetting)
	if err != nil {
		return err
	}
	log.Println("global.ServerSetting:",global.ServerSetting)
	log.Println("global.DataBaseSetting:",global.DataBaseSetting)
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil

}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return errors.Wrap(err,"db init connect fail")

	}
	return nil

}