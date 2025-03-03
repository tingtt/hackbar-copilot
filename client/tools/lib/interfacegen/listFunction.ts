import { convertFunction, listScalars } from "../graphqlschema"

export type functionType = {
  name: string
  args: functionArgType[]
  returnType: string
}
export type functionArgType = { name: string; argType: string }

export const listFunctionTypes = (
  schemaRaw: string,
  typeDefImportAs?: string,
): {
  queries: functionType[]
  mutations: functionType[]
} => {
  const scalars = listScalars(schemaRaw)

  return schemaRaw.split("\n").reduce(
    (acc, line) => {
      if (line.includes("type Query {")) {
        acc.state.in = "query"
        return acc
      }
      if (line.includes("type Mutation {")) {
        acc.state.in = "mutation"
        return acc
      }
      if (line.startsWith("}")) {
        acc.state.in = null
        return acc
      }
      if (acc.state.in === null || line === "") {
        return acc
      }

      // Process query or mutation
      const func = convertFunction(line.trim(), scalars, typeDefImportAs)
      acc.result[acc.state.in === "query" ? "queries" : "mutations"].push(func)

      return acc
    },
    {
      state: { in: null as "query" | "mutation" | null },
      result: { queries: [], mutations: [] } as {
        queries: functionType[]
        mutations: functionType[]
      },
    },
  ).result
}
