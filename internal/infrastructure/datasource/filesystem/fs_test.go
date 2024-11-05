package filesystem

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		fs, err := NewRepository(t.TempDir())

		assert.NoError(t, err)
		assert.NotNil(t, fs)
	})

	t.Run("may return error, if fail to create data files", func(t *testing.T) {
		t.Parallel()

		fs, err := NewRepository(path.Join(t.TempDir(), "path/to/not/writable"))

		assert.Error(t, err)
		assert.Nil(t, fs)
	})
}

func Test_filesystem_Find(t *testing.T) {
	t.Parallel()

	type fields struct {
		data data
	}
	tests := []struct {
		name    string
		fields  fields
		want    []recipes.RecipeGroup
		wantErr bool
	}{
		{
			name: "will return cached data",
			fields: fields{
				data: data{
					recipeGroups: []recipes.RecipeGroup{{
						Name:     "Phuket Sling",
						ImageURL: nil,
						Recipes: []recipes.Recipe{{
							Name:  "Cocktail",
							Type:  "build",
							Glass: "collins",
							Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
						}, {
							Name:  "Mocktail",
							Type:  "build",
							Glass: "collins",
							Steps: []string{"Peach syrup 20ml", "Blue curacao syrup 10ml", "Grapefruit juice 30ml", "Tonic water - Full up"},
						}},
					}},
					recipeTypes: nil,
					glassTypes:  nil,
				},
			},
			want: []recipes.RecipeGroup{{
				Name:     "Phuket Sling",
				ImageURL: nil,
				Recipes: []recipes.Recipe{{
					Name:  "Cocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
				}, {
					Name:  "Mocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []string{"Peach syrup 20ml", "Blue curacao syrup 10ml", "Grapefruit juice 30ml", "Tonic water - Full up"},
				}},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := &filesystem{
				data: tt.fields.data,
			}

			got, err := f.Find()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_filesystem_FindOne(t *testing.T) {
	t.Parallel()

	type fields struct {
		data data
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    recipes.RecipeGroup
		wantErr error
	}{
		{
			name: "may return match data from cache",
			fields: fields{
				data: data{
					recipeGroups: []recipes.RecipeGroup{{
						Name:     "Phuket Sling",
						ImageURL: nil,
						Recipes: []recipes.Recipe{{
							Name:  "Cocktail",
							Type:  "build",
							Glass: "collins",
							Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
						}, {
							Name:  "Mocktail",
							Type:  "build",
							Glass: "collins",
							Steps: []string{"Peach syrup 20ml", "Blue curacao syrup 10ml", "Grapefruit juice 30ml", "Tonic water - Full up"},
						}},
					}},
					recipeTypes: nil,
					glassTypes:  nil,
				},
			},
			args: args{
				name: "Phuket Sling",
			},
			want: recipes.RecipeGroup{
				Name:     "Phuket Sling",
				ImageURL: nil,
				Recipes: []recipes.Recipe{{
					Name:  "Cocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
				}, {
					Name:  "Mocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []string{"Peach syrup 20ml", "Blue curacao syrup 10ml", "Grapefruit juice 30ml", "Tonic water - Full up"},
				}},
			},
			wantErr: nil,
		},
		{
			name: "may return error, if recipeGroup not found",
			fields: fields{
				data: data{
					recipeGroups: []recipes.RecipeGroup{},
					recipeTypes:  nil,
					glassTypes:   nil,
				},
			},
			args: args{
				name: "Phuket Sling",
			},
			want:    recipes.RecipeGroup{},
			wantErr: usecaseutils.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := &filesystem{
				read:  nil,
				write: nil,
				data:  tt.fields.data,
			}

			got, err := f.FindOne(tt.args.name)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_filesystem_Save(t *testing.T) {
	t.Parallel()

	t.Run("may add new recipeGroup to cache", func(t *testing.T) {
		t.Parallel()

		f := &filesystem{
			read:  nil,
			write: nil,
			data: data{
				recipeGroups: []recipes.RecipeGroup{},
				recipeTypes:  map[string]model.RecipeType{},
				glassTypes:   map[string]model.GlassType{},
			},
		}
		newRecipeGroup := recipes.RecipeGroup{
			Name:     "Phuket Sling",
			ImageURL: nil,
			Recipes: []recipes.Recipe{{
				Name:  "Cocktail",
				Type:  "build",
				Glass: "collins",
				Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
			}, {
				Name:  "Mocktail",
				Type:  "build",
				Glass: "collins",
				Steps: []string{"Peach syrup 20ml", "Blue curacao syrup 10ml", "Grapefruit juice 30ml", "Tonic water - Full up"},
			}},
		}
		err := f.Save(newRecipeGroup)

		assert.Nil(t, err)
		assert.Equal(t, []recipes.RecipeGroup{newRecipeGroup}, f.data.recipeGroups)
	})

	t.Run("may replace exists recipeGroup in cache", func(t *testing.T) {
		t.Parallel()

		existsRecipeGroup := recipes.RecipeGroup{
			Name:     "Phuket Sling",
			ImageURL: nil,
			Recipes: []recipes.Recipe{{
				Name:  "Cocktail",
				Type:  "build",
				Glass: "collins",
				Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
			}},
		}
		f := &filesystem{
			read:  nil,
			write: nil,
			data: data{
				recipeGroups: []recipes.RecipeGroup{existsRecipeGroup},
				recipeTypes:  map[string]model.RecipeType{},
				glassTypes:   map[string]model.GlassType{},
			},
		}

		replaceRecipeGroup := existsRecipeGroup
		replaceRecipeGroup.Recipes = append(existsRecipeGroup.Recipes, recipes.Recipe{
			Name:  "Mocktail",
			Type:  "build",
			Glass: "collins",
			Steps: []string{"Peach syrup 20ml", "Blue curacao syrup 10ml", "Grapefruit juice 30ml", "Tonic water - Full up"},
		})
		err := f.Save(replaceRecipeGroup)

		assert.Nil(t, err)
		assert.Equal(t, []recipes.RecipeGroup{replaceRecipeGroup}, f.data.recipeGroups)
	})
}

func Test_filesystem_FindRecipeType(t *testing.T) {
	t.Parallel()

	type fields struct {
		data data
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]model.RecipeType
		wantErr bool
	}{
		{
			name: "will return cached data",
			fields: fields{
				data: data{
					recipeGroups: []recipes.RecipeGroup{},
					recipeTypes: map[string]model.RecipeType{
						"build": {
							Name:        "build",
							Description: nil,
						},
						"shake": {
							Name:        "shake",
							Description: nil,
						},
						"stir": {
							Name:        "stir",
							Description: nil,
						},
						"blend": {
							Name:        "blend",
							Description: nil,
						},
					},
					glassTypes: nil,
				},
			},
			want: map[string]model.RecipeType{
				"build": {
					Name:        "build",
					Description: nil,
				},
				"shake": {
					Name:        "shake",
					Description: nil,
				},
				"stir": {
					Name:        "stir",
					Description: nil,
				},
				"blend": {
					Name:        "blend",
					Description: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := &filesystem{
				data: tt.fields.data,
			}

			got, err := f.FindRecipeType()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_filesystem_FindGlassType(t *testing.T) {
	t.Parallel()

	type fields struct {
		data data
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]model.GlassType
		wantErr bool
	}{
		{
			name: "will return cached data",
			fields: fields{
				data: data{
					recipeGroups: []recipes.RecipeGroup{},
					recipeTypes:  map[string]model.RecipeType{},
					glassTypes: map[string]model.GlassType{
						"collins": {
							Name:        "collins",
							Description: nil,
						},
						"shot": {
							Name:        "shot",
							Description: nil,
						},
						"rock": {
							Name:        "rock",
							Description: nil,
						},
						"beer": {
							Name:        "beer",
							Description: nil,
						},
					},
				},
			},
			want: map[string]model.GlassType{
				"collins": {
					Name:        "collins",
					Description: nil,
				},
				"shot": {
					Name:        "shot",
					Description: nil,
				},
				"rock": {
					Name:        "rock",
					Description: nil,
				},
				"beer": {
					Name:        "beer",
					Description: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := &filesystem{
				data: tt.fields.data,
			}

			got, err := f.FindGlassType()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_filesystem_SaveRecipeType(t *testing.T) {
	t.Parallel()

	t.Run("may add new recipeType to cache", func(t *testing.T) {
		t.Parallel()

		f := &filesystem{
			read:  nil,
			write: nil,
			data: data{
				recipeGroups: []recipes.RecipeGroup{},
				recipeTypes:  map[string]model.RecipeType{},
				glassTypes:   map[string]model.GlassType{},
			},
		}
		newRecipeType := model.RecipeType{
			Name:        "build",
			Description: nil,
		}
		err := f.SaveRecipeType(newRecipeType)

		assert.Nil(t, err)
		assert.Equal(t, map[string]model.RecipeType{
			"build": {
				Name:        "build",
				Description: nil,
			},
		}, f.data.recipeTypes)
	})

	t.Run("may replace exists recipeType in cache", func(t *testing.T) {
		t.Parallel()

		existsRecipeType := model.RecipeType{
			Name:        "build",
			Description: nil,
		}
		f := &filesystem{
			read:  nil,
			write: nil,
			data: data{
				recipeGroups: []recipes.RecipeGroup{},
				recipeTypes:  map[string]model.RecipeType{"build": existsRecipeType},
				glassTypes:   map[string]model.GlassType{},
			},
		}

		replaceRecipeType := existsRecipeType
		replaceRecipeType.Description = func() *string {
			text := "recipe type description..."
			return &text
		}()
		err := f.SaveRecipeType(replaceRecipeType)

		assert.Nil(t, err)
		assert.Equal(t, map[string]model.RecipeType{"build": replaceRecipeType}, f.data.recipeTypes)
	})
}

func Test_filesystem_SaveGlassType(t *testing.T) {
	t.Parallel()

	t.Run("may add new glassType to cache", func(t *testing.T) {
		t.Parallel()

		f := &filesystem{
			read:  nil,
			write: nil,
			data: data{
				recipeGroups: []recipes.RecipeGroup{},
				recipeTypes:  map[string]model.RecipeType{},
				glassTypes:   map[string]model.GlassType{},
			},
		}
		newGlassType := model.GlassType{
			Name:        "collins",
			Description: nil,
		}
		err := f.SaveGlassType(newGlassType)

		assert.Nil(t, err)
		assert.Equal(t, map[string]model.GlassType{
			"collins": {
				Name:        "collins",
				Description: nil,
			},
		}, f.data.glassTypes)
	})

	t.Run("may replace exists glassType in cache", func(t *testing.T) {
		t.Parallel()

		existsGlassType := model.GlassType{
			Name:        "collins",
			Description: nil,
		}
		f := &filesystem{
			read:  nil,
			write: nil,
			data: data{
				recipeGroups: []recipes.RecipeGroup{},
				recipeTypes:  map[string]model.RecipeType{},
				glassTypes:   map[string]model.GlassType{"collins": existsGlassType},
			},
		}

		replaceGlassType := existsGlassType
		replaceGlassType.Description = func() *string {
			text := "recipe type description..."
			return &text
		}()
		err := f.SaveGlassType(replaceGlassType)

		assert.Nil(t, err)
		assert.Equal(t, map[string]model.GlassType{"collins": replaceGlassType}, f.data.glassTypes)
	})
}
