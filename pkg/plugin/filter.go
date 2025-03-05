package plugin

import "github.com/tri0mphe/gostash/internal/event"

type Filter interface {
	Process(event event.Event) event.Event
}
