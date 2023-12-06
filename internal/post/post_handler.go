package post

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type Handler struct {
	postService Service
}

func NewPostHandler(postService Service) *Handler {
	return &Handler{postService: postService}
}

func (ph *Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := ph.postService.GetPosts(context.Background())

	if err != nil {
		fmt.Fprintln(w, err)
	}

	response, err := json.Marshal(posts)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.Write(response)
}

func (ph *Handler) Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	idNum, err := strconv.Atoi(id)

	if err != nil {
		fmt.Fprintln(w, "Illegal id params")
		return
	}

	post, err := ph.postService.GetPostById(context.Background(), idNum)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	response, err := json.Marshal(post)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.Write(response)
}
