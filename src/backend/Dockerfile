# step 1: build go app
FROM golang:1.13.5-alpine3.11 as build-step

# for go mod download
RUN apk add --update --no-cache ca-certificates git

# in container
RUN mkdir /go-app
WORKDIR /go-app

# from host-pc to container
COPY go.mod .
COPY go.sum .

# in container
RUN go mod download

# from host-pc to container
COPY . .
RUN pwd

# in container
RUN cd server && CGO_ENABLED=0 go build -o /go/bin/go-app

# -----------------------------------------------------------------------------
# step 2: exec
FROM scratch
COPY --from=build-step /go/bin/go-app /go/bin/go-app
EXPOSE 5050
ENTRYPOINT ["/go/bin/go-app"]