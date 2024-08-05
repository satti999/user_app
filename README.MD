# User App

A simple user app demostrating the CRUD operation of user using fiber gorm and postgress.

## API Reference

#### Create user

```http
  POST /user
```

| Parameter | Type   | Description            |
| :-------- | :----- | :--------------------- |
|           | `POST` | Create the user record |

#### login

```http
  POST /login
```

| Parameter | Type   | Description                       |
| :-------- | :----- | :-------------------------------- |
|           | `POST` | It verfiy the user and then login |

#### Get user

```http
  GET /user/${id}
```

| Parameter | Type   | Description                                   |
| :-------- | :----- | :-------------------------------------------- |
| `id`      | `POST` | It get the single user record base on user Id |

```http
  GET /user/
```

| Parameter | Type     | Description                   |
| :-------- | :------- | :---------------------------- |
|           | `string` | It fetch all the users record |

#### Update user

```http
  PUT /user/${id}
```

| Parameter | Type  | Description                               |
| :-------- | :---- | :---------------------------------------- |
| `id`      | `PUT` | It Update the user record base on user Id |

#### Delete user

```http
  DELETE /user/${id}
```

| Parameter | Type     | Description                               |
| :-------- | :------- | :---------------------------------------- |
| `id`      | `DELETE` | It delete the user record base on user Id |

## Run Locally

Clone the project

```cmd
  git clone https://github.com/satti999/user_app.git
```

Go to the project directory

```cmd
  cd user_app
```

Install dependencies

```cmd
  go mod tidy
```

Start the server

```cmd
  go run main.go
```
