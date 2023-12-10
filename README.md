# chess-move-generator

A move generator for a chess engine implemented in Go.

This isn't meant to be efficient or fancy in any way. Just a little side project to practise Go.


## Good practices I've had to ignore

Unfortunately, calling a function billions of times means that sometimes I've had to do things that I'd usually avoid, no matter how much it hurts to do. Although there are many more, I've listed some of them below.

* In Go almost everything is returned as an error, this means checking for if err!= nil. THese checks all add up so where I'm calling code a lot I sometimes make assumptions about things without checking them. For example, when getting a square index from a bitboard I assume that the function NEVER receives a bitboard that represents multiple squares. If things like this are passed around incorrectly bad things will happen.

* For speed I've used global variables a lot. I've tried to limit this but sometimes it's been unavoidable.
