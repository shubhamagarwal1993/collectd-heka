// This is our plugin struct.
type collectdOutput struct {
    conn net.Conn
}

// This is our plugin's config struct
type collectdOutputConfig struct {
    Address string
}

// Provides pipeline.HasConfigStruct interface.
func (c *collectdOutputConfig) ConfigStruct() interface{} {
    return &collectdOutputConfig{"my.example.com:44444"}
}

// Initialize UDP connection
func (c *collectdOutput) Init(config interface{}) (err error) {
    conf := config.(*collectdOutputConfig) // assert we have the right config type
    var udpAddr *net.UDPAddr
    if udpAddr, err = net.ResolveUDPAddr("udp", conf.Address); err != nil {
        return fmt.Errorf("can't resolve %s: %s", conf.Address,
            err.Error())
    }
    if c.conn, err = net.DialUDP("udp", nil, udpAddr); err != nil {
        return fmt.Errorf("error dialing %s: %s", conf.Address,
            err.Error())
    }
    return
}
