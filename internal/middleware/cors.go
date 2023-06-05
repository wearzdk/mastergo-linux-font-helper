package middleware

import "net/http"

// CORS HTTP
func CORSMiddleware(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "https://mastergo.com")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length,Authorization,Accept,X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return false
	}
	return true
}

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !CORSMiddleware(w, r) {
			return
		}
		h.ServeHTTP(w, r)
	})
}
