public class User {
    String name;
    static int userCount;

    public User(String name) {
        this.name = name;
        this.userCount += 1;
    }
}

public static void main(String[] args) {
    User user1 = new User("Bob");
    User user2 = new User("Adam");
    User user3 = new User("Jacob");

    System.out.println(User.userCount);
}

// userCount is static so it's shared between instances
// hence its value is 3 * 1 = 3