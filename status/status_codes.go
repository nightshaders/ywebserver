package status
import (
	"net/http"
	"fmt"
	"io"
)


type Code int

const (
	StatusContinue Code = http.StatusContinue
	StatusSwitchingProtocols Code = http.StatusSwitchingProtocols

	StatusOK Code = http.StatusOK
	StatusCreated Code = http.StatusCreated
	StatusAccepted Code = http.StatusAccepted
	StatusNonAuthoritativeInfo Code = http.StatusNonAuthoritativeInfo
	StatusNoContent Code = http.StatusNoContent
	StatusResetContent Code = http.StatusResetContent
	StatusPartialContent Code = http.StatusPartialContent

	StatusMultipleChoices Code = http.StatusMultipleChoices
	StatusMovedPermanently Code = http.StatusMovedPermanently
	StatusFound Code = http.StatusFound
	StatusSeeOther Code = http.StatusSeeOther
	StatusNotModified Code = http.StatusNotModified
	StatusUseProxy Code = http.StatusUseProxy
	StatusTemporaryRedirect Code = http.StatusTemporaryRedirect

	StatusBadRequest Code = http.StatusBadRequest
	StatusUnauthorized Code = http.StatusUnauthorized
	StatusPaymentRequired Code = http.StatusPaymentRequired
	StatusForbidden Code = http.StatusForbidden
	StatusNotFound Code = http.StatusNotFound
	StatusMethodNotAllowed Code = http.StatusMethodNotAllowed
	StatusNotAcceptable Code = http.StatusNotAcceptable
	StatusProxyAuthRequired Code = http.StatusProxyAuthRequired
	StatusRequestTimeout Code = http.StatusRequestTimeout
	StatusConflict Code = http.StatusConflict
	StatusGone Code = http.StatusGone
	StatusLengthRequired Code = http.StatusLengthRequired
	StatusPreconditionFailed Code = http.StatusPreconditionFailed
	StatusRequestEntityTooLarge Code = http.StatusRequestEntityTooLarge
	StatusRequestURITooLong Code = http.StatusRequestURITooLong
	StatusUnsupportedMediaType Code = http.StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable Code = http.StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed Code = http.StatusExpectationFailed
	StatusTeapot Code = http.StatusTeapot

	StatusInternalServerError Code = http.StatusInternalServerError
	StatusNotImplemented Code = http.StatusNotImplemented
	StatusBadGateway Code = http.StatusBadGateway
	StatusServiceUnavailable Code = http.StatusServiceUnavailable
	StatusGatewayTimeout Code = http.StatusGatewayTimeout
	StatusHTTPVersionNotSupported Code = http.StatusHTTPVersionNotSupported
)

func StatusText(code Code) string {
	return http.StatusText(int(code))
}

type ErrorResult struct {
	Code Code
	Message string
	Err error
	StatusText string
}

func (e Code) Message(msg string) *ErrorResult {
	return &ErrorResult{
		Code: e,
		Message: msg,
		StatusText: StatusText(e),
	}
}

func (e *ErrorResult) Provided(err error) *ErrorResult {
	e.Err = err
	return e
}

func (e *ErrorResult) Error() string {
	msg := e.Message
	if e.Err != nil {
		msg = fmt.Sprintf("%s => %s", msg, e.Err)
	}
	return fmt.Sprintf("%d %s %s", e.Code, StatusText(e.Code), msg)
}

func (e *ErrorResult) Log(w io.Writer) {
	w.Write([]byte(e.Error()))
}