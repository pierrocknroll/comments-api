package internal

import (
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

// accessMiddleware checks if the request is authorized through an access key transmitted in the HTTP request header.
func accessMiddleware(next httprouter.Handle) httprouter.Handle {

	accessKey := viper.GetString("ACCESS_KEY")

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		currentAccessKey := r.Header.Get("AccessKey")

		if currentAccessKey == accessKey {
			// The request is authorized.
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next(w, r, ps)
		} else {
			log.Warning("unauthorized access")

			// Write a 401 error and stop the handler chain
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
