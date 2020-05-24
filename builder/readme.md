
# PROBLEM
- NOTE: Strongly advised to see Factory first then Builder.

- Builder is a more noisy/complicated case of the Factory; if in Factory
  case based on a key you quickly create the object "inline" then if to
  imagine that the "inline" operation would not have enough space or would
  take more lines to express the intricate details of object creation, you
  would tend to move that object creation into a separate function/method.

- Builder approach is exactly the above except it adds more abstraction and
  makes it a formal way to create objects for a specific functionality.
