import type { functionType } from "./listFunction"

export const generateInterface = (
  name: string,
  functions: functionType[],
): string => {
  return (
    functions.reduce((acc, { name, args, returnType }) => {
      acc += `  ${name}(`
      acc += args.map((arg) => `${arg.name}: ${arg.argType}`).join(", ")
      acc += `): ${returnType}\n`
      return acc
    }, `export interface ${name} {\n`) + `}\n`
  )
}
