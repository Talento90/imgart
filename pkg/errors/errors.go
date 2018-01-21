package errors

// Error applicational
type Error struct {
	Type    Type
	Message string
	cause   error
}

// Error message
func (e *Error) Error() string {
	return e.Message
}

// Cause of the original error
func (e *Error) Cause() string {
	if e.cause != nil {
		return e.cause.Error()
	}

	return ""
}

// Type defines the type of an error
type Type string

const (
	// Internal error
	Internal Type = "internal"
	// NotFound error means that a specific item does not exist
	NotFound Type = "not_found"
	// Malformed error represents data that not respect the standard format
	Malformed Type = "malformed"
	// Validation error
	Validation Type = "validation"
	// AlreadyExists error
	AlreadyExists Type = "already_exists"
)

func (t Type) String() string {
	switch t {
	case Internal:
		return "Internal Error"
	case NotFound:
		return "Item not found"
	case Malformed:
		return "Malformed error"
	case Validation:
		return "Validation error"
	case AlreadyExists:
		return "Item already exists"
	}

	return "Unknown error"
}

// New creates a new error
func New(t Type, msg string, err error) error {
	return &Error{
		Type:    t,
		Message: msg,
		cause:   err,
	}
}

// EValidation creates an error of type Validationn
func EValidation(msg string, err error) error {
	return New(Validation, msg, err)
}

// ENotExists creates an error of type NotExist
func ENotExists(msg string, err error) error {
	return New(NotFound, msg, err)
}

// EMalformed creates an error of type Malformed
func EMalformed(msg string, err error) error {
	return New(Malformed, msg, err)
}

// EAlreadyExists creates an error of type EAlreadyExistsl
func EAlreadyExists(msg string, err error) error {
	return New(AlreadyExists, msg, err)
}

// EInternal creates an error of type Internal
func EInternal(msg string, err error) error {
	return New(Internal, msg, err)
}

// Is method checks if an error is of a specific type
func Is(t Type, err error) bool {
	e, ok := err.(*Error)

	return ok && e.Type == t
}
