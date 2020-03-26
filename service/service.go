package service

import (
	"github.com/micro/go-micro/v2"
	logs "github.com/satoukick/webserver/log"
	"github.com/satoukick/webserver/proto"
)

func NewService() {
	ser := micro.NewService(
		micro.Name("todo"),
	)
	ser.Init()

	proto.RegisterDBToDoQueryHandler(ser.Server(), NewToDoRequestHandler())

	if err := ser.Run(); err != nil {
		logs.Error(err)
	}
}
