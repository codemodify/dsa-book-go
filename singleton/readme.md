
# PROBLEM
- Create an object and make sure it does not get created multiple times.
- Also, access that object from anywhere.

- Thoughts:
	- To create an object is easy.

	- Make the object globally accessible.

	- Next challenge is how to make sure multi-threaded/concurrent/parallel
		code won't step on each other and create the same object thinking the
		object does not exist because of race conditions.

