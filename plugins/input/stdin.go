package input

import (
	"bufio"
	"context"
	"fmt"
	"github.com/tri0mphe/gostash/internal/event"
	"github.com/tri0mphe/gostash/pkg/plugin"
	"github.com/tri0mphe/gostash/pkg/registry"
	"os"
	"time"
)

func init() {
	registry.RegisterInput("stdin", &StdinInputBuilder{})
}

type StdinInput struct {
}

func (s *StdinInput) Start(ctx context.Context, outChan chan<- event.Event) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("程序已启动，请输入:")
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			outChan <- event.Event{
				Timestamp: time.Now(),
				Message:   scanner.Text(),
				Tags:      []string{"stdin"},
			}
		}
	}
}

// 实现input 插件的构造器
type StdinInputBuilder struct{}

func (s *StdinInputBuilder) Build(conf map[string]interface{}) (plugin.Input, error) {
	return &StdinInput{}, nil
}
