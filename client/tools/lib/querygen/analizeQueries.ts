// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-ignore
import graphqlQueryGen from "graphql-query-gen"
import { convertFunction, listScalars } from "../graphqlschema"
import type { functionArgType } from "../interfacegen"

export type Query = {
  schemaFilename?: string
  type: "query" | "mutation"
  name: string
  query: string
  variables: functionArgType[]
}

export const analizeQueries = (
  schemaRaw: string,
  typeDefImportAs?: string,
): { queries: Query[]; mutations: Query[] } => {
  const generated = graphqlQueryGen.processSchema(schemaRaw, {
    inputVariables: true,
  }) as {
    operations: { name: string; options: { name: string; query: string }[] }[]
  }

  const queries = generated.operations
    .filter((operation) => operation.options.length !== 0)
    .flatMap(({ options: queries }) => {
      return queries.reduce(
        (acc, { name: queryNameRaw, query }) => {
          acc.push({ name: queryNameRaw, query })
          return acc
        },
        [] as { name: string; query: string }[],
      )
    })
    .reduce(
      (acc, query) => {
        acc[query.name] = query.query
        return acc
      },
      {} as { [key: string]: string },
    )

  const scalars = listScalars(schemaRaw)

  const result = schemaRaw.split("\n").reduce(
    (acc, line) => {
      if (line.includes(".graphqls")) {
        acc.state.filename = line.split(" ")[2]
        return acc
      }
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
      let query = queries[func.name]

      switch (acc.state.in) {
        case "query": {
          const queryName =
            "get" + func.name[0].toUpperCase() + func.name.slice(1)
          query = query.replace("query ", `query ${queryName} `)
          break
        }
        case "mutation": {
          query = query.replace("mutation ", `mutation ${func.name} `)
          break
        }
      }

      acc.result.push({
        schemaFilename: acc.state.filename,
        type: acc.state.in,
        name: func.name,
        query,
        variables: func.args,
      })
      return acc
    },
    {
      state: {
        filename: undefined as string | undefined,
        in: null as "query" | "mutation" | null,
      },
      result: [] as Query[],
    },
  ).result

  return result.reduce(
    (acc, query) => {
      switch (query.type) {
        case "query":
          acc.queries.push(query)
          break
        case "mutation":
          acc.mutations.push(query)
          break
      }
      return acc
    },
    { queries: [] as Query[], mutations: [] as Query[] },
  )
}
