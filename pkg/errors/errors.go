package errors

// Error applicational
type Error struct {
	ErrorType Type
	Message   string
	cause     error
}

// Error message
func (e Error) Error() string {
	return e.Message
}

// Cause returns the original error message
func (e Error) Cause() string {

	if e.cause != nil {
		return e.cause.Error()
	}

	return ""
}

// Type defines the type of an error
type Type int

const (
	_ Type = (iota * 10000) << 1
	// Internal error
	Internal
	// NotExist error means that a specific item does not exist
	NotExist
	// Malformed error represents data that not respect the standard format
	Malformed
	// Validation error
	Validation
)

func (t Type) String() string {
	switch t {
	case Internal:
		return "Internal Error"
	case NotExist:
		return "Item does not exist"
	case Malformed:
		return "Malformed error"
	case Validation:
		return "Validation error"
	}

	return "Unknown error"
}

// New creates a new error
func New(t Type, msg string, err error) error {
	return &Error{
		ErrorType: t,
		Message:   msg,
		cause:     err,
	}
}

// EValidation creates an error of type Validationn
func EValidation(msg string, err error) error {
	return New(Validation, msg, err)
}

// ENotExists creates an error of type NotExist
func ENotExists(msg string, err error) error {
	return New(NotExist, msg, err)
}

// EMalformed creates an error of type Malformed
func EMalformed(msg string, err error) error {
	return New(Malformed, msg, err)
}

// EInternal creates an error of type Internal
func EInternal(msg string, err error) error {
	return New(Internal, msg, err)
}

// Is method checks if an error is of a specific type
func Is(t Type, err error) bool {
	e, ok := err.(*Error)

	return ok && e.ErrorType == t
}
