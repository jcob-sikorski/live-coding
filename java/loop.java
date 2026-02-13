import java.util.ArrayList;

public static void main(String[]) {
    // String[] namesList = new String[3]; or...
    // String[] namesArray = {"Bob", "Charlie", "Adam"};

    ArrayList<String> namesList = new ArrayList<>();

    namesList.add("Bob");
    namesList.add("Charlie");
    namesList.add("Adam");

    for (int i = 0; i < namesList.size(); i++) {
        System.out.println(namesList.get(i));
    }

    for (String name : namesList) {
        System.out.println(name);
    }
}