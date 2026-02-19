package equipment

import (
	"testing"
)

func TestInitEquipmentProvider(t *testing.T) {
	ep, err := InitEquipmentProvider()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ep.equipmentMap) == 0 {
		t.Error("expected non-empty equipment map")
	}
}

func TestGetEquipmentDescription(t *testing.T) {
	mockMap := map[string]Equipment{
		"Shield": {
			Name:        "Shield",
			Cost:        "10g",
			Description: "Adds 1 armor.",
			Examples:    []string{"Wooden Shield", "Iron Shield"},
		},
		"Torch": {
			Name:        "Torch",
			Cost:        "2s",
			Description: "",
			Examples:    nil,
		},
		"Mystery Box": {
			Name:        "Mystery Box",
			Cost:        "",
			Description: "Contents unknown.",
			Examples:    nil,
		},
		"Empty Item": {
			Name:        "Empty Item",
			Cost:        "",
			Description: "",
			Examples:    nil,
		},
	}

	tests := []struct {
		name             string
		inputName        string
		givenDescription string
		want             string
	}{
		{
			name:             "found by exact name with cost and description",
			inputName:        "Shield",
			givenDescription: "",
			want:             "Shield (Cost: 10g, Description: Adds 1 armor.)",
		},
		{
			name:             "found by exact name with cost, no description",
			inputName:        "Torch",
			givenDescription: "",
			want:             "Torch (Cost: 2s)",
		},
		{
			name:             "found by exact name with no cost, has description",
			inputName:        "Mystery Box",
			givenDescription: "",
			want:             "Mystery Box (Cost: 0, Description: Contents unknown.)",
		},
		{
			name:             "found by exact name with no cost and no description",
			inputName:        "Empty Item",
			givenDescription: "",
			want:             "Empty Item (Cost: 0)",
		},
		{
			name:             "found by example name uses input name in output",
			inputName:        "Wooden Shield",
			givenDescription: "",
			want:             "Wooden Shield (Cost: 10g, Description: Adds 1 armor.)",
		},
		{
			name:             "not found with given description falls back to name plus description",
			inputName:        "Rope",
			givenDescription: "50 feet",
			want:             "Rope (50 feet)",
		},
		{
			name:             "not found with empty given description returns name only",
			inputName:        "Rope",
			givenDescription: "",
			want:             "Rope",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := providerWithMap(mockMap)
			got := ep.GetEquipmentDescription(tt.inputName, tt.givenDescription)
			if got != tt.want {
				t.Errorf("GetEquipmentDescription(%q, %q) = %q; want %q",
					tt.inputName, tt.givenDescription, got, tt.want)
			}
		})
	}
}

func providerWithMap(m map[string]Equipment) EquipmentProvider {
	return EquipmentProvider{equipmentMap: m}
}
