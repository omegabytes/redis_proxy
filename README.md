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

## Cache Operation Complexity

## Running and Testing

## Time Spent
- learning Docker, redis, APIs in golang: 5h 

## Omissions

## Resources
https://github.com/hauke96/tiny-http-proxy
https://github.com/mingrammer/go-todo-rest-api-example/