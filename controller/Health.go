package controller

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Version: 1.0.0\nHealthy: True"))
}
