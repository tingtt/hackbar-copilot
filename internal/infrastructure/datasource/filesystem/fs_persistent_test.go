package filesystem

import (
	"bytes"
	"errors"
	"hackbar-copilot/internal/infrastructure/datasource/filesystem/toml"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"
	"io"
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_loadData(t *testing.T) {
	t.Parallel()

	type args struct {
		fs fsR
	}
	tests := []struct {
		name    string
		args    args
		wantD   data
		wantErr bool
	}{
		{
			name: "may return empty data",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", mock.Anything).Return(&MockFile{&bytes.Buffer{}}, nil)
					return m
				}(),
			},
			wantD: data{
				recipeGroups: []recipes.RecipeGroup{},
				recipeTypes:  map[string]model.RecipeType{},
				glassTypes:   map[string]model.GlassType{},
			},
			wantErr: false,
		},
		{
			name: "may return error, if not expected",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", mock.Anything).Return(&MockFile{}, errors.New("error unexpected")).Once()
					return m
				}(),
			},
			wantD: data{
				recipeGroups: nil,
				recipeTypes:  nil,
				glassTypes:   nil,
			},
			wantErr: true,
		},
		{
			name: "may return error, if not expected",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", mock.Anything).Return(&MockFile{&bytes.Buffer{}}, nil).Once()
					m.On("Open", mock.Anything).Return(&MockFile{}, errors.New("error unexpected")).Once()
					return m
				}(),
			},
			wantD: data{
				recipeGroups: nil,
				recipeTypes:  nil,
				glassTypes:   nil,
			},
			wantErr: true,
		},
		{
			name: "may return error, if not expected",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", mock.Anything).Return(&MockFile{&bytes.Buffer{}}, nil).Twice()
					m.On("Open", mock.Anything).Return(&MockFile{}, errors.New("error unexpected")).Once()
					return m
				}(),
			},
			wantD: data{
				recipeGroups: nil,
				recipeTypes:  nil,
				glassTypes:   nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotD, err := loadData(tt.args.fs)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, gotD, tt.wantD)
		})
	}
}

func Test_loadRecipeGroups(t *testing.T) {
	t.Parallel()

	type args struct {
		fs fsR
	}
	tests := []struct {
		name    string
		args    args
		want    []recipes.RecipeGroup
		wantErr bool
	}{
		{
			name: "may load data successfully",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					r := bytes.NewBufferString(`[[recipe_group]]
Name = "Phuket Sling"

[[recipe_group.Recipes]]
Name = "Cocktail"
Type = "build"
Glass = "collins"
Steps = ["Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"]

[[recipe_group.Recipes]]
Name = "Mocktail"
Type = "build"
Glass = "collins"
Steps = ["Peach syrup 20ml", "Blue curacao syrup 10ml", "Grapefruit juice 30ml", "Tonic water - Full up"]
`)
					m.On("Open", "1_recipe_groups.toml").Return(&MockFile{r}, nil)
					return m
				}(),
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
		{
			name: "may return empty data, if file not exists",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", "1_recipe_groups.toml").Return(&MockFile{}, os.ErrNotExist)
					return m
				}(),
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "may return error, if unexpected",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", "1_recipe_groups.toml").Return(&MockFile{}, errors.New("unexpected error"))
					return m
				}(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "may return error, if unexpected",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					r := bytes.NewBufferString("invalid toml format")
					m.On("Open", "1_recipe_groups.toml").Return(&MockFile{r}, nil)
					return m
				}(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := loadRecipeGroups(tt.args.fs)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_loadRecipeTypes(t *testing.T) {
	t.Parallel()

	type args struct {
		fs fsR
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]model.RecipeType
		wantErr bool
	}{
		{
			name: "may load data successfully",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					r := bytes.NewBufferString(`[recipe_type]
[recipe_type.build]
Name = "build"
[recipe_type.shake]
Name = "shake"
[recipe_type.stir]
Name = "stir"
[recipe_type.blend]
Name = "blend"
`)
					m.On("Open", "2_recipe_types.toml").Return(&MockFile{r}, nil)
					return m
				}(),
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
		{
			name: "may return empty data, if file not exists",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", "2_recipe_types.toml").Return(&MockFile{}, os.ErrNotExist)
					return m
				}(),
			},
			want:    map[string]model.RecipeType{},
			wantErr: false,
		},
		{
			name: "may return error, if unexpected",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", "2_recipe_types.toml").Return(&MockFile{}, errors.New("unexpected error"))
					return m
				}(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "may return error, if unexpected",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					r := bytes.NewBufferString("invalid toml format")
					m.On("Open", "2_recipe_types.toml").Return(&MockFile{r}, nil)
					return m
				}(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := loadRecipeTypes(tt.args.fs)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_loadGlassTypes(t *testing.T) {
	t.Parallel()

	type args struct {
		fs fsR
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]model.GlassType
		wantErr bool
	}{
		{
			name: "may load data successfully",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					r := bytes.NewBufferString(`[glass_type]
[glass_type.collins]
Name = "collins"
[glass_type.shot]
Name = "shot"
[glass_type.rock]
Name = "rock"
[glass_type.beer]
Name = "beer"
`)
					m.On("Open", "3_glass_types.toml").Return(&MockFile{r}, nil)
					return m
				}(),
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
		{
			name: "may return empty data, if file not exists",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", "3_glass_types.toml").Return(&MockFile{}, os.ErrNotExist)
					return m
				}(),
			},
			want:    map[string]model.GlassType{},
			wantErr: false,
		},
		{
			name: "may return error, if unexpected",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", "3_glass_types.toml").Return(&MockFile{}, errors.New("unexpected error"))
					return m
				}(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "may return error, if unexpected",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					r := bytes.NewBufferString("invalid toml format")
					m.On("Open", "3_glass_types.toml").Return(&MockFile{r}, nil)
					return m
				}(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := loadGlassTypes(tt.args.fs)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

type MockFSW struct {
	mock.Mock
}

func (m *MockFSW) Create(name string) (io.WriteCloser, error) {
	args := m.Called(name)
	return args.Get(0).(io.WriteCloser), args.Error(1)
}

func Test_filesystem_SavePersistently(t *testing.T) {
	t.Parallel()

	t.Run("may write files successfully", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			data    data
			wantErr bool
		}{
			{
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
				wantErr: false,
			},
		}
		for i, tt := range tests {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				t.Parallel()

				ioWriters := struct {
					recipeGroups *MockFile
					recipeTypes  *MockFile
					glassTypes   *MockFile
				}{
					recipeGroups: &MockFile{&bytes.Buffer{}},
					recipeTypes:  &MockFile{&bytes.Buffer{}},
					glassTypes:   &MockFile{&bytes.Buffer{}},
				}
				m := new(MockFSW)
				m.On("Create", "1_recipe_groups.toml").Return(ioWriters.recipeGroups, nil)
				m.On("Create", "2_recipe_types.toml").Return(ioWriters.recipeTypes, nil)
				m.On("Create", "3_glass_types.toml").Return(ioWriters.glassTypes, nil)
				f := &filesystem{
					write: m,
					data:  tt.data,
				}

				err := f.SavePersistently()

				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
				{ // assert written data
					data := struct {
						recipeGroups map[string][]recipes.RecipeGroup
						recipeTypes  map[string]map[string]model.RecipeType
						glassTypes   map[string]map[string]model.GlassType
					}{
						recipeGroups: map[string][]recipes.RecipeGroup{},
						recipeTypes:  map[string]map[string]model.RecipeType{},
						glassTypes:   map[string]map[string]model.GlassType{},
					}
					assert.NoError(t, toml.Decode(ioWriters.recipeGroups, &data.recipeGroups))
					assert.NoError(t, toml.Decode(ioWriters.recipeTypes, &data.recipeTypes))
					assert.NoError(t, toml.Decode(ioWriters.glassTypes, &data.glassTypes))
					assert.Equal(t, tt.data.recipeGroups, data.recipeGroups["recipe_group"])
					assert.Equal(t, tt.data.recipeTypes, data.recipeTypes["recipe_type"])
					assert.Equal(t, tt.data.glassTypes, data.glassTypes["glass_type"])
				}
			})
		}
	})

	t.Run("may fail to write files", func(t *testing.T) {
		t.Parallel()
		// TODO: add test
	})
}

func Test_filesystem_saveFile(t *testing.T) {
	t.Parallel()

	type args struct {
		filename string
		data     any
	}
	tests := []struct {
		name    string
		fsW     fsW
		args    args
		wantErr bool
	}{
		{
			name: "may return error, if unexpected",
			fsW: func() fsW {
				m := new(MockFSW)
				m.On("Create", mock.Anything).Return(&MockFile{}, errors.New("unexpected error"))
				return m
			}(),
			args: args{
				filename: path.Join(t.TempDir(), "1"),
				data:     nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := &filesystem{
				read:  nil,
				write: tt.fsW,
				data:  data{},
			}

			err := f.saveFile(tt.args.filename, tt.args.data)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
