from messaging import messages


def get_type_id(game_message):

    match(type(game_message)):
        case messages.PlayerLogin:
            return messages.MType.PLAYER_LOGIN

        case messages.MouseEvent:
            return messages.MType.MOUSE_MOVE

        case _:
            raise Exception("Unsupported message type: %s" % type(game_message))


def build_packet(game_message):
    """
    This is close to the wire, called (only?) from a low layer (TCP/UDP) to
    build a tight packet:

    [1 byte packet length] + [1 byte packet type] + [N bytes packet data]

    :param game_message:
    :return:
    """
    gm_id = get_type_id(game_message)
    msg_bytes = bytes(game_message)
    header = bytes([len(msg_bytes), gm_id])
    return header + msg_bytes
