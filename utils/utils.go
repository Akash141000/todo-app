package utils

type Request struct {
	Status int
	Body   interface{}
}
type AuthRequest struct {
	Email    string
	Password string
}
