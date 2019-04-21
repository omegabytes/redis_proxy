#!/bin/bash
sleep 10
echo some very basic tests to show the cache in action:
sleep 1
curl -s GET http://172.17.0.1:$PORT/v1/api/post/first
curl -s GET http://172.17.0.1:$PORT/v1/api/post/second
curl -s GET http://172.17.0.1:$PORT/v1/api/post/third
curl -s GET http://172.17.0.1:$PORT/v1/api/post/fourth
curl -s GET http://172.17.0.1:$PORT/v1/api/post/second
curl -s GET http://172.17.0.1:$PORT/v1/api/post/third
curl -s GET http://172.17.0.1:$PORT/v1/api/post/second
echo waiting a few moments to demonstrate LRU expiry
sleep 5
curl -s GET http://172.17.0.1:$PORT/v1/api/post/second
curl -s GET http://172.17.0.1:$PORT/v1/api/post/second
echo -e "that's all folks! \033[0;31mctrl+c\033[0m and \033[0;31mmake down\033[0m to exit and clean up containers, and thank you!"
exit

