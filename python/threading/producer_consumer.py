# Producer-Consumer (Queue): Use queue.Queue to pass data between a "Producer"
# thread (generating numbers) and a "Consumer" thread (squaring them).

import threading
import queue
import time
import random

# 1. Initialize the thread-safe Queue
data_queue = queue.Queue(maxsize=5)

def producer():
    for i in range(1, 6):
        item = random.randint(1, 10)
        print(f"[Producer] Generating: {item}")
        
        # This will block if the queue is full (maxsize=5)
        data_queue.put(item) 
        
        time.sleep(random.random()) # Simulate work
    
    # Signal that the producer is done
    data_queue.put(None) 
    print("[Producer] Done.")

def consumer():
    while True:
        # This will block if the queue is empty
        item = data_queue.get()
        
        # Check for the stop signal
        if item is None:
            data_queue.task_done()
            break
            
        print(f"[Consumer] Squaring {item}: {item**2}")
        
        # Signal that the specific task is finished
        data_queue.task_done()
        time.sleep(random.random()) # Simulate processing time
    
    print("[Consumer] Done.")

# Create and start threads
t1 = threading.Thread(target=producer)
t2 = threading.Thread(target=consumer)

t1.start()
t2.start()

# Wait for threads to finish
t1.join()
t2.join()

print("All tasks completed.")