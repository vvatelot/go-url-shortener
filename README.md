# Go Url Shortener

This project aims to provide a simple solution to create short urls.

It is built in Golang and Fiber to provide great performance and be as light as possible and can be deployed using Docker or directly on your server.

## 🛠️ Tech Stack

- [Docker](https://www.docker.com/)
- [Docker Compose v2](https://docs.docker.com/compose/compose-v2/)
- [Golang](https://go.dev/)
- [Fiber](https://gofiber.io/)
- [Air](https://github.com/cosmtrek/air) (live reload)

## 🛠️ Install Dependencies

```bash
go mod download
```

## 🧑🏻‍💻 Usage

To start the project, you first need configure your `.env` file. You can use the `.env.dist` file as a template:

```bash
cp .env.dist .env
```

Then you can launch your project simply using air command:

```bash
air
```

You can also use docker compose to launch the project:

```bash
cp docker-compose.yml.dist docker-compose.yml
docker-compose up -d --build && docker-compose logs -f
```

> You can now reach your application instance on [http://localhost:3001](http://localhost:3001) depending on your configuration.

## 🔧 Configuration

### Environment variables

| Name             | Description                                      | Default value      |
| ---------------- | ------------------------------------------------ | ------------------ |
| `APP_PORT`       | Default application listen port                  | `3001`             |
| `DATABASE_URI`   | Database connection string                       | `url_shortener.db` |
| `ENV`            | Application environment (can be `dev` or `prod`) | `dev`              |
| `ADMIN_PASSWORD` | Admin password to access the admin panel         | `password`         |
| `LOCALE`         | Application locale                               | `fr`               |



> When using docker, you must set the environment variables in the `docker-compose-override.yml` file. You can override the default values of the `docker-compose.yml` file.


## [License](LICENSE)

## [Contributing](CONTRIBUTING.md)

## [Code of conduct](CODE_OF_CONDUCT.md)
