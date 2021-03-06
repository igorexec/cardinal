package rest

import (
	"fmt"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"runtime"
	"strings"
)

const (
	ErrPageValidation = 1
	ErrCollectFail    = 2
	ErrSaveFail       = 3
	ErrNoData         = 4
)

func SendErrorJSON(w http.ResponseWriter, r *http.Request, code int, err error, details string, errCode int) {
	log.Printf("[debug] %s", errDetailsMsg(r, code, err, details, errCode))
	render.Status(r, code)
	render.JSON(w, r, JSON{"error": err.Error(), "details": details, "code": errCode})
}

func errDetailsMsg(r *http.Request, code int, err error, details string, errCode int) string {
	remoteIP := r.RemoteAddr
	q := r.URL.Query()

	srcFileInfo := ""
	if pc, file, line, ok := runtime.Caller(2); ok {
		fnameElems := strings.Split(file, "/")
		funcNameElems := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		srcFileInfo = fmt.Sprintf("[caused by %s:%d %s]", strings.Join(fnameElems[len(fnameElems)-3:], "/"), line, funcNameElems[len(funcNameElems)-1])
	}
	return fmt.Sprintf("%s - %v - %d (%d) - %s - %s - %s", details, err, code, errCode, remoteIP, q, srcFileInfo)
}
