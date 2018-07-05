# Go Set 

[![Build Status](https://travis-ci.org/anandmishra01/go-sets.svg?branch=master)](https://travis-ci.org/anandmishra01/go-sets)

A simple and easy to use implementation of **Set** in golang, with most of the usefull operations on them. 

No need to worry about concurrency issues, this implementation is **thread safe**.

# Installation

```sh
    go get github.com/anandmishra01/go-sets
```

# In Nustshell

1. Create

```go
    s := set.New() // will create a new empty set
    s := set.New(1, 2, "some text", strcut{}{}) // will create a set of few objects
```

2. Remove element(s)

```go
    s.Remove(1, 2)
```

3. Enumerate

```go
    s.List() // returns a slice of all the set elements
```

4. Empty

```go
    s.Empty() // remove all the elements from the set
```

# To Do

1. Union

2. Intersection

3. Difference

4. Check Subset

5. Check Equality
