package sampinput

import (
	"fmt"
	"github.com/mozilla-services/heka/pipeline"
)

type SampleInputPlug struct {
	filePath string
	config   *SampleInputConfig
}

type SampleInputConfig struct {
}

func init() {
	fmt.Println("Init VOID001's Plugin Done")
	pipeline.RegisterPlugin("SampleInput", func() interface{} { return new(SampleInputPlug) })
}

func (t *SampleInputPlug) Init(config interface{}) error {
	return nil
}
