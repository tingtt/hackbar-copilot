# oauth2rbac - E2E test

## Setup OAuth2

### 1. Google Cloud

- [Create OAuth client ID](https://console.cloud.google.com/apis/credentials/oauthclient)

### 2. GitHub

- [Register a new OAuth application](https://github.com/settings/applications/new)

## Create `test/e2e/.env`

```sh
PORT=8080
PORT_PROXY=8081
JWT_SECRET="<secret>"
OAUTH2_GITHUB="<your-client-id>;<your-client-secret>"
OAUTH2_GOOGLE="<your-client-id>;<your-client-secret>"
```

## Run

### Develop stage

```sh
make e2e-up
```

http://localhost:8080
http://127.0.0.1:8080
