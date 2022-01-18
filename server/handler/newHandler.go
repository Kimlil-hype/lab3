package hendler

import (
	gs "github.com/Kimlil-hype/lab3/server/store"
)

func NewHandler(gs *gs.UniqueStore) *Handlers {
	return &Handlers{gs}
}
