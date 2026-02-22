package functions

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Roll(dimensions int) (int, error) {
	if dimensions < 1 {
		return 0, errors.New("Wrong dimensions")
	}

	rand.New(rand.NewSource(time.Now().Unix()))
	return rand.Intn(dimensions) + 1, nil
}

func RollMultipleDice(amount int, dimensions int) (int, error) {
	if amount < 1 {
		return 0, errors.New("Wrong amount")
	}

	sum := 0
	for range amount {
		result, err := Roll(dimensions)
		if err != nil {
			return 0, err
		}
		sum += result
	}
	return sum, nil
}

func RollDice(diceToRoll string) (int, error) {
	diceInfo := strings.Split(diceToRoll, "d")
	if len(diceInfo) != 2 {
		return 0, errors.New("Invalid format. Use NdM (e.g. 2d6)")
	}

	number := 1
	if diceInfo[0] != "" {
		var err error
		number, err = strconv.Atoi(diceInfo[0])
		if err != nil {
			return 0, errors.New("Invalid number. Use NdM (e.g. 2d6)")
		}
	}

	var err error
	dimensions, err := strconv.Atoi(diceInfo[1])
	if err != nil {
		return 0, errors.New("Invalid dimensions. Use NdM (e.g. 2d6)")
	}
	return RollMultipleDice(number, dimensions)
}
