# golang-memcached-example

- Javdal qo'shish
```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short_code VARCHAR(255) NOT NULL UNIQUE,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

- Memcached-ni docker bilan ishga tushirish
```bash
docker run -d --name memcached -p 11211:11211 memcached:latest memcached -m 128
```
