package env

import (
	"time"

	"github.com/caarlos0/env"
)

// Config is a struct containing config values.
type Config struct {
	// Env the current environment. Defaults to development
	Env Environment `env:"ENV" envDefault:"DEVELOPMENT"`

	// TfLURLs is a list of the URL's to fetch from TFL.
	TfLURLs []string `env:"TFL_URLS" envSeparator:","`
	// RefreshTime is the frequency of fetching.
	RefreshTime time.Duration `env:"TFL_REFRESH_TIME" envDefault:"10s"`

	ChromeCastDeviceName   string `env:"CC_DEVICE_NAME"`
	ChromeCastDeviceUUID   string `env:"CC_DEVICE_UUID"`
	ChromeCastDevice       string `env:"CC_DEVICE"`
	ChromeCastDisableCache bool   `env:"CC_DISABLE_CACHE"`
}

var conf *Config

// Get will return the Config and parse it if not already parsed.
func Get() Config {
	if conf == nil {
		conf = new(Config)
		if err := env.Parse(conf); err != nil {
			panic(err)
		}
	}
	return *conf
}
