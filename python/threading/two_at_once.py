# Two-at-Once: Launch two threads simultaneously: one that prints "Ping"
# every second, and one that prints "Pong" every second. Let them run for 3 seconds.

import threading
import time

def task(message: str, repeat: int):
    for _ in range(repeat):
        time.sleep(1)
        print(message)

thread1 = threading.Thread(target=task, args=("Ping", 3))
thread2 = threading.Thread(target=task, args=("Pong", 3))

thread1.start()
thread2.start()

thread1.join()
thread2.join()