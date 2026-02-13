import java.util.Arrays;
import java.util.stream.DoubleStream;

public class TemperatureLogger {
    public static void main(String[] args) {
        double[] weekTemps = new double[7];

        Arrays.setAll(weekTemps, i -> 20.0 + i);

        double maxTemp = Arrays.stream(weekTemps)
                               .max()
                               .orElse(0.0);

        System.out.println("Highest Temp: " + maxTemp);

        double[] reversedTemps = new double[weekTemps.length];
        for (int i = 0; i < weekTemps.length; i++) {
            reversedTemps[i] = weekTemps[weekTemps.length - 1 - i]
        }
    }
}