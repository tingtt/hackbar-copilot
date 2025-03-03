export type Maybe<T> = T | null
export type InputMaybe<T> = Maybe<T>
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K]
}
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>
}
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>
}
export type MakeEmpty<
  T extends { [key: string]: unknown },
  K extends keyof T,
> = { [_ in K]?: never }
export type Incremental<T> =
  | T
  | {
      [P in keyof T]?: P extends " $fragmentName" | "__typename" ? T[P] : never
    }
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string }
  String: { input: string; output: string }
  Boolean: { input: boolean; output: boolean }
  Int: { input: number; output: number }
  Float: { input: number; output: number }
  DateTime: { input: string; output: string }
}

export type Cashout = {
  __typename?: "Cashout"
  checkouts: Array<Checkout>
  revenue: Scalars["Float"]["output"]
  staffID: Scalars["String"]["output"]
  timestamp: Scalars["DateTime"]["output"]
}

export type CashoutInput = {
  checkoutIDs: Array<Scalars["String"]["input"]>
  staffID: Scalars["String"]["input"]
}

export type Checkout = {
  __typename?: "Checkout"
  customerID: Scalars["String"]["output"]
  diffs: Array<PaymentDiff>
  id: Scalars["String"]["output"]
  orderIDs: Array<Scalars["String"]["output"]>
  paymentType: CheckoutType
  timestamp: Scalars["DateTime"]["output"]
  totalPrice: Scalars["Float"]["output"]
}

export enum CheckoutType {
  Cash = "CASH",
  Credit = "CREDIT",
  Qr = "QR",
}

export type GlassType = {
  __typename?: "GlassType"
  description?: Maybe<Scalars["String"]["output"]>
  imageURL?: Maybe<Scalars["String"]["output"]>
  name: Scalars["String"]["output"]
}

export type InputAsMenuArgs = {
  flavor?: InputMaybe<Scalars["String"]["input"]>
}

export type InputAsMenuItemArgs = {
  imageURL?: InputMaybe<Scalars["String"]["input"]>
  price: Scalars["Float"]["input"]
}

export type InputCashoutQuery = {
  since: Scalars["DateTime"]["input"]
  until: Scalars["DateTime"]["input"]
}

export type InputCheckout = {
  customerID: Scalars["String"]["input"]
  diffs: Array<InputPriceDiff>
  orderIDs: Array<Scalars["String"]["input"]>
  paymentType: CheckoutType
}

export type InputGlassType = {
  description?: InputMaybe<Scalars["String"]["input"]>
  imageURL?: InputMaybe<Scalars["String"]["input"]>
  name: Scalars["String"]["input"]
  save?: InputMaybe<Scalars["Boolean"]["input"]>
}

export type InputOrder = {
  menuItemID: Scalars["String"]["input"]
}

export type InputOrderStatusUpdate = {
  id: Scalars["String"]["input"]
  status: OrderStatus
}

export type InputPriceDiff = {
  description?: InputMaybe<Scalars["String"]["input"]>
  price: Scalars["Float"]["input"]
}

export type InputRecipe = {
  asMenu?: InputMaybe<InputAsMenuItemArgs>
  glassType?: InputMaybe<InputGlassType>
  name: Scalars["String"]["input"]
  recipeType?: InputMaybe<InputRecipeType>
  steps?: InputMaybe<Array<InputStep>>
}

export type InputRecipeGroup = {
  asMenu?: InputMaybe<InputAsMenuArgs>
  imageURL?: InputMaybe<Scalars["String"]["input"]>
  name: Scalars["String"]["input"]
  recipes?: InputMaybe<Array<InputRecipe>>
}

export type InputRecipeType = {
  description?: InputMaybe<Scalars["String"]["input"]>
  name: Scalars["String"]["input"]
  save?: InputMaybe<Scalars["Boolean"]["input"]>
}

export type InputStep = {
  amount?: InputMaybe<Scalars["String"]["input"]>
  description?: InputMaybe<Scalars["String"]["input"]>
  material?: InputMaybe<Scalars["String"]["input"]>
}

export type InputStockUpdate = {
  in?: InputMaybe<Array<Scalars["String"]["input"]>>
  out?: InputMaybe<Array<Scalars["String"]["input"]>>
}

export type Material = {
  __typename?: "Material"
  inStock: Scalars["Boolean"]["output"]
  name: Scalars["String"]["output"]
}

export type MenuGroup = {
  __typename?: "MenuGroup"
  flavor?: Maybe<Scalars["String"]["output"]>
  imageURL?: Maybe<Scalars["String"]["output"]>
  items?: Maybe<Array<MenuItem>>
  minPriceYen: Scalars["Float"]["output"]
  name: Scalars["String"]["output"]
}

export type MenuItem = {
  __typename?: "MenuItem"
  imageURL?: Maybe<Scalars["String"]["output"]>
  materials?: Maybe<Array<Scalars["String"]["output"]>>
  name: Scalars["String"]["output"]
  outOfStock: Scalars["Boolean"]["output"]
  priceYen: Scalars["Float"]["output"]
  recipe?: Maybe<Recipe>
}

export type Mutation = {
  __typename?: "Mutation"
  cashout: Cashout
  checkout: Checkout
  order: Order
  saveRecipe: RecipeGroup
  updateOrderStatus: Order
  updateStock: Array<Material>
}

export type MutationCashoutArgs = {
  input: CashoutInput
}

export type MutationCheckoutArgs = {
  input: InputCheckout
}

export type MutationOrderArgs = {
  input: InputOrder
}

export type MutationSaveRecipeArgs = {
  input: InputRecipeGroup
}

export type MutationUpdateOrderStatusArgs = {
  input: InputOrderStatusUpdate
}

export type MutationUpdateStockArgs = {
  input: InputStockUpdate
}

export type Order = {
  __typename?: "Order"
  customerID: Scalars["String"]["output"]
  id: Scalars["String"]["output"]
  menuItemID: Scalars["String"]["output"]
  price: Scalars["Float"]["output"]
  status: OrderStatus
  timestamps: Array<OrderStatusUpdateTimestamp>
}

export enum OrderStatus {
  Canceled = "CANCELED",
  Checkedout = "CHECKEDOUT",
  Delivered = "DELIVERED",
  Ordered = "ORDERED",
  Prepared = "PREPARED",
  Unknown = "UNKNOWN",
}

export type OrderStatusUpdateTimestamp = {
  __typename?: "OrderStatusUpdateTimestamp"
  status: OrderStatus
  timestamp: Scalars["DateTime"]["output"]
}

export type PaymentDiff = {
  __typename?: "PaymentDiff"
  description?: Maybe<Scalars["String"]["output"]>
  price: Scalars["Float"]["output"]
}

export type Query = {
  __typename?: "Query"
  cashouts: Array<Cashout>
  checkouts: Array<Checkout>
  materials: Array<Material>
  menu: Array<MenuGroup>
  orders: Array<Order>
  recipes: Array<RecipeGroup>
}

export type QueryCashoutsArgs = {
  input: InputCashoutQuery
}

export type Recipe = {
  __typename?: "Recipe"
  glass?: Maybe<GlassType>
  name: Scalars["String"]["output"]
  steps?: Maybe<Array<Step>>
  type?: Maybe<RecipeType>
}

export type RecipeGroup = {
  __typename?: "RecipeGroup"
  imageURL?: Maybe<Scalars["String"]["output"]>
  name: Scalars["String"]["output"]
  recipes?: Maybe<Array<Recipe>>
}

export type RecipeType = {
  __typename?: "RecipeType"
  description?: Maybe<Scalars["String"]["output"]>
  name: Scalars["String"]["output"]
}

export type Step = {
  __typename?: "Step"
  amount?: Maybe<Scalars["String"]["output"]>
  description?: Maybe<Scalars["String"]["output"]>
  material?: Maybe<Scalars["String"]["output"]>
}
