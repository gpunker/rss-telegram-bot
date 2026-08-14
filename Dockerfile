FROM golang:1.26
RUN useradd -m -u 1000 user
USER 1000:1000

WORKDIR /usr/src/app

# COPY go.mod go.sum ./
# RUN go mod download

COPY . .
#RUN go build -v -o /usr/local/bin/app ./...
