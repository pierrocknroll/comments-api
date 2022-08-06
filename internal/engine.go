package internal

import (
	"comments-api/internal/core/services"
	"comments-api/internal/handlers"
	"comments-api/internal/repositories"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

// Run is the main entry for this application.
func Run() {

	// Initialize config file.
	configFilename := ".env"
	err := initializeConfig(configFilename)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "configFilename": configFilename}).
			Error("Cannot read config file")
		fmt.Println("Error: cannot read config file")
		return
	}

	// Initialize log file.
	logFilename := viper.GetString("LOG_FILENAME")
	logFile, err := initializeLog(logFilename)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "logFilename": logFilename}).
			Error("Cannot initialize log file")
		fmt.Println("Error: cannot initialize log file")
		return
	}
	defer func() {
		err := logFile.Close()
		if err != nil {
			log.
				WithFields(log.Fields{"error": err, "logFilename": logFilename}).
				Error("Cannot close log file")
			return
		}
	}()

	repository, err := repositories.NewCommentRepository(viper.GetString("COMMENTS_DATABASE_ADDRESSS"))
	if err != nil {
		log.
			WithFields(log.Fields{"error": err}).
			Error("Can't initialize comments repository")
		fmt.Println("Error: can't initialize comments repository")
		return
	}
	defer func() {
		repository.Close()
	}()

	commentsService := services.NewCommentService(repository)
	handlerManager := handlers.NewHTTPHandler(commentsService)

	router := createRoutes(handlerManager)

	log.Fatal(http.ListenAndServe(viper.GetString("PORT_NUMBER"), router))
}
