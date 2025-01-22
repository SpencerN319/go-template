FROM golang:1.23 AS build-stage

ARG TARGET

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app -ldflags '-w -s' /usr/src/app/cmd/${TARGET}


# release image
FROM gcr.io/distroless/base-debian12:nonroot AS release-stage

WORKDIR /

COPY --from=build-stage /usr/local/bin/app /app

USER nonroot:nonroot

ENTRYPOINT ["/app"]
