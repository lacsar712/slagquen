# slagquen

Blast-furnace slag water-quench granulation loop. No web UI.

## Requirements

- Go 1.22+ (container image uses golang:1.22)
- `GOTOOLCHAIN=local` recommended on host when using a pinned toolchain

## Build

```bash
export GOTOOLCHAIN=local
go build ./...
```

## Run

```bash
export GOTOOLCHAIN=local
go run ./cmd/graindry
```

## Test

```bash
export GOTOOLCHAIN=local
go test ./... -count=1
```

## Docker (benzhi)

Must build **linux/amd64** and **linux/arm64**:

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh slagquen linux/amd64
./build_benzhi_docker.sh slagquen linux/arm64
docker run -it slagquen:latest
# inside container:
export GOTOOLCHAIN=local
go version
go build ./...
go test ./... -count=1
```
