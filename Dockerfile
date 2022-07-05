FROM golang:buster AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD ./ ./
RUN go build -o /url-shortener


FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /url-shortener /url-shortener
ADD public/ /public/
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/url-shortener"]
