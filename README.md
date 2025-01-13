# hackbar-copilot

[Design Doc](.docs/DesignDoc.md)

## Usage

```sh
mkdir .data
go run cmd/registry/main.go -d ./.data/
```

[http://localhost:8080/recipes.v1graphql.Registry/playground](http://localhost:8080/recipes.v1graphql.Registry/playground)

```graphql
query {
  recipes {
    name
    imageURL
    recipes {
      name
      type {
        name
        description
      }
      glass {
        name
        description
      }
      steps {
        material
        amount
        description
      }
    }
  }
}
```

```graphql
mutation saveRecipe($input: InputRecipeGroup!) {
  saveRecipe(input: $input) {
    name
    imageURL
    recipes {
      name
      type {
        name
        description
      }
      glass {
        name
        description
      }
      steps {
        material
        amount
        description
      }
    }
  }
}
```

```json
{
  "input": {
    "name": "Phuket Sling",
    "recipes": [
      {
        "name": "Cocktail",
        "recipeType": {
          "name": "build"
        },
        "glassType": {
          "name": "collins"
        },
        "steps": [
          {
            "material": "Peach liqueur",
            "amount": "30ml"
          },
          {
            "material": "Blue curacao",
            "amount": "15ml"
          },
          {
            "material": "Grapefruit juice",
            "amount": "30ml"
          },
          {
            "material": "Tonic water",
            "amount": "Half up"
          },
          {
            "description": "Stir"
          },
          {
            "material": "Tonic water",
            "amount": "Full up"
          },
          {
            "description": "Stir (Lightly so as not to lose gas of soda.)"
          }
        ],
        "asMenu": {
          "price": 700
        }
      },
      {
        "name": "Mocktail",
        "recipeType": {
          "name": "build"
        },
        "glassType": {
          "name": "collins"
        },
        "steps": [
          {
            "material": "Peach syrup",
            "amount": "20ml"
          },
          {
            "material": "Blue curacao syrup",
            "amount": "10ml"
          },
          {
            "material": "Grapefruit juice",
            "amount": "30ml"
          },
          {
            "material": "Tonic water",
            "amount": "Half up"
          },
          {
            "description": "Stir"
          },
          {
            "material": "Tonic water",
            "amount": "Full up"
          },
          {
            "description": "Stir (Lightly so as not to lose gas of soda.)"
          }
        ],
        "asMenu": {
          "price": 700
        }
      }
    ],
    "asMenu": {
      "flavor": "Sweet"
    }
  }
}
```
