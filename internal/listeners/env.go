package listeners

import (
	"github.com/n7down/iota/internal/stores"
)

type Env struct {
	store *stores.InfluxStore
}

func NewEnv(store *stores.InfluxStore) *Env {
	return &Env{
		store: store,
	}
}
