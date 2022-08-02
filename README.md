# go-sysinfo-api

Small API to get basic system information

## Setup
copy `.env.example` as `.env` and change the **PASSWORD**

```bash
# build the docker container
docker-compose build

# start docker container
docker-compose up
```

## Usage

| endpoint | function |
| -------- | -------- |
| /ping    | pong                   |
| /status  | get system status JSON (takes 2 seconds) |