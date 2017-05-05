# Silo

Silo represents a different take on container orchestration. At least, that's the intention, as of writing this, it's at the inception stage.
Below sets out it's intended vision and the current plans to date. All work on this and associated repsitories is open sourced and with community driven development in-mind.

# What problem are we trying to solve?

Container's are totally awesome for many reasons, one of the keys to the proliferation of this technology is it's simplicity and ease-of-use.
When building and developing in the container space, you natrually at some stage encounter container orchestration platforms such as Kubernetes and Rancher.

These systems are really really powerful and are very appropriate to a lot of deployments. They represent a giant leap in the quality and features over what came before.

Sometimes though, you just want a host to come online, perhaps as part of an auto-scaling group and run a pre-prescribed set of containers. 
When using cloud services such as application load balancers, this simple model of scaling predictable compute makes a lot of sense to both those looking to keep things simple and those looking to ensure consistent performance.

Container orchestration tooling of course cater for such a model, but primarily such tools are built around the concept of handling the placement of containers themselves and the provision of value-add services that wrap around your services.
In reality you can find yourself in such situations using a hugely powerful orchestration and scheduling tool to execute what should be a simple task. 
The ability for these platforms to handle encrypted networks, provide auto-discovery is excellent, but at times it can be awkward when going against the grain.

*Surely there is a simpler way?*

You migh then think that configuration management tools such as Chef or Ansible might then factor into your design. 
These are perhaps more suited to this deployment model, but they don't always fit the jigsaw quite so well with other container tools, not do they give you good visibility.

# Silo

So what's Silo and how does it fit in? Whilst not by any means a new concept, providing a simple and easy to use way to upon the booting of a linux kernel, obtain a container run-list and start them up.
Instead of 'orchestrating', we focus on the most simple scheduling we can. This will be done by simply running a silo-agent container, servers can obtain their run-list and start the services on creation or boot. 
That 'run-list' would be something remotely available to the server (S3,git,http,mount). At it's simplest form, this agent/run-list represent the 'runtime' part of the stack.

This gives people an easy way to stand-up containers on hosts. Sounds great, but surely bash scripts or docker-compose could be used in such a fashion to acheive the same goals?
This is where the server part of the equation brings benefit. The intention is to have a central server, fully independant from the runtime.

This server would then be able to handle:

* Agent/Container State monitoring
* Metrics
* Ability to mark servers as down/degraded
* Integration with release/upgrade tooling

The above features will likely be subject to a high rate of change in the inital stages. Fundamentally we want to create something that's minimal and simple to reduce the overhead and complexity of running services this way.
