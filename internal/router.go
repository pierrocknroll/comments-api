package internal

import (
	"comments-api/internal/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func createRoutes(handler *handlers.HTTPHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET(
		"/comments/target/:target",
		accessMiddleware(
			func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
				handler.HandleIndexComments(w, r, ps)
			}))

	router.POST(
		"/comments/target/:target",
		accessMiddleware(
			func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
				handler.HandleStoreComment(w, r, ps)
			}))

	router.OPTIONS(
		"/comments/target/:target",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			handler.HandleOptionsTargetComments(w, r, ps)
		})

	router.PUT(
		"/comments/comment/:comment",
		accessMiddleware(
			func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
				handler.HandleUpdateComment(w, r, ps)
			}))

	router.DELETE(
		"/comments/comment/:comment",
		accessMiddleware(
			func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
				handler.HandleDeleteComment(w, r, ps)
			}))

	router.OPTIONS(
		"/comments/target/:comment",
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			handler.HandleOptionsComments(w, r, ps)
		})

	return router
}
