/**
 * client
 *
 ** Example: (for Next.js App Router)
 *   ```tsx
 *   //* `src/app/layout.tsx` or `app/layout.tsx
 *   import { HackbarCopilotAPIProvider, HackbarClient, useClient } from "@hackbar/copilot/client"
 *
 *   export default function RootLayout({
 *     children,
 *   }: Readonly<{
 *     children: React.ReactNode
 *   }>) {
 *     return (
 *      <html>
 *        <HackbarCopilotAPIProvider
 *          client={new HackbarClient("https://example.test/uri/to/backend")}
 *        >                            //* Add this line
 *                                     //* You can use `useClient()` in any child component.
 *          <body>{children}</body>
 *        </HackbarCopilotAPIProvider> //* Add this line
 *      </html>
 *     )
 *   }
 *   ```
 *
 ** Example: (for React)
 *   ```tsx
 *   //* `src/app/App.tsx`
 *   import { HackbarCopilotAPIProvider, HackbarClient, useClient } from "@hackbar/copilot/client"
 *
 *   export const App: React.FunctionComponent = () => {
 *     return (
 *       <HackbarCopilotAPIProvider
 *         client={new HackbarClient("https://example.test/uri/to/backend")}
 *       >                            //* Add this line
 *                                    //* You can use `useClient()` in any child component.
 *         <YourComponent />
 *       </HackbarCopilotAPIProvider> //* Add this line
 *     )
 *   }
 *   ```
 */
export * from "./client"
export * from "./Provider"
