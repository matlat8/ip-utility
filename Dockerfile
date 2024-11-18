FROM golang:1.23

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY *.go ./

# Build the Go app
RUN go build main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]