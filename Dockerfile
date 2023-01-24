FROM golang:1.18

# Set the working directory in the container
WORKDIR /go/src/sim-u

# Copy the local source code into the container
COPY . /go/src/sim-u

# Build the project
RUN go build -o main .

# Run the migrations
# RUN go run main.go migrate

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary
CMD ["./main server"]