
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

POST:-
````
curl --location --request POST 'http://localhost:8080/api/v1/comment' -v \
--header 'Content-Type: application/json' \
--data-raw '{"slug": "hello", "body": "body", "author": "me"}'
````

GET:-
````
curl --location --request GET 'http://localhost:8080/api/v1/comment/{id}'
````

PUT:-
````
curl --location --request PUT 'http://localhost:8080/api/v1/comment/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{"slug": "hello-update", "body": "body", "author": "me"}'
````

DELETE:-
````
curl --location --request DELETE 'http://localhost:8080/api/v1/comment/{id}'
````
### CRUD Queries with JWT Token (created on https://jwt.io/, using "mysigningkey" as Signing-Key):- ###
POST
````
curl --location --request POST 'http://localhost:8080/api/v1/comment' -v \
--header 'Content-Type: application/json' \
--header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.r0NrgliWg1MbKxWsM9pGhxeuAC_9Ctw5Lb4ttnQDNKg' \
--data-raw '{"slug": "hello", "body": "body", "author": "me"}'
````

GET:-
````
curl --location --request GET 'http://localhost:8080/api/v1/comment/{id}' \
--header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.r0NrgliWg1MbKxWsM9pGhxeuAC_9Ctw5Lb4ttnQDNKg'
````

PUT:- 
````
curl --location --request PUT 'http://localhost:8080/api/v1/comment/{id}' \
--header 'Content-Type: application/json' \
--header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.r0NrgliWg1MbKxWsM9pGhxeuAC_9Ctw5Lb4ttnQDNKg' \
--data-raw '{"slug": "hello-update", "body": "body", "author": "me"}'
````

DELETE:- 
````
curl --location --request DELETE 'http://localhost:8080/api/v1/comment/{id}' \
--header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.r0NrgliWg1MbKxWsM9pGhxeuAC_9Ctw5Lb4ttnQDNKg'
````

