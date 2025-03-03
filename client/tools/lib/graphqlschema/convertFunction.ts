import type { functionArgType, functionType } from "../interfacegen"
import { convertType } from "./convertType"

export const convertFunction = (
  functionLine: string,
  scalars: string[],
  typeDefImportAs?: string,
): functionType => {
  const [returnTypeRaw, remain] = (() => {
    const splitted = functionLine.trim().split(": ")
    if (splitted.length === 1 || splitted[splitted.length - 1].endsWith(")")) {
      return [null, splitted.join(": ")]
    }
    return [splitted[splitted.length - 1], splitted.slice(0, -1).join(": ")]
  })()
  const returnType = returnTypeRaw
    ? `Promise<${convertType(returnTypeRaw, scalars, "output", typeDefImportAs)}>`
    : "Promise<void>"

  const [functionName, argsRaw] = (() => {
    const splitted = remain.split("(")
    if (splitted.length === 1) {
      return [splitted[0], null]
    }
    return [splitted[0], splitted[1].slice(0, -1)]
  })()

  const args = argsRaw
    ? argsRaw.split(", ").map((arg): functionArgType => {
        const [name, typeRaw] = arg.split(": ")
        return {
          name,
          argType: convertType(typeRaw, scalars, "input", typeDefImportAs),
        }
      })
    : []

  return {
    name: functionName,
    args,
    returnType: returnType,
  }
}
