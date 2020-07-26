# Description

Multi-user chat system using asynchronous pub/sub pattern. Each user should be authenticated and authorized to receive certain messages in a system. 

There are 4 abstract authorization levels: `A`, `B`, `C`, `D`. User level `C` can see level `C` and `D` messages, user level `A` can see everyones' messages.

Each user before sending/receiving a message should prove his/her identity via authentication. Authentication is implemented via JWT.

## Tech stack

Backend and frontend parts are in the same repository (monorepo). 

The backend is written in Go and the frontend in React (javascript).

## How to install

clone repository

cd client

npm i

## How to run project

### API (backend)

cd api

go run server.go

api should be running on http://localhost:8000

### Client (frontend)

cd client

npm start

react should be running on http://localhost:3000

## How to run tests

cd api

go test *.go (go test -v *.go for verbose mode)
