// Task: Write a method findTopCustomers(List<Order> orders, int n) that returns a list
// of the top n customerIds, ranked by their total spending across all their orders.

// Constraints:

// You must group the orders by customer.

// You must sum the total value of all orders for each customer.

// The output should be a List<String> of customer IDs, sorted descending by total spend.

import java.time.LocalDate;
import java.util.List;

record Product(String id, String category, double price) {}

record OrderLine(Product product, int quantity) {
    public double lineTotal() { return product.price() * quantity; }
}

record Order(String id, LocalDate orderDate, List<OrderLine> lines, String customerId, String status) {}

List<String> findTopCustomers(List<Order> orders, int n) {
    // groupingBy

    return orders.stream()
                 .groupingBy(customerId)
}