package handlers

import (
	"comments-api/internal/core/domains"
	"comments-api/internal/core/ports"
	"comments-api/pkg/api"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HTTPHandler struct {
	service ports.CommentsService
}

func NewHTTPHandler(service ports.CommentsService) *HTTPHandler {
	return &HTTPHandler{
		service: service,
	}
}

func (handler *HTTPHandler) HandleIndexComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	target, _ := HandleTargetParameter(ps.ByName("target"))

	comments, err := handler.service.GetCommentsByTarget(domains.Target(target))
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "target": target}).
			Error("cannot get target comments")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	convertedComments := api.Comments{Comments: make([]api.Comment, len(comments))}
	for i, comment := range comments {
		convertedComments.Comments[i] = api.Comment(comment)
	}

	e := json.NewEncoder(w)
	err = e.Encode(convertedComments)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err}).
			Error("cannot encode comments")
	}
}

func (handler *HTTPHandler) HandleStoreComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	target, _ := HandleTargetParameter(ps.ByName("target"))
	userId, _ := HandleUserIdParameter(r.FormValue("user_id"))
	message, _ := HandleMessageParameter(r.FormValue("message"))

	comment, err := handler.service.StoreComment(domains.Target(target), domains.Author(userId), message)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "target": target, "userId": userId, "message": message}).
			Error("cannot store comment")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	err = e.Encode(api.Comment(*comment))
	if err != nil {
		log.
			WithFields(log.Fields{"error": err}).
			Error("cannot encode comment")
	}
}

func (handler *HTTPHandler) HandleUpdateComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	commentId, _ := HandleCommentIdParameter(ps.ByName("comment"))
	userId, _ := HandleUserIdParameter(r.FormValue("user_id"))
	message, _ := HandleMessageParameter(r.FormValue("message"))

	comment, err := handler.service.UpdateCommentMessage(commentId, domains.Author(userId), message)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "commentId": commentId, "userId": userId, "message": message}).
			Error("cannot update comment")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(w)
	err = e.Encode(api.Comment(*comment))
	if err != nil {
		log.
			WithFields(log.Fields{"error": err}).
			Error("cannot encode comment")
	}
}

func (handler *HTTPHandler) HandleDeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	commentId, _ := HandleCommentIdParameter(ps.ByName("comment"))
	userId, _ := HandleUserIdParameter(r.FormValue("user_id"))

	err := handler.service.DeleteComment(commentId, domains.Author(userId))
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "commentId": commentId, "userId": userId}).
			Error("cannot delete comment")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (handler *HTTPHandler) HandleOptionsTargetComments(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Add("Allow", "GET,POST")
}
func (handler *HTTPHandler) HandleOptionsComments(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Add("Allow", "PUT,DEL")
}
