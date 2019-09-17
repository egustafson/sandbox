FROM python:3

WORKDIR /app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

ENV DOCKER_TURD="Trinket from Dockerfile"

CMD [ "python", "./app.py" ]
