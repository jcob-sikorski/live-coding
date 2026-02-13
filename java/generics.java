public class Box<T> {
    private T contents;

    public Box(T contents) {
        this.contents = contents;
    }

    public T getContents() {
        return this.contents;
    }
}

public static void main(String[] args) {
    Box<String> nameBox = new Box<>("Gemini");
    Box<Integer> numberBox = new Box<>(100);

    System.out.println(nameBox.getContents());
    System.out.println(numberBox.getContents());
}