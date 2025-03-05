package plugin

import (
	"context"
	"github.com/tri0mphe/gostash/internal/event"
)

type Input interface {
	Start(ctx context.Context, outChan chan<- event.Event)
}
