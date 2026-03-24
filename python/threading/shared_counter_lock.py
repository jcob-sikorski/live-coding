import threading
import time

counter = 0
counter_lock = threading.Lock()

def increment():
    global counter
    for _ in range(100):
        # The Lock acts as a 'talking stick'
        with counter_lock: 
            current_value = counter
            time.sleep(0.000001) # Even with a sleep, the lock is HELD
            counter = current_value + 1

def run_experiment():
    global counter
    counter = 0
    threads = []
    num_threads = 10
    iterations = 100
    expected = num_threads * iterations

    for i in range(num_threads):
        t = threading.Thread(target=increment)
        threads.append(t)
        t.start()

    for t in threads:
        t.join()

    print(f"Final Counter Value: {counter}")
    print(f"Expected Value:      {expected}")
    print(f"Difference:          {expected - counter}")

if __name__ == "__main__":
    run_experiment()