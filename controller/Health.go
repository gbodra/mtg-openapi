package controller

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Version: 0.1\nHealthy: True"))
}
