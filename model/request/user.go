package request

type ReqUseRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type REqUserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type ReqUserUpdate struct {
	OldPassword string
	NewPassword string
}
