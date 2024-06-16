# Building the binary of the App
FROM golang:1.22-alpine AS build

# `boilerplate` should be replaced with your project name
WORKDIR /go/src/build

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o fiber-api .

# Moving the binary to the 'final Image' to make it smaller
FROM scratch

COPY --from=build /go/src/build/fiber-api .

# Exposes port 8080 because our program listens on that port
EXPOSE 8080

ENTRYPOINT ["/fiber-api"]