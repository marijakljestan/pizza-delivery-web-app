# pizza-delivery-web-app

## Features

Methods for customers:  \
List the menu \
Create order  \
Check order status  \
Cancel order if it’s status is not ready_to_be_delivered  \
Methods for admins - this methods are allowed only for users with role ADMIN (which is achieved using JWT) \
Add pizza to the menu \
Delete pizza from the menu \
Cancel order regardless the status \

CLI part should do the same as the HTTP requests, but it should be executable file, where you can pass your arguments


## Running application localy

### Server
Server is a Golang application. Source code is in the <i>./server</i> folder. It can be started as a Golang project in Goland IDE. In project configuration you have to set path to <i>/server</i> as a run directory.
MongoDB is needed for the backend, as well as a few environment variables: DB_HOST and DB_PORT.

### Client
Client is a CLI application, written in Go. If you want to run this application, first you have to be sure that server application is running, and after that, navigate to <i>/client</i> directory and run <code>go run client.go</code>.


## Running infrastructure using Docker
Whole infrastructure can be runned with <code>docker compose up --build</code> in root directory. 
