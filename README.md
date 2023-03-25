# zero-downtime-deploy-example


## call go api and websocket

Use curl and wscat to call a Go API and Websocket, respectively.
<pre><code>
curl --location 'http://127.0.0.1:8081/api/welcome'
</code></pre>

Use curl to call the Go API.

<pre><code>
wscat -c 'ws://127.0.0.1:8081/websocket'
</code></pre>

Use wscat to call the Websocket.

## how to drain connect

### STEP1. Start the Docker services:

<pre><code>
docker-compose up -d
</code></pre>

### STEP2. Enter the bash environment of the NGINX container using the docker exec command.
<pre><code>
docker exec -it nginx bash
</code></pre>

### STEP3. Use vim to edit the NGINX configuration file in the bash environment of the NGINX container.
<pre><code>
vim /etc/nginx/nginx.conf
</code></pre>

### STEP4. Modify the NGINX configuration file to gradually stop the node.


In the upstream section, change the original configuration:
<pre><code>
upstream node_cluster {
    server goservice1:18086;
    server goservice2:18086;
}
</code></pre>

to:

<pre><code>
upstream node_cluster {
    server goservice1:18086 down;
    server goservice2:18086;
}
</code></pre>



### STEP5. Test the new NGINX configuration file for correctness and then reload the NGINX configuration.
Use the nginx -t command to test the new NGINX configuration file for correctness:
<pre><code>
nginx -t && nginx -s reload
</code></pre>

### STEP6.
Use socket connection to test whether NGINX will interrupt the established connection during reload.

In the case of an established connection, try to maintain the connection and send messages during NGINX reload using a tool similar to wscat.

If the connection is not interrupted and the message is successfully delivered, it means that NGINX will not interrupt the established connection during reload.


### STEP7. Use the docker-compose command to stop the goservice1 container
Stop goservice1:

<pre><code>
docker-compose stop goservice1
</code></pre>


After completing the above steps, zero-downtime deployment is complete.