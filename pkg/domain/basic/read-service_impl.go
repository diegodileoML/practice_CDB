package basic

import "context"


func (s *service) GetByID(ctx context.Context, id string) (*User, error) {
	usuario, err := s.Storage.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}

func (s *service) Exists(ctx context.Context, id string) bool {
	return s.Storage.Exists(ctx, id)
}
