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

const QueryGetCheckouts = `
	query getCheckouts {
		checkouts  {
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
