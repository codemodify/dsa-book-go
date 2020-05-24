package main

//
// The Go way SOLUTION
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// This solution is leveraging Go modern principles and implements
// Object-Pool based on that.
//
// The example below is kept minimal to avoid code noise so you can capture
// the idea.
//

type apiConnection struct {
	// ...
}

func (ac apiConnection) Connect()     {}
func (ac apiConnection) MakeRequest() {}

type apiConnectionPool struct {
	connections         []apiConnection
	nextAvailConnection int
}

func newApiConnectionPool() apiConnectionPool {
	r := apiConnectionPool{
		connections: []apiConnection{
			// ...
		},
		nextAvailConnection: 0,
	}

	for _, c := range r.connections {
		c.Connect()
	}

	return r
}

func (cp apiConnectionPool) getConnection() apiConnection {
	return cp.connections[cp.nextAvailConnection]
}

func main() {
	apiConnectionPool := newApiConnectionPool()
	apiConnectionPool.getConnection().MakeRequest()
}
