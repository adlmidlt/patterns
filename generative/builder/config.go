package builder

import (
	"errors"
	"fmt"
	"log"
)

type ServerConfig struct {
	Host           string
	Port           int
	Protocol       string
	Timeout        int
	MaxConnections int
}

type ServerConfigBuilder struct {
	config *ServerConfig
}

func NewServerConfigBuilder() *ServerConfigBuilder {
	return &ServerConfigBuilder{config: &ServerConfig{}}
}

func (sc *ServerConfigBuilder) SetHost(host string) *ServerConfigBuilder {
	sc.config.Host = host

	return sc
}

func (sc *ServerConfigBuilder) SetPort(port int) *ServerConfigBuilder {
	sc.config.Port = port

	return sc
}

func (sc *ServerConfigBuilder) SetProtocol(protocol string) *ServerConfigBuilder {
	sc.config.Protocol = protocol

	return sc
}

func (sc *ServerConfigBuilder) SetTimeout(timeout int) *ServerConfigBuilder {
	sc.config.Timeout = timeout

	return sc
}

func (sc *ServerConfigBuilder) SetMaxConnections(max int) *ServerConfigBuilder {
	sc.config.MaxConnections = max

	return sc
}

func (sc *ServerConfigBuilder) Build() (*ServerConfig, error) {
	if sc.config.Host == "" {
		return nil, errors.New("host is required")
	}

	if sc.config.Port == 0 {
		return nil, errors.New("port is required")
	}
	return sc.config, nil
}

func StartBuilder1Pattern() {
	config, err := NewServerConfigBuilder().
		SetHost("localhost").
		SetPort(8080).
		SetProtocol("http").
		SetTimeout(30).Build()
	if err != nil {
		log.Fatal(err)
	}

	// без таймаута
	config1, err := NewServerConfigBuilder().
		SetHost("localhost").
		SetPort(8080).
		SetProtocol("http").
		Build()

	fmt.Printf("%+v\n", config)
	fmt.Printf("%+v\n", config1)
}
