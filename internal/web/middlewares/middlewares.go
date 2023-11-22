package middlewares

import (
	"net/http"
	"strings"

	"github.com/elnormous/contenttype"
	"github.com/go-chi/chi/v5/middleware"
)

func SetHtmlContent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func SetJsonContent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func FilterAcceptHeaderAndSetContentType(availableMediaTypes []contenttype.MediaType) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s := strings.Builder{}
			s.WriteString("\nthe available representations are ")
			for i, at := range availableMediaTypes {
				if i+1 == len(availableMediaTypes) {
					s.WriteString(at.MIME())
				} else {
					s.WriteString(at.MIME() + ", ")
				}
			}
			accepted, _, err := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
			if err == contenttype.ErrNoAcceptableTypeFound {
				http.Error(w, err.Error()+s.String(), http.StatusNotAcceptable)
				return
			} else if err == contenttype.ErrNoAvailableTypeGiven {
				panic(err)
			} else if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			middleware.SetHeader("Content-Type", accepted.MIME())(next).ServeHTTP(w, r)
		})
	}
}
