package sampinput

import (
	"fmt"
	"github.com/mozilla-services/heka/pipeline"
)

type SampleInputPlugin struct {
	filePath string
	config   *SampleInputPluginConfig
}

type SampleInputPluginConfig struct {
}

func (t *SampleInputPlugin) Init(config interface{}) error {
	fmt.Println("Init VOID001's Plugin Done")
	pipeline.RegisterPlugin("SampleInputPlugin", func() interface{} { return new(SampleInputPlugin) })
	return nil
}
