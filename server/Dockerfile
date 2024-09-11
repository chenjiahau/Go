# Official Golang image (You shouldn't use the `latest` version in production but I'm a bad guy)
FROM golang:latest

# Arguments
ARG APP_NAME
ARG APP_VERSION
ARG DB_HOST
ARG DB_USER
ARG DB_PASSWORD

# Working directory
WORKDIR /app

# Copy everything at /app
COPY . /app

# Build the go app
RUN go build -o main .

# Expose port
EXPOSE 8080

# Environment variables
ENV APP_NAME=${APP_NAME}
ENV APP_VERSION=${APP_VERSION}
ENV DB_HOST=${DB_HOST}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}

# Define the command to run the app
CMD ./main ${APP_NAME} ${APP_VERSION} ${DB_HOST} ${DB_USER} ${DB_PASSWORD}

