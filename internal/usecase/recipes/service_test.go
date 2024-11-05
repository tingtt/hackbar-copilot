package recipes

import (
	"errors"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewService(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		resolver := NewService(new(MockRepository))

		assert.NotNil(t, resolver)
	})
}

func Test_service_Find(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository *MockRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []RecipeGroup
		wantErr bool
	}{
		{
			name: "will return response from repository",
			fields: fields{
				repository: func() *MockRepository {
					m := new(MockRepository)
					m.On("Find").Return([]RecipeGroup{{
						Name:     "Phuket Sling",
						ImageURL: nil,
						Recipes: []Recipe{{
							Name:  "Phuket Sling",
							Type:  "build",
							Glass: "collins",
							Steps: []string{
								"Peach liqueur 30ml",
								"Blue curacao 10ml",
								"Grapefruit juice 30ml",
								"Tonic water - Full up",
							},
						}},
					}}, nil)
					return m
				}(),
			},
			want: []RecipeGroup{{
				Name:     "Phuket Sling",
				ImageURL: nil,
				Recipes: []Recipe{{
					Name:  "Phuket Sling",
					Type:  "build",
					Glass: "collins",
					Steps: []string{
						"Peach liqueur 30ml",
						"Blue curacao 10ml",
						"Grapefruit juice 30ml",
						"Tonic water - Full up",
					},
				}},
			}},
			wantErr: false,
		},
		{
			name: "will return response from repository",
			fields: fields{
				repository: func() *MockRepository {
					m := new(MockRepository)
					m.On("Find").Return([]RecipeGroup{}, errors.New("error"))
					return m
				}(),
			},
			want:    []RecipeGroup{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &service{
				repository: tt.fields.repository,
			}

			got, err := s.Find()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_service_FindRecipeType(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository *MockRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]model.RecipeType
		wantErr bool
	}{
		{
			name: "will return response from repository",
			fields: fields{
				repository: func() *MockRepository {
					m := new(MockRepository)
					m.On("FindRecipeType").Return(map[string]model.RecipeType{
						"build": {
							Name:        "build",
							Description: nil,
						},
					}, nil)
					return m
				}(),
			},
			want: map[string]model.RecipeType{
				"build": {
					Name:        "build",
					Description: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "will return response from repository",
			fields: fields{
				repository: func() *MockRepository {
					m := new(MockRepository)
					m.On("FindRecipeType").Return(map[string]model.RecipeType{}, errors.New("error"))
					return m
				}(),
			},
			want:    map[string]model.RecipeType{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &service{
				repository: tt.fields.repository,
			}

			got, err := s.FindRecipeType()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_service_FindGlassType(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository *MockRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]model.GlassType
		wantErr bool
	}{
		{
			name: "will return response from repository",
			fields: fields{
				repository: func() *MockRepository {
					m := new(MockRepository)
					m.On("FindGlassType").Return(map[string]model.GlassType{
						"collins": {
							Name:        "collins",
							Description: nil,
						},
					}, nil)
					return m
				}(),
			},
			want: map[string]model.GlassType{
				"collins": {
					Name:        "collins",
					Description: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "will return response from repository",
			fields: fields{
				repository: func() *MockRepository {
					m := new(MockRepository)
					m.On("FindGlassType").Return(map[string]model.GlassType{}, errors.New("error"))
					return m
				}(),
			},
			want:    map[string]model.GlassType{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &service{tt.fields.repository}

			got, err := s.FindGlassType()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_service_Register(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository *MockRepository
	}
	type args struct {
		input model.InputRecipeGroup
	}
	tests := []struct {
		name    string
		fields  fields
		before  func(fields)
		args    args
		want    RecipeGroup
		wantErr bool
		after   func(*testing.T, fields)
	}{
		{
			name: "may add new recipeGroup",
			fields: fields{
				repository: new(MockRepository),
			},
			before: func(f fields) {
				f.repository.On("FindOne", "NewRecipeGroup").Return(RecipeGroup{}, usecaseutils.ErrNotFound)
				f.repository.On("SaveRecipeType", mock.Anything).Return(nil)
				f.repository.On("SaveGlassType", mock.Anything).Return(nil)
				f.repository.On("Save", mock.Anything).Return(nil)
			},
			args: args{
				input: model.InputRecipeGroup{Name: "NewRecipeGroup"},
			},
			want:    RecipeGroup{Name: "NewRecipeGroup"},
			wantErr: false,
			after: func(t *testing.T, f fields) {
				f.repository.AssertCalled(t, "Save", RecipeGroup{Name: "NewRecipeGroup"})
			},
		},
		{
			name: "may update recipeGroup imageURL",
			fields: fields{
				repository: new(MockRepository),
			},
			before: func(f fields) {
				f.repository.On("FindOne", "ExistsRecipeGroup").Return(RecipeGroup{
					Name:    "ExistsRecipeGroup",
					Recipes: []Recipe{},
				}, nil)
				f.repository.On("Save", mock.Anything).Return(nil)
			},
			args: args{
				input: model.InputRecipeGroup{
					Name: "ExistsRecipeGroup",
					ImageURL: func() *string {
						text := "https://example.com/path/to/image"
						return &text
					}(),
				},
			},
			want: RecipeGroup{
				Name: "ExistsRecipeGroup",
				ImageURL: func() *string {
					text := "https://example.com/path/to/image"
					return &text
				}(),
				Recipes: nil,
			},
			wantErr: false,
			after: func(t *testing.T, f fields) {
				f.repository.AssertCalled(t, "Save", RecipeGroup{
					Name: "ExistsRecipeGroup",
					ImageURL: func() *string {
						text := "https://example.com/path/to/image"
						return &text
					}(),
				})
			},
		},
		{
			name: "may add new recipeGroup, recipeType and glassType",
			fields: fields{
				repository: new(MockRepository),
			},
			before: func(f fields) {
				f.repository.On("FindOne", "NewRecipeGroup").Return(RecipeGroup{}, usecaseutils.ErrNotFound)
				f.repository.On("FindRecipeType").Return(map[string]model.RecipeType{}, nil)
				f.repository.On("FindGlassType").Return(map[string]model.GlassType{}, nil)
				f.repository.On("SaveRecipeType", mock.Anything).Return(nil)
				f.repository.On("SaveGlassType", mock.Anything).Return(nil)
				f.repository.On("Save", mock.Anything).Return(nil)
			},
			args: args{
				input: model.InputRecipeGroup{
					Name: "NewRecipeGroup",
					Recipes: []*model.InputRecipe{{
						Name: "NewRecipe",
						RecipeType: &model.InputRecipeType{
							Name: "NewRecieType",
						},
						GlassType: &model.InputGlassType{
							Name: "NewGlassType",
						},
						Steps: []string{"Step 1", "Step 2"},
					}},
				},
			},
			want: RecipeGroup{
				Name: "NewRecipeGroup",
				Recipes: []Recipe{{
					Name:  "NewRecipe",
					Type:  "NewRecieType",
					Glass: "NewGlassType",
					Steps: []string{"Step 1", "Step 2"},
				}},
			},
			wantErr: false,
			after: func(t *testing.T, f fields) {
				f.repository.AssertCalled(t, "Save", RecipeGroup{
					Name: "NewRecipeGroup",
					Recipes: []Recipe{{
						Name:  "NewRecipe",
						Type:  "NewRecieType",
						Glass: "NewGlassType",
						Steps: []string{"Step 1", "Step 2"},
					}},
				})
				f.repository.AssertCalled(t, "SaveRecipeType", model.RecipeType{Name: "NewRecieType"})
				f.repository.AssertCalled(t, "SaveGlassType", model.GlassType{Name: "NewGlassType"})
			},
		},
		{
			name: "may add new recipeGroup specify exists recipeType and glassType",
			fields: fields{
				repository: new(MockRepository),
			},
			before: func(f fields) {
				f.repository.On("FindOne", "NewRecipeGroup").Return(RecipeGroup{}, usecaseutils.ErrNotFound)
				f.repository.On("FindRecipeType").Return(map[string]model.RecipeType{"ExistsRecieType": {Name: "ExistsRecieType"}}, nil)
				f.repository.On("FindGlassType").Return(map[string]model.GlassType{"ExistsGlassType": {Name: "ExistsGlassType"}}, nil)
				f.repository.On("SaveRecipeType", mock.Anything).Return(nil)
				f.repository.On("SaveGlassType", mock.Anything).Return(nil)
				f.repository.On("Save", mock.Anything).Return(nil)
			},
			args: args{
				input: model.InputRecipeGroup{
					Name: "NewRecipeGroup",
					Recipes: []*model.InputRecipe{{
						Name: "NewRecipe",
						RecipeType: &model.InputRecipeType{
							Name: "ExistsRecieType",
						},
						GlassType: &model.InputGlassType{
							Name: "ExistsGlassType",
						},
						Steps: []string{"Step 1", "Step 2"},
					}},
				},
			},
			want: RecipeGroup{
				Name: "NewRecipeGroup",
				Recipes: []Recipe{{
					Name:  "NewRecipe",
					Type:  "ExistsRecieType",
					Glass: "ExistsGlassType",
					Steps: []string{"Step 1", "Step 2"},
				}},
			},
			wantErr: false,
			after: func(t *testing.T, f fields) {
				f.repository.AssertCalled(t, "Save", RecipeGroup{
					Name: "NewRecipeGroup",
					Recipes: []Recipe{{
						Name:  "NewRecipe",
						Type:  "ExistsRecieType",
						Glass: "ExistsGlassType",
						Steps: []string{"Step 1", "Step 2"},
					}},
				})
				f.repository.AssertNotCalled(t, "SaveRecipeType")
				f.repository.AssertNotCalled(t, "SaveGlassType")
			},
		},
		{
			name: "may update recipeGroup, recipeType and glassType",
			fields: fields{
				repository: new(MockRepository),
			},
			before: func(f fields) {
				f.repository.On("FindOne", "ExistsRecipeGroup").Return(RecipeGroup{
					Name:    "ExistsRecipeGroup",
					Recipes: []Recipe{},
				}, nil)
				f.repository.On("FindRecipeType").Return(map[string]model.RecipeType{"ExistsRecieType": {Name: "ExistsRecieType"}}, nil)
				f.repository.On("FindGlassType").Return(map[string]model.GlassType{"ExistsGlassType": {Name: "ExistsGlassType"}}, nil)
				f.repository.On("SaveRecipeType", mock.Anything).Return(nil)
				f.repository.On("SaveGlassType", mock.Anything).Return(nil)
				f.repository.On("Save", mock.Anything).Return(nil)
			},
			args: args{
				input: model.InputRecipeGroup{
					Name: "ExistsRecipeGroup",
					Recipes: []*model.InputRecipe{{
						Name: "ExistsRecipe",
						RecipeType: &model.InputRecipeType{
							Name: "ExistsRecieType",
							Description: func() *string {
								text := "description"
								return &text
							}(),
							Save: func() *bool {
								b := true
								return &b
							}(),
						},
						GlassType: &model.InputGlassType{
							Name: "ExistsGlassType",
							Description: func() *string {
								text := "description"
								return &text
							}(),
							Save: func() *bool {
								b := true
								return &b
							}(),
						},
						Steps: []string{"Step 1", "Step 2"},
					}},
				},
			},
			want: RecipeGroup{
				Name: "ExistsRecipeGroup",
				Recipes: []Recipe{{
					Name:  "ExistsRecipe",
					Type:  "ExistsRecieType",
					Glass: "ExistsGlassType",
					Steps: []string{"Step 1", "Step 2"},
				}},
			},
			wantErr: false,
			after: func(t *testing.T, f fields) {
				f.repository.AssertCalled(t, "Save", RecipeGroup{
					Name: "ExistsRecipeGroup",
					Recipes: []Recipe{{
						Name:  "ExistsRecipe",
						Type:  "ExistsRecieType",
						Glass: "ExistsGlassType",
						Steps: []string{"Step 1", "Step 2"},
					}},
				})
				f.repository.AssertCalled(t, "SaveRecipeType", model.RecipeType{
					Name: "ExistsRecieType",
					Description: func() *string {
						text := "description"
						return &text
					}(),
				})
				f.repository.AssertCalled(t, "SaveGlassType", model.GlassType{
					Name: "ExistsGlassType",
					Description: func() *string {
						text := "description"
						return &text
					}(),
				})
			},
		},
		{
			name: "may return error on updating exists recipeType without true save flag",
			fields: fields{
				repository: new(MockRepository),
			},
			before: func(f fields) {
				f.repository.On("FindOne", "ExistsRecipeGroup").Return(RecipeGroup{
					Name: "ExistsRecipeGroup",
					Recipes: []Recipe{{
						Name:  "ExistsRecipe",
						Type:  "ExistsRecieType",
						Glass: "ExistsGlassType",
						Steps: []string{"Step 1", "Step 2"},
					}},
				}, nil)
				f.repository.On("FindRecipeType").Return(map[string]model.RecipeType{"ExistsRecieType": {Name: "ExistsRecieType"}}, nil)
				f.repository.On("FindGlassType").Return(map[string]model.GlassType{"ExistsGlassType": {Name: "ExistsGlassType"}}, nil)
				f.repository.On("SaveRecipeType", mock.Anything).Return(nil)
				f.repository.On("SaveGlassType", mock.Anything).Return(nil)
				f.repository.On("Save", mock.Anything).Return(nil)
			},
			args: args{
				input: model.InputRecipeGroup{
					Name: "ExistsRecipeGroup",
					Recipes: []*model.InputRecipe{{
						Name: "ExistsRecipe",
						RecipeType: &model.InputRecipeType{
							Name: "ExistsRecieType",
							Description: func() *string {
								text := "description"
								return &text
							}(),
							Save: func() *bool {
								b := false
								return &b
							}(),
						},
						GlassType: &model.InputGlassType{Name: "ExistsGlassType"},
					}},
				},
			},
			want:    RecipeGroup{},
			wantErr: true,
		},
		{
			name: "may return error on updating exists recipeType without true save flag",
			fields: fields{
				repository: new(MockRepository),
			},
			before: func(f fields) {
				f.repository.On("FindOne", "ExistsRecipeGroup").Return(RecipeGroup{
					Name: "ExistsRecipeGroup",
					Recipes: []Recipe{{
						Name:  "ExistsRecipe",
						Type:  "ExistsRecieType",
						Glass: "ExistsGlassType",
						Steps: []string{"Step 1", "Step 2"},
					}},
				}, nil)
				f.repository.On("FindRecipeType").Return(map[string]model.RecipeType{"ExistsRecieType": {Name: "ExistsRecieType"}}, nil)
				f.repository.On("FindGlassType").Return(map[string]model.GlassType{"ExistsGlassType": {Name: "ExistsGlassType"}}, nil)
				f.repository.On("SaveRecipeType", mock.Anything).Return(nil)
				f.repository.On("SaveGlassType", mock.Anything).Return(nil)
				f.repository.On("Save", mock.Anything).Return(nil)
			},
			args: args{
				input: model.InputRecipeGroup{
					Name: "ExistsRecipeGroup",
					Recipes: []*model.InputRecipe{{
						Name:       "ExistsRecipe",
						RecipeType: &model.InputRecipeType{Name: "ExistsRecieType"},
						GlassType: &model.InputGlassType{Name: "ExistsGlassType",
							Description: func() *string {
								text := "description"
								return &text
							}(),
							Save: func() *bool {
								b := false
								return &b
							}()},
					}},
				},
			},
			want:    RecipeGroup{},
			wantErr: true,
		},
		{
			name: "may return error, if got unexpected error from repository",
			fields: fields{
				repository: new(MockRepository),
			},
			before: func(f fields) {
				f.repository.On("FindOne", "test").Return(RecipeGroup{}, errors.New("unexpected error"))
			},
			args: args{
				input: model.InputRecipeGroup{Name: "test"},
			},
			want:    RecipeGroup{},
			wantErr: true,
		},
		{
			name: "may return error, if got unexpected error from repository",
			fields: fields{
				repository: new(MockRepository),
			},
			before: func(f fields) {
				f.repository.On("FindOne", mock.Anything).Return(RecipeGroup{}, nil)
				f.repository.On("Save", mock.Anything).Return(errors.New("unexpected error"))
			},
			args: args{
				input: model.InputRecipeGroup{},
			},
			want:    RecipeGroup{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &service{tt.fields.repository}

			if tt.before != nil {
				tt.before(tt.fields)
			}

			got, err := s.Register(tt.args.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
			if tt.after != nil {
				tt.after(t, tt.fields)
			}
		})
	}
}
