FROM golang:latest AS builder

COPY .env .

# Copy and download dependency using go mod
RUN go get github.com/joho/godotenv
RUN go get github.com/rs/cors
RUN go get github.com/gorilla/mux
RUN go get github.com/xfrr/goffmpeg/transcoder
RUN go get github.com/common-nighthawk/go-figure

# Set necessary environmet variables needed for our image
ENV CGO_ENABLED=0 \
    GOOS=linux 

# Move to working directory /build
RUN mkdir /build
ADD . /build/
WORKDIR /build

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
RUN mkdir /dist
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM scratch

COPY --from=builder /dist/main /

# Command to run
ENTRYPOINT ["/main"]