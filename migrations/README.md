# Migrations

## Create migrations
```
migrate create -ext sql -dir db/migrations -seq create_users_table
```

## Run migrations
```
migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS up
```

## Forcing your database version
```
migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force VERSION
```

## Source

https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md