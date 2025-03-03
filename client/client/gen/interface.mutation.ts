// Code generated by script, DO NOT EDIT.

import * as types from "./types"

export interface MutationClient {
  cashout(input: types.CashoutInput): Promise<types.Cashout>
  checkout(input: types.InputCheckout): Promise<types.Checkout>
  order(input: types.InputOrder): Promise<types.Order>
  updateOrderStatus(input: types.InputOrderStatusUpdate): Promise<types.Order>
  saveRecipe(input: types.InputRecipeGroup): Promise<types.RecipeGroup>
  updateStock(input: types.InputStockUpdate): Promise<types.Material[]>
}
