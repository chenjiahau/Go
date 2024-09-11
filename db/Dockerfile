FROM postgres:12.4-alpine

# Arguments
ARG DB_NAME
ARG DB_USER
ARG DB_PASSWORD

# Environment variables
ENV POSTGRES_DB=${DB_NAME}
ENV POSTGRES_USER=${DB_USER}
ENV POSTGRES_PASSWORD=${DB_PASSWORD}

# Volume
VOLUME ["/var/lib/postgresql/data"]

# Expose port
EXPOSE 5432
