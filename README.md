# Urlify

Сервис сокращения URL

## Routes

| HTTP method | Route           |  Description       |
|-------------|-----------------|--------------------|
| POST        | /api/url/create | Create short link  |
| GET         | /api/url/<hash> | Get link from hash |

## Migrations

### Make

```shell
docker-compose run migrator create -dir /opt/app/urlify/db/migrations -ext sql -seq create_references_table
```