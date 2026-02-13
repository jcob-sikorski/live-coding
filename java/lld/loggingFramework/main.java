// logger is going to be a singleton

// for the behavior of the logger we're going to use a strategy pattern

// when it comes to logging messages with three different types i'd either use
// strategy (to swap the algorithm for an entire format)
// or a decorator pattern to decorate the logs

// when it comes to outputting to multiple destination we are going to use
// observer pattern where the subject is the logger and observers
// are the destinations

// logging needs to be threat

// for extensibility we're going to use a chain of responsibility pattern