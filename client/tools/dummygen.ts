import yargs from "yargs"
import { HackbarClient } from "../client/client"
import fs from "fs"
import type { Order } from "../client/gen/types"

// Options:
//   --uri
//   --base      Base schema file.
//   -o, --dest  Output destination directory for generated dummy data. (JSON formatted)
const args = await yargs(process.argv.slice(2)).options({
  uri: {
    type: "string",
    default: "http://localhost:8080/recipes.v1graphql.Registry/",
    description: "URI for the GraphQL server.",
  },
  token: {
    type: "string",
    description: "JWT token for authorization.",
  },
  dest: {
    alias: "o",
    type: "string",
    default: "client/dummy/data/",
    description:
      "Output destination directory for generated dummy data. (JSON formatted)",
  },
}).argv

const client = new HackbarClient(args.uri, args.token)

console.log(`Writing generated dummy data to '${args.dest}'.`)

client
  .menu()
  .then((menu) => {
    const data = JSON.stringify(menu, null, 2)
    console.log(`- menu.json`)
    fs.writeFileSync(`${args.dest}/menu.json`, data)
  })
  .catch((e) => {
    console.error(`failed to fetch menu: ${e}`)
    console.log(`- menu.json (empty)`)
    fs.writeFileSync(`${args.dest}/menu.json`, "[]")
  })
client
  .orders()
  .then((orders) => {
    const maskedOrders = orders.map(
      (order): Order => ({
        ...order,
        customerID: "user@example.test",
      }),
    )
    const data = JSON.stringify(maskedOrders, null, 2)
    console.log(`- orders.json`)
    fs.writeFileSync(`${args.dest}/orders.json`, data)
  })
  .catch((e) => {
    console.error(`failed to fetch orders: ${e}`)
    console.log(`- orders.json (empty)`)
    fs.writeFileSync(`${args.dest}/orders.json`, "[]")
  })
client
  .recipes()
  .then((recipes) => {
    const data = JSON.stringify(recipes, null, 2)
    console.log(`- recipes.json`)
    fs.writeFileSync(`${args.dest}/recipes.json`, data)
  })
  .catch((e) => {
    console.error(`failed to fetch recipes: ${e}`)
    console.log(`- recipes.json (empty)`)
    fs.writeFileSync(`${args.dest}/recipes.json`, "[]")
  })
client
  .materials()
  .then((materials) => {
    const data = JSON.stringify(materials, null, 2)
    console.log(`- materials.json`)
    fs.writeFileSync(`${args.dest}/materials.json`, data)
  })
  .catch((e) => {
    console.error(`failed to fetch materials: ${e}`)
    console.log(`- materials.json (empty)`)
    fs.writeFileSync(`${args.dest}/materials.json`, "[]")
  })
