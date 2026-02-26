package providers

import (
	"testing"

	"github.com/maslovpi/odd-character-htmx/models"
)

func TestInitStarterProvider(t *testing.T) {
	sp, err := InitStarterProvider()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sp.starterMap) == 0 {
		t.Error("expected non-empty starter map")
	}
}

func TestGenerateStarter(t *testing.T) {
	mockStarterMap := map[Key]Starter{
		{HitProtection: 1, Max: 9}: {
			Content: []Content{{Name: "Torch", ExtraInfo: ""}},
			Arcana:  false,
			Pet:     "",
			Hire:    "",
		},
		{HitProtection: 1, Max: 12}: {
			Content: []Content{{Name: "Shield", ExtraInfo: ""}},
			Arcana:  true,
			Pet:     "",
			Hire:    "",
		},
		{HitProtection: 2, Max: 9}: {
			Content: []Content{{Name: "Rope", ExtraInfo: "50 feet"}},
			Arcana:  false,
			Pet:     "Mutt",
			Hire:    "",
		},
		{HitProtection: 2, Max: 12}: {
			Content: []Content{{Name: "Lantern", ExtraInfo: ""}},
			Arcana:  false,
			Pet:     "",
			Hire:    "Guard",
		},
	}

	mockArcanaProvider := ArcanaProvider{
		arcanaSlice: []Arcana{
			{Name: "The Star", Description: "Hope and inspiration."},
		},
	}

	mockPetProvider := PetProvider{
		petMap: map[string]Pet{
			"Mutt": {Type: "Mutt", Cost: "5s", Strength: "d6", Attack: "Bite d6"},
		},
	}

	mockEquipmentProvider := EquipmentProvider{
		equipmentMap: map[string]Equipment{
			"Torch":  {Name: "Torch", Cost: "2s", Description: "", Examples: nil},
			"Shield": {Name: "Shield", Cost: "10g", Description: "Adds 1 armor.", Examples: nil},
			"Lantern": {
				Name:        "Lantern",
				Cost:        "5s",
				Description: "Burns for 6 hours.",
				Examples:    nil,
			},
		},
	}

	mockHireProvider := HireProvider{
		hireMap: map[string]Hire{
			"Guard": {
				Type:          "Guard",
				CostPerDay:    "3s",
				HitProtection: "4",
				Strength:      "",
				Equipment:     []string{},
				AbilityScores: 15,
			},
		},
		equipmentProvider: &mockEquipmentProvider,
	}

	getProvider := func() StarterProvider {
		return StarterProvider{
			starterMap:        mockStarterMap,
			equipmentProvider: &mockEquipmentProvider,
			arcanaProvider:    &mockArcanaProvider,
			petProvider:       &mockPetProvider,
			hireProvider:      &mockHireProvider,
		}
	}

	t.Run("should return error when key not found", func(t *testing.T) {
		sp := getProvider()
		_, err := sp.GenerateStarter(99, 99)
		if err == nil {
			t.Error("expected error for missing key, got nil")
		}
	})

	t.Run("should return description with content when found", func(t *testing.T) {
		sp := getProvider()
		got, err := sp.GenerateStarter(1, 8)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(got.Content) == 0 {
			t.Error("expected non-empty content")
		}
	})

	t.Run("should clamp maxStat below 9 to 9", func(t *testing.T) {
		sp := getProvider()
		got, err := sp.GenerateStarter(1, 5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(got.Content) == 0 {
			t.Error("expected non-empty content for clamped maxStat")
		}
	})

	t.Run("should not populate arcana when arcana is false", func(t *testing.T) {
		sp := getProvider()
		got, err := sp.GenerateStarter(1, 9)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		for _, item := range got.Content {
			if item.Type == models.Arcana {
				t.Error("expected arcana to be empty")
			}
		}
	})

	t.Run("should populate arcana when arcana is true", func(t *testing.T) {
		sp := getProvider()
		got, err := sp.GenerateStarter(1, 12)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		for _, item := range got.Content {
			if item.Type == models.Arcana {
				if item.Name == "" {
					t.Error("expected arcana name to be non-empty")
				}
				return
			}
		}
		t.Error("expected arcana to be populated")
	})

	t.Run("should populate pet when pet is set", func(t *testing.T) {
		sp := getProvider()
		got, err := sp.GenerateStarter(2, 9)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		for _, item := range got.Content {
			if item.Type == models.Pet {
				if item.Name != "Mutt" {
					t.Errorf("expected pet name %q, got %q", "Mutt", item.Name)
				}
				return
			}
		}
		t.Error("expected pet to be populated")
	})

	t.Run("should not populate pet when pet is empty", func(t *testing.T) {
		sp := getProvider()
		got, err := sp.GenerateStarter(1, 9)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		for _, item := range got.Content {
			if item.Type == models.Pet {
				t.Error("expected pet to be empty")
			}
		}
	})

	t.Run("should populate hire when hire is set", func(t *testing.T) {
		sp := getProvider()
		got, err := sp.GenerateStarter(2, 12)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		for _, item := range got.Content {
			if item.Type == models.Hire {
				if item.Name != "Guard" {
					t.Errorf("expected hire name %q, got %q", "Guard", item.Name)
				}
				return
			}
		}
		t.Error("expected hire to be populated")
	})

	t.Run("should not populate hire when hire is empty", func(t *testing.T) {
		sp := getProvider()
		got, err := sp.GenerateStarter(1, 9)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		for _, item := range got.Content {
			if item.Type == models.Hire {
				t.Error("expected hire to be empty")
			}
		}
	})
}
