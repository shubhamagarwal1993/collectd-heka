package collectd

import (
	"fmt"
	. "github.com/mozilla-services/heka/collectd"
	"github.com/mozilla-services/heka/message"
	"github.com/mozilla-services/heka/pipeline"
	"github.com/mozilla-services/heka/plugins/tcp"
	"github.com/mozilla-services/heka/sandbox/lua"
	"os"
	"path/filepath"
	"time"
)

// collectd config struct
type collectdInputConfig struct {
	Queue string
}

type collectdInput struct {
	CONN_HOST string
	CONN_PORT string
	CONN_TYPE string

	hosts  map[string]bool
	output string
}

func (c *collectdInput) ConfigStruct() interface{} {
	return NewcollectdConfig(c.pConfig.Globals)
}

func (c *collectdInput) SetPipelineConfig(pConfig *pipeline.PipelineConfig) {
	c.pConfig = pConfig
}

func (c *collectdInput) Init(config interface{}) error {
	var (
		hostsConf  interface{}
		hosts      []interface{}
		host       string
		outputConf interface{}
		ok         bool
	)

	conf := config.(pipeline.PluginConfig)
	if hostsConf, ok = conf["hosts"]; ok {
		return errors.New("No 'hosts' setting specified.")
	}
	if hosts, ok = hostsConf.([]interface{}); !ok {
		return errors.New("'hosts' setting not a sequence.")
	}
	if outputConf, ok = conf["output"]; !ok {
		return errors.New("No 'output' setting specified.")
	}
	if f.output, ok = outputConf.(string); !ok {
		return errors.New("'output' setting not a string value.")
	}
	f.hosts = make(map[string]bool)
	for _, h := range hosts {
		if host, ok = h.(string); !ok {
			return errors.New("Non-string host value.")
		}
		f.hosts[host] = true
	}
	return nil
}

//	c.conf.CONN_HOST = ""
//	c.conf.CONN_PORT = "6003"
//	c.conf.CONN_TYPE = "tcp"
//	fmt.Println("init everything")
//}

func (c *collectdInput) Run(ir pipeline.InputRunner, h pipeline.PluginHelper) (err error) {

}

func (c *collectdInput) Stop() {

}

func (c *collectdInput) ReportMsg(msg *message.Message) error {

}

func init() {
	RegisterPlugin("collectdInput", func() interface{} { return new(UdpInput) })
}
