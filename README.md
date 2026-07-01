# dog-go

A small Go project for serving and exploring friendly dog profiles from either a CLI or HTTP API.

## Requirements

- Go 1.22 or newer

## Usage

Print a random dog profile:

```bash
go run ./cmd/dog-go
```

Find a dog by name:

```bash
go run ./cmd/dog-go -name Mochi
```

Start the HTTP API:

```bash
go run ./cmd/dog-go -serve -addr :8080
```

Available endpoints:

- `GET /healthz`
- `GET /dogs`
- `GET /dogs/random`
- `GET /dogs/{name}`

## Development

```bash
go test ./...
```
