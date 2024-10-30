# hackbar-copilot

[Design Doc](.docs/DesignDoc.md)

## Usage

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
    "name": "Go",
    "recipes": [
      {
        "name": "カクテル",
        "recipeType": {
          "name": "build"
        },
        "glassType": {
          "name": "collins"
        },
        "steps": [
          "ピーチリキュール 20ml",
          "ブルーキュラソー 15ml",
          "グレープフルーツジュース 20ml",
          "ステア",
          "トニックウォーター full up"
        ]
      },
      {
        "name": "モクテル",
        "recipeType": {
          "name": "build"
        },
        "glassType": {
          "name": "collins"
        },
        "steps": [
          "ピーチシロップ 20ml",
          "ブルーキュラソーシロップ 15ml",
          "グレープフルーツジュース 20ml",
          "ステア",
          "トニックウォーター full up",
          "ステア"
        ]
      }
    ]
  }
}
```
