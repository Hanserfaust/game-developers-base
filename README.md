# Python Game Developers starting point
Last tested and working: January 2023.

Zero to hero suit of examples for the aspiring Python Game Developer. Seem to work well on
Python 3.10, all installed and contained in a virtual environment. Server components (when needed) 
are started using docker-compose.

The intention of this
repository is mostly educational or to provide a skeleton for further development and to provide
a complete and well-defined starting point, developer setup instructions to get started.

I [browsed for the available game frameworks](https://geekflare.com/python-game-development-libraries-frameworks/) 
and decided to first go for cocos2d as it seemed to be the best maintained and documented as of 2023.

So for now focus will go to showing how to use that. At a later point I may show similar examples
for another framework.

Coupled with the ImGUI library, it seems to be what I needed for now, and integration with that is
showed as a first example.

Honorable mentions are pandas-3D as well as pygame. I quickly evaluated Ogre which seemes like
a very capable C++ framework for graphics, but lacked in examples and documentation for Python
especially for the later versions (2023).

## Useful links for the game developer

https://gameprogrammingpatterns.com/

## Dev setup

### MAC
I recommend installing pyenv using brew (Homebrew).

    brew install pyenv

### Ubuntu
Like so:

    apt install pyenv

Set up a new virtual environment called 'pyg' using pyenv
 
    pyenv install 3.10
    pyenv virtualenv 3.10 pyg
    pyenv activate pyg

Install requrements needed for the examples

    pip install -r requirements.txt

And to run the first example:

    python ./coco2d_examples/example1/example1_imgui.py

I use Pycharm for development, but any IDE or editor of choice will do.

## Coco-2D Examples

Based on the library https://pypi.org/project/cocos2d/


### Example 1: Coco and ImGUI

Source: [example1_imgui.py](coco2d_examples/example1/example1_imgui.py).

This example is based upon this: https://github.com/pyimgui/pyimgui/blob/master/doc/examples/integrations_cocos2d.py

with a fix to the broken ImGUI base class.

Purpose, to:
- Show a very basic Coco application skeleton opening a window.
- Show how to integrate ImGUI with coco to make settings to the application dynamically.


### Example 2: Coco Sprites

This example is based on this: https://github.com/los-cocos/cocos/blob/master/samples/demo_sprites.py

The same [repository](https://github.com/los-cocos/cocos/blob/master/samples/) contains a number of other useful examples
that you could check out.


### Example 3: Networking with a RedisDB backend
...
