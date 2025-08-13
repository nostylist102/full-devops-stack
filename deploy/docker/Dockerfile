#билд
FROM golang:1.22-alpine as build
WORKDIR /src
COPY go.mod go.sum ./
RUN go.mod download
COPY ..
RUN CGO_ENABLED=0 go build -0 full-devops-stack .




#рантайм
FROM grc.io/distoroless/static-debian12
COPY --from=build /src/full-devops-stack /full-devops-stack
EXPOSE 8080
ENTRYPOINT ["/full-devops-stack"]