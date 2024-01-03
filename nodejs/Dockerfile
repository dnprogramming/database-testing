# Use the latest version of Node.js as the base image
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json files to the working directory
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the application code to the working directory
COPY . .

# Expose a port (if needed)
EXPOSE 8080

# Define the command to run the program
CMD [ "npm", "start" ]
