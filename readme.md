## Vue Events Bulletin Board

This is the Go backend code for the Vue.js [tutorial on Scotch.io](https://scotch.io/tutorials/build-a-single-page-time-tracking-app-with-vue-js-introduction). In the tutorial we build a events bulletin board application and cover the basics of [Vue](http://vuejs.org/).

Please read the `server.go` file for elaborate comments. 

If you are new to Go, read [my book](http://github.com/thewhitetulip/web-dev-golang-anti-textbook) or read the working code of the project which is taught in the book [here](http://github.com/thewhitetulip/Tasks).

## Building the app

1. install Go
2. `go get github.com/thewhitetulip/go-vue-events`
3. go run server.go
4. Open localhost:8080 in your browser

## Build in Docker
1. [Install Docker](https://docs.docker.com/get-docker/)
2.```
docker build --tag go-vue-events:1.0 .
```
3.```
docker run --publish 8000:8080 --detach --name gve go-vue-events:1.0
```
### Remove docker image
```
docker rm --force gve
```