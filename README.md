# Twitter : Distributed System Project

Twitter-like distributed and stateless web application (scalable) with microservices that communicate with the webserver via gRPC. The services persist their state in a raft replicated data store to provide consistency. The services include login, logout, register, follow user, unfollow user, tweet and profile dashboard. The raft cluster has three nodes and we can observe that the application runs only when majority(two) of the nodes are running.

## Pre-requisites:

	Go-lang

## Execution:

Terminal 1:

	git clone https://github.com/swumar/twitter.git
	export GOPATH=$HOME/twitter
	cd twitter/src/web/main
	go run web.go

Terminal 2:

	export GOPATH=$HOME/twitter
	cd twitter/src/backend/service
	go run service.go

Terminal 3:

	export GOPATH=$HOME/twitter
	cd twitter/src/go.etcd.io/etcd
	export PATH=$GOPATH/src/go.etcd.io/etcd/bin:$PATH
	goreman start

Terminal 4:

	export GOPATH=$HOME/twitter
	cd twitter/src/go.etcd.io/etcd
	export PATH=$GOPATH/src/go.etcd.io/etcd/bin:$PATH
	goreman run stop etcd1 (to stop a node)
	goreman run restart etcd1 (to restart a node)
	goreman run status(to see status of each node)

Open http://localhost:8000/ on a browser (client)

Click on register to register -> Login with credentials -> Profile

Functionality:

    1) Enter emailid to follow a user
    2) Enter emailid to unfollow a user
    3) Enter tweet message to tweet
    4) Click on feed to see five recent tweets of your followers
    5) Signout -> Destroys the cookie and removes the token for user
