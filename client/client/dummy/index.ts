import type { QueryClient } from "../gen/interface.client"
import type { MutationClient } from "../gen/interface.mutation"
import type {
  MenuItem,
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
  User,
} from "../gen/types"

import dummyMenuData from "./data/menu.json"
import dummyOrdersData from "./data/orders.json"
import dummyRecipesData from "./data/recipes.json"
import dummyMaterialsData from "./data/materials.json"

export class DummyHackbarClient implements QueryClient, MutationClient {
  async getMenu() {
    return dummyMenuData as { data: MenuItem[]; error: null }
  }
  async getOrders() {
    return dummyOrdersData as { data: Order[]; error: null }
  }
  async getRecipes() {
    return dummyRecipesData as { data: RecipeGroup[]; error: null }
  }
  async getMaterials() {
    return dummyMaterialsData as { data: Material[]; error: null }
  }
  async getCheckouts(): Promise<
    { data: null; error: string } | { data: Checkout[]; error: null }
  > {
    throw new Error("Method not implemented.")
  }
  async getCashouts(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: InputCashoutQuery,
  ): Promise<{ data: null; error: string } | { data: Cashout[]; error: null }> {
    throw new Error("Method not implemented.")
  }
  async getUserInfo() {
    return {
      data: { email: "john.doe@example.test", name: "John Doe" } as User,
      error: null,
    }
  }

  async order(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: InputOrder,
  ): Promise<{ data: null; error: string } | { data: Order; error: null }> {
    throw new Error("order() not implemented.")
  }
  async updateOrderStatus(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: InputOrderStatusUpdate,
  ): Promise<{ data: null; error: string } | { data: Order; error: null }> {
    throw new Error("updateOrderStatus() not implemented.")
  }
  async saveRecipe(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: InputRecipeGroup,
  ): Promise<
    { data: null; error: string } | { data: RecipeGroup; error: null }
  > {
    throw new Error("saveRecipe() not implemented.")
  }
  async updateStock(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: InputStockUpdate,
  ): Promise<
    { data: null; error: string } | { data: Material[]; error: null }
  > {
    throw new Error("updateStock() not implemented.")
  }
  async checkout(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: InputCheckout,
  ): Promise<{ data: null; error: string } | { data: Checkout; error: null }> {
    throw new Error("Method not implemented.")
  }
  async cashout(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: CashoutInput,
  ): Promise<{ data: null; error: string } | { data: Cashout; error: null }> {
    throw new Error("Method not implemented.")
  }
}
