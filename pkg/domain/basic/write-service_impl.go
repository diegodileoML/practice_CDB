package basic

import "context"

func (s *service) Store(ctx context.Context, u User) (User, error) {
	return s.Storage.Store(ctx, u)
}

func (s *service) Update(ctx context.Context, u User) error {
	return s.Storage.Update(ctx, u)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.Storage.Delete(ctx, id)
}
