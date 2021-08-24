package context

type key int

const (
	// HTTPCode is used to return a special context to represent http status code.
	HTTPCode key = 1
)
