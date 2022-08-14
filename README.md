# Auto-scalable Github Runner
Simple auto-scaling on self-hosted runner using Docker for Github Actions.

> **Warning**: In development! Please don't use it in live environment yet.

---

## Usage

### How to start

- Go to https://github.com/settings/tokens/new
  - Expiration
    - No expiration
  - Select scopes
    - [x] repo
      - [x] repo:status
      - [x] repo_deployment
      - [x] public_repo
      - [x] repo:invite
      - [x] security_events
- Go to _GITHUB_REPOSITORY_/settings/hooks/new
  - Payload URL
    - http://localhost:8080/api/hooks 
  - Content type
    - application/json
  - Which events would you like to trigger this webhook?
    - [x] Let me select individual events.
      - [x] Workflow jobs (Workflow job queued, requested or completed on a repository.)
  - [x] Active

### Setup
> The [.env.example](.env.example) file contains example environment variables.

**docker**
```shell
docker build . -t auto-scalable-github-runner
docker run --env-file ./.env -d auto-scalable-github-runner
```

**docker-compose**
```shell
docker compose up --build -d
```

### Config
.github/workflows/main.yml
```diff
jobs:
  process:
+    runs-on: self-hosted
```

## License
This project is licensed under the terms of the [MIT license](LICENSE).
