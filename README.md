
### Below are the Dependencies required to run this on your Local ###
1. Install [Golang](https://go.dev/doc/install)
2. Install [Docker](https://docs.docker.com/get-docker/)
3. Install [Task](https://taskfile.dev/installation/)
4. Install [golangci-lint](https://golangci-lint.run/usage/install/)

### Steps to follow for accessing your tables in Docker ###
1. `docker ps`
2. `docker exec -it 'CONTAINER ID' bash`
3. `psql -U postgres`
4. `\dt`
5. `\d+ comments`
