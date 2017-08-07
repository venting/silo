# Silo
[![](https://images.microbadger.com/badges/version/venting/silo-agent.svg)](http://microbadger.com/images/venting/silo-agent "Get your own version badge on microbadger.com") [![](https://images.microbadger.com/badges/image/venting/silo-agent.svg)](http://microbadger.com/images/venting/silo-agent "Get your own image badge on microbadger.com")
[![Go Report Card](https://goreportcard.com/badge/github.com/venting/silo)](https://goreportcard.com/report/github.com/venting/silo)
[![GoDoc](https://godoc.org/github.com/venting/silo?status.svg)](https://godoc.org/github.com/venting/silo)

<img src="http://siloproject.io/images/logo.png" width="400">

Silo represents a different take on container orchestration, one with simplicity as the focus and intenteded to fit neatly into existing tooling. Silo and any associated repositories are open source, community driven development is our goal. Read more on the [official silo documentation](https://siloproject.io).

So what's Silo and how does it fit in? Silo is a distributed container scheduling tool built for the cloud. Relying on well tested concepts, silo provides a simpler layer of functionality to enable it's users to schedule and deploy containers without the complexity and maintenance overheads of full-fat orchestration platforms.

The core component in-play here is the silo-agent. A simple and lightweight container that can be started on boot using 'user data' or similar scripts that execute on the start of a virtual machine. Once started, the agent obtains a run-list and starts the containers. That 'run-list' can be something available to the server over a number of mediums (S3, git, http, volume mount etc).

This gives you an easy way to stand-up container stacks. Sounds great... surely bash scripts or docker-compose could be used in such a fashion to achieve the same goals? 
You absoloutley could of course, but above and beyond scheduling the containers, silo gives you:

* Agent/Container monitoring via a prometheus compatible `/metrics` endpoint
* Overall node health - a customisable health for all services on this stack. This can be used to drive loadbalancer health checks
* Ability to easily control these remote hosts using the Silo CLI, useful for tasks such as upgrading stacks or restarting containers
* Integration with release/upgrade tooling, interactaction with common loadbalancers such as the Amazon ALB and mark the server as down when performing upgrades or maintenance. 

For now, support for Silo is still experimental. The features will likely be subject to a high rate of change in the initial stages. 
Fundamentally though, we want to create something that's small and simple. With the aim of reducing the overhead and complexity of running services this way.
