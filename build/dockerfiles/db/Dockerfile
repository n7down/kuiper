FROM mysql:8.0.18

ENV MYSQL_ROOT_PASSWORD=password

ADD schema.sql /docker-entrypoint-initdb.d

EXPOSE 3306
