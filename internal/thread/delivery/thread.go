package delivery

import (
	"encoding/json"
	"forums/internal/models"
	threadModel "forums/internal/thread"
	"forums/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	threadUcase threadModel.ThreadUsecase
}

func (h Handler) AddPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slugOrId := vars["slugOrId"]
	posts := new(models.Posts)
	err := json.NewDecoder(r.Body).Decode(&posts.Posts)
	if err != nil {
		sendErr := utils.NewError(http.StatusBadRequest, err.Error())
		w.WriteHeader(sendErr.Code())
		return
	}
	defer r.Body.Close()
	if len(posts.Posts) == 0 {
		utils.NewResponse(http.StatusCreated, posts.Posts).SendSuccess(w)
		return
	}
	responsePosts, err := h.threadUcase.AddPosts(*posts, slugOrId)

	if err != nil {
		switch err.Error() {
		case "404":
			w.WriteHeader(http.StatusNotFound)
			return
		case "409":
			utils.NewResponse(http.StatusConflict, responsePosts.Posts).SendSuccess(w)
			return
		}
	}

	utils.NewResponse(http.StatusCreated, responsePosts.Posts).SendSuccess(w)
}

func NewThreadHandler(threadUcase threadModel.ThreadUsecase) threadModel.ThreadHandler {
	handler := &Handler{
		threadUcase: threadUcase,
	}

	return handler
}
