import yargs from "yargs"
import fs from "fs"
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-ignore
import graphqlQueryGen from "graphql-query-gen"
import { listFunctionTypes } from "./lib/querygen/listFunction"
import { WithDonoteditHeader } from "./lib/donoteditHeader"

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

console.log(`Generating query from '${args.schema}'.`)

const generated = graphqlQueryGen.processSchema(schemaRaw) as {
  operations: { name: string; options: { name: string; query: string }[] }[]
}

console.log(`Writing generated query to '${args.dest}'.`)
generated.operations
  .filter(
    (operation) => operation.options.length !== 0 && operation.name === "Query",
  )
  .forEach(({ name, options: queries }) => {
    const filename = name.toLowerCase()
    console.log(`- ${filename}.ts`)

    const importsRaw = `import { gql } from "@apollo/client/core"\n`
    const tsRaw = queries.reduce((acc, { name: queryNameRaw, query }) => {
      const queryName =
        "get" + queryNameRaw[0].toUpperCase() + queryNameRaw.slice(1)
      console.log(`  - ${queryName}`)
      acc += `\n`
      acc += `export const ${queryName} = gql\`\n`
      acc += query.replace("query {", `query ${queryName} {`) + `\n`
      acc += `\`\n`
      return acc
    }, WithDonoteditHeader(importsRaw))

    fs.writeFileSync(`${args.dest}/${filename}.ts`, tsRaw)
  })

console.log(`Generating interface from '${args.schema}'.`)

const IMPORT_TYPE_AS = "types"
const interfaceAST = listFunctionTypes(schemaRaw, IMPORT_TYPE_AS)

console.log(`Writing generated interface to '${args.dest}/interface.*.ts'.`)

const importsRaw = `import * as ${IMPORT_TYPE_AS} from "./types"\n`

const tsRawQueryClientInterface =
  interfaceAST.queries.reduce(
    (acc, { name, args, returnType }) => {
      acc += `  ${name}(`
      acc += args.map((arg) => `${arg.name}: ${arg.argType}`).join(", ")
      acc += `): ${returnType}\n`
      return acc
    },
    WithDonoteditHeader(importsRaw + `\nexport interface QueryClient {\n`),
  ) + `}\n`
fs.writeFileSync(`${args.dest}/interface.client.ts`, tsRawQueryClientInterface)

const tsRawMutationClientInterface =
  interfaceAST.mutations.reduce(
    (acc, { name, args, returnType }) => {
      acc += `  ${name}(`
      acc += args.map((arg) => `${arg.name}: ${arg.argType}`).join(", ")
      acc += `): ${returnType}\n`
      return acc
    },
    WithDonoteditHeader(importsRaw + `\nexport interface MutationClient {\n`),
  ) + `}\n`
fs.writeFileSync(
  `${args.dest}/interface.mutation.ts`,
  tsRawMutationClientInterface,
)
