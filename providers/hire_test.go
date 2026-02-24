package providers

import (
	"strings"
	"testing"
)

func TestInitHireProvider(t *testing.T) {
	ep, err := InitEquipmentProvider()
	if err != nil {
		t.Fatalf("unexpected error initializing equipment provider: %v", err)
	}
	hp, err := InitHireProvider(ep)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(hp.hireMap) == 0 {
		t.Error("expected non-empty hire map")
	}
}

func TestGetHireDescription(t *testing.T) {
	ep := getEquipmentProviderWithMap(map[string]Equipment{})

	mockMap := map[string]Hire{
		"Guard": {
			Type:          "Guard",
			CostPerDay:    "3s",
			HitProtection: "4",
			Strength:      "",
			Equipment:     []string{},
			AbilityScores: 15,
		},
		"Mage": {
			Type:          "Mage",
			CostPerDay:    "10g",
			HitProtection: "3",
			Strength:      "d6",
			Equipment:     []string{"Torch"},
			AbilityScores: 15,
		},
	}

	t.Run("should return empty string and no error when hire not found", func(t *testing.T) {
		hp := getHireProviderWithMap(mockMap, ep)
		got, err := hp.GetHireDescription("Wizard")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != "" {
			t.Errorf("expected empty string, got %q", got)
		}
	})

	t.Run("should return description with cost and hit protection when found", func(t *testing.T) {
		hp := getHireProviderWithMap(mockMap, ep)
		got, err := hp.GetHireDescription("Guard")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !strings.Contains(got, "Cost (per day): 3s") {
			t.Errorf("expected cost in description, got %q", got)
		}
		if !strings.Contains(got, "Hit protection: 4") {
			t.Errorf("expected hit protection in description, got %q", got)
		}
	})

	t.Run("should return description with ability scores when found", func(t *testing.T) {
		hp := getHireProviderWithMap(mockMap, ep)
		got, err := hp.GetHireDescription("Guard")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !strings.Contains(got, "Strength:") {
			t.Errorf("expected strength in description, got %q", got)
		}
		if !strings.Contains(got, "Dexterity:") {
			t.Errorf("expected dexterity in description, got %q", got)
		}
		if !strings.Contains(got, "Willpower:") {
			t.Errorf("expected willpower in description, got %q", got)
		}
	})

	t.Run("should return description with rolled pregen strength for mage", func(t *testing.T) {
		hp := getHireProviderWithMap(mockMap, ep)
		for range 20 {
			got, err := hp.GetHireDescription("Mage")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !strings.Contains(got, "Strength:") {
				t.Errorf("expected strength in description, got %q", got)
			}
			if !strings.Contains(got, "Torch") {
				t.Errorf("expected equipment in description, got %q", got)
			}
		}
	})

	t.Run("should return error for invalid dice format in strength", func(t *testing.T) {
		badMap := map[string]Hire{
			"Ghost": {
				Type:          "Ghost",
				CostPerDay:    "0",
				HitProtection: "1",
				Strength:      "invalid",
				Equipment:     []string{},
				AbilityScores: 10,
			},
		}
		hp := getHireProviderWithMap(badMap, ep)
		_, err := hp.GetHireDescription("Ghost")
		if err == nil {
			t.Error("expected error for invalid dice format, got nil")
		}
	})
}

func getHireProviderWithMap(m map[string]Hire, ep EquipmentProvider) HireProvider {
	return HireProvider{hireMap: m, equipmentProvider: ep}
}
