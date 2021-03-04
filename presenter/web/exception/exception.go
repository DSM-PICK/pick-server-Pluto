package exception

const (
	DefaultBadRequestException    = "d400"
	InvalidAuthorizationException = "a401"
)

type Exception map[string]string