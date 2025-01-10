# golang-memcached-example

- SQL code to add a table
```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short_code VARCHAR(255) NOT NULL UNIQUE,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

- Running Memcached with docker
```bash
docker run -d --name memcached -p 11211:11211 memcached:latest memcached -m 128
```
