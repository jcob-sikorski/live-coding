import threading
import time

def task():
    time.sleep(2)
    print("Hello World!")

thread = threading.Thread(target=task)

thread.start()

thread.join()