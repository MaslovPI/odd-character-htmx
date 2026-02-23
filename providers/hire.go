package providers

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/maslovpi/odd-character-htmx/functions"
)

//go:embed data/hires.json
var hireJSON []byte

type (
	Hire struct {
		Type          string   `json:"type"`
		CostPerDay    string   `json:"cost_per_day"`
		HitProtection string   `json:"hp"`
		Strength      string   `json:"str"`
		Equipment     []string `json:"equipment"`
		AbilityScores int      `json:"ability_scores"`
	}
	HireProvider struct {
		hireMap           map[string]Hire
		equipmentProvider EquipmentProvider
	}
)

func InitHireProvider(equipmentProvider EquipmentProvider) (HireProvider, error) {
	hireMap, err := getHireMap()
	return HireProvider{hireMap: hireMap, equipmentProvider: equipmentProvider}, err
}

func getHireMap() (map[string]Hire, error) {
	var hireSlice []Hire
	err := json.Unmarshal(hireJSON, &hireSlice)
	if err != nil {
		return nil, err
	}

	hireMap := make(map[string]Hire, len(hireSlice))

	for _, item := range hireSlice {
		hireMap[item.Type] = item
	}
	return hireMap, nil
}

func (h *HireProvider) GetHireDescription(hireType string) (string, error) {
	item, exists := h.hireMap[hireType]
	if !exists {
		return "", nil
	}

	var description strings.Builder
	fmt.Fprintf(&description, "Cost (per day): %s\n", item.CostPerDay)
	fmt.Fprintf(&description, "Hit protection: %s\n", item.HitProtection)
	availableScore := item.AbilityScores
	pregenStrength := 0
	var err error
	if item.Strength != "" {
		pregenStrength, err = functions.RollDice(item.Strength)
	}
	if err != nil {
		return "", err
	}

	strength, dexterity, willpower := generateHireAbilityScores(availableScore, pregenStrength)

	fmt.Fprintf(&description,
		"Strength: %d\nDexterity: %d\nWillpower: %d\n",
		strength,
		dexterity,
		willpower,
	)

	for _, equipment := range item.Equipment {
		fmt.Fprintf(
			&description,
			"%s\n",
			h.equipmentProvider.GetEquipmentDescription(equipment, ""),
		)
	}
	return description.String(), nil
}

func generateHireAbilityScores(
	availableScore, pregenStrength int,
) (strength, dexterity, willpower int) {
	strength = pregenStrength
	if strength == 0 {
		strength = generateScore(availableScore - 2)
	}
	availableScore -= strength
	dexterity = generateScore(availableScore - 1)
	availableScore -= dexterity
	willpower = max(availableScore, 1)
	return
}

func generateScore(availableScore int) int {
	if availableScore < 2 {
		return 1
	}
	rand.New(rand.NewSource(time.Now().Unix()))
	return rand.Intn(availableScore) + 1
}
