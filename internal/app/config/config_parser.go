package config

import (
	"time"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

type Parser struct {
}

func (*Parser) Parse() (*Config, bool, error) {
	config := &Config{}
	config.StartTimeCallback = func(date string) {
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
		switch err.(type) {
		case *flags.Error:
			isHelpCommandCalled := err.(*flags.Error).Type == flags.ErrHelp
			if isHelpCommandCalled {
				return nil, false, nil
			}
		default:
			return nil, false, err
		}
	}

	if config.StartTime.IsZero() {
		config.StartTime = time.Now()
	}

	return config, config.Verbose, nil
}
