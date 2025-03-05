package output

import (
	"context"
	"fmt"
	"github.com/tri0mphe/gostash/internal/event"
	"github.com/tri0mphe/gostash/pkg/plugin"
	"github.com/tri0mphe/gostash/pkg/registry"
)

func init() {
	registry.RegisterOutput("stdout", &StdoutOuputBuilder{})
}

type StdoutOutput struct {
}

func (s *StdoutOutput) Start(ctx context.Context, inChan <-chan event.Event) {
	for {
		select {
		case <-ctx.Done():
			return
		default:

			event, ok := <-inChan
			if !ok {
				return
			}
			fmt.Printf("Timestamp:%s,Message:%s,Fields:%s\n",
				event.Timestamp,
				event.Message,
				event.Fields)
		}
	}
}

type StdoutOuputBuilder struct {
}

func (s *StdoutOuputBuilder) Build(conf map[string]interface{}) (plugin.Output, error) {
	return &StdoutOutput{}, nil
}
