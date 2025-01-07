package recipe

import "testing"

type ValidateGlassTypeTest struct {
	GlassType GlassType
	Valid     bool
}

var ValidateGlassTypeTests = []ValidateGlassTypeTest{
	{GlassType{Name: "collins"}, true},
	{GlassType{Name: ""}, false},
}

func TestGlassType_Validate(t *testing.T) {
	t.Parallel()
	for _, tt := range ValidateGlassTypeTests {
		t.Run("will return nil when valid", func(t *testing.T) {
			t.Parallel()
			err := tt.GlassType.Validate()
			if tt.Valid && err != nil {
				t.Errorf("GlassType.Validate() error = %v, want nil", err)
			}
		})
	}
}
