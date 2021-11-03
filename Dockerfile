FROM golang:1.16-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o comidit-app .
FROM scratch
COPY --from=builder /build/comidit-app /app/
WORKDIR /app
CMD ["./comidit-app"]
