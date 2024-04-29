# Use the official golang image as base image
FROM golang:latest AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules and build files to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the executable file from the build stage to the current working directory in the container
COPY --from=build /app/main .

# Expose the port that the application listens on
EXPOSE 8888

# Command to run the executable
CMD ["./main"]
