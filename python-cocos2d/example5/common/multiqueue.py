"""
Allow multiple queues to be waited upon.

queue,value = multiq.select(list_of_queues)
"""
import queue
import threading


class QueueReader(threading.Thread):
    def __init__(self, inq, sharedq):
        threading.Thread.__init__(self)
        self.inq = inq
        self.sharedq = sharedq

    def run(self):
        while True:
            data = self.inq.get()
            print("thread reads data=", data)
            result = (self.inq, data)
            self.sharedq.put(result)


class MultiQueue(queue.Queue):
    def __init__(self, list_of_queues):
        queue.Queue.__init__(self)
        for q in list_of_queues:
            qr = QueueReader(q, self)
            qr.start()


def select(list_of_queues):
    outq = queue.Queue()
    for q in list_of_queues:
        qr = QueueReader(q, outq)
        qr.start()
    return outq.get()
