package kvs

import "github.com/diegodileoML/practice_CDB/pkg/domain/basic"

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Dni         int    `json:"dni"`
	BirthDate   string `json:"birth_date"`
	Email       string `json:"email"`
	Nacionality string `json:"nacionality"`
	Address     string `json:"address"`
}

func userFromDomain(usr basic.User) *User{
	u := User{}
	u.ID = usr.ID
	u.FirstName = usr.FirstName
	u.LastName = usr.LastName
	u.Dni = usr.Dni
	u.BirthDate = usr.BirthDate
	u.Email = usr.Email
	u.Nacionality = usr.Nacionality
	u.Address = usr.Address

	return &u
}

func (u *User) ToDomain() *basic.User{
	usr := &basic.User{
		ID: u.ID,
		FirstName: u.FirstName,
		LastName: u.LastName,
		Dni: u.Dni,
		BirthDate: u.BirthDate,
		Email: u.Email,
		Nacionality: u.Nacionality,
		Address: u.Address,
	}
	return usr
}