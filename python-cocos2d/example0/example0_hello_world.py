import cocos


# Subclass a Layer and define the logic of you program here:
class HelloWorld(cocos.layer.Layer):

    def __init__(self):
        super(HelloWorld, self).__init__()

        label = cocos.text.Label(
            'Hello, world',
            font_name='Times New Roman',
            font_size=32,
            anchor_x='center', anchor_y='center'
        )

        label.position = 320, 240

        # Since Label is a subclass of CocosNode it can be added as a child. All CocosNode objects
        # know how to render itself, perform actions and transformations. To add it as a layer’s
        # child, use the CocosNode.add method:
        self.add(label)


def main():

    # Init director
    cocos.director.director.init()

    # Instantiate the HelloWorld layer
    hello_layer = HelloWorld()

    # Create a scene
    main_scene = cocos.scene.Scene(hello_layer)

    # And let the director run the scene
    cocos.director.director.run(main_scene)


if __name__ == "__main__":
    main()
