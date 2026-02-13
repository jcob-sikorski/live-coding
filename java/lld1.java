import java.util.*;

enum Coin {
    NICKEL(5), DIME(10), QUARTER(25);

    private final int value;

    // The only time the constructor is called is inside
    // the enum itself when it defines the constants
    // (NICKEL, DIME, QUARTER). Once the program starts,
    // those are the only "Coin" objects that will ever exist.
    Coin(int value) {
        this.value = value;
    }

    public int getValue() {
        return value;
    }
}

public class Item {
    private final String name;
    private final int price;

    public Item(String name, int price) {
        this.name = name;
        this.price = price;
    }

    public String getName() {
        return this.name;
    }

    public int getPrice() {
        return this.price;
    }
}

public class Inventory<T> {
    private Map<T, Integer> inventory = new HashMap<>();

    public void add(T item) {
        int count = inventory.getOrDefault(item, 0);
        inventory.put(item, count + 1);
    }

    public int getQuantity(T item) {
        return inventory.getOrDefault(item, 0);
    }
}

public class Bucket<T1, T2> {
    private final T1 item;
    private final T2 change;

    public Bucket(T1 item, T2 change) {
        this.item = item;
        this.change = change;
    }

    public T1 getFirst() {
        return this.item;
    }

    public T2 getSecond() {
        return this.change;
    }
}

interface VendingMachine {
    long selectItemAndGetPrice(Item item);
    void insertCoin(Coin coin);
    List<Coin> refund();
    Bucket<Item, List<Coin>> collectItemAndChange();
}

public class PhysicalVendingMachine implements VendingMachine {
    private Inventory<Item> itemInventory;
    private Inventory<Coin> cashInventory;
    private long currentBalance;
    private Item currentItem;

    @Override
    public void insertCoin(Coin coin) {
        throw new UnsupportedOperationException("Not implemented yet.");
    }
}