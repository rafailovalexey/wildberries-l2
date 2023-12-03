### Test task for "Wildberries-L2-Dev11"

```bash
docker compose --file docker-compose-postgresql.yml up --detach --build
```

```bash
docker compose --file docker-compose-microservices.yml up --detach --build
```

```bash
curl \
    --location '192.168.50.100:3000/v1/events/events_for_day?user_id=bdfff233-ada6-41b5-954e-7b81e8f8813f&date=2019-01-01' \
    --header 'authentication-token: token'
```

```bash
curl \
    --location '192.168.50.100:3000/v1/events/events_for_week?user_id=bdfff233-ada6-41b5-954e-7b81e8f8813f&date=2019-01-01' \
    --header 'authentication-token: token'
```

```bash
curl \
    --location '192.168.50.100:3000/v1/events/events_for_month?user_id=bdfff233-ada6-41b5-954e-7b81e8f8813f&date=2019-01-01' \
    --header 'authentication-token: token'
```

```bash
curl \
    --location '192.168.50.100:3000/v1/events/create_event' \
    --header 'Content-Type: application/json' \
    --header 'authentication-token: token' \
    --data '{
        "user_id": "bdfff233-ada6-41b5-954e-7b81e8f8813f",
        "date": "2019-01-01"
    }'
```

```bash
curl \
    --location '192.168.50.100:3000/v1/events/update_event' \
    --header 'Content-Type: application/json' \
    --header 'authentication-token: token' \
    --data '{
        "id": 1,
        "user_id": "bdfff233-ada6-41b5-954e-7b81e8f8813f",
        "date": "2017-01-01"
    }'
```

