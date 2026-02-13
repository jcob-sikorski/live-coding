public class Calculator {
    // we can also use static to show that we are 
    // not modifying the internal state of the class and these are just helper methods
    public int add(int a, int b) {
        return a + b;
    }

    public int add(int a, int b, int c) {
        return a + b + c;
    }

    public double add(double a, double b) {
        return a + b;
    }
}

public void static main(String[] args) {
    Calculator calculator = new Calculator();

    a = calculator.add(1, 2);
    b = calculator.add(1, 2, 3);
    c = calculator.add(10.5, 20.2);
}