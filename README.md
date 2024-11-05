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
      steps
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
      steps
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
          "Peach liqueur 30ml",
          "Blue curacao 15ml",
          "Grapefruit juice 30ml",
          "Stir",
          "Tonic water - Full up"
        ]
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
          "Peach syrup 20ml",
          "Blue curacao syrup 10ml",
          "Grapefruit juice 30ml",
          "Stir",
          "Tonic water - Full up",
          "Stir"
        ]
      }
    ]
  }
}
```
