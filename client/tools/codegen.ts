import type { CodegenConfig } from "@graphql-codegen/cli"

const config: CodegenConfig = {
  overwrite: true,
  schema: "hackbar-copilot.graphqls",
  generates: {
    "client/gen/types.ts": {
      plugins: ["typescript"],
      config: {
        scalars: {
          DateTime: "string",
        },
      },
    },
  },
}

export default config
