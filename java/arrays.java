import java.util.ArrayList;

public static void main(String[] args) {
    String[] namesArray = new String[3];

    namesArray[0] = "Alice";
    namesArray[1] = "Bob";
    namesArray[2] = "Charlie";

    // namesArray[3] = "Daisy"; // This will throw an ArrayIndexOutOfBoundsException!

    ArrayList<String> namesList = new ArrayList<>();

    namesList.add("Alice");
    namesList.add("Bob");
    namesList.add("Charlie");

    namesList.add("Daisy"); // This works perfectly!
}