FROM louislam/uptime-kuma:latest

WORKDIR /app

CMD ["sh", "-c", "node server/server.js | tee -a /logs/uptimekuma.log"]

EXPOSE 3001
