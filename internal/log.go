package internal

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func initializeLog(logFilename string) (*os.File, error) {

	var err error
	var f *os.File

	// Test if the logs must be written to the standard output or to a file
	if logFilename == "" {
		f = os.Stdout
	} else {
		// Open the log file
		f, err = os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.
				WithFields(log.Fields{"error": err, "logFilename": logFilename}).
				Debug("Can't open log file")
			return nil, err
		}
	}

	log.SetOutput(f)

	// Load the level of log from the config file
	logLevelConfig := viper.GetString("LOG_LEVEL")
	logLevel, err := log.ParseLevel(logLevelConfig)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "logFilename": logFilename, "logLevel": logLevelConfig}).
			Debug("Unknown log level config")
		return nil, err
	}

	log.SetLevel(logLevel)

	if logLevel > log.InfoLevel {
		// Add the calling method as a field when the log level is DEBUG or TRACE.
		log.SetReportCaller(true)
	}
	log.Printf("--- COMMENTS API v%s START --------------------------------------------------------", PROJECT_VERSION)

	return f, nil
}
