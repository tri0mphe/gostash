package plugin

import (
	"context"
	"github.com/tri0mphe/gostash/internal/event"
)

type Output interface {
	Start(ctx context.Context, inChan <-chan event.Event)
}
