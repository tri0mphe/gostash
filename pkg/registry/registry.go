package registry

import "github.com/tri0mphe/gostash/pkg/plugin"

/***
*  The global plugins is initialized by the init() function of each plugin-in when the comand start
 ***/
var (
	InputPlugins  = make(map[string]InputBuilder)
	FilterPlugins = make(map[string]FilterBuilder)
	OutputPlugins = make(map[string]OutputBuilder)
)

// register func
func RegisterInput(name string, builder InputBuilder) {
	InputPlugins[name] = builder
}

func RegisterFilter(name string, builder FilterBuilder) {
	FilterPlugins[name] = builder
}

func RegisterOutput(name string, builder OutputBuilder) {
	OutputPlugins[name] = builder
}

// Builder interface
type InputBuilder interface {
	Build(conf map[string]interface{}) (plugin.Input, error)
}
type FilterBuilder interface {
	Build(conf map[string]interface{}) (plugin.Filter, error)
}
type OutputBuilder interface {
	Build(conf map[string]interface{}) (plugin.Output, error)
}
