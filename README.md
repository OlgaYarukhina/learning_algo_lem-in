## LEM-IN

### Author
[Olha Yarukhina](https://01.kood.tech/git/oyarukhi)

### Running the program
1. Run in terminal: 
`go run . example00.txt` (all examples in folder "examples")
2. [Check here if result correct](https://github.com/01-edu/public/tree/master/subjects/lem-in/audit)


___

## What is LEM-IN?
### Objectives
This project is meant to make a digital game of an ant farm. Its goal is to  successfully finding the quickest path, from start point to the end point in the anthill.

So the program is designed to find the maximum number of vertex-independent paths, to move 'ants' from point A to point B in the least number of moves, provided that 'ants' can go to the neighboring vertex in one move, but cannot move several pieces along the same edge, and also cannot be at the same vertex at the same time (except for the start vertex and the finish vertex).

### About lem-in
The program receives information about the graph (anthill) as input:

 - the number of 'ants';
 - the name of the vertices and their coordinates. The ##start and ##end commands specify the start and end vertices of the path.
 - edges connecting vertices.

At the beginning from standart output the program recieves information about number of ants had to be leaded through path, room names (can be numbers, words, characters etc.)

First comes the number of ants (in the example above, 5). Next comes the description of the vertices (its NAME + its COORDINATES (x and y)).
The ##start and ##end commands mean that they will be followed by a start node and a finish node, respectively. After the peaks comes the information
about which vertices are connected.

Here is an example:


```console
        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/
```

More example in the folder "examples"


### Validation
The following conditions are checked during the validation stage:
1) The number of ants is not negative.
2) Exactly one vertex is marked as a start and as a finish.
3) There is at least one path connecting start and finish, i.e. the graph is connected.
4) There are no vertices with the same names.
5) There are no vertices with the same coordinates (the same x and y).
6) Edges connect only existing vertices.
7) No vertices starting with 'L'
8) A graph can have multi-edges (2 vertices can be connected by more than 1 edge) and loops (edges connecting a vertex to itself).
9) These edges will not cause errors in the algorithm.