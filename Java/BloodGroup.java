import java.util.*;
public class BloodGroup{

Map<String,Float> bloodMap = new HashMap<String,Float>();

public static void main(String[] args){
	new BloodGroup().bloodGroupPercentage();
}

public void bloodGroupPercentage(){
String bloodGroupString = "O O A B A O A A A O B O B O O A O O A A A A AB A B A A O O A O O A A A O A O O AB";
String[] bloodGroups = bloodGroupString.split("\\s");
for(String group : bloodGroups){
	Float a = bloodMap.get(group);
	if(null == a){
		bloodMap.put(group,1F);
		continue;
	}
	a++;
	bloodMap.put(group,a);
}
int total = bloodGroups.length;
//System.out.println(bloodMap.values().stream().mapToInt(Number::intValue).sum());
bloodMap.forEach( (k,v) -> System.out.println(k+"|"+v+"|"+(v/total*100)) );

}

}
