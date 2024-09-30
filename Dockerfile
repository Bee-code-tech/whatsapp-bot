# Use the official Golang image for building the app
FROM golang:1.19-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o whatsapp-bot ./cmd/main.go

# Expose the application port if needed
EXPOSE 8080

# Run the Go app
CMD ["./whatsapp-bot"]
