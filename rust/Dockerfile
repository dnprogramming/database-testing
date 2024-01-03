# Use the official Rust image as the base image
FROM rust:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Cargo.toml and Cargo.lock files to the container
COPY Cargo.toml Cargo.lock ./

# Build the dependencies separately to leverage Docker layer caching
RUN cargo build --release

# Copy the source code to the container
COPY src ./src

# Build the application
RUN cargo build --release

EXPOSE 8080

# Set the entry point for the container
CMD ["./target/release/rust"]
