export const convertType = (
  typeRaw: string,
  scalars: string[],
  mode: "input" | "output",
  typeDefImportAs?: string,
): string => {
  const converted = convert(typeRaw, scalars, mode, typeDefImportAs)
    .replaceAll("<", "[")
    .replaceAll(">", "]")
  if (converted.startsWith("(") && converted.endsWith(")")) {
    return converted.slice(1, -1)
  }
  return converted
}

const convert = (
  typeRaw: string,
  scalars: string[],
  mode: "input" | "output",
  typeDefImportAs?: string,
): string => {
  if (typeRaw.endsWith("!")) {
    return unwrapUndefined(
      convert(typeRaw.slice(0, -1), scalars, mode, typeDefImportAs),
    )
  }
  if (typeRaw.endsWith("]")) {
    return wrapUndefined(
      convert(typeRaw.slice(1, -1), scalars, mode, typeDefImportAs) + "<>",
    )
  }
  if (scalars.includes(typeRaw)) {
    if (typeDefImportAs) {
      return wrapUndefined(
        `${typeDefImportAs}.Scalars<"${typeRaw}"><"${mode}">`,
      )
    }
    return wrapUndefined(`Scalars<"${typeRaw}"><"${mode}">`)
  }
  if (typeDefImportAs) {
    return wrapUndefined(`${typeDefImportAs}.${typeRaw}`)
  }
  return wrapUndefined(typeRaw)
}

const wrapUndefined = (type: string): string => {
  return `(${type} | undefined)`
}

const unwrapUndefined = (type: string): string => {
  return type.slice(1).replace(" | undefined)", "")
}
