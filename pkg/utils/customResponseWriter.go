package utils

import "net/http"

type CustomResponseWriter struct {
	responseWriter http.ResponseWriter
	StatusCode     int
}

func ExtendResponseWriter(w http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{w, 0}
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	return w.responseWriter.Write(b)
}

func (w *CustomResponseWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	// receive status code from this method
	w.StatusCode = statusCode
	w.responseWriter.WriteHeader(statusCode)

	return
}

func (w *CustomResponseWriter) Done() {
	// if the `w.WriteHeader` wasn't called, set status code to 200 OK
	if w.StatusCode == 0 {
		w.StatusCode = http.StatusOK
	}

	return
}
