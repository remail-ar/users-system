docker run -v ../sql:/home/sql \
	--name postgres \
	-e POSTGRES_PASSWORD=password \
	-p 5432:5432 \
	-d docker.io/postgres:16.4-bookworm
