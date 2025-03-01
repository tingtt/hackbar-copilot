// Code generated by script, DO NOT EDIT.

import * as types from "./types"

export interface MutationClient {
  order(input: types.InputOrder): Promise<types.Order | undefined>
  updateOrderStatus(
    input: types.InputOrderStatusUpdate,
  ): Promise<types.Order | undefined>
  saveRecipe(input: types.InputRecipeGroup): Promise<types.RecipeGroup>
  updateStock(input: types.InputStockUpdate): Promise<types.Material[]>
}
