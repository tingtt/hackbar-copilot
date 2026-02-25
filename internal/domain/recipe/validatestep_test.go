package recipe

import "testing"

type ValidateStepTest struct {
	Step  Step
	Valid bool
}

var validateStepTests = []ValidateStepTest{
	{
		Step: Step{
			Material: new("Peach Liqueur"),
			Amount:   new("30ml"),
		},
		Valid: true,
	},
	{
		Step: Step{
			Material:    nil,
			Amount:      nil,
			Description: nil,
		},
		Valid: false,
	},
	{
		Step: Step{
			Material:    nil,
			Amount:      new("30ml"),
			Description: nil,
		},
		Valid: false,
	},
}

func TestStep_Validate(t *testing.T) {
	t.Parallel()

	for _, tt := range validateStepTests {
		t.Run("will return nil when valid", func(t *testing.T) {
			t.Parallel()
			err := tt.Step.Validate()
			if tt.Valid && err != nil {
				t.Errorf("Step.Validate() error = %v, want nil", err)
			}
		})
	}
}
