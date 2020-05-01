# static-html-server
Static http server designed for vue-router [history mode](https://router.vuejs.org/guide/essentials/history-mode.html) (if file not exists return index.html).

##How to use
###with docker
```shell script
 docker run -v /path/to/project:/var/www -p 8123:8080 static-http-server   
```
8123 - your port

###with docker-compose
```yaml
version: '3.1'

services:
  httpserver:
    image: aquaminer/static-html-server
    ports:
      - 8080:8080
    volumes:
    - ./vue-dist:/var/www
```

##TODO
* [ ] minimize docker image, compile external, run with alpine 