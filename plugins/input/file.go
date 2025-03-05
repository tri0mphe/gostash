package input

import (
	"bufio"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tri0mphe/gostash/internal/event"
	"github.com/tri0mphe/gostash/pkg/plugin"
	"github.com/tri0mphe/gostash/pkg/registry"
	"os"
	"time"
)

func init() {
	registry.RegisterInput("file", &FileInputBuilder{})
}

type FileInput struct {
	path string //file path
}

func (f *FileInput) Start(ctx context.Context, outChan chan<- event.Event) {
	file, err := os.Open(f.path)
	if err != nil {
		log.Panic().Msg("FileInput openfile error:" + err.Error())
		panic(err)
	}
	defer file.Close()
	fields := make(map[string]interface{})
	fields["timestamp"] = time.Now()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			outChan <- event.Event{
				Timestamp: time.Now(),
				Message:   scanner.Text(),
				Fields:    fields,
			}
		}
	}
	return
}

// 实现input 插件的构造器
type FileInputBuilder struct{}

func (f *FileInputBuilder) Build(conf map[string]interface{}) (plugin.Input, error) {
	path, ok := conf["path"].(string)
	if !ok {
		return nil, fmt.Errorf("missing path in file input config")
	}
	return &FileInput{path: path}, nil
}
