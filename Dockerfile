FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /template -ldflags '-w -s' /app/cmd/app


# release image
FROM gcr.io/distroless/base-debian12 AS release-stage

WORKDIR /

COPY --from=build-stage /template /template

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/template"]
