package handlers

import (
	"encoding/json"
	"net/http"

	"matiks-leaderboard-backend/services"
)

func LeaderboardHandler(service *services.LeaderboardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entries := service.GetTopWithRanks(100)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entries)
	}
}
