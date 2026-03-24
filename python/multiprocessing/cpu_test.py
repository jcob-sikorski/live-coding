import time
from multiprocessing import Process

def cpu_burner(n=10_000_000):
    """Calculates the sum of squares up to n."""
    count = 0
    for i in range(n):
        count += i**2
    return count

def run_sequential(iterations):
    print(f"--- Starting Sequential Run ({iterations} times) ---")
    start_time = time.perf_counter()
    for _ in range(iterations):
        cpu_burner()
    end_time = time.perf_counter()
    print(f"Sequential Duration: {end_time - start_time:.4f} seconds\n")

def run_multiprocessing(iterations):
    print(f"--- Starting Multiprocessing Run ({iterations} processes) ---")
    processes = []
    start_time = time.perf_counter()

    # Create and start processes
    for _ in range(iterations):
        p = Process(target=cpu_burner)
        processes.append(p)
        p.start()

    # Wait for all processes to finish
    for p in processes:
        p.join()

    end_time = time.perf_counter()
    print(f"Multiprocessing Duration: {end_time - start_time:.4f} seconds\n")

if __name__ == "__main__":
    # 1. Run 4 times sequentially
    run_sequential(4)

    # 2. Run 4 times using multiprocessing
    run_multiprocessing(4)