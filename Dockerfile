# BUILD
FROM golang:1.23.3-alpine3.19 AS build

# Update packages and clear cache
RUN apk update && apk add --no-cache git curl && rm -rf /var/cache/apk/*
# Set work directory
WORKDIR /app
# Create binary directory
RUN mkdir /app/bin -p
# Create golang migrate util directory
RUN mkdir /bin/golang-migrate -p
# Add migrate tool
ADD ./tools/migrate-unix /bin/golang-migrate/migrate
RUN chmod +x /bin/golang-migrate/migrate
# Add main files to app
ADD . .

RUN printenv

# Download go depences
RUN go mod download
# Build app
RUN GOOS=linux go build -o bin ./...


# APP
FROM alpine:3.19 AS app

# Install packages
RUN apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*
# Create home directory
WORKDIR /app
# Copy build file
COPY --from=build /app/bin/app ./app
# Copy migration dir
COPY --from=build /app/migrations ./migrations
# Install migrate tool
COPY --from=build /bin/golang-migrate /usr/local/bin
# CMD
CMD ["./app"]


