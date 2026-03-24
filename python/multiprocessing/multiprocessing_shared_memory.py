# Introduced in Python 3.8, this is a lower-level approach.
# It maps a block of raw bytes into the memory space of both processes.
# This is much faster for large datasets but requires you to manage the data types and synchronization manually.

import multiprocessing
from multiprocessing import shared_memory
import struct

def increment_shm(name, loops, lock):
    existing_shm = shared_memory.SharedMemory(name=name)
    
    # FIX: Tell Pylance we are sure 'buf' is not None
    assert existing_shm.buf is not None
    
    for _ in range(loops):
        with lock:
            # Pylance is now happy to let you subscript this
            current_value = struct.unpack('i', existing_shm.buf[:4])[0]
            new_value = current_value + 1
            existing_shm.buf[:4] = struct.pack('i', new_value)
            
    existing_shm.close()

if __name__ == "__main__":
    shm = shared_memory.SharedMemory(create=True, size=4)
    
    # FIX: Guard the initialization too
    assert shm.buf is not None
    shm.buf[:4] = struct.pack('i', 0)
    
    lock = multiprocessing.Lock()

    p1 = multiprocessing.Process(target=increment_shm, args=(shm.name, 1000, lock))
    p2 = multiprocessing.Process(target=increment_shm, args=(shm.name, 1000, lock))

    p1.start()
    p2.start()
    p1.join()
    p2.join()

    # FIX: One last check for the final print
    assert shm.buf is not None
    final_val = struct.unpack('i', shm.buf[:4])[0]
    print(f"Final SHM value: {final_val}")

    shm.close()
    shm.unlink()