public class BinarySearch {
	
	
	public static void main(String[] args) {
		
		int[] values = { 1,2,2,3,3,4,4,4,5,5,5,5,5,5,6,6,6,7};
		int key = 6;

		System.out.println(findKey(values,key));
		System.out.println(findFirstApperanceOfKey(values,key));
		System.out.println(findLastApperanceOfKey(values,key));
		System.out.println(countOfOccurence(values,key));
		
			
	}
	
	
	static int findKey(int[] values,int key){
		int low = 0;
		int high = values.length -1;
		
		while(low <= high){
			int middle = (low + high) / 2;
			
			if(key > values[middle])
				low = middle+1;
			else if(key < values[middle])
				high = middle-1;
			else 
				return middle;
			
		}	
		return -1;
	}
	
	
	static int findFirstApperanceOfKey(int[] values,int key){
		int low = 0;
		int high = values.length -1 ;
		
		while(low <= high){
			int mid = (low+high) / 2;
			if(key == values[mid] && (mid == 0 || key != values[mid-1]))
				return mid;
			else if(key <= values[mid])
				high = mid - 1;
			else
				low = mid + 1;	
		}
		return -1;
	}
	
	static int findLastApperanceOfKey(int[] values,int key){
		int low = 0;
		int high = values.length -1 ;
		int arrayLength = high; 
		
		while(low <= high){
			int mid = (low+high) / 2;			
			if(key == values[mid] && (mid == arrayLength || key != values[mid+1]))
				return mid;
			else if(key >= values[mid])
				low = mid + 1;
			else
				high = mid - 1;		
		}
		return -1;
	}
	
	static int countOfOccurence(int[] values,int key){
		int first = findFirstApperanceOfKey(values,key);
		if(first == -1) return -1;
		
		int last = findLastApperanceOfKey(values,key);
		return last - first + 1;
		
	}

}