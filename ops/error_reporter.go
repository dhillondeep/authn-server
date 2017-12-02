package ops

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrorReporter is a thing that exports details about errors and panics to another service. Care
// must be taken by each implementation to ensure that passwords are not leaked.
type ErrorReporter interface {
	ReportError(err error)
	ReportRequestError(err error, r *http.Request)
}

// PanicHandler returns a http.Handler that will recover any panics and report them as request
// errors. If a panic is caught, the handler will return HTTP 500.
func PanicHandler(r ErrorReporter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			val := recover()
			switch err := val.(type) {
			case nil:
				return
			case error:
				r.ReportRequestError(err, req)
				w.WriteHeader(http.StatusInternalServerError)
			default:
				r.ReportRequestError(errors.New(fmt.Sprint(err)), req)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, req)
	})
}
