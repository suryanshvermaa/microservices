FROM postgres:15.4

COPY ./services/order/up.sql /docker-entrypoint-initdb.d/1.sql

EXPOSE 5432
CMD ["postgres"]