package data

import (
	"math/rand"
	"strings"
	"time"

	"matiks-leaderboard-backend/models"
)

func SeedUsers(count int) []models.User {
	rand.Seed(time.Now().UnixNano())
	users := make([]models.User, count)

	separators := []string{"_", ".", ""}

	for i := 0; i < count; i++ {
		first := FirstNames[rand.Intn(len(FirstNames))]
		last := LastNames[rand.Intn(len(LastNames))]
		sep := separators[rand.Intn(len(separators))]

		username := first + sep + last

		// occasional suffix for uniqueness
		if rand.Float32() < 0.3 {
			username += string('a' + rune(rand.Intn(26)))
		}

		users[i] = models.User{
			ID:       i + 1,
			Username: strings.ToLower(username),
			Rating:   rand.Intn(4901) + 100,
		}
	}

	return users
}
