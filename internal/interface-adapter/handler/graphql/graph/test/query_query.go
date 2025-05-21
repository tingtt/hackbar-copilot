package test

const QueryGetUserInfo = `
	query getUserInfo {
		userInfo {
			email
			name
			nameConfirmed
		}
	}
`

const QueryGetRecipes = `
	query getRecipes {
		recipes  {
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
	}
`

const QueryGetMenu = `
	query getMenu {
		menu {
			name
			imageURL
			flavor
			options {
				name
				category
				imageURL
				materials
				outOfStock
				priceYen
				recipe {
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
			minPriceYen
		}
	}
`

const QueryGetUncheckedOrdersCustomer = `
	query getUncheckedOrdersCustomer {
		uncheckedOrdersCustomer  {
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

const QueryGetUncheckedOrders = `
	query getUncheckedOrders {
		uncheckedOrders  {
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

const QueryGetUncashedCheckouts = `
	query getUncashedCheckouts {
		uncashedCheckouts  {
			id
			customerEmail
			orders {
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
			diffs {
				price
				description
			}
			totalPrice
			paymentType
			timestamp
		}
	}
`

const QueryGetCashouts = `
	query getCashouts ($input: InputCashoutQuery!) {
		cashouts (
			input: $input
		) {
			checkouts {
				id
				customerEmail
				orders {
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
				diffs {
					price
					description
				}
				totalPrice
				paymentType
				timestamp
			}
			revenue
			timestamp
			staffID
		}
	}
`
