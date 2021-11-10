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
- [ ] - Like
- [ ] - Unlike


## Requirements
* Docker 20.10.8 (or superior)
* Docker-Compose 2.0.0 (or superior)
* SGBD (your preference)

** the required versions are the ones that were used

## installation instructions

1. in the project folder
```
docker-compose up -d
```
2. connect with database and run the script that is in the [sql.sql](https://github.com/piovani/go_api/blob/master/example3/sql/sql.sql)

## How To Test
