# WSO-Go Backend Demo
This is the backend demo for the proposed WSO backend rewrite, WSO-Go (proposal found [here](https://github.com/WilliamsStudentsOnline/wso-on-rails/wiki/Proposal:-WSO-Backend-Rewrite)).

## Running
To run the server, simply do `go run . -env=development`.

## API Endpoints
Get All Users:
```http request
GET localhost:8080/api/v1/user
```
Get User:
```http request
GET localhost:8080/api/v1/user/:user_id
```
Update User:
```http request
PUT localhost:8080/api/v1/user/:user_id
{
    "visible": true,
    "dorm_visible": true,
    "home_visible": true,
    "pronoun": "",
    "off_cycle": false
}
```
Authenticate/Login:
```http request
POST localhost:8080/api/v1/auth/login
{
    "unix_id": "admin",
    "password": "doesnt matter"
}
```
Refresh JWT Token:
```http request
GET localhost:8080/api/v1/auth/refresh_token
```

### Authentication Flow
We use something called a [JWT](jwt.io), or JSON Web Token for the API. This allows us to keep sessions and verify user identities without cookies or database queries. It works like this:
1. A user will request a token from the `auth/login` endpoint. They will pass in their login credentials, which will be checked with LDAP (not implemented yet).
1. If the user is verified, the server will then pull their user from the DB and create a payload. This payload will consist of the user's ID and the scopes the user is allowed (e.g. if the user is a senior, they can go to ephcatch; if the user is an admin, they can do other queries; if the user is not signed in but on school wifi, they can be read only).
1. The server will then take this payload and sign it with its secret key, before handing the JWT back to the user.
1. The user now can add the header `Authorization: Bearer <JWT GOES HERE>` to any request and be authenticated and allowed to access other API endpoints (like `user`)
1. The JWT has a one hour timeout (we can change this). After an hour, the JWT becomes invalid and the user must sign in again.
1. Alternatively, before the hour is up, a user can query the `auth/refresh_token` endpoint to get a new token without having to sign in again.

## Structure

- `config/` contains configurations for server
  - `auth_middleware.go` is the JWT authentication middleware that runs on _all_ API calls. Learn more about the authentication process above
  - `config.go` loads the environment config files into go
  - `database.go` loads the database connection/configuration
  - `scope_middleware.go` is the scope-based role authorization package running on specific API calls. Learn more above
  - `environment/*.yml` are configuration yaml files named by the environment it is run in
- `controllers/` contains all controllers tied to the server
  - `base_controller.go` every container inherits the base container properties; include global code for containers here
  - `user_controller.go` runs user-related requests. Currently, that's Get Users, Get User, and Update User
- `models/` contains database models
  - `base_model.go` every model inherits the base model; it contains important properties/funcs for all models
  - `department_schema.go` is the database schema for the departments table
  - `user.go` is the model for a user; it includes all database-side functions we run from controllers
  - `user_schema.go` is the database schema for the users table
- `test.db` the database of generated data the demo server uses
- `main.go` the entry-point of the code; contains all routing information