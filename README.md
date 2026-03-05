# Docker 

## run locally 
docker compose up --build

## build 
docker compose build

## Tag 
// Accidentally commited to 'frontend' instead of 'website'
docker tag simple-todo-app-frontend:latest markusnilssn/simple-todo-app:frontend
docker tag simple-todo-app-backend:latest markusnilssn/simple-todo-app:backend

## push
docker push markusnilssn/simple-todo-app:backend
docker push markusnilssn/simple-todo-app:frontend

# Node.js
node website/main.js

in 'website/' directory: 
npm install
npm start

# Golang
go run backend/bin/main.go
go mod init EXAMPLE/EXAMPLE

# To Run ...