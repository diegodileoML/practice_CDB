package kvs

import (
	"context"

	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
)
/*
const (
	SAVE ="INSERT INTO User(first_name,last_name,dni,birth_date,email, nacionality,address) VALUES (?,?,?,?,?,?,?)"
	GET ="SELECT id,first_name,last_name,dni,birth_date,email, nacionality,address FROM User WHERE id=?"
)
*/
/*
var users []basic.User = []basic.User{
	{
		ID:          1,
		FirstName:   "Diego",
		LastName:    "Di Leo",
		Dni:         40369842,
		BirthDate:   "03/05/97",
		Email:       "diegodileo@gmail.com",
		Nacionality: "argentino",
		Address:     "dr montes de oca",
	},
	{
		ID:          2,
		FirstName:   "Sole",
		LastName:    "Lujan",
		Dni:         37521322,
		BirthDate:   "31/12/93",
		Email:       "sol93lujan@gmail.com",
		Nacionality: "argentino",
		Address:     "brigante",
	},
}

 */


/*
func (r repository) GetAll(ctx context.Context) ([]basic.User, error) {
	return users, nil
}
 */
func (r repository) GetByID(ctx context.Context, id string) (basic.User, error) {
	user := &basic.User{}
	err := r.repo.Get(ctx, id, user)
	if err!=nil{
		return basic.User{}, err
	}

	return *user,nil

	/*
	row:=r.db.QueryRow(GET,id)
	u:=basic.User{}
	err:=row.Scan(&u.ID,&u.FirstName,&u.LastName,&u.Dni,&u.BirthDate,&u.Email,&u.Nacionality,&u.Address)
	if err!=nil{
		return basic.User{},err
	}
	return u,nil
	 */
}
func (r repository) Exists(ctx context.Context, id string) bool {
	err := r.repo.Get(ctx, id,nil)
	if err!=nil {
		return false
	}
	return true
}
func (r repository) Store(ctx context.Context, u basic.User) error {
	/*
	id := len(users) + 1
	u.ID = id
	users = append(users, u)
	return u, nil

	stmt, err:= r.db.Prepare(SAVE)
	if err!=nil {
		return basic.User{},err
	}
	res ,err:=stmt.Exec(&u.FirstName,&u.LastName,&u.Dni,&u.BirthDate,&u.Email,&u.Nacionality,&u.Address)
	if err!=nil{
		return basic.User{},err
	}
	id,err:=res.LastInsertId()
	if err!=nil{
		return basic.User{},err
	}
	user,_:=r.GetByID(ctx,int(id))

	return user,nil

	 */
	return r.repo.Set(ctx,u.ID,u)

}
func (r repository) Update(ctx context.Context, u basic.User) error {

	/*
	for i, usuarios := range users {
		if usuarios.ID == u.ID {
			users[i] = u
			return nil
		}
	}
	return &web.Error{Status: 404, Code: "404", Message: "Not Updated"}

	 */
	return r.repo.Set(ctx,u.ID,u)
}
func (r repository) Delete(ctx context.Context, id string) error {
	/*for i, usuarios := range users {
		if usuarios.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return &web.Error{Status: 404, Code: "404", Message: "Not Deleted"}

	 */
	return r.repo.Delete(ctx,id)
}
