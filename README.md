## Simple User App
Simple user application which demonstrates how user registration works. This app also provides user log in and log out mechanism and all its limitation to access several endpoints.

## Requirement

- [Go](https://golang.org/)
- [PotgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)

## Installation

1. Clone this repository.

```sh
git clone https://github.com/joshiaj7/simple-user-app.git
```

2. Navigate to the project root directory.

```sh
cd simple-user-app
```

3. Run the app.

Using docker compose is recommended
```sh
go mod vendor
docker-compose up --build
```

Alternatively, you can just run the main.go
```sh
go run main.go
```

## How to Use

Before we start, please refer to the following available endpoints in this application:

|  Endpoint  | Method | Request Body Required | Log In Required | Description                                |
| ---------- | ------ | --------------------- | --------------- | ------------------------------------------ |
| /create    | POST   | Yes                   | No              | Register new user                          |
| /get       | GET    | No                    | Yes             | Get a list of existing users               |
| /update    | PUT    | Yes                   | Yes             | Update user data by its user id            |
| /delete    | DELETE | Yes                   | Yes             | Delete a user by its user id               |
| /login     | POST   | Yes                   | No              | Log in to app using user name and password |
| /logout    | POST   | No                    | Yes             | Log out from app                           |


1. Create a new user

Create request body example:
```json
{
    "email": "example@example.com",
    "user_name": "admin",
    "address": "Indonesia",
    "password": "password123"
}
```

2. Use your uuid as bearer token

You can find your uuid from the response once you created a new user. We use uuid as bearer token to show that you are currently logged in.

3. Try other APIs!
Try `/update`, `/delete`, and `/get`

Update request body example:
```json
{
    "user_id": 1,
    "user_name": "not_admin",
    "email": "my_email@example.com",
    "address": "Singapore",
    "password": "mypassword"
}
```

Delete request body example:
```json
{
    "user_id": 2
}
```

4. After you finished, log out from the app
Once you logged out, you can't perform update, delete, and get user. Please log in if you want to perform these actions.

Log in request body example:
```json
{
    "user_name": "not_admin",
    "password": "mypassword"
}
```

## Details

For more details, visit my [github page](https://github.com/joshiaj7/simple-user-app)