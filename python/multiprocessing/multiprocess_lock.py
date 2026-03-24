# This is the high-level, standard way to share a single variable.
# It includes a built-in lock to prevent "race conditions" (where two
# processes try to update the number at the exact same time and lose data).

import multiprocessing

def increment_counter(shared_num, loops):
    for _ in range(loops):
        # We must use a lock to ensure the increment is atomic
        with shared_num.get_lock():
            shared_num.value += 1

if __name__ == "__main__":
    # 'i' stands for signed integer. Initial value is 0.
    counter = multiprocessing.Value('i', 0)
    
    p1 = multiprocessing.Process(target=increment_counter, args=(counter, 1000))
    p2 = multiprocessing.Process(target=increment_counter, args=(counter, 1000))

    p1.start()
    p2.start()
    p1.join()
    p2.join()

    print(f"Final counter value: {counter.value}")