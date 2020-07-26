# Description

Multi-user chat system using asynchronous pub/sub pattern. Each user should be authenticated and authorized to receive certain messages in a system. 

There are 4 abstract authorization levels: `A`, `B`, `C`, `D`. User level `C` can see level `C` and `D` messages, user level `A` can see everyones' messages.

Each user before sending/receiving a message should prove his/her identity via authentication. Authentication is implemented via JWT.

# Tech stack

Backend and frontend parts are in the same repository (monorepo). 

The backend is written in Go and the frontend in React (javascript).

# How to run this project on a local environment

## Prerequisites:
Go: [https://golang.org/dl/](https://golang.org/dl/)<br/>
NodeJs: [https://nodejs.org/](https://nodejs.org/)<br/>

## Linux/MacOS

### API (backend)

run the following commands on terminal:

`cd api`<br/>
`go run server.go`</br>

api should be running on http://localhost:8000

### Client (frontend)

run the following commands on terminal:

`cd client`<br/>
`npm i` (first time only)<br/>
`npm start`

react should be running on http://localhost:3000

# Login credentials

There are four test users (one for each level): A, B, C and D

### Level A user:

username: A

password: 123

### Level B user:

username: B

password: 123

### Level C user:

username: C

password: 123

### Level D user:

username: D

password: 123

# How to run tests

run the following commands on terminal:

`cd api`<br/>
`go test *.go` (`go test -v *.go` for verbose mode)
