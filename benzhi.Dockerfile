FROM golang:1.22 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/slagquen ./cmd/slagquen

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/slagquen /slagquen
ENTRYPOINT ["/slagquen"]
