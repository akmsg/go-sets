#go-sets

Gives you an implementation of Set, with most of the usefull operations on them. No need to worry about concurrency issues it has been kept in mind while designing.

#In Nustshell

1. Create

        s := set.New() // will create a new empty set
	s := set.New(1, 2, "some text", strcut{}{}) // will create a set of few objects

2. Remove element(s)

        s.Remove(1, 2)

3. Enumerate

        s.List() // returns a slice of all the set elements

4. Empty

        s.Empty() // remove all the elements from the set

#ToDo
