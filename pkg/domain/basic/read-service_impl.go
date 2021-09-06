package basic

import "context"

func (s *service) GetAll(ctx context.Context) ([]User, error) {

	usr, err := s.Storage.GetAll(ctx)
	if err != nil {
		return []User{}, err
	}

	usrToMap := make(map[int]int)
	for _, value := range usr {
		_, exists := usrToMap[value.ID]
		if exists {
			err = s.Storage.Delete(ctx, value.ID)
			if err != nil {
				return []User{}, err
			}
			usr, err = s.GetAll(ctx)
		} else {
			usrToMap[value.ID] = value.ID
		}
	}
	return usr, err
}

func (s *service) GetByID(ctx context.Context, id int) (User, error) {
	s.GetAll(ctx)
	usuario, err := s.Storage.GetByID(ctx, id)
	if err != nil {
		return User{}, err
	}
	return usuario, nil
}

func (s *service) Exists(ctx context.Context, id int) bool {
	return s.Storage.Exists(ctx, id)
}
