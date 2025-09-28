FROM postgres:15.4

COPY ./services/account/up.sql /docker-entrypoint-initdb.d/1.sql

EXPOSE 5432
CMD ["postgres"]