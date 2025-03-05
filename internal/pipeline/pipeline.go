package pipeline

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tri0mphe/gostash/internal/event"
	"github.com/tri0mphe/gostash/pkg/plugin"
	"github.com/tri0mphe/gostash/pkg/registry"
	"sync"
)

// 管道管理器
type Pipeline struct {
	inputs  []plugin.Input
	filters []plugin.Filter
	outputs []plugin.Output
}

func (p *Pipeline) Run(ctx context.Context) {
	var (
		inputWg  sync.WaitGroup
		filterWg sync.WaitGroup
		outputWg sync.WaitGroup
	)

	inputChan := make(chan event.Event, 100)
	outputChan := make(chan event.Event, 100)

	//启动输入插件
	for _, input := range p.inputs {
		inputWg.Add(1)
		go func(i plugin.Input) {
			defer inputWg.Done()
			i.Start(ctx, inputChan)
		}(input)
	}
	//输入完成后关闭inputChan
	go func() {
		inputWg.Wait()
		close(inputChan)
	}()

	//启动过滤插件
	filterWg.Add(1)
	go func() {
		defer filterWg.Done()
		for event := range inputChan {

			processed := event
			for _, filter := range p.filters {
				processed = filter.Process(processed)
			}
			select {
			case outputChan <- processed:
			case <-ctx.Done():
				return
			}

		}
	}()
	//过滤完成后关闭outputChan
	go func() {
		filterWg.Wait()
		close(outputChan)
	}()

	//启动输出插件
	for _, output := range p.outputs {
		outputWg.Add(1)
		go func(o plugin.Output) {
			defer outputWg.Done()
			o.Start(ctx, outputChan)
		}(output)
	}
	outputWg.Wait()
	log.Info().Timestamp().Msg("执行到这里俄")
}

func NewPipeline(cfg *PipelineConfig) (*Pipeline, error) {
	p := &Pipeline{}
	//pluginconfig
	for _, pc := range cfg.Inputs {
		log.Info().Msg(pc.Name)
		builder, exist := registry.InputPlugins[pc.Name]
		if !exist {
			log.Error().Timestamp().Msgf("input plugin %s not registered", pc.Name)
			return nil, fmt.Errorf("input plugin %s not registered", pc.Name)
		}
		input, err := builder.Build(pc.Config)
		if err != nil {
			return nil, err
		}
		p.inputs = append(p.inputs, input)
	}
	// 初始化过滤器
	for _, pc := range cfg.Filters {
		builder, exist := registry.FilterPlugins[pc.Name]
		if !exist {
			return nil, fmt.Errorf("filter plugin %s not registered", pc.Name)
		}

		filter, err := builder.Build(pc.Config)
		if err != nil {
			return nil, err
		}
		p.filters = append(p.filters, filter)
	}

	// 初始化输出
	for _, pc := range cfg.Outputs {
		builder, exist := registry.OutputPlugins[pc.Name]
		if !exist {
			return nil, fmt.Errorf("output plugin %s not registered", pc.Name)
		}

		output, err := builder.Build(pc.Config)
		if err != nil {
			return nil, err
		}
		p.outputs = append(p.outputs, output)
	}

	for _, pc := range cfg.Outputs {
		log.Info().Msg(pc.Name)
	}
	return p, nil
}
