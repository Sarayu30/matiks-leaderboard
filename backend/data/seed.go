package data

import (
	"math/rand"
	"strconv"
	"time"

	"matiks-leaderboard-backend/models"
)

func SeedUsers(count int) []models.User {
	rand.Seed(time.Now().UnixNano())

	users := make([]models.User, count)

	for i := 0; i < count; i++ {
		users[i] = models.User{
			ID:       i + 1,
			Username: "user_" + strconv.Itoa(i+1),
			Rating:   rand.Intn(4901) + 100,
		}
	}

	return users
}
