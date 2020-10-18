//Try creating an array with dynamic size by taking the size as input from cli

fn array(){
    let a = [9,10,3,5,10,20];
    for i in 0..a.len(){
        println!("[{}]={}",i,a[i]);
    }
}

fn sum(a: &[u32]) -> u32{
    let mut sum = 0;
    for i in 0..a.len() {
        sum += a[i];
    }
    sum
}

fn main(){
    array();
    let a = [10,20,30];
    println!("The array is {:?}",a);
    println!("The sum is {}", sum(&a));

    let slice1 = &a[1..];
    //let slice1 = &a[1..10]; //create a panick at runtime - by giving out of bound index
    println!("The slice of array a is {:?}", slice1);

    let slice2 = &a;
    let first = slice2.get(0);
    let last = slice2.get(10);
    println!("Slice of index 0 is {:?}", first); 
    println!("The non-panick way, return option value -  {:?}", last);
    
    println!("first {} {}", first.is_some(), first.is_none());
    println!("last {} {}", last.is_some(), last.is_none());
    println!("first value {}", first.unwrap());
    //println!("last value {}", last.unwrap()); // Will panick as the value is None
    
    let may_be_last = slice2.get(10);
    let last = if may_be_last.is_some() {
        *may_be_last.unwrap()
    }else{
        0
    };
    println!("The tenth index value is {}", last);

    //The above logic can be replaced
    let last = slice2.get(1).unwrap_or(&0);
    println!("The tenth index is {}", last);


    
    
}