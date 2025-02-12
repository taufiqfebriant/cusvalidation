package post

type ReqBody struct {
	Name     string `json:"name" validate:"required" label:"Name"`
	Email    string `json:"email" validate:"required,email" label:"Email"`
	Password string `json:"password" validate:"required,min=6" label:"Password"`
}
