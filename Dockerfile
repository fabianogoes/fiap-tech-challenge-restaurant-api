# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

#Expose ports
EXPOSE 8080

# Build the Go application
RUN go build -o web app/web/main.go

#Run the web usecases on container startup
CMD ["./web"]