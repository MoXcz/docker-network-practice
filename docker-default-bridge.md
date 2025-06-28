# `docker` `bridge`

To reset the state of the containers at any point just stop them and start them again:

```sh
docker stop srv srv-2 # stop and delete containers (due to --rm flag)
```

The default `bridge` network uses the private range of the `172.17.0.0/16` subnet (with a broadcast address of `172.17.255.255`, use `ip a`). Container with no specific network specified will have a dynamically assigned IP from this subnet, it's easy enough to see this when running different containers:

```sh
docker run -d --name srv --rm -p 8080:8080 srv-test
# change the mapping of the port
docker run -d --name srv-2 --rm -p 8081:8080 srv-test

# prints the config of the bridge (including the containers IP addresses)
docker network inspect bridge
```

If no other container was running while running the commands above the IP assigned to them should have been `172.17.0.2` and `171.17.0.3`, which are `bridge`d to the host through the *port mapping* done with the `-p` flag. And, because both containers are *under the same network* they can communicate with one another:

```
docker exec -it srv sh
# Inside the container
/app # curl 172.17.0.2:8080
Hi from :8080 (hostname: 3a2c981d68fe IP: [172.17.0.2])
/app # curl 172.17.0.3:8080
Hi from :8080 (hostname: 81f121cce518 IP: [172.17.0.3])
```

At this point it might still not be obvious what's happening. Docker `bridge` network is using the whole subnet of `172.17.0.0/16` (by default) to assign IPs to containers so that when they are mapped they are forwarded to the host.

*Any container inside this network can communicate with other containers under this same network*, and because they are also being forwarded to the host they can also communicate to the *world*:

```
# Outside the containers
curl localhost:8080
Hi from :8080 (hostname: 3a2c981d68fe IP: [172.17.0.2])
curl localhost:8081
Hi from :8080 (hostname: 81f121cce518 IP: [172.17.0.3])
```
