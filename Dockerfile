# Use the official Ubuntu image
FROM docker.arvancloud.ir/ubuntu:latest

# Install necessary dependencies
RUN apt-get update && \
    apt-get install -y curl git && \
    rm -rf /var/lib/apt/lists/*

# Install Golang
RUN apt-get update && \
    apt-get install -y wget gnupg gnupg2 gnupg1 lsb-release && \
    wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz

# Set Go environment variables
ENV GOPATH=/go
ENV PATH=/usr/local/go/bin:$PATH
ENV PATH=$GOPATH/bin:$PATH

# Set the working directory inside the container
WORKDIR /app

# Copy init_db.sh script
COPY ./scripts/init_db.sh .

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the application
RUN go build -o sf_delivery_report .

# Expose port 8080
EXPOSE 8080

# Make init_db.sh executable
RUN chmod +x init_db.sh

# Make init_db.sh executable
RUN chmod +x ./main.go

# Command to run PostgreSQL initialization script and then run the application
CMD ["./init_db.sh"]
