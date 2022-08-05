# Example 3

This application aims to simulate a social network, only texts

## Check List

### Step 1

- [X] - CRUD users
- [X] - Follow another user
- [X] - Stop follow another user
- [X] - Search all users that follow
- [X] - Search all users that are followed
- [X] - Update password

### Step 2

- [X] - CRUD Publications
- [X] - Search for publications according to the users who follow
- [X] - Like
- [X] - Unlike


## Requirements
* Docker 20.10.8 (or superior)
* Docker-Compose 2.0.0 (or superior)
* SGBD (your preference)

** the required versions are the ones that were used

## installation instructions

1. In the project folder
```
docker-compose up -d
```
2. Connect with database 

a. Run the script that is in the [sql.sql](https://github.com/piovani/go_api/blob/master/example3/environment/sql.sql)

OR

b. Restore dump in the [dump.sql](https://github.com/piovani/go_api/blob/master/example3/environment/dump.sql)


## How To Test

1. In your [Postman](https://www.postman.com/home) or [Insomnia](https://insomnia.rest/) restore collection [Exemple3.postman_collection.json](https://github.com/piovani/go_api/blob/master/example3/environment/Exemple3.postman_collection.json)

2. Run
```
docker exec go-api-exemple-3-api go run main.go
```

3. Exec the routers

<b>ps:</br>
* if your port 5000 is in use just change the .env file to one that is free

* in the variables of the route collection it will also be necessary for the port of your choice