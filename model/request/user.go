package request

type ReqUseRegister struct {
	Name     string
	Emali    string
	Password string
}
type REqUserLogin struct {
	Name     string
	password string
}
type ReqUserUpdate struct {
	OldPassword string
	NewPassword string
}
