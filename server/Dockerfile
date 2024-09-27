# Official Golang image (You shouldn't use the `latest` version in production but I'm a bad guy)
FROM golang:latest

# Working directory
WORKDIR /app

# Copy everything at /app
COPY . /app

# Build the go app
RUN go build -o main .

# Copy .env at /app again
COPY .env /app

# Expose port
EXPOSE 8080

# Define the command to run the app
CMD ./main
