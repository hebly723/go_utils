package utils

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	BEGIN_LINE    = "┍-*-_-_-_-S-T-A-R-T-_-_-_-*-┑"
	DURATION_LINE = "|>>-~~--~~~--~+~--~~~--~~-<<|"
	END_LINE      = "┕-^^^^^^^^^-E-N-D-^^^^^^^^^-┘"
)

type ResponseWriter struct {
	writer   http.ResponseWriter
	response string
	code     int
}

type printHandler struct {
	next http.Handler
}

func (r *ResponseWriter) Header() http.Header {
	return r.writer.Header()
}

func (r *ResponseWriter) Write(b []byte) (int, error) {
	r.response = string(b)
	return r.writer.Write(b)
}

func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.code = statusCode
	r.writer.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(printHandler{
		next: next,
	}.handler)
}

func (p printHandler) handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var reqBody string
	if r.Method == http.MethodPost {
		var buf bytes.Buffer
		teeReader := io.TeeReader(r.Body, &buf)
		dataByte, _ := io.ReadAll(teeReader)
		reqBody = string(dataByte)
		r.Body = io.NopCloser(&buf)
	}
	sw := &ResponseWriter{
		writer: w,
	}
	p.next.ServeHTTP(sw, r)
	log.Printf("\n%v\n"+
		"#Time:   %v\n"+
		"#Method: %v\n"+
		"#URL:    %v\n"+
		"#Body:   %v\n%v\n"+
		"#Code:   %v\n"+
		"#Body:   %v\n%v\n",
		BEGIN_LINE, time.Since(start),
		r.Method, r.RequestURI,
		reqBody, DURATION_LINE,
		sw.code, sw.response, END_LINE)
}
