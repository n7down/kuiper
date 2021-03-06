package servers

import (
	"container/list"

	"github.com/n7down/kuiper/internal/common/listeners"
	"github.com/sirupsen/logrus"
)

type ListenersServer struct {
	listenerList *list.List
}

func NewListenersServer() *ListenersServer {
	return &ListenersServer{
		listenerList: list.New(),
	}
}

func (i *ListenersServer) AddListener(listener *listeners.Listener) {
	i.listenerList.PushBack(listener)
}

func (i *ListenersServer) Connect() {
	for l := i.listenerList.Front(); l != nil; l = l.Next() {
		err := l.Value.(*listeners.Listener).Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}
	}
}
