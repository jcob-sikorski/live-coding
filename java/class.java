class Book {
    String title;
    double price;

    public Book(String title, double price) {
        this.title = title;
        this.price = price;
    }
}

public static void main(String[] args) {
    Book myBook = new Book("Java Basics", 29.99)

    System.out.println("Title: " + myBook.title);
    System.out.println("Price: " + myBook.price);
}