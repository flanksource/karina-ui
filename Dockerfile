
FROM golang:1.13.6 as builder
WORKDIR /app
COPY pkg/ /app
COPY main.go /app
COPY go.mod /app
ARG NAME
ARG VERSION
# RUN make static
RUN GOOS=linux GOARCH=amd64 go build -o server -ldflags "-X \"main.version=$VERSION\""  main.go

FROM node:lts-alpine as ui
WORKDIR /app
COPY package.json .
RUN npm install
RUN npm install @vue/cli -g
COPY . .
RUN npm run build

FROM ubuntu:bionic
WORKDIR /app
COPY --from=builder /app/server /app
COPY --from=ui /app/dist /app/dist
ENTRYPOINT ["/app/server"]
