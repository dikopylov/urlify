# Urlify

URL shortener

## How to run
```shell
cp example.env .env
make build
make up
```

## Commands
* Build: `make build`
* Run: `make up`
* Tests: `make test`

## Routes

| HTTP method | Route            | Description               |
|-------------|------------------|---------------------------|
| POST        | /api/url         | Create short link         |
| GET         | /api/<HASH_HERE> | Get link from hash        |
| GET         | /<HASH_HERE>     | Redirect to original link |

## Migrations

### Make

```shell
docker-compose run migrator create -dir /opt/app/urlify/database/migrations -ext sql -seq create_references_table
```
