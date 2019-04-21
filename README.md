# redis proxy

## Architecture Overview

```
           network            proxy             cache
         .---------.       .---------.       .---------.
    <----|- - < - -|---<---|- - < - -|---<---|- < -.   |
you ---->|- - > - -|--->---|- -,- > -|--->---|- > -|   |
         |         |       |   |(*)  |       |     |   |
         |    ,-< -|---<---|< -'     |       |     |   |
         |    , ,->|--->---|- - > - -|--->---|- > -'   |
         `----+-+--´       `---------´       `---------´
              ' '
              '_'
            redis
```

## Code Summary
The frontent application ("app") is a simple golang api layer implementing GET and POST commands for a redis backend. Please note that the POST endpoint has been provided for testing purposes only.
It contains a cache which is implemented as a doubly-linked queue indexed by "title" and ordered by least recently used entry. 
Entries after the cache reaches the configurable max size cause the least recently used element to be evicted. 
Cache entries have an expiration time applied when added to the cache, and "expired" entries are replaced through a GET request.
Configurable properties for this service include the cache size (`SIZE`, default 5), cache duration (`TTL`, default 5), and port the application can be accessed on (`PORT`, default 8080). 

The redis backend ("redis") runs in a separate container with default configuration (default port 6379). You may configure the port by setting `REDIS_PORT`.

## Cache Operation Complexity
The cache implements a doubly-linked list, which grants O(1) insertion and deletion time. It further implements a hash for indexing, which allows for O(1) search.

## Running and Testing
Please ensure your environment has the following installed:
- make
- docker
- docker-compose
- Bash

### Running the service
Installation:
```
git clone https://github.com/omegabytes/redis_proxy.git
cd redis_proxy/
```

Run service:
```
make up
```

Set environment variable `DEMO=true` to pre-populate the data store with some sample data to perform GET requests against.
An example GET request of `curl -s GET http://172.17.0.1:8080/v1/api/post/first` would return ` {"title":"first","body":"one"}`

### Testing
Run example tests:
```
make test
```

Run same tests with configurations:
```
REDIS_PORT=6378 TTL=10 SIZE=10 PORT=8081 make test
```

Run go unit tests:
```
make unit_test
```
Please note that the go unit tests are provided to as an example of TDD. Your system requires `go` to be installed to run these, so please consider them as "extra" and not a part of the main requirements.

Teardown:
```
make down
```
### Configuration
- Address of the backing Redis service (`REDIS_PORT`)
- Cache expiry time (`TTL`)
- Capacity (number of keys) (`SIZE`)
- TCP/IP port number the proxy listens on (`PORT`)
- Flag to init database with values (`DEMO`)

## Time Spent
- Docker, redis ramp up: 5h
- API implementation: 3h
- Test suite: 2h
- Cache: 1h

## Resources
A short list of some materials I referenced while completing this challenge:

- https://github.com/hauke96/tiny-http-proxy
- https://github.com/mingrammer/go-todo-rest-api-example/
- https://github.com/Lebonesco/go_lru_cache/blob/master/main.go
- https://github.com/wunderlist/ttlcache