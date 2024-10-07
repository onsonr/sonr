package router

type Status string

const (
	StatusOK                  Status = "ok"
	StatusError               Status = "error"
	StatusNotFound            Status = "not-found"
	StatusUnauthorized        Status = "unauthorized"
	StatusForbidden           Status = "forbidden"
	StatusInternalServerError Status = "internal-server-error"
	StatusBadRequest          Status = "bad-request"
)
