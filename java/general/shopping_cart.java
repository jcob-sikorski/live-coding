import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

// Exercise 1: The "Smart" Shopping Cart
// Goal: Test encapsulation, logic placement, and use of BigDecimal.

// Java Knowledge Check: Why should you use BigDecimal instead of double for currency? How do you handle a list of items efficiently?

class Product {
    // Requirement: Create a Product class with a name and a price.
    private final String name;
    private final BigDecimal price;

    public Product(String name, BigDecimal price) {
        this.name = name;
        this.price = price;
    }

    public BigDecimal getPrice() {
        return this.price;
    }
}

class Cart {
    // Requirement: Create a Cart class that allows adding products.
    private final List<Product> products = List<>();

    // The cart must calculate the total price, but it must apply a 10% discount if the total exceeds $100.
    public void addProduct(Product product) {
        this.products.add(product);
    }

    public List<Product> getProducts() {
        return Collections.unmodifiableList(products);
    }

    public double calculateRawTotal() {
        return this.products.stream()
                            .map(Product::getPrice)
                            .reduce(BigDecimal.ZERO, BigDecimal::add);
    }
}

class DiscountService {
    private static final BigDecimal DISCOUNT_THRESHOLD = new BigDecimal("100.00");
    private static final BigDecimal DISCOUNT_RATE = new BigDecimal("0.90");

    public static applyDiscount(Cart cart) {
        // if sum of all products in the cart > $100; return a 0.9 * price else price
        BigDecimal total = cart.calculateRawTotal();

        if (total.compareTo(DISCOUNT_THRESHOLD) > 0) {
            return total.multiply(DISCOUNT_RATE).setScale(2, RoundingMode.HALF_UP);
        }
        return total.setScale(2, RoundingMode.HALF_UP);
    }
}