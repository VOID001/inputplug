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
	stopMsg  chan interface{}
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
	t.stopMsg = make(chan interface{})
	return nil
}

func (t *SampleInputPlug) Stop() {
	fmt.Println("Stop Plugin")
	t.stopMsg <- struct{}{}
	return
}

func (t *SampleInputPlug) Run(ir pipeline.InputRunner, h pipeline.PluginHelper) (err error) {
	//Access Data From Outside world
	if _, err := os.Open(t.filePath); err != nil {
		return fmt.Errorf("Open File %s error: %s", t.filePath, err.Error())
	}
	file, _ := os.Open(t.filePath)
	sr := ir.NewSplitterRunner("")
	//Send it to heka pipeline
	for {
		sr.SplitStream(file, nil)
		fmt.Printf("%s\n", "run")
		select {
		case <-t.stopMsg:
			fmt.Println("Runner get the Stop Message")
			return nil
		default:
		}
	}
	return nil
}
