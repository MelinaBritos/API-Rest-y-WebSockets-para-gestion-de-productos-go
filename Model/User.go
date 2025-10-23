package Model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	RoleId   int    `json:"role_id"`
}
