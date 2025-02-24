export const listScalars = (schemaRaw: string): string[] => {
  const scalars = schemaRaw
    .split("\n")
    .filter((line) => line.includes("scalar"))
    .map((line) => line.split(" ")[1])
  scalars.push("String", "Int", "Float", "Boolean", "ID")
  return scalars
}
