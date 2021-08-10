FROM python:3

WORKDIR /app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 5000/tcp
ENV FLASK_APP=flask_service
CMD ["flask", "run", "--host=0.0.0.0"]
