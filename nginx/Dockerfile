# Official Nginx image (Yes, in this article I always use the `latest`. Kill me!)
FROM nginx:latest

# Copy Nginx configuration file to the container
COPY nginx.conf /etc/nginx/conf.d

# Copy Verify Domain Configuration file to the container
RUN mkdir -p /usr/share/nginx/html
COPY www /usr/share/nginx/html
RUN chown -R www-data:www-data /usr/share/nginx/html/.well-known/

# Expose port 80
EXPOSE 80
