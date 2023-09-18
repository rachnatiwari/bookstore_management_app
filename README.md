# BOOKSTORE MANAGEMENT APP
A golang backend application to manage bookstore inventory

### TECHLOGY USED
- Go
- Docker
- Kubernetes

### SET-UP
- For local setup (MacOS)
  1. Start mysql service <br>
     `brew services start mysql`
  2. Export the environment variables
     - `export DB_SOURCE="<db_user>:<db_password>@tcp(<host>:<port>)/"` <br>
       eg: `export DB_SOURCE="root:password@tcp(localhost:3306)/"`
     - `export DB_SOURCE_DB: <db_user>:<db_password>@tcp(<host>:<port>)/bookStore?charset=utf8mb4&parseTime=True&loc=Local` <br>
       eg: `export DB_SOURCE_DB: root:password@tcp(localhost:3306)/bookStore?charset=utf8mb4&parseTime=True&loc=Local`
  3. Download required Go modules <br>
    `go mod downlaod`
  4. Start Go Server <br>
     `go run cmd/main/main.go`<br> <br>
- For deploying on minikube 
  1. Build docker image <br>
     `docker build -t bookstore:latest .`
  2. Load docker image in minikube <br>
     `minikube image load bookstore:latest` 
  3. Apply the deployment on minikube <br>
    `kubectl apply -f deployment`
  
