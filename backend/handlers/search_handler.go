package handlers

import (
	"encoding/json"
	"net/http"

	"matiks-leaderboard-backend/services"
)

func SearchHandler(service *services.LeaderboardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")

		if query == "" {
			http.Error(w, "query parameter is required", http.StatusBadRequest)
			return
		}

		results := service.SearchUsers(query)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}
