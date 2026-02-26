package providers

import (
	"fmt"
	"strings"
	"testing"
)

func TestInitPetProvider(t *testing.T) {
	pp, err := InitPetProvider()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(pp.petMap) == 0 {
		t.Error("expected non-empty pet map")
	}
}

func TestGetPetDescription(t *testing.T) {
	mockMap := map[string]Pet{
		"Mutt": {
			Type:     "Mutt",
			Cost:     "5s",
			Strength: "d6",
			Attack:   "Bite d6",
		},
	}

	t.Run("should return empty string and no error when pet not found", func(t *testing.T) {
		pp := getPetProviderWithMap(mockMap)
		got, err := pp.GetPetDescription("Dragon")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != "" {
			t.Errorf("expected empty string, got %q", got)
		}
	})

	t.Run(
		"should return description with cost, strength and attack when found",
		func(t *testing.T) {
			pp := getPetProviderWithMap(mockMap)
			got, err := pp.GetPetDescription("Mutt")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !strings.Contains(got, "Cost: 5s") {
				t.Errorf("expected cost in description, got %q", got)
			}
			if !strings.Contains(got, "Attack: Bite d6") {
				t.Errorf("expected attack in description, got %q", got)
			}
			if !strings.Contains(got, "Strength:") {
				t.Errorf("expected strength in description, got %q", got)
			}
		},
	)

	t.Run("should roll strength within dice range", func(t *testing.T) {
		pp := getPetProviderWithMap(mockMap)
		for range 20 {
			got, err := pp.GetPetDescription("Mutt")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			var strength int
			fmt.Sscanf(got, "Cost: 5s, Strength: %d, Attack: Bite d6", &strength)
			if strength < 1 || strength > 6 {
				t.Errorf("strength %d out of d6 range [1,6]", strength)
			}
		}
	})

	t.Run("should return error for invalid dice format", func(t *testing.T) {
		badMap := map[string]Pet{
			"Ghost": {
				Type:     "Ghost",
				Cost:     "0",
				Strength: "invalid",
				Attack:   "Spook",
			},
		}
		pp := getPetProviderWithMap(badMap)
		_, err := pp.GetPetDescription("Ghost")
		if err == nil {
			t.Error("expected error for invalid dice format, got nil")
		}
	})
}

func getPetProviderWithMap(m map[string]Pet) PetProvider {
	return PetProvider{petMap: m}
}
