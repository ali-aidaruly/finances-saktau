package errcode

type Code uint

const (
	Undefined Code = iota
	InvalidArgument
	BadRequest
	NotFound
	PermissionDenied
	Internal
	Unauthenticated
)
