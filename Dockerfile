# step 1: build go app
FROM golang:1.13.5-alpine3.11 as build-step

# for go mod download
RUN apk add --update --no-cache ca-certificates git

RUN mkdir /go-app
WORKDIR /go-app
COPY backend/go.mod .
COPY backend/go.sum .

RUN go mod download
COPY backend .

RUN cd server && CGO_ENABLED=0 go build -o /go/bin/go-app

# -----------------------------------------------------------------------------
# step 2: exec
FROM scratch
COPY --from=build-step /go/bin/go-app /go/bin/go-app
EXPOSE 5050
ENTRYPOINT ["/go/bin/go-app"]