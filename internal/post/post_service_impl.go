package post

import "context"

type postService struct {
	postRepository PostRepository
}

func NewPostService(postRepository PostRepository) PostService {
	return &postService{postRepository: postRepository}
}

func (ps *postService) GetPosts(ctx context.Context) ([]Post, error) {
	posts, err := ps.postRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
