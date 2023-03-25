# zero-downtime-deploy-example


## call go api and websocket
curl --location 'http://127.0.0.1:8081/api/welcome'
wscat -c 'ws://127.0.0.1:8081/websocket'


## how to drain connect

### step1.
docker exec -it nginx bash
### step2.
vim /etc/nginx/nginx.conf

### step3.


### step4.
nginx -t && nginx -s reload

### step5.
docker-compose stop goservice1
