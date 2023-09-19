FROM python:3.9

WORKDIR /app

COPY server/requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY server/src/ .

CMD ["python", "main.py"]