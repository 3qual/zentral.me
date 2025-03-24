FROM postgres:latest


RUN mkdir -p /var/log/postgresql && \
    chown -R postgres:postgres /var/log/postgresql && \
    chmod 700 /var/log/postgresql


COPY ./config/postgresql.conf /etc/postgresql/postgresql.conf

RUN chown -R postgres:postgres /etc/postgresql && \
    chmod 600 /etc/postgresql/postgresql.conf

CMD ["postgres", "-c", "config_file=/etc/postgresql/postgresql.conf"]

EXPOSE 5432
