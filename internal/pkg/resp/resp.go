package resp

import "net/http"

type H map[string]interface{}

// Http错误返回
func HttpError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
