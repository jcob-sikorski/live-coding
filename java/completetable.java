CompletableFuture.supplyAsync(() -> "hello")
                 .thenApply(String::toUppercase)
                 .thenApply(x -> x + " WORLD")
                 .thenAccept(System.out::println);


ExecutorService executor = Executors.newFixedThreadPool(2);
long startTime = System.currentTimeMillis();

CompletableFuture<String> grinding = CompletableFuture.supplyAsync(() -> {
    try { Thread.sleep(2000); } catch (Exception e) {}
    return "Ground Coffee";
}, executor);

// 1. Create the boiling future
// 2. Combine them
// 3. Print result and (System.currentTimeMillis() - startTime)