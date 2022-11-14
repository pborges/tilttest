FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD dbpass
ADD db/backup.sql /docker-entrypoint-initdb.d

EXPOSE 3306