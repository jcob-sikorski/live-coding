class Car {
    private String brand;
    private int speed;

    public Car(String brand, int speed) {
        this.brand = brand;

        if (speed > 0) {
            this.speed = speed;
        } else {
            this.speed = 0;
        }
    }

    public String getBrand() {
        return brand;
    }

    public void setSpeed(int newSpeed) {
        if (newSpeed > 0) {
            this.speed = newSpeed;
        } else {
            System.out.println("Error: Speed cannot be negative.");
        }
    }
}

public static void main(String[] args) {
    Car car = new Car("Chevrolet", -50);

    car.setSpeed(-50);

    System.out.println("Brand: " + car.brand);
    System.out.println("Speed: " + car.speed);
}