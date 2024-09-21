# Official Golang image (You shouldn't use the `latest` version in production but I'm a bad guy)
FROM golang:latest

# Arguments
ARG DOMAIN
ARG APP_NAME
ARG APP_VERSION
ARG DB_HOST
ARG DB_NAME
ARG DB_USER
ARG DB_PASSWORD
ARG EMAIL_HOST
ARG EMAIL_PORT
ARG EMAIL_USER
ARG EMAIL_PASSWORD

# Working directory
WORKDIR /app

# Copy everything at /app
COPY . /app

# Build the go app
RUN go build -o main .

# Expose port
EXPOSE 8080

# Environment variables
ENV DOMAIN=${DOMAIN}
ENV APP_NAME=${APP_NAME}
ENV APP_VERSION=${APP_VERSION}
ENV DB_HOST=${DB_HOST}
ENV DB_NAME=${DB_NAME}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV EMAIL_HOST=${EMAIL_HOST}
ENV EMAIL_PORT=${EMAIL_PORT}
ENV EMAIL_USER=${EMAIL_USER}
ENV EMAIL_PASSWORD=${EMAIL_PASSWORD}

# Define the command to run the app
CMD ./main ${DOMAIN} ${APP_NAME} ${APP_VERSION} ${DB_HOST} ${DB_NAME} ${DB_USER} ${DB_PASSWORD} ${EMAIL_HOST} ${EMAIL_PORT} ${EMAIL_USER} ${EMAIL_PASSWORD}