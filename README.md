# Python Game Developers starting point
Last tested and working: January 2023.

Zero to hero suit of examples for the aspiring Python Game Developer. I
[browsed for the available game frameworks](https://geekflare.com/python-game-development-libraries-frameworks/) 
and decided to go for cocos2d as it seemed to be the best maintained and documented as of 2023.

So for now focus will go to showing how to use that. At a later point I may show another one.

Coupled with the ImGUI library, it seems to be what I needed for now, and integration with that is
showed as a first example.

Honorable mentions are pandas-3D as well as pygame. I quickly evaluated Ogre which seemes like
a very capable C++ framework for graphics, but lacked in examples and documentation for Python
especially for the later versions (2023).

## Doc sources
Based on the library https://pypi.org/project/cocos2d/

### Other useful links

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
    pip install -r requirements.txt

And to run the first example:

    python ./example1_imgui.py

I use Pycharm for development, but any IDE or editor of choice will do.

## Examples

Subsequent examples involving networking are extensions of that.

### Example 1: Coco and ImGUI

[examples/example1/](Source).

This example is based upon this: https://github.com/pyimgui/pyimgui/blob/master/doc/examples/integrations_cocos2d.py

with a fix to the broken ImGUI base class.

Purpose, to:
- Show a very basic Coco application skeleton opening a window.
- Show how to integrate ImGUI with coco to make settings to the application dynamically.


### Example 2: Coco Sprites

This example is based on this: https://github.com/los-cocos/cocos/blob/master/samples/demo_sprites.py

The same reopository contains a number of other useful examples as well


### Example 3: Networking with a RedisDB backend
...
