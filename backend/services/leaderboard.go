package services

import (
	"sort"
	"sync"

	"matiks-leaderboard-backend/models"
)
import "strings"
import (
	"math/rand"
	"time"
)


type LeaderboardService struct {
	Users []models.User
	mu    sync.RWMutex
}

func NewLeaderboardService(users []models.User) *LeaderboardService {
	return &LeaderboardService{Users: users}
}

func (l *LeaderboardService) GetTopWithRanks(limit int) []models.LeaderboardEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()

	// Sort users by rating DESC
	sort.Slice(l.Users, func(i, j int) bool {
		return l.Users[i].Rating > l.Users[j].Rating
	})

	result := []models.LeaderboardEntry{}

	rank := 1
	prevRating := -1
	usersBefore := 0

	for i, user := range l.Users {
		if i == 0 {
			rank = 1
		} else if user.Rating < prevRating {
			rank = usersBefore + 1
		}

		result = append(result, models.LeaderboardEntry{
			Rank:     rank,
			Username: user.Username,
			Rating:   user.Rating,
		})

		prevRating = user.Rating
		usersBefore++

		if len(result) == limit {
			break
		}
	}

	return result
}

func (l *LeaderboardService) SearchUsers(query string) []models.LeaderboardEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()

	// Users are already sorted by rating DESC
	result := []models.LeaderboardEntry{}

	rank := 1
	prevRating := -1
	usersBefore := 0

	for i, user := range l.Users {
		// Rank calculation (same logic as leaderboard)
		if i == 0 {
			rank = 1
		} else if user.Rating < prevRating {
			rank = usersBefore + 1
		}

		// Search condition
		if strings.Contains(strings.ToLower(user.Username), strings.ToLower(query)) {
			result = append(result, models.LeaderboardEntry{
				Rank:     rank,
				Username: user.Username,
				Rating:   user.Rating,
			})
		}

		prevRating = user.Rating
		usersBefore++
	}

	return result
}

func (l *LeaderboardService) StartRatingUpdates() {
	go func() {
		rand.Seed(time.Now().UnixNano())

		for {
			l.mu.Lock()

			// Update ratings for random users
			for i := 0; i < 20; i++ { // simulate many players
				index := rand.Intn(len(l.Users))
				change := rand.Intn(201) - 100 // -100 to +100

				newRating := l.Users[index].Rating + change
				if newRating < 100 {
					newRating = 100
				}
				if newRating > 5000 {
					newRating = 5000
				}

				l.Users[index].Rating = newRating
			}

			l.mu.Unlock()

			time.Sleep(1 * time.Second)
		}
	}()
}
