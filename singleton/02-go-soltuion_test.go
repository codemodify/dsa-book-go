package main

import (
	"sync"
	"testing"
)

//
// The Go way SOLUTION
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// This solution is leveraging Go modern principles and implements
// Singleton based on that. It is a more compact version of the classic version.
//

func TestGoSolution(t *testing.T) {
	type ObjectBlueprint struct {
		Field1 int
		Field2 string
	}

	var objectInstance *ObjectBlueprint

	// allow safe concurrent global access to the object
	var objectInstanceOnce sync.Once
	var getInstance = func() *ObjectBlueprint {
		objectInstanceOnce.Do(func() {
			objectInstance = &ObjectBlueprint{
				Field1: 1000,
				Field2: "1000",
			}
		})

		return objectInstance
	}
}
