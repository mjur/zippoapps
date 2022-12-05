package configuration

import "github.com/mjur/zippo/pkg/configuration/log"

const (
	ServiceName string = "configuration"
	Version     string = "v1.0.0"
)

// Config is used for injection into various services.
type Config struct {
	Log     *log.Logger
	Host    string
	Port    string
	TTL     int
	Timeout int
}
