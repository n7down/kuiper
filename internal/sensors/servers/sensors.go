package servers

import (
	"container/list"

	"github.com/n7down/kuiper/internal/sensors/listeners"
	"github.com/sirupsen/logrus"
)

type SensorsServer struct {
	listenerList *list.List
}

func NewSensorsServer() *SensorsServer {
	return &SensorsServer{
		listenerList: list.New(),
	}
}

func (i SensorsServer) AddListener(listener *listeners.Listener) {
	i.listenerList.PushBack(listener)
}

func (i SensorsServer) Connect() {
	for l := i.listenerList.Front(); l != nil; l = l.Next() {
		err := l.Value.(*listeners.Listener).Connect()
		if err != nil {
			logrus.Fatal(err.Error())
		}
	}
}
