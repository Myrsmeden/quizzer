# Quizzer
Quizzer is an application written i Go and Vue.js 
This version of Quizzer is intended to be used when studying for a test in the History of Mathematics.

## Technologies
Quizzer uses Golang for backend and Vue.js for the client application.
The backend is using Echo, a golang framework for easy API writing

## Endpoints
- `/persons` for persons
- `/attributes` for attributes
- `/question` for a random question

## How it works
Via the `/manage`-page one can add persons and attributes.
Each attribute is connected to one person, and one person only.
The question endpoint serves a question to the user, by randomly selecting one of the question types, i.e. one person question that asks which attribute that corresponds with the given person.

## Installation
To start using Quizzer on your own machine, install the neccessary packages with the following commands:
```
go get github.com/labstack/echo
go get github.com/mattn/go-sqlite3
```
Build the code with
```
go build quiz.go
```
and run with
```
./quiz
```
