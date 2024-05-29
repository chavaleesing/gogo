package entities

type UserRegisterReq struct {
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

type UserRegisterRes struct {
	Id       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type UserService interface {
	Register(req *UserRegisterReq) (*UserRegisterRes, error)
}

type UserRepository interface {
	Register(req *UserRegisterReq) (*UserRegisterRes, error)
}
