public class Palindrome {
	
	
	public static void main(String[] args) {
		
		String[] inputs = {"MOM","DAD","MOMMY","DADDY"};
		for(String input : inputs){
			System.out.print(palindrome(input)+" "+palindromeBufferReverse(input)+" "+palindromeCharAt(input));
			System.out.println();
		}
		
		
	}
	
	
	static boolean palindrome(String input){
		char[] inputChars = input.toCharArray();
		StringBuffer br =  new StringBuffer();
		for(int i = inputChars.length - 1; i >= 0; i--){
			br.append(inputChars[i]);
		}		
		return input.equals(br.toString());
	}
	
	static boolean palindromeCharAt(String input){
		StringBuilder sb = new StringBuilder();
		for(int i = input.length() -1 ; i >= 0;i--){
			sb.append(input.charAt(i));
		}
		return input.equals(sb.toString());
		
	}
	
	static boolean palindromeBufferReverse(String input){
		return input.equals(new StringBuffer(input).reverse().toString());
	}
	

}