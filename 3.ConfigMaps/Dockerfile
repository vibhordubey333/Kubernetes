FROM golang:1.18-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum ./cmd/ReadEnv.go /build/
WORKDIR /build
RUN go build ReadEnv.go

FROM alpine
WORKDIR /build
RUN adduser -S -D -H -h /app appuser
USER appuser
RUN ls /build/
COPY --from=builder /build/ReadEnv /app/
#COPY views/ /app/views
WORKDIR /app
CMD ["./ReadEnv"]
# ---
# FROM golang:1.18-alpine
# RUN mkdir /build
# WORKDIR /build
# COPY go.mod go.sum ./cmd/ReadEnv.go /build/
# CMD ["./ReadEnv"]
