// Rust code for linearly searching x in arr[]. If x  
// is present then return its location, otherwise  
// return -1  

fn search(arr: [i32;5],x: i32) -> i32 {
    for (index,item) in arr.iter().enumerate(){
        if *item == x {
            return index as i32;
        }
    }
    return -1;
}

fn main(){
    let arr: [i32;5] = [ 2, 3, 4, 10, 40 ]; 
    let x: i32 = 10;
    let result = search(arr, x);
    if result == -1 {
        println!("Element is not present in array");
    }else{
        println!("Element is present at index {}",result);
    }
}

// This code is contributed by pranavek
