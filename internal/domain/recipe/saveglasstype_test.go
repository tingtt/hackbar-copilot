package recipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_saverLister_SaveGlassType(t *testing.T) {
	t.Parallel()

	t.Run("may return validation error", func(t *testing.T) {
		t.Parallel()

		argGlassType := GlassType{Name: ""}

		s := &saverLister{Repository: nil}
		err := s.SaveGlassType(argGlassType)
		assert.Error(t, err)
	})

	t.Run("will call Repository.SaveGlassType with sanitized glass type", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name                     string
			arg                      GlassType
			wantArgCallSaveGlassType GlassType
		}{
			{"sanitized field/ImageURL",
				GlassType{
					Name:     "collins",
					ImageURL: ptr(""),
				},
				GlassType{
					Name:     "collins",
					ImageURL: nil,
				},
			},
			{"sanitized field/Description",
				GlassType{
					Name:        "collins",
					Description: ptr(""),
				},
				GlassType{
					Name:        "collins",
					Description: nil,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				mockRepository := new(MockRepository)
				mockRepository.On("SaveGlassType", mock.Anything).Return(nil)

				s := &saverLister{mockRepository}
				err := s.SaveGlassType(tt.arg)
				assert.NoError(t, err)
				mockRepository.AssertCalled(t, "SaveGlassType", tt.wantArgCallSaveGlassType)
			})
		}
	})
}
