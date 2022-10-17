FROM golang:buster AS build

WORKDIR /app
ADD ./ ./
RUN go mod download
RUN go build -o /url-shortener


FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /url-shortener /url-shortener
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/url-shortener"]
