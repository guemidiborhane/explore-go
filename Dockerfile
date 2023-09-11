ARG GO_VERSION
ARG NODEJS_VERSION=18.12
FROM node:${NODEJS_VERSION}-alpine AS node

WORKDIR /app
RUN apk add --no-cache yarn

COPY . /app
RUN yarn install

RUN NODE_ENV=production yarn run build

ARG GO_VERSION=1.21
FROM golang:${GO_VERSION}-alpine AS base-build

WORKDIR /app
RUN apk add --no-cache dumb-init

COPY . /app
RUN go mod download

COPY --from=node /app/modules/ui/build /app/modules/ui/build
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch

ENV PORT=3000

COPY --from=base-build /go/bin/explore-go /application
COPY --from=base-build /usr/bin/dumb-init /usr/bin/init

CMD ["/application"]
EXPOSE 3000

ENTRYPOINT ["/usr/bin/init", "--"]
