## Problem 1 : Find the number of islands.
Detail of the question can be found [here](https://www.geeksforgeeks.org/find-number-of-islands/).

### Solution
the matix in question isnt an [adjecency matrix](https://en.wikipedia.org/wiki/Adjacency_matrix). It is supposed to be a visual representation of connected nodes/land.

For example
```
1 1 0 0 0
0 1 0 0 1
1 0 0 1 1
0 0 0 0 0
1 0 1 0 1 
```

Here, there are 5 disconnected nodes / piece of land
```
1 1
  1     1
      1 1

1   1   1      
```

Aim is to programatically determine number of distinct nodes/islands given an input matrix.