import java.util.ArrayList;
import java.lang.reflect.Field;

public class Main {
    public static void main(String[] args) throws Exception {
        ArrayList<Integer> elements = new ArrayList<>(10);

        // 1. Get the Field object first
        Field dataField = ArrayList.class.getDeclaredField("elementData");
        dataField.setAccessible(true);

        for (int i = 0; i < 15; i++) {
            elements.add(i);
            
            // 2. Access the internal array *after* adding the element
            Object[] internalArray = (Object[]) dataField.get(elements);
            
            // 3. Use '+' to concatenate strings in println
            System.out.println("Size: " + elements.size() + 
                               " | Internal Capacity: " + internalArray.length);
        }
    }
}