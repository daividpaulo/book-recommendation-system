package httpdelivery

import "net/http"

func NewRouter(handlers *Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("POST /api/v1/books", handlers.CreateBook)
	mux.HandleFunc("GET /api/v1/books", handlers.ListBooks)
	mux.HandleFunc("POST /api/v1/users", handlers.CreateUser)
	mux.HandleFunc("GET /api/v1/users", handlers.ListUsers)
	mux.HandleFunc("POST /api/v1/recommendations/train", handlers.TriggerTraining)
	mux.HandleFunc("GET /api/v1/recommendations/{userId}", handlers.GetRecommendations)
	mux.HandleFunc("POST /api/v1/purchases", handlers.CreatePurchase)
	mux.HandleFunc("GET /api/v1/users/{userId}/purchases", handlers.ListPurchasesByUser)
	return withJSONContentType(mux)
}

func withJSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
