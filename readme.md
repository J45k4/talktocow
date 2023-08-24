# Application for talking to cow
i

## Development

```
cp example.env .env
```

## Install dependecies

```
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go get github.com/volatiletech/sqlboiler/v4
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