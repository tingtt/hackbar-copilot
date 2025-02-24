import { convertType } from "./convertType"
import { listScalars } from "./listScalars"

type functionType = {
  name: string
  args: { name: string; argType: string }[]
  returnType: string
}

export const listFunctionTypes = (
  schemaRaw: string,
  IMPORT_TYPE_AS: string,
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
      if (acc.state.in !== null) {
        line = line.trim()
        const [returnTypeRaw, remain] = (() => {
          const splitted = line.split(": ")
          if (
            splitted.length === 1 ||
            splitted[splitted.length - 1].endsWith(")")
          ) {
            return [null, splitted.join(": ")]
          }
          return [
            splitted[splitted.length - 1],
            splitted.slice(0, -1).join(": "),
          ]
        })()
        const [functionName, argsRaw] = (() => {
          const splitted = remain.split("(")
          if (splitted.length === 1) {
            return [splitted[0], null]
          }
          return [splitted[0], splitted[1].slice(0, -1)]
        })()
        const args = argsRaw
          ? argsRaw.split(", ").map((arg) => {
              const [name, typeRaw] = arg.split(": ")
              return {
                name,
                argType: convertType(typeRaw, scalars, "input", IMPORT_TYPE_AS),
              }
            })
          : []
        const func: functionType = {
          name: functionName,
          args,
          returnType: returnTypeRaw
            ? `Promise<${convertType(returnTypeRaw, scalars, "output", IMPORT_TYPE_AS)}>`
            : "Promise<void>",
        }
        acc.result[acc.state.in === "query" ? "queries" : "mutations"].push(
          func,
        )
      }
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
