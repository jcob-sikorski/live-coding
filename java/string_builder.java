public static void main(String[] args) {
    StringBuilder s = new StringBuilder();

    for (int num = 0; num < 10; num++) {
        // You can append the integer directly!
        s.append(num);
    }

    String result = s.toString();
    System.out.println("Final String: " + result); 
    // Output: 0123456789
}