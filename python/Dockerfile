# Use the latest Python Docker image
FROM python:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the requirements file to the container
COPY requirements.txt .

# Install the project dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the entire project directory to the container
COPY . .

# Set the entry point command to run the Python project
CMD ["python", "main.py"]
