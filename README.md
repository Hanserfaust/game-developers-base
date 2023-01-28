# Game Developers Examples
Last tested and working: January 2023.

This repository contains a set of code examples that I set up while learning game development.

Starting with Python for fun and easy programing moving towards high-performance game-server
using Go and Node.js interacting with Unity and possibly Unreal.

## Useful links for the game developer

Patterns in general: https://gameprogrammingpatterns.com/

The examples will not build upon the ECS pattern, but this is recommended for for more complex
designs. Worth to keep in mind when/if we see flaws in a simpler design:

ECS architectural pattern: https://en.wikipedia.org/wiki/Entity_component_system


## 1. Python
Since my main programming language at the time of writing this was Python I decided to see what
it had to offer even though Pyton is not the language of choice for most professionals. 

The set of examples and documentation could still very well be used do develop simpler (or even
complex) games for educational purposes.

The Python-based examples can be found [here](./python-cocos2d/README.md).

Typical use-cases, but not limited to:

- Fun Python development, events, graphics, 2D etc.
- Educational for game (client) development.
- Cocos 2D development in general.


## 2. Python client, Go server
This a smaller set of examples, building on the basic Python examples above but working
against a simple but custom made [Go](https://go.dev/) server-side. This part contains only one
or a few examples for experimental purposes. I may expand this once my Go skills improves.

The Python-Go examples

Use-cases:

- Go server development in general.
- Scalability, multiprocessing.
- Server coding pattern experiments.
- Reducing latency etc.


## 3. Unity client, Go server
This part represents a more professional starting point. The examples here build on an existing
Go-based game backend called [GoWorld](https://github.com/xiaonanln/goworld) and its Unity examples 
as well as looking to other existing open-source game-server solutions, such as the node.js-based
[Pomelo](https://github.com/NetEase/pomelo) engine, that contains a host of useful information regarding
design choices made.

Use-cases:
- Learning professional large scale game server architecture.
- Interaction with networking Unity-based games.
- Deployment and Ops aspects of game servers.