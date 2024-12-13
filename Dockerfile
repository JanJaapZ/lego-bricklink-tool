FROM golang:1.20-alpine as builder

# Set the working directory
WORKDIR /app

# Copy the application code into the container
COPY ./code /app

# Copy and install dependencies
RUN go mod init lego && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o lego lego.go
RUN ls -la /app

FROM golang:1.20-alpine

RUN apk add --no-cache ca-certificates
# Set the working directory
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/lego /app/lego

RUN chmod +x /app/lego

RUN ls -ls /app

# Expose the port the app runs on
EXPOSE 5000

# Command to run the application
CMD ["/app/lego"]