ARG NODEJS_VERSION
FROM node:${NODEJS_VERSION}-alpine

RUN apk add --update --no-cache git fish curl less build-base tzdata vim yarn

WORKDIR /app
VOLUME /app

USER node
CMD ["yarn", "run", "dev"]
