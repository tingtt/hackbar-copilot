import type * as types from "./gen/types"
import * as query from "./gen/query"
import * as mutation from "./gen/mutation"
import type { QueryClient } from "./gen/interface.client"
import type { MutationClient } from "./gen/interface.mutation"

export class HackbarClient implements QueryClient, MutationClient {
  constructor(
    private uri: string,
    private jwt?: string,
  ) {}

  async fetch<T>(
    payload: {
      query: string
      variables?: { [key in string]: unknown }
    },
    init?: RequestInit,
  ): Promise<{ data: T; error: null } | { data: null; error: string }> {
    const res = await fetch(this.uri, {
      method: "POST",
      ...init,
      headers: {
        "Content-Type": "application/json",
        Authorization: this.jwt ? `Bearer ${this.jwt}` : "",
        ...init?.headers,
      },
      body: JSON.stringify(payload),
    })
    try {
      const json: {
        data: T | null
        errors: { message: string; path: string[] }[]
      } = await res.json()
      if (json.data === null) {
        throw new Error(
          json.errors.reduce((acc, err) => {
            return (
              acc + `failed to query (${err.path.join(",")}): ${err.message}\n`
            )
          }, ""),
        )
      }
      return { data: json.data, error: null }
    } catch (err: unknown) {
      if (err instanceof Error) {
        return { data: null, error: err.message }
      } else {
        return { data: null, error: `unknown error: ${err}` }
      }
    }
  }

  async getMenu() {
    const res = await this.fetch<{ menu: types.MenuItem[] }>(query.getMenu())
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.menu, error: null }
  }
  async getUncheckedOrdersCustomer() {
    const res = await this.fetch<{ uncheckedOrdersCustomer: types.Order[] }>(
      query.getUncheckedOrdersCustomer(),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.uncheckedOrdersCustomer, error: null }
  }
  async getUncheckedOrders() {
    const res = await this.fetch<{ uncheckedOrders: types.Order[] }>(
      query.getUncheckedOrders(),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.uncheckedOrders, error: null }
  }
  async getRecipes() {
    const res = await this.fetch<{ recipes: types.RecipeGroup[] }>(
      query.getRecipes(),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.recipes, error: null }
  }
  async getMaterials() {
    const res = await this.fetch<{ materials: types.Material[] }>(
      query.getMaterials(),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.materials, error: null }
  }
  async getUncashedoutCheckouts() {
    const res = await this.fetch<{ uncashedoutCheckouts: types.Checkout[] }>(
      query.getUncashedoutCheckouts(),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.uncashedoutCheckouts, error: null }
  }
  async getCashouts(input: types.InputCashoutQuery) {
    const res = await this.fetch<{ cashouts: types.Cashout[] }>(
      query.getCashouts({ input }),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.cashouts, error: null }
  }
  async getUserInfo() {
    const res = await this.fetch<{ userInfo: types.User }>(query.getUserInfo())
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.userInfo, error: null }
  }

  async order(input: types.InputOrder) {
    const res = await this.fetch<{ order: types.Order }>(
      mutation.order({ input }),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.order, error: null }
  }
  async updateOrderStatus(input: types.InputOrderStatusUpdate) {
    const res = await this.fetch<{ order: types.Order }>(
      mutation.updateOrderStatus({ input }),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.order, error: null }
  }
  async saveRecipe(input: types.InputRecipeGroup) {
    const res = await this.fetch<{ saveRecipe: types.RecipeGroup }>(
      mutation.saveRecipe({ input }),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.saveRecipe, error: null }
  }
  async updateStock(input: types.InputStockUpdate) {
    const res = await this.fetch<{ updateStock: types.Material[] }>(
      mutation.updateStock({ input }),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.updateStock, error: null }
  }
  async checkout(input: types.InputCheckout) {
    const res = await this.fetch<{ checkout: types.Checkout }>(
      mutation.checkout({ input }),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.checkout, error: null }
  }
  async cashout(input: types.CashoutInput) {
    const res = await this.fetch<{ cashout: types.Cashout }>(
      mutation.cashout({ input }),
    )
    if (res.error !== null) {
      return { data: null, error: res.error }
    }
    return { data: res.data.cashout, error: null }
  }
}
