## UCTest

## Sample Requests
### Post Invoices
#### Correct Cases
```shell
curl -XPOST -H 'content-type: application/json' localhost:8080/api/invoices -d '{"publisher": "0be62c928261", "receiver": "bdb2476fd525", "payment_amount": 10000}'
```

### Get Invoices
#### Correct Cases
```shell
curl -H 'content-type: application/json' localhost:8080/api/invoices\?period_start_at='2024-11-28'
```

```shell
curl -H 'content-type: application/json' localhost:8080/api/invoices\?period_end_at='2025-11-28'
```

```shell
curl -H 'content-type: application/json' localhost:8080/api/invoices\?period_start_at='2024-11-30'\&period_end_at='2024-12-31'
```

#### Validation Error Case
```shell
curl -H 'content-type: application/json' localhost:8080/api/invoices\?period_start_at='2024-11-30'\&period_end_at='2024-11-28'
```
