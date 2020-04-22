FROM mysql:8.0.18

ENV MYSQL_USER=root
ENV MYSQL_ROOT_PASSWORD=password

ADD schema.sql /docker-entrypoint-initdb.d/1.sql

EXPOSE 3306
