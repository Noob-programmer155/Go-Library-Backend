package configuration

import (
	"net/http"
	"os"
)

func webConfigHeader(w http.ResponseWriter, cors bool) {
	w.Header().Add("Expires", "0")
	w.Header().Add("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
	w.Header().Add("X-Frame-Options", "DENY")
	w.Header().Add("X-XSS-Protection", "1; mode=block")

	//cors
	if cors {
		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Access-Control-Expose-Headers", "Set-Cookie")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "*")
	}
}
