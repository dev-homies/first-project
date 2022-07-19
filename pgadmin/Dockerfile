FROM dpage/pgadmin4:6.11

COPY servers.json /pgadmin4/servers.json
COPY pgpass /pgadmin4/pgpass

RUN cp /pgadmin4/pgpass /var/lib/pgadmin/pgpass
RUN chmod 600 /var/lib/pgadmin/pgpass

EXPOSE 80
