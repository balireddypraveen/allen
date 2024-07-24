package common_utils

import (
	"crypto/rand"
	"math/big"
	"time"
)

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result += string(charset[index.Int64()])
	}

	return result, nil
}

func GetStartAndEndTimeForDeal() (*time.Time, *time.Time, error) {
	// todo: read time from config
	startTime := time.Now()
	endTime := time.Now().Local().Add(time.Minute * time.Duration(15))
	return &startTime, &endTime, nil
}
