import java.util.*;

interface Borrowable {
    void borrowItem(String title) throws Exception;
}

public class Library implements Borrowable {
    // BUG 1: Should this be static? 
    // If it is, every library in the world shares the same books!
    private ArrayList<String> books = new ArrayList<>();

    public void addBook(String title) {
        this.books.add(title);
    }

    public int getSize() {
        return this.books.size();
    }

    @Override
    public void borrowItem(String title) throws Exception {
        if (this.books.contains(title)) {
            this.books.remove(title);
            System.out.println("Successfully borrowed: " + title);
        } else {
            throw new Exception("Book not found: " + title);
        }
    }
}

public class Main {
    public static void main(String[] args) {
        Library myLibrary = new Library();

        myLibrary.addBook("Java 101");
        myLibrary.addBook(String.valueOf(12345)); // BUG 2: This shouldn't be allowed! 

        // BUG 3: Direct access to 'books' is dangerous.
        System.out.println("Books count: " + myLibrary.getSize());

        // BUG 4: What if we want a 'Member' who can 'borrow'?
        // Let's implement an Interface for 'Borrowable'
        try {
            System.out.println("Books count: " + myLibrary.getSize());

            myLibrary.borrowItem("Java 101");

            myLibrary.borrowItem("Python Basics"); 

        } catch (Exception e) {
            // This catches the "Book not found" error!
            System.out.println("Error: " + e.getMessage());
        }

        System.out.println("Final Books count: " + myLibrary.getSize());
    }
}