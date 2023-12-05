package post

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PostHandler struct {
	postService PostService
}

func NewPostHandler(postService PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (ph *PostHandler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := ph.postService.GetPosts(context.Background())

	if err != nil {
		fmt.Fprintln(w, err)
	}

	for _, post := range posts {
		fmt.Fprintln(w, post)
	}
}
