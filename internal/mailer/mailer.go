package mailer

const (
	FromName   = "GoBuzz"
	maxRetries = 3
)

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) error
}
