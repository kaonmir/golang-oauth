package forms

type UserSignup struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Password string `json:"password" binding:"required"`
}
