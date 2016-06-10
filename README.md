# Discovery agent for docker exposed ports

A simple in go written agent that exposes a minimal REST api so a container can find the exposed port for a container port. The agent is just a workaround for the problem mentioned in https://github.com/docker/docker/issues/3778


## Usages
Start the agent
```
docker run -p 5555:8080 -v /var/run/docker.sock:/var/run/docker.sock --name docker-discovery-agent npalm/docker-discovery-agent
```

Test the agent
```
docker run -p :80 --add-host="dockerhost:<docker-host-ip / your machine> --rm test centos /bin/bash
```
Now you are on the shell of the container. You can look up the exposed port as follow.
```
# find the container id
DOCKER_CONTAINER_ID=$(cat /proc/self/cgroup  | grep "cpu:/" | sed 's/\([0-9]\):cpu:\/docker\///g')

# request the port binding for mapped port 80
curl "http://dockerhost:5555/container/${DOCKER_CONTAINER_ID}/portbinding?port=8080&protocol=tcp"
```
The response will look like:
```[{"HostIp":"0.0.0.0","HostPort":"12345"}]```


## REST API
Only a minimal API is implemented

## Health check
Once the service is running the health check should response with a 200 status code.
```resource:  <host:port>
   response: 200 OK
```

## Port binding
 ```
resource: <host:port>/container/{id}/portbinding?port={port}&protocol={protocol}
response: A JSON resprenting the exposded port for container with id
          {
            "HostIp": "",
            "HostPort": ""
          }
```