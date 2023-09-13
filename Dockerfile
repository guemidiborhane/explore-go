ARG GO_VERSION=1.21
ARG NODEJS_VERSION=18.12
FROM node:${NODEJS_VERSION}-alpine AS node

RUN apk add --no-cache yarn

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn install

COPY . .
RUN NODE_ENV=production yarn run build

ARG GO_VERSION=1.21
FROM golang:${GO_VERSION}-alpine AS base-build

WORKDIR /app
RUN apk add --no-cache dumb-init

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=node /app/build ./build
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch

ENV PORT=3000

COPY --from=base-build /go/bin/explore-go /application
COPY --from=base-build /usr/bin/dumb-init /usr/bin/init

CMD ["/application"]
EXPOSE 3000

ENTRYPOINT ["/usr/bin/init", "--"]
