abstract class Shape {
    void displayName() {
        System.out.println("This is shape.");
    }

    abstract double calculateArea();
}

public class Circle extends Shape {
    private double radius;

    public Circle(double radius) {
        this.radius = radius;
    }

    @Override
    public double calculateArea() {
        return Math.PI * radius * radius;
    }
}

public class Main {
    public static void main(String[] args) {
        Circle circle = new Circle(5.0);
        circle.displayName();

        double area = circle.calculateArea();
        System.out.println("Area: " + area);
    }
}