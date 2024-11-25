FROM cimg/go:1.22.5
WORKDIR /app
COPY . .
CMD ["go", "run", "/app/."] >> /dev/stdout 2>&1