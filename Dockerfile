# Stage 1: Install Golang Image
FROM golang:1.22.4-alpine3.20

WORKDIR /app

# Copy the source code into the container
COPY . .

# Download all dependencies
RUN go mod download

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Expose port 10000
EXPOSE ${SERVER_PORT}

# Command to run the Server
CMD ["./main", "run", ".", "s"]