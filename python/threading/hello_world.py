import threading

def task():
    print("Hello World!")

thread = threading.Thread(target=task)

thread.start()

thread.join()