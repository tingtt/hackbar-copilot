package filesystem

import (
	"bytes"
	"errors"
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/menu/menutest"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"hackbar-copilot/internal/domain/stock/stocktest"
	"hackbar-copilot/internal/infrastructure/datasource/filesystem/toml"
	"io"
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/lithammer/dedent"
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
				recipeGroups:         nil,
				recipeTypes:          nil,
				glassTypes:           nil,
				menuItems:            nil,
				stocks:               nil,
				uncheckedOrders:      nil,
				uncashedoutCheckouts: nil,
			},
			wantErr: false,
		},
		{
			name: "may return empty data/no file",
			args: args{
				fs: func() fsR {
					m := new(MockFSR)
					m.On("Open", mock.Anything).Return(&MockFile{&bytes.Buffer{}}, os.ErrNotExist)
					return m
				}(),
			},
			wantD: data{
				recipeGroups:         nil,
				recipeTypes:          nil,
				glassTypes:           nil,
				menuItems:            nil,
				stocks:               nil,
				uncheckedOrders:      nil,
				uncashedoutCheckouts: nil,
			},
			wantErr: false,
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

type MockFSW struct {
	mock.Mock
}

func (m *MockFSW) Create(name string) (io.WriteCloser, error) {
	args := m.Called(name)
	return args.Get(0).(io.WriteCloser), args.Error(1)
}

func ptr[T any](v T) *T {
	return &v
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
					recipeGroups: []recipe.RecipeGroup{{
						Name:     "Phuket Sling",
						ImageURL: nil,
						Recipes: []recipe.Recipe{{
							Name:  "Cocktail",
							Type:  "build",
							Glass: "collins",
							Steps: []recipe.Step{
								{
									Material: ptr("Peach liqueur"),
									Amount:   ptr("30ml"),
								},
								{
									Material: ptr("Blue curacao"),
									Amount:   ptr("15ml"),
								},
								{
									Material: ptr("Grapefruit juice"),
									Amount:   ptr("30ml"),
								},
								{
									Description: ptr("Stir"),
								},
								{
									Material: ptr("Tonic water"),
									Amount:   ptr("Full up"),
								},
							},
						}, {
							Name:  "Mocktail",
							Type:  "build",
							Glass: "collins",
							Steps: []recipe.Step{
								{
									Material: ptr("Peach syrup"),
									Amount:   ptr("20ml"),
								},
								{
									Material: ptr("Blue curacao syrup"),
									Amount:   ptr("15ml"),
								},
								{
									Material: ptr("Grapefruit juice"),
									Amount:   ptr("30ml"),
								},
								{
									Description: ptr("Stir"),
								},
								{
									Material: ptr("Tonic water"),
									Amount:   ptr("Full up"),
								},
							},
						}},
					}},
					recipeTypes: map[string]recipe.RecipeType{
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
					glassTypes: map[string]recipe.GlassType{
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
					menuItems: []menu.Item{
						{
							Name:     "Phuket Sling",
							ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
							Flavor:   ptr("Sweet"),
							Options: []menu.ItemOption{
								{
									Name:       "Cocktail",
									ImageURL:   ptr("https://example.com/path/to/image/cocktail"),
									Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
									OutOfStock: false,
									Price:      700,
								},
							},
						},
					},
					stocks: map[string]bool{
						"Peach liqueur":    true,
						"Blue curacao":     true,
						"Grapefruit juice": true,
						"Tonic water":      true,
					},
					uncheckedOrders: []order.Order{
						{
							ID:            "",
							CustomerEmail: "",
							MenuItemID:    order.MenuItemID{},
							Timestamps:    []order.StatusUpdateTimestamp{},
							Status:        "",
							Price:         0,
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
					users        *MockFile
					recipeGroups *MockFile
					recipeTypes  *MockFile
					glassTypes   *MockFile
					menuGroups   *MockFile
					stocks       *MockFile
					orders       *MockFile
					checkouts    *MockFile
				}{
					users:        &MockFile{&bytes.Buffer{}},
					recipeGroups: &MockFile{&bytes.Buffer{}},
					recipeTypes:  &MockFile{&bytes.Buffer{}},
					glassTypes:   &MockFile{&bytes.Buffer{}},
					menuGroups:   &MockFile{&bytes.Buffer{}},
					stocks:       &MockFile{&bytes.Buffer{}},
					orders:       &MockFile{&bytes.Buffer{}},
					checkouts:    &MockFile{&bytes.Buffer{}},
				}
				m := new(MockFSW)
				m.On("Create", "0_user.toml").Return(ioWriters.users, nil)
				m.On("Create", "1_recipe_groups.toml").Return(ioWriters.recipeGroups, nil)
				m.On("Create", "2_recipe_types.toml").Return(ioWriters.recipeTypes, nil)
				m.On("Create", "3_glass_types.toml").Return(ioWriters.glassTypes, nil)
				m.On("Create", "4_menu_items.toml").Return(ioWriters.menuGroups, nil)
				m.On("Create", "5_stocks.toml").Return(ioWriters.stocks, nil)
				m.On("Create", "6_orders.toml").Return(ioWriters.orders, nil)
				m.On("Create", "7_checkouts.toml").Return(ioWriters.checkouts, nil)
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
						recipeGroups map[string][]recipe.RecipeGroup
						recipeTypes  map[string]map[string]recipe.RecipeType
						glassTypes   map[string]map[string]recipe.GlassType
						menuGroups   map[string][]menu.Item
						stocks       map[string]map[string]bool
						orders       map[string][]order.Order
					}{
						recipeGroups: map[string][]recipe.RecipeGroup{},
						recipeTypes:  map[string]map[string]recipe.RecipeType{},
						glassTypes:   map[string]map[string]recipe.GlassType{},
						menuGroups:   map[string][]menu.Item{},
						stocks:       map[string]map[string]bool{},
						orders:       map[string][]order.Order{},
					}
					assert.NoError(t, toml.Decode(ioWriters.recipeGroups, &data.recipeGroups))
					assert.NoError(t, toml.Decode(ioWriters.recipeTypes, &data.recipeTypes))
					assert.NoError(t, toml.Decode(ioWriters.glassTypes, &data.glassTypes))
					assert.NoError(t, toml.Decode(ioWriters.menuGroups, &data.menuGroups))
					assert.NoError(t, toml.Decode(ioWriters.stocks, &data.stocks))
					assert.NoError(t, toml.Decode(ioWriters.orders, &data.orders))
					assert.Equal(t, tt.data.recipeGroups, data.recipeGroups["recipe_group"])
					assert.Equal(t, tt.data.recipeTypes, data.recipeTypes["recipe_type"])
					assert.Equal(t, tt.data.glassTypes, data.glassTypes["glass_type"])
					assert.Equal(t, tt.data.menuItems, data.menuGroups["menu_items"])
					assert.Equal(t, tt.data.stocks, data.stocks["stock"])
					assert.Equal(t, tt.data.uncheckedOrders, data.orders["order"])
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

type loadTest struct {
	name string
	key  string
	raw  string
	want any
}

var loadTests = []loadTest{
	{
		name: "may load recipe groups",
		key:  "recipe_group",
		raw: dedent.Dedent(`
			[[recipe_group]]
			Name = "Phuket Sling"
			ImageURL = "https://example.com/path/to/image/phuket-sling"

			[[recipe_group.Recipes]]
			Name = "Cocktail"
			Category = "Cocktail"
			Type = "build"
			Glass = "collins"

			[[recipe_group.Recipes.Steps]]
			Material = "Peach liqueur"
			Amount = "30ml"

			[[recipe_group.Recipes.Steps]]
			Material = "Blue curacao"
			Amount = "15ml"

			[[recipe_group.Recipes.Steps]]
			Material = "Grapefruit juice"
			Amount = "30ml"

			[[recipe_group.Recipes.Steps]]
			Description = "Stir"

			[[recipe_group.Recipes.Steps]]
			Material = "Tonic water"
			Amount = "Full up"

			[[recipe_group]]
			Name = "Passoamoni"
			ImageURL = "https://example.com/path/to/image/passoamoni"

			[[recipe_group.Recipes]]
			Name = "Cocktail"
			Category = "Cocktail"
			Type = "build"
			Glass = "collins"

			[[recipe_group.Recipes.Steps]]
			Material = "Passoa"
			Amount = "45ml"

			[[recipe_group.Recipes.Steps]]
			Material = "Grapefruit juice"
			Amount = "30ml"

			[[recipe_group.Recipes.Steps]]
			Description = "Stir"

			[[recipe_group.Recipes.Steps]]
			Material = "Tonic water"
			Amount = "Full up"

			[[recipe_group]]
			Name = "Blue Devil"
			ImageURL = "https://example.com/path/to/image/passoamoni"

			[[recipe_group.Recipes]]
			Name = "Cocktail"
			Category = "Cocktail"
			Type = "shake"
			Glass = "cocktail"

			[[recipe_group.Recipes.Steps]]
			Description = "Chill shaker and glass."

			[[recipe_group.Recipes.Steps]]
			Description = "Put ingredients in a shaker."

			[[recipe_group.Recipes.Steps]]
			Material = "Gin"
			Amount = "30ml"

			[[recipe_group.Recipes.Steps]]
			Material = "Blue curacao"
			Amount = "15ml"

			[[recipe_group.Recipes.Steps]]
			Material = "Lemon juice"
			Amount = "15ml"

			[[recipe_group.Recipes.Steps]]
			Description = "Put ice in a shaker."

			[[recipe_group.Recipes.Steps]]
			Description = "Shake."

			[[recipe_group.Recipes.Steps]]
			Description = "Pour into a glass."
		`),
		want: recipetest.ExampleRecipeGroups,
	},
	{
		name: "may load recipe types",
		key:  "recipe_type",
		raw: dedent.Dedent(`
			[recipe_type]
			[recipe_type.shake]
			Name = "shake"
			Description = "shake description"
			[recipe_type.build]
			Name = "build"
			Description = "build description"
			[recipe_type.stir]
			Name = "stir"
			Description = "stir description"
			[recipe_type.blend]
			Name = "blend"
			Description = "blend description"
		`),
		want: recipetest.ExampleRecipeTypesMap,
	},
	{
		name: "may load glass types",
		key:  "glass_type",
		raw: dedent.Dedent(`
			[glass_type]
			[glass_type.collins]
			Name = "collins"
			ImageURL = "https://example.com/path/to/image/collins"
			Description = "collins glass description"
			[glass_type.cocktail]
			Name = "cocktail"
			ImageURL = "https://example.com/path/to/image/cocktail"
			Description = "cocktail glass description"
			[glass_type.shot]
			Name = "shot"
			ImageURL = "https://example.com/path/to/image/shot"
			Description = "shot glass description"
			[glass_type.rock]
			Name = "rock"
			ImageURL = "https://example.com/path/to/image/rock"
			Description = "rock glass description"
			[glass_type.beer]
			Name = "beer"
			ImageURL = "https://example.com/path/to/image/beer"
			Description = "beer glass description"
		`),
		want: recipetest.ExampleGlassTypesMap,
	},
	{
		name: "may load menu groups",
		key:  "menu_items",
		raw: dedent.Dedent(`
			[[menu_items]]
			Name = "Phuket Sling"
			ImageURL = "https://example.com/path/to/image/phuket-sling"
			Flavor = "Sweet"

			[[menu_items.Options]]
			Name = "Cocktail"
			Category = "Cocktail"
			ImageURL = "https://example.com/path/to/image/phuket-sling/cocktail"
			Materials = ["Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"]
			OutOfStock = false
			Price = 700

			[[menu_items.Options]]
			Name = "Mocktail"
			Category = "Mocktail"
			ImageURL = "https://example.com/path/to/image/phuket-sling/mocktail"
			Materials = ["Peach syrup", "Blue curacao syrup", "Grapefruit juice", "Tonic water"]
			OutOfStock = false
			Price = 500

			[[menu_items]]
			Name = "Passoamoni"
			ImageURL = "https://example.com/path/to/image/passoamoni"
			Flavor = "Fruity"

			[[menu_items.Options]]
			Name = "Cocktail"
			Category = "Cocktail"
			ImageURL = "https://example.com/path/to/image/passoamoni"
			Materials = ["Passoa", "Grapefruit juice", "Tonic water"]
			OutOfStock = false
			Price = 700

			[[menu_items]]
			Name = "Blue Devil"
			ImageURL = "https://example.com/path/to/image/blue-devil"
			Flavor = "Medium sweet and dry"

			[[menu_items.Options]]
			Name = "Cocktail"
			Category = "Cocktail"
			ImageURL = "https://example.com/path/to/image/blue-devil"
			Materials = ["Gin", "Blue curacao", "Lemon juice"]
			OutOfStock = false
			Price = 700
		`),
		want: menutest.ExampleItems,
	},
	{
		name: "may load stocks",
		key:  "stock",
		raw: dedent.Dedent(`
			[stock]
			"Blue curacao" = true
      "Gin" = true
			"Grapefruit juice" = true
			"Lemon juice" = true
			"Passoa" = true
			"Peach liqueur" = true
			"Tonic water" = true
		`),
		want: stocktest.ExampleMaterialsMap,
	},
}

func Test_loadFromToml(t *testing.T) {
	t.Parallel()

	t.Run("may load data successfully", func(t *testing.T) {
		t.Parallel()

		for _, tt := range loadTests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				m := new(MockFSR)
				r := bytes.NewBufferString(tt.raw)
				m.On("Open", mock.Anything).Return(&MockFile{r}, nil)
				var got any
				var err error
				switch tt.key {
				case "recipe_group":
					typedGot := []recipe.RecipeGroup{}
					err = loadFromToml(m, "data.toml", tt.key, &typedGot)
					got = typedGot
				case "recipe_type":
					typedGot := map[string]recipe.RecipeType{}
					err = loadFromToml(m, "data.toml", tt.key, &typedGot)
					got = typedGot
				case "glass_type":
					typedGot := map[string]recipe.GlassType{}
					err = loadFromToml(m, "data.toml", tt.key, &typedGot)
					got = typedGot
				case "menu_items":
					typedGot := []menu.Item{}
					err = loadFromToml(m, "data.toml", tt.key, &typedGot)
					got = typedGot
				case "stock":
					typedGot := map[string]bool{}
					err = loadFromToml(m, "data.toml", tt.key, &typedGot)
					got = typedGot
				}

				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}
