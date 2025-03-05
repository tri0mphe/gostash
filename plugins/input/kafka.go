package input

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/tri0mphe/gostash/internal/event"
	"github.com/tri0mphe/gostash/pkg/plugin"
	"github.com/tri0mphe/gostash/pkg/registry"
	"sync"
	"time"
)

/*
**
 */

func init() {
	registry.RegisterInput("kafka", &KafkaInputBuilder{})
}

// 原始的可以读取topics，这里修改只能读取一个topic
type KafkaInput struct {
	kafkaConfig *kafka.ConfigMap
	topic       string
	concurrency int
}

func (k *KafkaInput) Start(ctx context.Context, outChan chan<- event.Event) {
	var wg sync.WaitGroup

	for i := 0; i < k.concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			consumer, err := kafka.NewConsumer(k.kafkaConfig)
			if err != nil {
				log.Fatal().Msgf("Error creating consumer: %v\n", err)
				return
			}
			defer consumer.Close()
			topics := []string{k.topic}
			consumer.SubscribeTopics(topics, nil)
			for {
				m, err := consumer.ReadMessage(-1)
				if err != nil {
					log.Error().Msgf("kafka read message error: %s (%v)\n", err.Error(), m)
				}

				select {
				case <-ctx.Done():
					return
				default:
					outChan <- event.Event{
						Timestamp: time.Now(),
						Message:   string(m.Value),
						Fields:    make(map[string]interface{}),
						Tags:      []string{"kafka"},
					}
				}

			}
		}()
	}
	wg.Wait()

}

// 实现input 插件的构造器
type KafkaInputBuilder struct{}

func (k *KafkaInputBuilder) Build(conf map[string]interface{}) (plugin.Input, error) {
	var kafkaConfig kafka.ConfigMap

	if v, ok := conf["kafka_config"]; !ok {
		log.Fatal().Msg("kafka input must have kafka_config")
	} else {
		if err := mapstructure.Decode(v, &kafkaConfig); err != nil {
			return nil, err
		}
		// official json marshal: unsupported type: map[interface {}]interface {}
	}
	for k, v := range kafkaConfig {
		log.Info().Msgf("%v,%v", k, v)
	}
	topic, ok := conf["topic"].(string)
	if !ok {
		log.Fatal().Msg("kafka input config must have key:toipc")
	}
	concurrency, ok := conf["concurrency"].(int)
	if !ok {
		concurrency = 10
	}

	var kafkaInput = KafkaInput{
		kafkaConfig: &kafkaConfig,
		topic:       topic,
		concurrency: concurrency,
	}
	return &kafkaInput, nil
}
