### 1. Go Basics & Control Flow
* **Hello World (`hello`):** Write a basic Go program that prints a string to the console to verify your environment is working.
* **The "Who Are You?" (`scan_input`):** Use `fmt.Scanln()` to capture user input from the terminal and print a personalized greeting. *Bonus:* Ask for their age and calculate how many years until they turn 100.
* **The Bouncer (`is_adult`):** Write an `if/else` statement that checks an integer variable to see if a user is 21 or older.
* **The Short-Circuit (`short_statement`):** Practice Go's idiomatic `if` syntax by initializing a variable and checking its condition on the exact same line.
* **The Weekender (`basic_switch`):** Build a `switch` statement that takes numbers 1-7 and prints the corresponding day of the week, including a `default` fallback.
* **Grade Calculator (`tagless_switch`):** Write a switch statement without a specific variable (acting like a clean `if/else if` chain) to return letter grades based on integer scores.
* **Type Explorer (`explain_type`):** Use a type switch `switch v := i.(type)` to determine if an `interface{}` value is a string, int, or boolean.
* **The Countdown (`for_loop`):** Write a standard 3-part `for` loop that counts backward from 10 to 1.
* **The "While" Sum (`while_double`):** Go doesn't have a `while` keyword. Write a `for` loop with only a condition to double a number until it exceeds 1,000.
* **FizzBuzz (`fizz_buzz`):** The classic interview question. Loop 1 to 50. Print "Fizz" for multiples of 3, "Buzz" for 5, and "FizzBuzz" for multiples of both.
* **Default Values (`default_values`):** Declare an `int`, `float64`, `byte`, and `rune` without assigning values to observe Go's automatic zero-values.
* **ASCII/Byte Loop (`ascii_byte_loop`):** Loop through numbers 65 to 90 and print both the integer and its ASCII character representation to understand how bytes translate to text.

### 2. Slices & Arrays
* **The Empty Canvas (`slice_literal`):** Initialize a slice of strings with 3 fruits using a literal, and print its length.
* **The Grow Operation (`append`):** Start with an empty integer slice and use a loop + `append()` to fill it with numbers 1 through 10.
* **The Window View (`window`):** Slice an existing array using `[start:end]` syntax. Modify the new sub-slice and observe how it affects the original array.
* **Pre-allocation (`allocate`):** Use `make([]int, 0, 5)` to pre-allocate capacity. Append items and print `len()` and `cap()` to watch how Go manages the underlying memory.
* **The Copy Cat (`copy`):** Use the built-in `copy()` function to safely duplicate a slice so that changing the copy does not alter the original.
* **The Filter Challenge (`modify_int_place`):** Write a function that takes a slice and returns only the even numbers. *Bonus:* Try doing this by modifying the slice in-place (reslicing) to avoid allocating memory for a new array.
* **Modify Slice Trick (`modify_slice`):** Pass a slice into a function. Modify an index, and then `append` to it. Observe why the index change reflects in `main`, but the appended item does not (slice headers pass by value).

### 3. Maps (Key-Value Stores)
* **Name Tag (`name_tag`):** Create a map tying string names to integer ages. Zip two existing slices together into this map.
* **The Attendance Sheet (`check_attendance`):** Practice the "comma ok" idiom (`value, ok := map[key]`). Look up a student and safely handle the case where they don't exist in the map.
* **Word Frequency Counter (`freq_counter`):** Loop through a slice of string words and use a map to count how many times each word appears.
* **The Eraser (`eraser`):** Populate a map with 5 items, use the `delete()` function to remove two of them, and print the final length.
* **Dictionary Alphabetizer (`sorted_map`):** Maps are explicitly unordered in Go. Extract map keys into a slice, sort the slice, and use it to print the map's contents alphabetically.
* **Invert the Map (`invert_map`):** Take a map of `[username]team` and flip it to create a new map of `[team][]username` (where values become keys, and keys are grouped into slices).
* **Pointer Map Mutation (`pointer_map`):** Pass a map to a function and modify it. Notice that you *don't* need a pointer (`*map`) for the changes to persist.

### 4. Structs, Methods & Interfaces
* **Basic Struct (`book_struct`):** Define a `Book` struct. Initialize it with named fields and print it using `%+v` to see the field names.
* **Zero-Value Audit (`user_struct`):** Declare a `User` struct without giving it values. Print it to see how Go handles zero-values for complex types.
* **The Anonymous Quickie (`anonymous_struct`):** Define and instantiate a one-off struct for 2D coordinates inside your `main` function without defining a type globally.
* **The Constructor (`construct_user`):** Go has no classes. Write a `NewUser()` function that acts as a constructor, returning a pointer to a freshly initialized struct.
* **Composition (`composition_metadata`):** Embed a `Vehicle` struct directly inside a `Car` struct to see how Go "promotes" the vehicle's fields to be accessible directly on the car.
* **The Nested Manager (`nested_manager`):** Create a `Department` that holds a slice of `Employee` structs. Write a method to calculate total payroll.
* **The Value Receiver (`method`):** Write a method attached to a struct (e.g., `Describe()`) that uses a value receiver to format and return its data.
* **Value vs Pointer Receiver (`struct_receivers` / `update_struct`):** Write two "TakeDamage" or "UpdateAge" methods—one using a value receiver and one using a pointer. Call them and observe why only the pointer receiver permanently changes the struct's data.
* **The Universal Filter (`universal_filter`):** Practice Go Generics (`[T any]`) by writing a single filter function that works on any data type based on a custom test function.
* **The Interface Bridge (`interface_bridge`):** Create a `Shape` interface with an `Area()` method. Have both `Circle` and `Square` implement it, then write a function that accepts any `Shape`.

### 5. Pointers & Memory
* **The Identity Check (`pointer_check`):** Create a pointer. Print the variable's address (`&`), the pointer's value, and the dereferenced value (`*`).
* **The Remote Control (`swap_value` / `modifyValue`):** Write a function that accepts a pointer to an integer and squares/doubles it by dereferencing it.
* **Nil Safety (`nil_pointer_dereference`):** Write a function taking a string pointer. Explicitly check `if ptr == nil` before trying to read it to prevent a panic.
* **Multi-level Indirection (`pointer_pointer`):** Create a pointer to a pointer (`**int`). Dereference it twice to change the base integer's value.
* **Pass by Value (`copy_value`):** Try to mutate a simple integer by passing it to a function normally. Observe how its memory address changes inside the function (proving it's a copy).
* **Escape Analysis (`escape_analysis` / `escape_to_heap`):** Return a pointer to a local variable from a function. Use `go build -gcflags="-m"` to watch the Go compiler explicitly move it from the stack to the heap.
* **Stack vs Heap Limits (`large_array` / `stack_overflow`):** Force data onto the heap by allocating a massive array, and trigger a stack overflow via infinite recursion.
* **Dynamic Resizing (`grow_slice_test`):** Write benchmark tests comparing slice appends with and without pre-allocating capacity to see the performance cost of dynamic memory growth.
* **Explicit Converter (`convert_types_explicit` / `multiply_large_nums`):** Practice explicit type conversions. Safely multiply two `int16` values by upcasting them to `int32` first to avoid integer overflow.

### 6. Defer & Error Handling
* **LIFO Defer (`defer_lifo`):** Use `defer` inside a loop to print characters of a string. Watch them print in reverse order (Last-In, First-Out).
* **Defer Snapshots (`snapshot`):** Notice how `defer fmt.Println(x)` captures the value of `x` at the exact moment the defer is *declared*, not when it is *executed*.
* **Safe File Copier (`copy_file`):** Emulate real-world resource management. Open a file, immediately `defer file.Close()`, and copy data safely.
* **Safe Wrapper (`safe_wrapper`):** Write a function that uses `defer` and `recover()` to catch a forced `panic`, allowing the application to survive instead of crashing.

### 7. Goroutines & Channels
* **The Ghost Routine (`ghost_goroutine`):** Launch a basic `go func()`. Notice why `main` exits before the goroutine can print anything unless you add a sleep.
* **WaitGroup Synchronizer (`waitgroup_synchronizer`):** Fix the above problem the "right" way using `sync.WaitGroup` (`Add`, `Done`, and `Wait`) instead of sleeping.
* **The Loop Trap (`loop_trap` / `solve_loop_trap`):** Launch goroutines inside a `for` loop. Observe the classic bug where they all print the last loop index, then fix it by passing the index into the goroutine as an argument.
* **Ping-Pong (`ping_pong`):** Create a basic unbuffered channel. Have `main` send "ping" into it, and a separate goroutine receive and print it.
* **Buffered Buffer (`buffered_buffer`):** Create a channel with a capacity of 3. Send 3 items into it from `main` without blocking, proving that buffered channels are asynchronous until full.
* **The Range Loop (`close_channel`):** Send values into a channel and then `close()` it. In `main`, use `for val := range ch` to safely read until the channel is empty.
* **Directional Gatekeeper (`sender_receiver`):** Write functions that accept `chan<-` (send-only) and `<-chan` (receive-only) to enforce strict data-flow directions at compile time.
* **The Signal to Stop (`stop_goroutine` / `quit_signal`):** Use an empty struct channel (`chan struct{}`) to cleanly signal an infinite-loop worker goroutine to pack up and shut down.
* **Basic Listener (`basic_listener`):** Use a `select` statement to listen to multiple channels simultaneously, printing whichever one receives data first.
* **Timeout Pattern (`timeout_pattern`):** Inside a `select`, use `time.After()` to abort a channel read if it takes longer than 500ms.
* **Randomized Fairness (`randomized_fairness`):** Read from two fully-loaded channels via a `select` loop to observe Go's built-in pseudo-random fairness algorithm when multiple cases are ready.