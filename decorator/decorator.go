package decorator

import (
	"net/http"
	"github.com/nightshaders/ywebserver/status"
	"os"
	"fmt"
	"runtime/debug"
)


type Decorator func(h WebHandler) WebHandler

func NewPipeline() Decorator {
	return func(h WebHandler) WebHandler {
		return h
	}
}
func (d Decorator) Next(next Decorator) Decorator {
	return func(w WebHandler) WebHandler {
		return d(next(w))
	}
}
func (d Decorator) Using(pd ParamsDecorator) Decorator {
	return d.Next(pd.Handle())
}

func (d Decorator) Handle(handler Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		p := NewParams(res, req)
		h := d(handler.Handle)
		err := h.Handle(p)

		if err != nil {
			status.StatusInternalServerError.
				Message("Creating new params").
				Provided(err).
				Log(os.Stdout)
			http.Error(p.Response(), err.Error(), int(status.StatusInternalServerError))
		}
	}
}

func ProvideErrorHandling(w WebHandler) WebHandler {
	return func(p Params) (err error) {
		defer func() {
			var result *status.ErrorResult
			val := recover()
			if val == nil {
				return
			}
			e, ok := val.(error)
			if ok {
				result = status.StatusInternalServerError.Message("Captured Error").Provided(e)
			} else {
				result = status.StatusInternalServerError.Message("Non-Error ErrorHandling")
			}
			msg := fmt.Sprintf("%s %s", string(debug.Stack()), result.Error())
			http.Error(p.Response(), msg, int(result.Code))
		}()
		return w.Handle(p)
	}
}


