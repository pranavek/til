use std::cmp;

fn main() {

    //Create a vector of boolean
    let len = 2;
    let vec = vec![false; len];
    for x in vec {
        println!("{}",x);
    }

    //Get char at an index
    let str = "abc";
    let mut chars = str.chars();
    println!("{}", chars.nth(0).unwrap());

    // Use index in the loop
    for(i,_j) in (0..6).enumerate() {
        for(k,_l) in (i..6).enumerate() {
            println!("i={} and k={}",i,k);
        }        
    }

    let mut res = 0;
    res = cmp::max(1,5);
    println!("The max value is {}",res);
    assert_eq!(5, res);

}

