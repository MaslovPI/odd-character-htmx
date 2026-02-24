package providers

import "testing"

func TestInitArcanaProvider(t *testing.T) {
	ap, err := InitArcanaProvider()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ap.arcanaSlice) == 0 {
		t.Error("expected non-empty arcana slice")
	}
}

func TestGetRandomArcana(t *testing.T) {
	mockSlice := []Arcana{
		{Name: "The Fool", Description: "New beginnings."},
		{Name: "The World", Description: "Completion."},
	}

	t.Run("should return an arcana that exists in the slice", func(t *testing.T) {
		ap := getArcanaProviderWithSlice(mockSlice)
		got := ap.GetRandomArcana()
		found := false
		for _, a := range mockSlice {
			if a.Name == got.Name && a.Description == got.Description {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("GetRandomArcana() returned %+v which is not in the slice", got)
		}
	})

	t.Run("should always return a valid arcana over multiple calls", func(t *testing.T) {
		ap := getArcanaProviderWithSlice(mockSlice)
		for range 20 {
			got := ap.GetRandomArcana()
			found := false
			for _, a := range mockSlice {
				if a.Name == got.Name {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("GetRandomArcana() returned %+v which is not in the slice", got)
			}
		}
	})
}

func getArcanaProviderWithSlice(s []Arcana) ArcanaProvider {
	return ArcanaProvider{arcanaSlice: s}
}
