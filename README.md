# gin-postgres-rest-service

REST API service, built with [Go](https://go.dev/),
[gin-gonic/gin](https://gin-gonic.com/) and
[Postgres](https://www.postgresql.org/), that can be used as a template for
creating other REST APIs.

## Goals

Started originally as a "take-home" assignment for a job interview.

## Create the database

`docker-compose up`

## Connect to the database

`psql -h localhost -p 5432 -U postgres -W`

## Create the tables and seed them with test data

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS product (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    version_count INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_name on product(name);

CREATE TABLE IF NOT EXISTS product_version (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  product_id UUID,
  version VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  CONSTRAINT fk_product
    FOREIGN KEY(product_id) 
    REFERENCES product(id)
);

CREATE INDEX IF NOT EXISTS idx_product_id on product_version(product_id);

INSERT INTO product (name, description, version_count)
VALUES
  ('product 1', 'product 1 does neat stuff', 1),
  ('product 2', 'product 2 does more neat stuff', 1),
  ('product 3', 'product 3 does even more neat stuff', 3);

INSERT INTO product_version (product_id, version, description)
VALUES
  ((SELECT id from product where name='product 1'), 'product 1 version 1', 'does neat stuff'),
  ((SELECT id from product where name='product 2'), 'product 2 version 1', 'does more neat stuff'),
  ((SELECT id from product where name='product 3'), 'product 3 version 1', 'does even more neat stuff'),
  ((SELECT id from product where name='product 3'), 'product 3 version 2', 'does even more neat stuff'),
  ((SELECT id from product where name='product 3'), 'product 3 version 3', 'does even more neat stuff');
```

## Build the service

```shell
go build
```

## Run the service

```shell
./gin-postgres-rest-service
```

## Run the tests

```shell
go test ./...
```

## Example requests for local development

```shell
curl --location 'http://localhost:8080/products' \
--header 'Content-Type: application/json'
```

```shell
curl --location 'http://localhost:8080/products?name=service&start=2&page_size=2' \
--header 'Content-Type: application/json'
```

```shell
curl --location 'http://localhost:8080/products/{:product_id}' \
--header 'Content-Type: application/json'
```

```shell
curl --location 'http://localhost:8080/products/{:product_id}' \
--header 'Content-Type: application/json'
```
