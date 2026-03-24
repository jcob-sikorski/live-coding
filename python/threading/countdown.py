# The Timer Challenge: Write a script that starts a "Countdown" thread (10 to 1)
# and a "Main" thread that asks the user for their name. See if you can finish
# typing your name before the countdown hits 0.

import threading
import time
import sys

# A flag to let the countdown know the user finished
user_finished = False

def countdown():
    global user_finished
    print("\n[Timer] Countdown started! You have 10 seconds...\n")
    
    for i in range(10, 0, -1):
        if user_finished:
            return  # Stop counting if the user finished
        
        # Using sys.stdout.write to try and keep the prompt visible
        sys.stdout.write(f"\n[ {i} ] ... ")
        sys.stdout.flush()
        time.sleep(1)

    if not user_finished:
        print("\n\n[!!!] TIME IS UP! You lost the challenge.")

# Create the thread
timer_thread = threading.Thread(target=countdown)
timer_thread.daemon = True # This ensures the thread dies if the main program exits
timer_thread.start()

# Main thread input
name = input("QUICK! Type your name: ")
user_finished = True

if name:
    print(f"\nSuccess! Nice to meet you, {name}.")
else:
    print("\nYou submitted an empty name? That's cheating (or a very short name)!")