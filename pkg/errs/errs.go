package errs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/ali-aidaruly/finances-saktau/pkg/errs/errcode"

	"github.com/lib/pq"
	"go.uber.org/zap/zapcore"
)

type Error struct {
	Code    errcode.Code `json:"code,omitempty"`
	Message string       `json:"message"`
	Params  Params       `json:"params,omitempty"`
}

type Params map[string]string

func (p Params) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	for key, value := range p {
		encoder.AddString(key, value)
	}

	return nil
}

func (e *Error) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("message", e.Message)
	encoder.AddUint("code", uint(e.Code))

	if err := encoder.AddObject("params", e.Params); err != nil {
		return err
	}

	return nil
}

func (e *Error) WithParam(key string, value string) *Error {
	e.AddParam(key, value)

	return e
}

func (e *Error) WithParams(params map[string]string) *Error {
	for key, value := range params {
		e.AddParam(key, value)
	}

	return e
}

func (e *Error) Error() string {
	data, _ := json.Marshal(e)

	return string(data)
}

func (e *Error) Is(tgt error) bool {
	target, ok := tgt.(*Error)
	if !ok {
		return false
	}

	return reflect.DeepEqual(e, target)
}

func (e *Error) SetCode(code errcode.Code) {
	e.Code = code
}

func (e *Error) SetMessage(message string) {
	e.Message = message
}

func (e *Error) SetParams(params Params) {
	e.Params = params
}

func (e *Error) AddParam(key string, value string) {
	e.Params[key] = value
}

func NewError(code errcode.Code, message string) *Error {
	return &Error{Code: code, Message: message, Params: Params{}}
}

func NewUnexpectedBehaviorError(details string) *Error {
	return &Error{
		Code:    errcode.Unauthenticated,
		Message: "Unexpected behavior.",
		Params:  Params{"details": details},
	}
}

func NewInvalidFormError() *Error {
	return NewError(
		errcode.InvalidArgument,
		"The form sent is not valid, please correct the errors below.",
	)
}

// TODO: still experimental
func FromPostgresError(err error) *Error {
	e := &Error{
		Code:    errcode.Internal,
		Message: "Unexpected behavior.",
		Params:  Params{"error": err.Error()},
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		e.AddParam("details", pqErr.Detail)
		e.AddParam("message", pqErr.Message)
		e.AddParam("postgres_code", fmt.Sprint(pqErr.Code))

		if pqErr.Code == "23505" {
			e = NewInvalidFormError().WithParam(pqErr.Constraint, pqErr.Detail)
		}
	}

	if errors.Is(err, sql.ErrNoRows) {
		e = NewEntityNotFound()
	}

	return e
}

func NewEntityNotFound() *Error {
	return &Error{
		Code:    errcode.NotFound,
		Message: "Entity not found.",
		Params:  Params{},
	}
}

func NewPermissionDenied() *Error {
	return &Error{
		Code:    errcode.PermissionDenied,
		Message: "Permission denied.",
		Params:  Params{},
	}
}

func NewBadToken() *Error {
	return &Error{
		Code:    errcode.Unauthenticated,
		Message: "Bad token.",
		Params:  Params{},
	}
}
