// Exercise 2: The Notification Dispatcher
// Goal: Test interfaces, polymorphism, and the "Open/Closed" principle.

// The Task: Design a system that sends notifications to users.

// Java Knowledge Check: How do you iterate through the list? Can you use the Stream API to make it more concise?

record User(String name) {}

interface NotificationService {
    // Requirement: Define an interface NotificationService with a method sendNotification(String message, User user).

    void sendNotification(String message, User user);
}

class EmailService implements NotificationService {
    // Requirement: Implement two versions: EmailService and SmsService. (Just use System.out.println to simulate the send).

    @Override
    void sendNotification(String message, User user) {
        System.out.println("email: " + message + " sent to user: " + user.name());
    }
}

class SmsService implements NotificationService {
    // Requirement: Implement two versions: EmailService and SmsService. (Just use System.out.println to simulate the send).

    @Override
    void sendNotification(String message, User user) {
        System.out.println("sms: " + message + " sent to user: " + user.name());
    }
}

class NotificationManager {
    // Requirement: Create a NotificationManager that takes a List of services and sends a message through all of them at once.

    private final List<NotificationService> services = new List<>();

    public NotificationManager(List<NotificationService> services) {
        this.services = services;
    }

    public void notifyAll(String message, User user) {
        // TODO: Use a loop or a Stream to call sendNotification on every service in 'services'
        this.services.forEach(service -> service.sendNotification(message, user));
    }
}