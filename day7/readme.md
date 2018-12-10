

I think this is called a topological sort

see here: https://courses.cs.washington.edu/courses/cse326/03wi/lectures/RaoLect20.pdf


C -> A -> F -> None
A -> B -> D -> None
B -> E -> None
D -> E -> None
F -> E -> None

- get list of all the nodes on the right side and add them up
That is the degree of the node ...
Look for a node that has degree 0
I think this is the entry point into the whole thing


