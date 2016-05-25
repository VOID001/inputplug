package sampinput

import (
	"errors"
	"fmt"
	"github.com/mozilla-services/heka/pipeline"
	"os"
)

type SampleInputPlug struct {
	filePath string
	config   *SampleInputConfig
}

type SampleInputConfig struct{}

func init() {
	pipeline.RegisterPlugin("SampleInput", func() interface{} { return new(SampleInputPlug) })
}

func (t *SampleInputPlug) Init(config interface{}) error {
	var ok bool
	conf := config.(pipeline.PluginConfig)
	if t.filePath, ok = conf["filepath"].(string); !ok {
		return errors.New("Now 'filepath' supported")
	}
	fmt.Println("Init Complete")
	return nil
}

func (t *SampleInputPlug) Stop() {
	return
}

func (t *SampleInputPlug) Run(ir pipeline.InputRunner, h pipeline.PluginHelper) (err error) {
	//Access Data From Outside world
	if _, err := os.Open(t.filePath); err != nil {
		return fmt.Errorf("Open File %s error: %s", t.filePath, err.Error())
	}
	file, _ := os.Open(t.filePath)
	for {
		sr := ir.NewSplitterRunner("")
		sr.SplitStream(file, nil)
	}
	return nil
}
