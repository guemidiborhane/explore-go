FROM golang:1.21-alpine AS base-build

WORKDIR /app
COPY . .
RUN go mod download \
  && go build -o /app/tmp/main

FROM alpine:3.18

COPY --from=base-build /app/tmp/main /app
CMD ["/app"]

