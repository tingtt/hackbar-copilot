import {
  ApolloClient,
  createHttpLink,
  InMemoryCache,
  type NormalizedCacheObject,
} from "@apollo/client/core"
import { setContext } from "@apollo/client/link/context"
import type * as types from "./gen/types"
import * as query from "./gen/query"
import * as mutation from "./gen/mutation"
import type { QueryClient } from "./gen/interface.client"
import type { MutationClient } from "./gen/interface.mutation"

export class HackbarClient implements QueryClient, MutationClient {
  private client: ApolloClient<NormalizedCacheObject>

  constructor(uri: string, jwt?: string) {
    const link = setContext((_, { headers }) => {
      return {
        headers: {
          ...headers,
          Authorization: jwt ? `Bearer ${jwt}` : "",
        },
      }
    }).concat(createHttpLink({ uri }))
    this.client = new ApolloClient({ link, cache: new InMemoryCache() })
  }

  async getMenu(): Promise<types.MenuItem[]> {
    const res = await this.client.query<{ menu: types.MenuItem[] }>(
      query.getMenu(),
    )
    return res.data.menu
  }
  async getOrders(): Promise<types.Order[]> {
    const res = await this.client.query<{ orders: types.Order[] }>(
      query.getOrders(),
    )
    return res.data.orders
  }
  async getRecipes(): Promise<types.RecipeGroup[]> {
    const res = await this.client.query<{ recipes: types.RecipeGroup[] }>(
      query.getRecipes(),
    )
    return res.data.recipes
  }
  async getMaterials(): Promise<types.Material[]> {
    const res = await this.client.query<{ materials: types.Material[] }>(
      query.getMaterials(),
    )
    return res.data.materials
  }
  async getCheckouts(): Promise<types.Checkout[]> {
    const res = await this.client.query<{ checkouts: types.Checkout[] }>(
      query.getCheckouts(),
    )
    return res.data.checkouts
  }
  async getCashouts(input: types.InputCashoutQuery): Promise<types.Cashout[]> {
    const res = await this.client.query<{ cashouts: types.Cashout[] }>(
      query.getCashouts({ input }),
    )
    return res.data.cashouts
  }
  async getUserInfo(): Promise<types.User> {
    const res = await this.client.query<{ user: types.User }>(
      query.getUserInfo(),
    )
    return res.data.user
  }

  async order(input: types.InputOrder): Promise<types.Order> {
    const res = await this.client.mutate<{ order: types.Order }>(
      mutation.order({ input }),
    )
    return res.data!.order
  }
  async updateOrderStatus(
    input: types.InputOrderStatusUpdate,
  ): Promise<types.Order> {
    const res = await this.client.mutate<{ order: types.Order }>(
      mutation.updateOrderStatus({ input }),
    )
    return res.data!.order
  }
  async saveRecipe(input: types.InputRecipeGroup): Promise<types.RecipeGroup> {
    const res = await this.client.mutate<{ recipe: types.RecipeGroup }>(
      mutation.saveRecipe({ input }),
    )
    return res.data!.recipe
  }
  async updateStock(input: types.InputStockUpdate): Promise<types.Material[]> {
    const res = await this.client.mutate<{ materials: types.Material[] }>(
      mutation.updateStock({ input }),
    )
    return res.data!.materials
  }
  async checkout(input: types.InputCheckout): Promise<types.Checkout> {
    const res = await this.client.mutate<{ checkout: types.Checkout }>(
      mutation.checkout({ input }),
    )
    return res.data!.checkout
  }
  async cashout(input: types.CashoutInput): Promise<types.Cashout> {
    const res = await this.client.mutate<{ cashout: types.Cashout }>(
      mutation.cashout({ input }),
    )
    return res.data!.cashout
  }
}
