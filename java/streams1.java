List<String> names = Arrays.asList("Alice", "Bob", "Charlie", "David", "Edward", "Anna");

List<String> upperNames = names.stream()
     .filter(x -> x.startsWith("A"))
     .peek(System.out::println)       // "Peeks" at the data and prints it
     .map(String::toUpperCase)        // Transforms it
     .collect(Collectors.toList());   // Actually finishes the stream

long countNames = names.stream()
                        .filter(x -> x.length() > 4)
                        .count()

List<String> threeNames = names.stream()
                               .sorted()
                               .limit(3)
                               .toList()