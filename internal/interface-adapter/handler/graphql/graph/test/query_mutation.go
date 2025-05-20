package test

const QuerySaveRecipe = `
	mutation saveRecipe ($input: InputRecipeGroup!) {
		saveRecipe (
			input: $input
		) {
			__typename
			... on RecipeGroup {
				name
				imageURL
				recipes {
					name
					category
					type {
						name
						description
					}
					glass {
						name
						imageURL
						description
					}
					steps {
						material
						amount
						description
					}
				}
			}
			... on RemovedRecipeGroup {
				name
			}
		}
	}
`

const QueryOrder = `
	mutation order ($input: InputOrder!) {
		order (
			input: $input
		) {
			__typename
			id
			customerEmail
			customerName
			menuID {
				itemName
				optionName
			}
			timestamps {
				status
				timestamp
			}
			status
			price
		}
	}
`
