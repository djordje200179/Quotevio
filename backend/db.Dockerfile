FROM mysql

ADD ./db.sql /docker-entrypoint-initdb.d

EXPOSE 3306