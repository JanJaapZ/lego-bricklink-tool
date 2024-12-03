FROM python

# Set the working directory
WORKDIR /code

# Copy the application code into the container
COPY code/ /code

# Copy and install dependencies
COPY requirements.txt /code/requirements.txt
RUN apt-get update && apt-get install -y python3-tk
RUN pip install --no-cache-dir -r requirements.txt
RUN pip install --no-cache-dir Pillow
RUN pip install flask

EXPOSE 5000

ENTRYPOINT ["python lego.py"]