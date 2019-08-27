![project unmaintained](https://img.shields.io/badge/project-unmaintained-red.svg)

# Discovery agent for docker exposed ports

A simple in go written agent that exposes a minimal REST api so a container can find the exposed port for a container port. The agent is just a workaround for the problem mentioned in https://github.com/docker/docker/issues/3778


## Usages
Start the agent
```
docker run -p 5555:8080 -v /var/run/docker.sock:/var/run/docker.sock \
  --name docker-discovery-agent npalm/docker-discovery-agent
```

Test the agent
```
export DOCKERHOST=$(ifconfig | grep -E "([0-9]{1,3}\.){3}[0-9]{1,3}" | \
  grep -v 127.0.0.1 | awk '{ print $2 }' | cut -f2 -d: | head -n1)
docker run -it -p 80 --add-host="dockerhost:$(DOCKERHOST)" --rm centos /bin/bash
```
Now you are on the shell of the container. You can look up the exposed port as follow.
```
# find the container id
CONTAINER_ID=$(cat /proc/self/cgroup  | grep "cpu:/" | sed 's/\([0-9]\):cpu:\/docker\///g')

# request the port binding for mapped port 80
curl "http://dockerhost:5555/container/${CONTAINER_ID}/portbinding?port=80&protocol=tcp"
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
