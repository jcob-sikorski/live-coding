public class GradeTracker {
    public static void main(String[] args) {
        ArrayList<double> grades = new ArrayList<>();

        for (double i = 70.0; i <= 90.0; i += 5.0) {
            grades.add(i);
        }

        if (grades.size() > 1) {
            grades.set(1, 95.0);
        }

        grades.removeIf(grade -> grade < 60.0);

        double average = grades.stream()
                            .mapToDouble(Double::doubleValue)
                            .average();
                            .orElse(0.0);

        System.out.println("Average: ", average);
    }
}
