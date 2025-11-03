start:
  mysql -u root --socket ./data/mysql.sock

init:
  mysqld --initialize-insecure --datadir ./data

sql-server:
  mysqld --datadir ./data --socket ./mysql.sock &

stop:
  mysqladmin -u root --socket ./data/mysql.sock shutdown

go-server:
  cd backend && go build -o main *.go && ./main

dev:
  cd frontend && npm run dev

run:
  just go-server &
  just dev
