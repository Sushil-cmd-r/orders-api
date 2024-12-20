package application

import "net/http"

func (a *App) loadRoutes() {
	router := http.NewServeMux()

	router.HandleFunc("GET /health", check)
	a.router = router
}

func check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
