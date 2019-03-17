FROM postgres:10.3

## copy sql file into the container 
COPY cars-table.sql /docker-entrypoint-initdb.d/1.sql

## NOTE: SQL files will be executed in alphabetical order
CMD ["postgres"]

# Bug at https://github.com/docker-library/postgres/issues/203#issuecomment-255200501