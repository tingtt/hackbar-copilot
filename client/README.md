# hackbar-copilot

## Usage (Next.js App Router)

```ts
//* next.config.ts
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  transpilePackages: ["@tingtt/hackbar-copilot"],
};

export default nextConfig;
```

```tsx
//* `src/app/layout.tsx` or `app/layout.tsx
import { HackbarCopilotAPIProvider, HackbarClient, useClient } from "@tingtt/hackbar-copilot/client"

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html>
      <HackbarCopilotAPIProvider
        client={new HackbarClient("https://example.test/uri/to/backend")}
      >                            {/* Add this line */}
                                   {/* You can use `useClient()` in any child component. */}
        <body>{children}</body>
      </HackbarCopilotAPIProvider> {/* Add this line */}
    </html>
  )
}
```

## Usage (React)

```tsx
//* `src/app/App.tsx`
import { HackbarCopilotAPIProvider, HackbarClient, useClient } from "@tingtt/hackbar-copilot/client"

export const App: React.FunctionComponent = () => {
  return (
    <HackbarCopilotAPIProvider
      client={new HackbarClient("https://example.test/uri/to/backend")}
    >                            {/* Add this line */}
                                 {/* You can use `useClient()` in any child component. */}
      <YourComponent />
    </HackbarCopilotAPIProvider> {/* Add this line */}
  )
}
```

## Usage (Dummy data)

```diff
- import { HackbarClient } from "@tingtt/hackbar-copilot/client"
+ import { DummyHackbarClient } from "@tingtt/hackbar-copilot/client/dummy"

- new HackbarClient("https://example.test/uri/to/backend")
+ new DummyHackbarClient()
```

## Sync graphql schema and generate codes

```sh
bun generate
```

## Sync dummy data with responses

```sh
# Run hackbar-copilot (backend)

# Update client/dummy/data/*.json
bun dummygen -- --uri "http://localhost:8080/recipes.v1graphql.Registry/" --token "<JWT token>"
```
