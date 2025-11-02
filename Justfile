start:
  mysql -u root --socket ./data/mysql.sock

init:
  mysqld --initialize-insecure --datadir ./data

server:
  mysqld --datadir ./data --socket ./mysql.sock &

stop:
  mysqladmin -u root --socket ./data/mysql.sock shutdown;
