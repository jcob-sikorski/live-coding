interface Electronic {
    void turnOn();
}

public class Phone implements Electronic {
    public void turnOn() {
        System.out.println("Phone is booting up...");
    }
}

public class TV implements Electronic {
    public void turnOn() {
        System.out.println("Welcome to Samsung...");
    }
}

public static void main(String[] args){
    Phone phone = new Phone();

    phone.turnOn();
}

// Extends creates an "is-a" relationship (A Dog is-an Animal).
// Implements creates a "can-do" relationship (A Phone can-do electronic things).