# Application for talking to cow
i

## Development

```
cp example.env .env
```

## Install dependecies

```
go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-psql@latest
go get github.com/aarondl/sqlboiler/v4
```

## Run migrations

```
go run scripts/migrate/migrate.go
```

## Database

Start database with docker compose
```
docker compose up
```
## Backend e2e tests

The backend e2e suite boots the Gin router in-process and drives real HTTP requests against the configured Postgres database.
It is opt-in because it creates temporary users/files/diary entries in the configured DB.

```
TALKTOCOW_E2E=1 go test . -run TestBackendE2E
```
