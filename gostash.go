package main

import (
	"context"
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/tri0mphe/gostash/internal/logger"
	"github.com/tri0mphe/gostash/internal/pipeline"
	_ "github.com/tri0mphe/gostash/plugins/input"
	_ "github.com/tri0mphe/gostash/plugins/output"
)

var filePath = flag.String("f", "./config/fileinput_test.yaml", "file")

func main() {

	flag.Parse()
	//日志配置
	cfg := logger.Config{
		Level:       "debug",
		Format:      "json",
		OutputPaths: []string{"./log/app.log"},
		EnableColor: false,
	}
	logger.Init(&cfg)
	log.Info().Msg("Hello gostash")

	//加载配置文件
	pcfg, err := pipeline.ParseConfig(*filePath)
	if err != nil {
		log.Panic().Timestamp().Msg(err.Error())
	}

	//初始化管道
	pipeline, err := pipeline.NewPipeline(pcfg)
	if err != nil {
		log.Panic().Timestamp().Msg(err.Error())
		panic(err)
	}

	// 启动管道
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pipeline.Run(ctx)

}
