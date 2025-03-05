package pipeline

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"os"
)

/***
Example:

inputs:
  - name: File
    config:
      path: "/var/log/app.log"

filters:
  - name: Add
    config:
      key: "environment"
      value: "production"

outputs:
  - name: Stdout
    config:
      pretty_print: true
*/

// YAML配置结构
type PipelineConfig struct {
	Inputs  []PluginConfig `yaml:"inputs"`
	Filters []PluginConfig `yaml:"filters"`
	Outputs []PluginConfig `yaml:"outputs"`
}

// YAML 配置解析
type PluginConfig struct {
	Name   string                 `yaml:"name"`
	Config map[string]interface{} `yaml:"config"`
}

// 解析Yaml配置文件
func ParseConfig(path string) (*PipelineConfig, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		log.Error().Timestamp().Msg("read file error when loadconfig: " + err.Error())
		return nil, err
	}
	var cfg PipelineConfig

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Error().Msg("yaml unmarshal error :" + err.Error())
		return nil, err
	}
	return &cfg, nil
}
