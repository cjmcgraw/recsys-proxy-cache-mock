# RecsysProxyCache Mock

[At dockerhub](https://hub.docker.com/repository/docker/qarlm/recsys-proxy-cache-mock)

This project is the implementation of a simple mock container to replace the
[RecsysProxyCache](https://github.com/cjmcgraw/recsys-proxy-cache) in projects that may
be using it, or closely tied to it.

The purpose of this is mainly to support "System Tests". i.e., tests that execute through
multiple containers to ensure functionality is working as expected.

## How do I use it?

```bash
$ docker run -p 50051:50051 'qarlm/recsys-proxy-cache-mock'
```

This will start the container locally on port 50051 for you
However that is rarely useful, and instead it is highly recommended
that you include it in your project's docker-compose.yml to weave the
stack together for system testing

```
#docker-compose.yml

version: "3.9"

services:

    my_service:
        ...

    recsys_proxy_cache:
        image: qarlm/recsys-proxy-cache-mock
```

Doing this then your dependencies can be satisfied easily and you can utilize the mock
for development purposes directly.

## Yeah.. but why?

On my team at Accretive Technology Group we consider System Testing to be extremely critical.
From this importance we find dependencies can sometimes be difficult to manage. The [RecsysProxyCache](https://github.com/cjmcgraw/recsys-proxy-cache)
is one such dependency. In order to run that project you must have machine learning models up, load balancers,
and the project itself which is relatively heavy weight.

Most of the time tests don't care that their dependency functions exactly how they'd like, they only care that
the contract is maintained and it operates how they expect.

With this in mind this mock exists. Its purposes is simple. When you send a request to this mock it returns to you
a repsonse of random scores. As if it was the RecsysProxyCache, but without all the heavyweight implementation details.
This allows people closely tied to that system to test easily in a local stack with strong tests, without having the
cumbersome problem of satisfiying other dependencies.

## Couldn't I just write unit tests instead?

Did I say not to write them. Unit tests definitely have their purpose. But from my experience 1 system test is worth
1000 unit tests. 

In this world of docker and dependency management as containers, why limit yourself?

