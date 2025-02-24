import {
  ApolloClient,
  createHttpLink,
  InMemoryCache,
  type NormalizedCacheObject,
} from "@apollo/client/core"
import { setContext } from "@apollo/client/link/context"
import type * as types from "./gen/types"
import * as query from "./gen/query"
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

  async menu(): Promise<types.MenuGroup[]> {
    const res = await this.client.query<{ menu: types.MenuGroup[] }>({
      query: query.getMenu,
    })
    return res.data.menu
  }
  async orders(): Promise<types.Order[]> {
    const res = await this.client.query<{ orders: types.Order[] }>({
      query: query.getOrders,
    })
    return res.data.orders
  }
  async recipes(): Promise<types.RecipeGroup[]> {
    const res = await this.client.query<{ recipes: types.RecipeGroup[] }>({
      query: query.getRecipes,
    })
    return res.data.recipes
  }
  async materials(): Promise<types.Material[]> {
    const res = await this.client.query<{ materials: types.Material[] }>({
      query: query.getMaterials,
    })
    return res.data.materials
  }

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async order(input: types.InputOrder): Promise<types.Order | undefined> {
    throw new Error("Method not implemented.")
  }
  async updateOrderStatus(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    input: types.InputOrderStatusUpdate,
  ): Promise<types.Order | undefined> {
    throw new Error("Method not implemented.")
  }
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async saveRecipe(input: types.InputRecipeGroup): Promise<types.RecipeGroup> {
    throw new Error("Method not implemented.")
  }
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async updateStock(input: types.InputStockUpdate): Promise<types.Material[]> {
    throw new Error("Method not implemented.")
  }
}
