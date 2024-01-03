# Use the official Gradle image as the base image
FROM gradle:latest AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the build.gradle and settings.gradle files to the container
COPY build.gradle .
COPY settings.gradle .

# Copy the source code to the container
COPY src/ src/

# Build the Java app using Gradle
RUN gradle build

# Use a lightweight Java image as the base image
FROM openjdk:latest

# Set the working directory in the container
WORKDIR /app

# Copy the built JAR file from the builder stage to the container
COPY --from=builder /app/build/libs/myapp.jar .

EXPOSE 8080

# Set the command to run the Java app
CMD ["java", "-jar", "myapp.jar"]
