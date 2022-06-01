DB_URL=postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable
UP:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up
UPP:
	migrate -path db/migrations -database "postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable" -verbose up
INIT:
	migrate create -ext sql -dir db/migrations -seq business_file
# 编译 会编译多个版本
GORE:
	goreleaser release --snapshot --rm-dist
# 生成docker镜像
DockerB:
	docker build -t demo_docker:v1.0.1 .
# 运行docker镜像
DockerRUN:
	docker run -d --name demo_docker_container -p 8008:8008 demo_docker:v1.0.1