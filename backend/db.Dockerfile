FROM mysql

ARG DB_PASSWORD

ENV MYSQL_ROOT_PASSWORD=$DB_PASSWORD
ENV MYSQL_DATABASE=quotes

ADD ./db.sql /docker-entrypoint-initdb.d

HEALTHCHECK --interval=3s --timeout=3s --retries=10 \
  CMD mysqladmin ping -uroot -p$MYSQL_ROOT_PASSWORD

EXPOSE 3306