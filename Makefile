postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
	# docker exec -it postgres12 /bin/sh # 进入postgres12容器
	# docker exec -it postgres12 psql -U root # 进入postgres12容器的SQL CLi
	# docker exec -it postgres12 psql -U root -d simple_bank # 进入 postgres12 容器的 simple_bank 数据库

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema
	# 别忘记把 sql 文件内容拷贝到 db/migration/*.up.sql 里面为 migrateup 做数据准备
	# migrate create -ext sql -dir db/migration -seq add_users

migrateup:
	migrate -path=db/migration -database=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable -verbose up
migrateup1:
	migrate -path=db/migration -database=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable -verbose up 1

migratedown:
	migrate -path=db/migration -database=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable -verbose down
migratedown1:
	migrate -path=db/migration -database=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination  db/mock/store.go techschool/samplebank/db/sqlc Store

image:
	docker build -t simplebank:latest .

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1 image
# 伪目标 PHONY 的作用：当 make xxx 时候，如果存在可以运行的 xxx 的命令，那 Makefile 中的 xxx 将不会被运行，加上 PHONY 就可以运行 Makefile 中的指令。