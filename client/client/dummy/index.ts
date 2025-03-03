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
  Cashout,
  Checkout,
  CashoutInput,
  InputCheckout,
  InputCashoutQuery,
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
  async checkouts(): Promise<Checkout[]> {
    throw new Error("Method not implemented.")
  }
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async cashouts(input: InputCashoutQuery): Promise<Cashout[]> {
    throw new Error("Method not implemented.")
  }

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async order(input: InputOrder): Promise<Order> {
    throw new Error("order() not implemented.")
  }
  async updateOrderStatus(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: InputOrderStatusUpdate,
  ): Promise<Order> {
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
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async checkout(input: InputCheckout): Promise<Checkout> {
    throw new Error("Method not implemented.")
  }
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async cashout(input: CashoutInput): Promise<Cashout> {
    throw new Error("Method not implemented.")
  }
}
