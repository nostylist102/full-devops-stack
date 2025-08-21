FROM golang:1.22-alpine AS build
WORKDIR /src
COPY app/go.mod app/go.sum ./
RUN go mod download
COPY app/. ./
RUN CGO_ENABLED=0 go build -o full-devops-stack .

FROM alpine:latest
RUN apk --no-cache add net-tools
COPY --from=build /src/full-devops-stack /full-devops-stack
EXPOSE 8080
EXPOSE 8081
ENTRYPOINT ["/full-devops-stack"]