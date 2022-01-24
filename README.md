# Urlify

Сервис сокращения URL

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