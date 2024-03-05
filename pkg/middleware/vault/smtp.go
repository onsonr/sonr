package vault

import (
	"os"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/sonrhq/sonr/internal/components/views/auth/register"
)

func Recovery(c echo.Context) recovery {
	return recovery{
		Context: c,
	}
}

type recovery struct {
	echo.Context
}

func (r recovery) getServer() (*mail.SMTPClient, error) {
	server := mail.NewSMTPClient()
	SmtpHost := os.Getenv("SMTP_HOST")
	SmtpUser := os.Getenv("SMTP_USER")
	SmtpPassword := os.Getenv("SMTP_PASS")
	portStr := os.Getenv("SMTP_PORT")
	if portStr != "" {
		portInt, err := strconv.Atoi(portStr)
		if err == nil {
			server.Port = portInt
		}
	}

	server.Host = SmtpHost
	server.Username = SmtpUser
	server.Password = SmtpPassword
	return server.Connect()
}

func (r recovery) SendConfirmationMail(to string, url string) error {
	smtp, err := r.getServer()
	if err != nil {
		return r.String(500, err.Error())
	}
	email := register.ConfirmationEmail(url)
	msg := mail.NewMSG()
	msg.SetFrom("foundation@sonr.id")
	msg.SetSubject("[Sonr ID] Confirm your email")
	tmp, err := templ.ToGoHTML(r.Request().Context(), email)
	if err != nil {
		return r.String(500, err.Error())
	}
	str := string(tmp)
	msg.SetBody(mail.TextHTML, str)
	msg.Send(smtp)
	return nil
}
