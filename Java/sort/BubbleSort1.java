public class BubbleSort1 {

	public static void main(String[] args) {
		int[] input = { 9, 3, 12, 6 };

		// Ascending
		for(int i = 0; i < input.length; i++){
			for(int j = 1 ;j< input.length - i ;j++){
				if(input[j-1] > input[j]){
					int temp = input[j];
					input[j] = input[j-1];
					input[j-1] = temp;
				}
			}
		}
		
		
		for (int i : input) {
			System.out.print(i + "\t");
		}

	}

}