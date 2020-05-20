package main

import (
	"fmt"
	"sync"
	"testing"
)

//
// CLASSIC SOLUTION
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// This solution is valid and works on any programming language and it shows
// how singleton is implemented regardless of the language.
//
// How:
//    1. Create an object.
//
//    2. Make the object globally accessible.
//
//    3. Next challenge is how to make sure multi-threaded/concurrent/parallel
//       code won't step on each other and create the same object thinking the
//       object does not exist because of race conditions.
//
//    4. Final and slimmed down version shows all three points in one package
//

func TestClassicSolution_1_CreateObject(t *testing.T) {
	type ObjectBlueprint struct {
		Field1 int
		Field2 string
	}

	var objectInstance *ObjectBlueprint

	// create object
	var createObject = func{
		objectInstance = &ObjectBlueprint{
			Field1: 1000,
			Field2: "1000",
		}
	}
}

func TestClassicSolution_2_MakeObjectGloballyAccessible(t *testing.T) {
	type ObjectBlueprint struct {
		Field1 int
		Field2 string
	}

	var objectInstance *ObjectBlueprint

	// create object
	var createObject = func{
		objectInstance = &ObjectBlueprint{
			Field1: 1000,
			Field2: "1000",
		}
	}

	// make object globally accessible
	var getInstance = func() {
		return objectInstance
	}
}

func TestClassicSolution_3_FixRaceConfitions(t *testing.T) {
	type ObjectBlueprint struct {
		Field1 int
		Field2 string
	}

	var objectInstance *ObjectBlueprint

	// create object
	var createObject = func{
		objectInstance = &ObjectBlueprint{
			Field1: 1000,
			Field2: "1000",
		}
	}

	// make object globally accessible
	var getInstance = func() {
		return createObjectAndAvoidRaceCondition()
	}

	// fix race condition
	var mutex sync.Mutex
	var createObjectAndAvoidRaceCondition = func() *ObjectBlueprint {
		mutex.Lock()
		if objectInstance == nil {
			createObject()
		}
		mutex.Unlock()

		return objectInstance
	}
}

func TestClassicSolution_4_Final(t *testing.T) {
	type ObjectBlueprint struct {
		Field1 int
		Field2 string
	}

	var objectInstance *ObjectBlueprint

	// allow safe concurrent global access to the object
	var mutex sync.Mutex
	var getInstance = func() *ObjectBlueprint {
		mutex.Lock()
		if objectInstance == nil {
			objectInstance = &ObjectBlueprint{
				Field1: 1000,
				Field2: "1000",
			}
		}
		mutex.Unlock()

		return objectInstance
	}
}
