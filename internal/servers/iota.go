package servers

import (
	"container/list"

	"github.com/n7down/iota/internal/listeners"
	"github.com/sirupsen/logrus"
)

type IotaServer struct {
	listenerList *list.List
}

func NewIotaServer() *IotaServer {
	return &IotaServer{
		listenerList: list.New(),
	}
}

func (i IotaServer) AddListener(listener *listeners.Listener) {
	i.listenerList.PushBack(listener)
}

func (i IotaServer) Connect() {
	for l := i.listenerList.Front(); l != nil; l = l.Next() {
		err := l.Value.(*listeners.Listener).Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}
	}
}
