FROM python:3.10-slim

WORKDIR /app

RUN pip install --no-cache-dir --upgrade pip

CMD ["sh", "-c", "python main.py | awk '{ print strftime(\"[%Y-%m-%d %H:%M:%S]\"), $0 }' | tee -a /logs/python.log"]

EXPOSE 8082
