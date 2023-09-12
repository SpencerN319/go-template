FROM golang:1.21-alpine

RUN addgroup -S nonroot && \
    adduser -S nonroot -G nonroot

WORKDIR /usr/src/app

COPY --chown=nonroot:nonroot go.* .
RUN go mod download && go mod verify

COPY --chown=nonroot:nonroot . .
RUN CGO_ENABLED=0 go build -o /usr/local/bin/app -ldflags '-w -s' .

USER nonroot

CMD ["app"]
