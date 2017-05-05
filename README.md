# Summary

Silo represents a different take on container orchestration. At least, that's the intention, it's early days.
Below sets out the intended vision and the current plans to date. Silo and any associated repositories are open source, community driven development is our goal.

# What problem are we trying to solve?

Container's are totally awesome for many reasons, one of the keys to their success, is their simplicity and ease-of-use.
When building and developing on containers, at some stage you encounter container orchestration platforms such as Kubernetes and Rancher.

These systems are really, really powerful and are very appropriate to a lot of deployments. They represent a giant leap in the quality and feature sets over what came before.

Sometimes though, you just want a host to come online and run a pre-prescribed set of containers, perhaps as part of an auto-scaling group or similar.

When using services available in the cloud such as load balancers, this simple way of scaling compute in a predictable fashion, makes a lot of sense. You might be looking to keep things simple or maybe your looking to ensure consistent performance for peak periods.

Container orchestrators cater for this type of deployment, primarily though, such tools are built around the concept of handling the placement of containers themselve, provision of value-add components that wrap around your services.
In reality you can find yourself in such situations using a hugely powerful orchestration and scheduling tool to execute what should be a simple task.  The ability for these platforms to handle encrypted networks, provide auto-discovery is excellent, but at times it can feel awkward when going against the grain of their typical use-cases.

*Surely there is a simpler way?*

You might then think that configuration management tools, Chef or Ansible might be a good option. 
Whilst very capable of handling such types of deployment, they don't always fit the jigsaw quite so well with other container tools, nor do they give you good visibility of your services.

# Silo

So what's Silo and how does it fit in? Whilst not by any means a new concept, providing a simple and easy to use way to upon the boot, obtain a container run-list and start the containers. Instead of 'orchestrating', we focus on the most simple scheduling we can. This will be done by simply running a silo-agent container, servers can obtain their run-list and start the services on creation or boot. 

That 'run-list' would be something remotely available to the server (S3, git, http, volume mount). At it's simplest form, this agent/run-list represent the 'runtime' part of the stack.

This gives people an easy way to stand-up containers on hosts. Sounds great... surely bash scripts or docker-compose could be used in such a fashion to achieve the same goals? This is where the server part of the equation brings benefit. The intention is to have a central server, fully independent from the runtime.

This optional server component would then be able to handle:

* Agent/Container State monitoring
* Metrics
* Ability to mark servers as down/degraded
* Integration with release/upgrade tooling

The features will likely be subject to a high rate of change in the initial stages. Fundamentally though, we want to create something that's small and simple. With the aim of reducing the overhead and complexity of running services this way.
