import threading
import time

# The shared global variable
counter = 0

def increment():
    global counter
    for _ in range(1000): # Lowered range so it finishes fast with sleep
        current_value = counter
        # Force a context switch right here!
        time.sleep(0.000001) 
        counter = current_value + 1

def run_experiment():
    global counter
    counter = 0
    threads = []

    # Create 10 threads
    for i in range(10):
        t = threading.Thread(target=increment)
        threads.append(t)
        t.start()

    # Wait for all threads to finish
    for t in threads:
        t.join()

    print(f"Final Counter Value: {counter}")
    print(f"Expected Value:      1,000,000")
    print(f"Difference:          {1000000 - counter}")

if __name__ == "__main__":
    run_experiment()