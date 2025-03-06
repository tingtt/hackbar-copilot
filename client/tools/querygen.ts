import yargs from "yargs"
import fs from "fs"
import { WithDonoteditHeader } from "./lib/donoteditHeader"
import { generateInterface, listFunctionTypes } from "./lib/interfacegen"
import { analizeQueries } from "./lib/querygen"

// Options:
//   -i, --schema   Path to the schema file.
//   -o, --dest     Output destination directory for generated codes.
const args = await yargs(process.argv.slice(2)).options({
  schema: {
    alias: "i",
    type: "string",
    default: "hackbar-copilot.graphqls",
    description: "Path to the schema file.",
  },
  dest: {
    alias: "o",
    type: "string",
    default: "client/gen/",
    description: "Output destination directory for generated codes.",
  },
}).argv

const schemaRaw = fs.readFileSync(args.schema, "utf-8")

/**
 * Generate query from schema.
 */

console.log(`Generating query from '${args.schema}'.`)

const IMPORT_TYPE_AS = "types"
const { queries, mutations } = analizeQueries(schemaRaw, IMPORT_TYPE_AS)

console.log(`Writing generated query to '${args.dest}'.`)

{
  const importsRaw = `import { gql } from "@apollo/client/core"\nimport * as ${IMPORT_TYPE_AS} from "./types"\n`

  const queryTSRaw = queries.reduce((acc, query) => {
    const queryName = "get" + query.name[0].toUpperCase() + query.name.slice(1)
    acc += `\n`
    acc += `export const ${queryName} = (`
    if (query.variables.length !== 0) {
      acc += `variables: {\n`
      acc += query.variables.reduce((acc, variable) => {
        acc += `  ${variable.name}: ${variable.argType}\n`
        return acc
      }, "")
      acc += `}`
    }
    acc += `) => ({\n`
    acc += `  query: gql\`\n`
    acc += query.query + `\`,\n`
    if (query.variables.length !== 0) {
      acc += `  variables,\n`
    }
    acc += `})\n`
    return acc
  }, "")

  fs.writeFileSync(
    `${args.dest}/query.ts`,
    WithDonoteditHeader(importsRaw + "\n" + queryTSRaw),
  )

  const mutationTSRaw = mutations.reduce((acc, query) => {
    acc += `\n`
    acc += `export const ${query.name} = (`
    if (query.variables.length !== 0) {
      acc += `variables: {\n`
      acc += query.variables.reduce((acc, variable) => {
        acc += `  ${variable.name}: ${variable.argType}\n`
        return acc
      }, "")
      acc += `}`
    }
    acc += `) => ({\n`
    acc += `  mutation : gql\`\n`
    acc += query.query + `\`,\n`
    if (query.variables.length !== 0) {
      acc += `  variables,\n`
    }
    acc += `})\n`
    return acc
  }, "")

  fs.writeFileSync(
    `${args.dest}/mutation.ts`,
    WithDonoteditHeader(importsRaw + "\n" + mutationTSRaw),
  )
}

/**
 * Generate interface from schema.
 */

console.log(`Generating interface from '${args.schema}'.`)

const interfaceAST = listFunctionTypes(schemaRaw, IMPORT_TYPE_AS)

console.log(`Writing generated interface to '${args.dest}/interface.*.ts'.`)

const importsRaw = `import * as ${IMPORT_TYPE_AS} from "./types"\n`

fs.writeFileSync(
  `${args.dest}/interface.client.ts`,
  WithDonoteditHeader(
    importsRaw +
      "\n" +
      generateInterface(
        "QueryClient",
        interfaceAST.queries.map((query) => {
          query.name = "get" + query.name[0].toUpperCase() + query.name.slice(1)
          return query
        }),
      ),
  ),
)

fs.writeFileSync(
  `${args.dest}/interface.mutation.ts`,
  WithDonoteditHeader(
    importsRaw +
      "\n" +
      generateInterface("MutationClient", interfaceAST.mutations),
  ),
)
