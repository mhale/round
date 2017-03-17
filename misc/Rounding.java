import java.math.*;

public class Rounding {
	public static final int MIN_WIDTH = 4;  // Minimum width to fit "-x.y" output

	public static void main(String[] args) {
		double[] values = {5.5, 2.5, 1.6, 1.1, 1.0, -1.0, -1.1, -1.6, -2.5, -5.5};

		System.out.print("   x");
		for (RoundingMode mode : RoundingMode.values()) {
				System.out.format("  %4s", mode);
		}
		System.out.println();

		for (double x : values) {
			System.out.format("%4.1f", x);
			for (RoundingMode mode : RoundingMode.values()) {
				try {
					BigDecimal bd = new BigDecimal(x).setScale(0, mode);
					int width = Math.max(mode.toString().length(), MIN_WIDTH);
					System.out.format("  %" + width + ".1f", bd);
				} catch (Exception ex) {
					String name = ex.getClass().getSimpleName();
					System.out.format("  " + name);
				}
			}
			System.out.println();
		}

	}
}
