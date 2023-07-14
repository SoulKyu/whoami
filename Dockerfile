# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20.5-alpine3.18 AS build-stage

# Définit le répertoire de travail dans le conteneur
WORKDIR /app

# Copie les fichiers de code source dans le conteneur
COPY . .

RUN apk add --update-cache gcc musl-dev

# Compile l'application Go
RUN go build -o whoami

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

WORKDIR /app

RUN apk add --update-cache gcc musl-dev

COPY --from=build-stage /app/whoami /app/whoami
COPY --from=build-stage /app/templates /app/templates
COPY --from=build-stage /app/static /app/static

#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

EXPOSE 8080

ENTRYPOINT ["/app/whoami"]