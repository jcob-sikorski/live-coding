// Task 1: Write a method calculateTotalRevenueForCategory(List<Order> orders, String category)
// that returns the total revenue (sum of price * quantity) for all products in a specific category,
// but only for orders that have the status "DELIVERED".

// Constraints:

// Do not use a for loop.

// Handle the case where no products match (return 0.0).

import java.time.LocalDate;
import java.util.List;

record Product(String id, String category, double price) {}

record OrderLine(Product product, int quantity) {
    public double lineTotal() { return product.price() * quantity; }
}

record Order(String id, LocalDate orderDate, List<OrderLine> lines, String customerId, String status) {}

double calculateTotalRevenueForCategory(List<Order> orders, String category) {
    return orders.stream()
                 .filter(order -> "DELIVERED".equals(order.status()))
                 .flatMap(order -> order.lines().stream())
                 .filter(line -> line.product().category().equals(category))
                 .mapToDouble(line -> line.lineTotal())
                 .sum()
}