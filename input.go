package sampinput

import (
	"fmt"
	"github.com/mozilla-services/heka/pipeline"
)

type SampleInput struct {
	filePath string
	config   *SampleInputConfig
}

type SampleInputConfig struct {
}

func init() {
	fmt.Println("Init VOID001's Plugin Done")
	pipeline.RegisterPlugin("SampleInput", func() interface{} { return new(SampleInput) })
}

func (t *SampleInput) Init(config interface{}) error {
	fmt.Println("Init VOID001's Plugin Done")
	pipeline.RegisterPlugin("SampleInput", func() interface{} { return new(SampleInput) })
	return nil
}
