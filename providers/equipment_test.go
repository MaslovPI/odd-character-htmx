package providers

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
			name:             "should find by exact name with cost and description",
			inputName:        "Shield",
			givenDescription: "",
			want:             "Shield (Cost: 10g, Description: Adds 1 armor.)",
		},
		{
			name:             "should find by exact name with cost, no description",
			inputName:        "Torch",
			givenDescription: "",
			want:             "Torch (Cost: 2s)",
		},
		{
			name:             "should find by exact name with no cost, has description",
			inputName:        "Mystery Box",
			givenDescription: "",
			want:             "Mystery Box (Cost: 0, Description: Contents unknown.)",
		},
		{
			name:             "should find by exact name with no cost and no description",
			inputName:        "Empty Item",
			givenDescription: "",
			want:             "Empty Item (Cost: 0)",
		},
		{
			name:             "should find by example name uses input name in output",
			inputName:        "Wooden Shield",
			givenDescription: "",
			want:             "Wooden Shield (Cost: 10g, Description: Adds 1 armor.)",
		},
		{
			name:             "should not find and should fall back to name plus description",
			inputName:        "Rope",
			givenDescription: "50 feet",
			want:             "Rope (50 feet)",
		},
		{
			name:             "should not find and should return name only",
			inputName:        "Rope",
			givenDescription: "",
			want:             "Rope",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := getEquipmentProviderWithMap(mockMap)
			got := ep.GetEquipmentDescription(tt.inputName, tt.givenDescription)
			if got != tt.want {
				t.Errorf("GetEquipmentDescription(%q, %q) = %q; want %q",
					tt.inputName, tt.givenDescription, got, tt.want)
			}
		})
	}
}

func getEquipmentProviderWithMap(m map[string]Equipment) EquipmentProvider {
	return EquipmentProvider{equipmentMap: m}
}
