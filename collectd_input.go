package collectd

import (
	"fmt"
	//	. "github.com/mozilla-services/heka/collectd"
	//"github.com/mozilla-services/heka/message"
	"github.com/mozilla-services/heka/pipeline"
	//"github.com/mozilla-services/heka/plugins/tcp"
	//"github.com/mozilla-services/heka/sandbox/lua"
	//"os"
	//"path/filepath"
	//"time"
	//"net"
	//"errors"
)

type CollectdInput struct {
	conf    *CollectdInputConfig
	running bool
}

// collectd config struct
type CollectdInputConfig struct {
	CONN_HOST string `toml:"CONN_HOST"`
	CONN_PORT string `toml:"CONN_PORT"`
	CONN_TYPE string `toml:"CONN_TYPE"`
}

func (c *CollectdInput) ConfigStruct() interface{} {
	return &CollectdInputConfig{}
}

func (c *CollectdInput) Init(config interface{}) (err error) {
	c.running = true
	c.conf = config.(*CollectdInputConfig)

	if c.conf.CONN_HOST == "" {
		return fmt.Errorf("CONN_HOST must be set")
	}

	if c.conf.CONN_PORT == "" {
		return fmt.Errorf("CONN_PORT must be set")
	}

	if c.conf.CONN_TYPE == "" {
		return fmt.Errorf("CONN_TYPE must be set")
	}

	return
}

func (c *CollectdInput) Run() (err error) {
	fmt.Println("inside Run function")

	return nil
}

func (c *CollectdInput) Stop() {
	c.running = false
}

func init() {
	pipeline.RegisterPlugin("CollectdInput", func() interface{} {
		return new(CollectdInput)
	})
}

