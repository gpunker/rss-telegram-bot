FROM golang:1.26
ARG SERVER_PORT
ARG NODE_VERSION=26.7.0

RUN apt-get update && apt-get install -y --no-install-recommends \
        curl \
        ca-certificates \
        xz-utils

# установка nodejs и localtunel для имитации внешнего сервера
#   чтобы работали веб-хуки Telegram
RUN curl -fsSL "https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-x64.tar.xz" \
        -o /tmp/node.tar.xz \
    tar -xJf /tmp/node.tar.xz -C /usr/local --strip-components=1 \
    && rm -rf /tmp/node.tar.xz \
    && apt-get purge -y --auto-remove xz-utils \
    && rm -rf /var/lib/apt/lists/*
RUN npm install -g localtunnel

RUN useradd -m -u 1000 user
USER 1000:1000
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
#RUN go build -v -o /usr/local/bin/app ./...

EXPOSE $SERVER_PORT