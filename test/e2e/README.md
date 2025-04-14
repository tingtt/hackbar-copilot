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

## Create config.yml

1. Copy example config

  ```sh
  cp test/e2e/config.yml.example test/e2e/config.yml
  ```

2. Edit for your environment

  ```diff
    proxies:
      - external_url: "http://localhost:8081/"
        target: "http://app:80"

    acl:
      "http://localhost:8081":
        paths:
          "/":
            - methods: ["*"]
  -           emails: ["<your email>"]
  +           emails: ["ting.taku@gmail.com"]
        roles:
  -       "bartender": ["<your email>"]
  +       "bartender": ["ting.taku@gmail.com"]
  ```

## Run

### Develop stage

```sh
make e2e-up
```

- http://localhost:8081 for requesting via oauth2rbac
- http://localhost:8080 for direct
