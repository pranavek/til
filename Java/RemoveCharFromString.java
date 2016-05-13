public class RemoveCharFromString {
	
	public static void main(String[] args) {
		String value = "Pranav";
		String removeChar = "a";
		
		System.out.println(replaceAll(value,removeChar));
		System.out.println(charAtReplace(value,removeChar.charAt(0)));
		
		
		
	}
	
	
	static String charAtReplace(String value,char removeChar){
		StringBuffer sb = new StringBuffer();
		for(int i= 0;i<=value.length()-1;i++){
			char val = value.charAt(i);
			if(removeChar == val)
				continue;
			sb.append(val);
		}
		return sb.toString();
	}
	
	static String replaceAll(String value,String removeChar){
		return  value.replaceAll(removeChar,"");
	}

}