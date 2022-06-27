
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

### Sample Curl CRUD Queries ###
````
POST:-
curl --location --request POST 'http://localhost:8080/api/v1/comment' \
--header 'Content-Type: application/json' \
--data-raw '{"slug": "hello", "body": "body", "author": "me"}'
````

````
GET:-
curl --location --request GET 'http://localhost:8080/api/v1/comment/{id}'
````

````
PUT:- 
curl --location --request PUT 'http://localhost:8080/api/v1/comment/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{"slug": "hello-update", "body": "body", "author": "me"}'
````

````
DELETE:- 
curl --location --request DELETE 'http://localhost:8080/api/v1/comment/{id}'
````