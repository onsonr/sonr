package authorize

type AuthorizeRequest struct {
	Subject string
	Action  string
	Origin  string
}
