package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ardanlabs/conf"
)

// Configuration object parsed from env.
type Config struct {
	Production bool `conf:"default:false"`
	DB         struct {
		URL string `conf:"required"`
	}
	Google struct {
		SheetID string `conf:"required"`
		Sheet   string `conf:"required"`
	}
	IcalURL string `conf:"required"`
}

// Get a new configuration object parsed from the environment.
func NewConfig() (Config, error) {
	var cfg Config

	if err := conf.Parse(os.Args[1:], "", &cfg); err != nil {
		if !errors.Is(err, conf.ErrHelpWanted) {
			return cfg, fmt.Errorf("error parsing config from env: %s \n -h to print usage", err.Error())
		}

		if err = printUsage(cfg); err != nil {
			return cfg, err
		}
	}

	return cfg, nil
}

// Print the configuration usage to stdout.
func printUsage(cfg Config) error {
	var usage string
	usage, err := conf.Usage("", &cfg)
	if err != nil {
		return fmt.Errorf("error generating config usage: %s", err.Error())
	}
	log.Println(usage)
	os.Exit(0)

	return nil
}
