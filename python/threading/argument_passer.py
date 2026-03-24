# The Argument Passer: Create a function that accepts a username. Launch 5 threads,
# each passing a different name to the function to be printed.

import threading

def task(name: str):
    print(name)

names = ["A", "B", "C", "D", "E"]
threads = []

for i in range(5):
    t = threading.Thread(target=task, args=(names[i]))
    threads.append(t)
    
for i in range(5):
    threads[i].start()
    
for i in range(5):
    threads[i].join()