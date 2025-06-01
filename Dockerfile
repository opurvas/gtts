# Start from official Go image
FROM golang:1.24

# Install Python & pip
RUN apt-get update && apt-get install -y python3 python3-pip

# Install gTTS
RUN pip3 install gTTS

# Set working directory
WORKDIR /app

# Copy project files
COPY . .

# Build Go app (optional, we’re running with go run for dev simplicity)
# RUN go build -o app

# Expose port
EXPOSE 9000

# Run the server
CMD ["go", "run", "main.go"]
