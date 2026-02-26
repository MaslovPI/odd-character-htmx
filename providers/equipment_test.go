package providers

import (
	"reflect"
	"testing"

	"github.com/maslovpi/odd-character-htmx/models"
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
			Description: "Adds 1 armor",
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
			Description: "Contents unknown",
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
		want             models.NamedItem
	}{
		{
			name:             "should find by exact name with cost and description",
			inputName:        "Shield",
			givenDescription: "",
			want: models.NamedItem{
				Name:        "Shield",
				Description: "Cost: 10g, Description: Adds 1 armor",
			},
		},
		{
			name:             "should find by exact name with cost, no description",
			inputName:        "Torch",
			givenDescription: "",
			want: models.NamedItem{
				Name:        "Torch",
				Description: "Cost: 2s",
			},
		},
		{
			name:             "should find by exact name with no cost, has description",
			inputName:        "Mystery Box",
			givenDescription: "",
			want: models.NamedItem{
				Name:        "Mystery Box",
				Description: "Cost: 0, Description: Contents unknown",
			},
		},
		{
			name:             "should find by exact name with no cost and no description",
			inputName:        "Empty Item",
			givenDescription: "",
			want: models.NamedItem{
				Name:        "Empty Item",
				Description: "Cost: 0",
			},
		},
		{
			name:             "should find by example name uses input name in output",
			inputName:        "Wooden Shield",
			givenDescription: "",
			want: models.NamedItem{
				Name:        "Wooden Shield",
				Description: "Cost: 10g, Description: Adds 1 armor",
			},
		},
		{
			name:             "should not find and should fall back to name plus description",
			inputName:        "Rope",
			givenDescription: "50 feet",
			want: models.NamedItem{
				Name:        "Rope",
				Description: "50 feet",
			},
		},
		{
			name:             "should not find and should return name only",
			inputName:        "Rope",
			givenDescription: "",
			want: models.NamedItem{
				Name: "Rope",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := getEquipmentProviderWithMap(mockMap)
			got := ep.GetEquipmentDescription(tt.inputName, tt.givenDescription)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEquipmentDescription(%v, %v) = %v; want %v",
					tt.inputName, tt.givenDescription, got, tt.want)
			}
		})
	}
}

func getEquipmentProviderWithMap(m map[string]Equipment) EquipmentProvider {
	return EquipmentProvider{equipmentMap: m}
}
