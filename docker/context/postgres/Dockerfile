FROM postgres:latest

# Устанавливаем postgresql-contrib (включает пакет uuid-ossp, dblink, pgcrypto и другие)
RUN apt-get update && apt-get install -y --no-install-recommends postgresql-contrib && rm -rf /var/lib/apt/lists/*

# Копируем конфиг в контейнер
COPY ./config/postgresql.conf /etc/postgresql/postgresql.conf
RUN chown -R postgres:postgres /etc/postgresql && \
    chmod 600 /etc/postgresql/postgresql.conf

# Ставим переменные окружения (при необходимости)
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=mydb

# Копируем скрипты, которые будут выполнены при инициализации контейнера
COPY ./create_extensions.sql /docker-entrypoint-initdb.d/create_extensions.sql

# Открываем порт
EXPOSE 5432

# Запускаем Postgres с вашим конфигом
CMD ["postgres", "-c", "config_file=/etc/postgresql/postgresql.conf"]
