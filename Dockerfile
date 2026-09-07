FROM golang:1.25-bookworm AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=1 go build -buildvcs=false -o /url-shortener


FROM gcr.io/distroless/base-debian12

WORKDIR /
COPY --from=build /url-shortener /url-shortener
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/url-shortener"]
