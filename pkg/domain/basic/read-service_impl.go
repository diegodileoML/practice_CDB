package basic

import "context"

func (s *service) GetAll(ctx context.Context) ([]User, error) {
	return s.Storage.GetAll(ctx)
}

func (s *service) GetByID(ctx context.Context, id int) (User, error) {
	return s.Storage.GetByID(ctx, id)
}

func (s *service) Exists(ctx context.Context, id int) bool {
	return s.Storage.Exists(ctx, id)
}
