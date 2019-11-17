# palang

Simple authentication service, to help create application faster. Main Idea is help developer not write again auth and user module when create new application.

## Feature

- Register
- Login
- Detail User
- Validate Token

## Project Structure

Basic idea for foldering we breakdown breakdown to 3 folder, Service, Manager and Delivery

- *Service*
in this folder contain logic and functions of each domain, and each domain must be single responsibility
- *Manager*
in this folder contain flow/usecase for apps (can be orchestration multiple services). example function to Register that will involve user servics and email services.
- *Delivery*
in this it's delivery protocol, so can be multiple delivery (grpc, rest API, Graphql)

## How To Run

### Http Delivery

Please go delivery/http and run main.go file `go run main.go`, don't forgot to change db config

```gotemplate
database.Init("localhost", 5432, "postgres", "postgres", "palang", true)
```

#### Endpoint

list of endpoint :

- /
- /register
- /login
- /user/me
- /validate-token

you can check also inside `__http-test__`, to run and test all andpoint

## Roadmap

- [x] add test on service and manager
- [ ] grpc delivery.
- [ ] serverless delivery.
- [ ] activation code
- [ ] reset password
