public class BinarySearch {
	
	
	public static void main(String[] args) {
		
		int[] values = IntStream.range(0, 100).toArray();
		int[] keys = {50,23,34,12};
		
		BinarySearch bs = new BinarySearch();
		for(int key: keys){
			System.out.println("************************");
			bs.findKey(values,key);
		}
			
	}
	
	
	int findKey(int[] values,int key){
		int low = 0;
		int high = values.length -1;
		
		while(low <= high){
			int middle = (low + high) / 2;
			
			System.out.println(middle+" "+low+" "+high);
			
			if(key > values[middle])
				low = middle + 1;
			else if(key < values[middle])
				high = middle - 1;
			else 
				return middle;
			
		}
		
		return -1;
		
	}

}
