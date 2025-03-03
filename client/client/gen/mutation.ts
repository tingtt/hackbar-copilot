// Code generated by script, DO NOT EDIT.

import { gql } from "@apollo/client/core"
import * as types from "./types"

export const cashout = (variables: { input: types.CashoutInput }) => ({
  mutation: gql`
    mutation cashout($input: CashoutInput!) {
      cashout(input: $input) {
        __typename
        checkouts {
          id
          customerID
          orderIDs
          diffs {
            price
            description
          }
          totalPrice
          paymentType
          timestamp
        }
        revenue
        timestamp
        staffID
      }
    }
  `,
  variables,
})

export const checkout = (variables: { input: types.InputCheckout }) => ({
  mutation: gql`
    mutation checkout($input: InputCheckout!) {
      checkout(input: $input) {
        __typename
        id
        customerID
        orderIDs
        diffs {
          price
          description
        }
        totalPrice
        paymentType
        timestamp
      }
    }
  `,
  variables,
})

export const order = (variables: { input: types.InputOrder }) => ({
  mutation: gql`
    mutation order($input: InputOrder!) {
      order(input: $input) {
        __typename
        id
        customerID
        menuItemID
        timestamps {
          status
          timestamp
        }
        status
        price
      }
    }
  `,
  variables,
})

export const updateOrderStatus = (variables: {
  input: types.InputOrderStatusUpdate
}) => ({
  mutation: gql`
    mutation updateOrderStatus($input: InputOrderStatusUpdate!) {
      updateOrderStatus(input: $input) {
        __typename
        id
        customerID
        menuItemID
        timestamps {
          status
          timestamp
        }
        status
        price
      }
    }
  `,
  variables,
})

export const saveRecipe = (variables: { input: types.InputRecipeGroup }) => ({
  mutation: gql`
    mutation saveRecipe($input: InputRecipeGroup!) {
      saveRecipe(input: $input) {
        __typename
        name
        imageURL
        recipes {
          name
          type {
            name
            description
          }
          glass {
            name
            imageURL
            description
          }
          steps {
            material
            amount
            description
          }
        }
      }
    }
  `,
  variables,
})

export const updateStock = (variables: { input: types.InputStockUpdate }) => ({
  mutation: gql`
    mutation updateStock($input: InputStockUpdate!) {
      updateStock(input: $input) {
        __typename
        name
        inStock
      }
    }
  `,
  variables,
})
