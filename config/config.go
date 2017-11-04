// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"time"
)

type Config struct {
	Address   string   `config:"address"`
	Namespace []string `config:"namespace"`

	TimestampField  string `config:"timestamp_field"`
	TimestampFormat string `config:"timestamp_format"`
}

var DefaultConfig = Config{
	Address:   "127.0.0.1:15170",
	Namespace: []string{},

	TimestampField:  "@timestamp",
	TimestampFormat: time.RFC3339,
}
