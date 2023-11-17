# Calcserv

A microservice that performs the following calculations to a list of prices:
- `total`
- `minimum`
- `maximum`
- `average`

## Getting started

To start this service, execute the binary file `calcserv` located within this project's root directory.

```bash
./calcserv
```

The service will run on port `3000` by default. To use a different port, pass a `PORT` environment variable when running the binary.

```bash
PORT=4000 ./calcserv
```

To check if the service is running properly, call the `/health` endpoint.

```bash
curl http://localhost:3000/health
```

## Endpoints

### `POST /calculate`

Returns the total, minimum, maximum, and average of a list of prices. Rounds all values to two decimal places.

#### Request body schema

- `data`: An array of float numbers. It must contain at least one item. The items must be greater than 0.0.

##### Example
```json
{
    "data": [
        2.33,
        0.22,
        10,
        5.1
    ]
}
```

#### Response schema

- `total`: Total sum of all values
- `min`: Minimum value in array
- `max`: Maximum value in array
- `avg`: Average of all values

##### Example
```json
{
    "total": 17.65,
    "min": 0.22,
    "max": 10,
    "avg": 4.41
}
```
