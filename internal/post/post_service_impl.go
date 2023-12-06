package post

import "context"

type postService struct {
	postRepository Repository
}

func NewPostService(postRepository Repository) Service {
	return &postService{postRepository: postRepository}
}

func (ps *postService) GetPosts(ctx context.Context) ([]Post, error) {
	posts, err := ps.postRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (ps *postService) GetPostById(ctx context.Context, id int) (*Post, error) {
	post, err := ps.postRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return post, nil
}
