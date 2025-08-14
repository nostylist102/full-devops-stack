FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o full-devops-stack .

FROM gcr.io/distroless/static-debian12
COPY --from=build /src/full-devops-stack /full-devops-stack
EXPOSE 8080
ENTRYPOINT ["/full-devops-stack"]