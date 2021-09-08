package basic

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Dni         int    `json:"dni"`
	BirthDate   string `json:"birth_date"`
	Email       string `json:"email"`
	Nacionality string `json:"nacionality"`
	Address     string `json:"address"`
	Productos   []Producto
}

/*
func Registrar() {
	usr:= User{Productos:[]Producto{{ID:1},{ID:2}}}
	//usr.Productos= []Producto{{ID:1},{ID:2}}
	fmt.Println(usr)
}

 */

type Producto struct{
	ID int
	User_ID int
}