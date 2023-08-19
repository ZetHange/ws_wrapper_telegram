package ws

import (
	"context"
	"nhooyr.io/websocket"
)

type Connect struct {
	conn      *websocket.Conn
	ctxGlobal *context.Context
	tgId      int
}

var Connections []Connect

func GetConnByTg(tgId int) Connect {
	for _, conn := range Connections {
		if conn.tgId == tgId {
			return conn
		}
	}
	return Connect{}
}

func RemoveConnById(tgId int) {
	for i, conn := range Connections {
		if conn.tgId == tgId {
			Connections = append(Connections[:i], Connections[i+1:]...)
			return
		}
	}
}
