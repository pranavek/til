import java.util.Arrays;

public class BubbleSort2 {

	public static void main(String[] args) {
		int[] input = { 9, 3, 12, 6 };
		boolean flag = true;
		while (flag) {
			flag = false;
			for (int i = 0; i < input.length - 1; i++) {
				if (input[i] > input[i + 1]) {
					int temp = input[i];
					input[i] = input[i + 1];
					input[i + 1] = temp;
					flag = true;
				}
			}
		}

		Arrays.stream(input).forEach(System.out::println);

	}

}