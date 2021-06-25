package config

import (
	"time"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

type Parser struct {
}

func (*Parser) Parse() (*Config, error) {
	config := &Config{}
	config.startTimeCallback = func(date string) {
		var err error
		config.StartTime, err = time.Parse("2006-01-02", date)
		if err != nil {
			if _, ok := err.(*time.ParseError); ok {
				log.Errorf("Start time parsing error. Must be YYYY-MM-DD, given \"%s\"", date)
				return
			}

			log.WithError(err).Errorf("Unexpected date parsing error")
			return
		}
	}

	_, err := flags.Parse(config)
	if err != nil {
		return nil, err
	}

	if config.StartTime.IsZero() {
		config.StartTime = time.Now()
	}

	return config, nil
}
