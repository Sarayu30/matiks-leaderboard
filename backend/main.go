package main

import (
	"log"
	"net/http"

	"matiks-leaderboard-backend/data"
	"matiks-leaderboard-backend/handlers"
	"matiks-leaderboard-backend/services"
)

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Seed users
	users := data.SeedUsers(10000)

	// Create leaderboard service
	leaderboardService := services.NewLeaderboardService(users)

	// Start live rating updates
	leaderboardService.StartRatingUpdates()

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/leaderboard", handlers.LeaderboardHandler(leaderboardService))
	mux.HandleFunc("/search", handlers.SearchHandler(leaderboardService))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", enableCORS(mux)))
}
