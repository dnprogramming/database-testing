# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files to the working directory
COPY . .

# Build the Go project
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the entry point for the container
CMD ["./main"]
