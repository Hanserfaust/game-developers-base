import cocos

from layers import MouseDisplay

from comm import send_to_server, start_socket_thread
from messages import PlayerLogin


def main():
    # Start the background threads needed before c
    socket_thread = start_socket_thread()

    # Login
    send_to_server(PlayerLogin("player", "s3cret"))

    # Init director
    cocos.director.director.init()

    mouse_display = MouseDisplay()

    main_scene = cocos.scene.Scene(mouse_display)

    # And let the director run the scene
    cocos.director.director.run(main_scene)

    socket_thread.join()


if __name__ == "__main__":
    main()
