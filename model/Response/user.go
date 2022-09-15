package Response

type ResUserLogin struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
