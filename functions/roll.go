package functions

import (
	"errors"
	"math/rand"
	"time"
)

func roll(dimensions int) (int, error) {
	if dimensions < 1 {
		return 0, errors.New("Wrong dimensions")
	}

	rand.New(rand.NewSource(time.Now().Unix()))
	return rand.Intn(dimensions) + 1, nil
}

func RollMultipleDice(amount int, dimensions int) (int, error) {
	if amount < 1 {
		return 0, errors.New("Wrong number")
	}

	sum := 0
	for range amount {
		result, err := roll(dimensions)
		if err != nil {
			return 0, err
		}
		sum += result
	}
	return sum, nil
}
