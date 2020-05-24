
# PROBLEM
- Create a set of objects that implements specific methods but each object nature is different
- Each object

- Thoughts:
	- To create an object is easy.

	- Make the object globally accessible.

	- Next challenge is how to make sure multi-threaded/concurrent/parallel
		code won't step on each other and create the same object thinking the
		object does not exist because of race conditions.
