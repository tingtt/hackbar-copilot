import type { QueryClient } from "../gen/interface.client"
import type { MutationClient } from "../gen/interface.mutation"
import type {
  MenuGroup,
  Order,
  RecipeGroup,
  Material,
  InputOrder,
  InputOrderStatusUpdate,
  InputRecipeGroup,
  InputStockUpdate,
} from "../gen/types"

import dummyMenuData from "./data/menu.json"
import dummyOrdersData from "./data/orders.json"
import dummyRecipesData from "./data/recipes.json"
import dummyMaterialsData from "./data/materials.json"

export class DummyHackbarClient implements QueryClient, MutationClient {
  async menu(): Promise<MenuGroup[]> {
    return dummyMenuData as MenuGroup[]
  }
  async orders(): Promise<Order[]> {
    return dummyOrdersData as Order[]
  }
  async recipes(): Promise<RecipeGroup[]> {
    return dummyRecipesData as RecipeGroup[]
  }
  async materials(): Promise<Material[]> {
    return dummyMaterialsData as Material[]
  }

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async order(input: InputOrder): Promise<Order | undefined> {
    throw new Error("order() not implemented.")
  }
  async updateOrderStatus(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: InputOrderStatusUpdate,
  ): Promise<Order | undefined> {
    throw new Error("updateOrderStatus() not implemented.")
  }
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async saveRecipe(input: InputRecipeGroup): Promise<RecipeGroup> {
    throw new Error("saveRecipe() not implemented.")
  }
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async updateStock(input: InputStockUpdate): Promise<Material[]> {
    throw new Error("updateStock() not implemented.")
  }
}
