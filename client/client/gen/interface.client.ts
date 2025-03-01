// Code generated by script, DO NOT EDIT.

import * as types from "./types"

export interface QueryClient {
  menu(): Promise<types.MenuGroup[]>
  orders(): Promise<types.Order[]>
  recipes(): Promise<types.RecipeGroup[]>
  materials(): Promise<types.Material[]>
}
