FROM golang AS base
WORKDIR /app
COPY . . 

FROM base AS development
RUN go install github.com/cosmtrek/air@latest
COPY go.* ./
RUN go mod download
CMD ["air", "-c", ".air.toml"]

FROM base AS builder
WORKDIR /build
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app cmd/*.go

FROM alpine:latest AS production
EXPOSE 8080
WORKDIR /prod
COPY --from=builder /build/app .
CMD [ "./app" ]