# palang
Simple authentication service, to help create application faster. Main Idea is help developer not write again auth and user module when create new application.

## Feature
- Register
- Activation
- Login
- Reset Password
- Detail User
- Validate Token

## Foldering
Basic idea for foldering we breakdown breakdown to 3 folder, Service, Manager and Delivery
- *Service*
in this folder contain logic and functions of each domain, and each domain must be single responsibility
- *Manager*
in this folder contain flow/usecase for apps (can be orchestration multiple services). example function to Register that will involve user servics and email services. 
- *Delivery*
in this it's delivery protocol, so can be multiple delivery (grpc, rest API, Graphql)
