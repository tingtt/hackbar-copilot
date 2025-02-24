import { createContext, useContext } from "react"
import type { QueryClient } from "./gen/interface.client"
import type { MutationClient } from "./gen/interface.mutation"

type Client = QueryClient & MutationClient

const clientContext = createContext({} as Client)

export const useClient = () => {
  return useContext(clientContext)
}

export const HackbarCopilotAPIProvider = ({
  children,
  client,
}: React.PropsWithChildren & { client: Client }) => {
  return (
    <clientContext.Provider value={client}>{children}</clientContext.Provider>
  )
}
