# zero-downtime-deploy-example


## call go api and websocket

Use curl and wscat to call a Go API and Websocket, respectively.

curl --location 'http://127.0.0.1:8081/api/welcome'
Use curl to call the Go API.

wscat -c 'ws://127.0.0.1:8081/websocket'
Use wscat to call the Websocket.

## how to drain connect

### step1. Enter the bash environment of the NGINX container using the docker exec command.
docker exec -it nginx bash

### step2. Use vim to edit the NGINX configuration file in the bash environment of the NGINX container.
vim /etc/nginx/nginx.conf

### step3. Modify the NGINX configuration file to gradually stop the node.


In the upstream section, change the original configuration:

upstream node_cluster {
    server goservice1:18086;
    server goservice2:18086;
}

to:

upstream node_cluster {
    server goservice1:18086 down;
    server goservice2:18086;
}




### step4. Test the new NGINX configuration file for correctness and then reload the NGINX configuration.
Use the nginx -t command to test the new NGINX configuration file for correctness:
nginx -t && nginx -s reload


### step5.
Use socket connection to test whether NGINX will interrupt the established connection during reload.

In the case of an established connection, try to maintain the connection and send messages during NGINX reload using a tool similar to wscat.

If the connection is not interrupted and the message is successfully delivered, it means that NGINX will not interrupt the established connection during reload.


### step6. Use the docker-compose command to stop the goservice1 container
Stop goservice1:

docker-compose stop goservice1

After completing the above steps, zero-downtime deployment is complete.