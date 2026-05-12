FROM node:24-alpine AS frontend-builder

WORKDIR /frontend

COPY frontend/package.json frontend/package-lock.json ./

RUN npm ci

COPY frontend/index.html frontend/tsconfig.json frontend/vite.config.ts frontend/eslint.config.mjs ./
COPY frontend/pages pages
COPY frontend/src src
COPY frontend/public public
COPY frontend/styles styles

RUN npm run build

FROM golang:1.26.2 AS backend-builder

WORKDIR /source

ADD go.mod go.mod
ADD go.sum go.sum

RUN go mod download

ADD main.go main.go
COPY auth auth
COPY config config
COPY models models
COPY chatroom chatroom
COPY migrations migrations
COPY routes routes
COPY eventbus eventbus
COPY bot bot

RUN CGO_ENABLED=0 go build -o /bin/talktocow

FROM debian:trixie-slim

WORKDIR /app

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates imagemagick libheif1 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=backend-builder /bin/talktocow /app/talktocow
COPY --from=backend-builder /source/migrations /app/migrations
COPY --from=frontend-builder /frontend/dist /app/frontend/dist

ENV FILE_STORAGE_PATH=/data/files
ENV FRONTEND_DIST_PATH=/app/frontend/dist
VOLUME ["/data/files"]

CMD ["/app/talktocow"]
