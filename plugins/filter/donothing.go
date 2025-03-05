package filter

import (
	"github.com/tri0mphe/gostash/internal/event"
	"github.com/tri0mphe/gostash/pkg/plugin"
	"github.com/tri0mphe/gostash/pkg/registry"
)

func init() {
	registry.RegisterFilter("donothing", &DoNothingFilterBuilder{})
}

type DoNothingFilter struct {
}

func (d *DoNothingFilter) Process(event event.Event) event.Event {
	return event
}

type DoNothingFilterBuilder struct {
}

func (d *DoNothingFilterBuilder) Build(conf map[string]interface{}) (plugin.Filter, error) {
	return &DoNothingFilter{}, nil
}
