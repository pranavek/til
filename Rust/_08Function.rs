use std::f64::consts;

fn cosine()-> f64{
    let x = 2.0 * consts::PI;
    let abs_difference = (x.cos() - 1.0).abs();
    assert!(abs_difference < 1e-10);
    abs_difference
}

fn sqrt(x: f64) -> f64 {
    x * x[]
}

// What is the complexity of recursive function
// Implement this fn using non-recursive function
fn factorial(x: u64) -> u64 {
    if x == 0 {
        1
    }else{
        x * factorial(x - 1)
    }
}

fn add(x: &u32) -> u32 {
    *x + 1
}

fn main(){
    let s = sqrt(4.0);
    println!("The square root is {}", s);
    println!("The factorial is {}", factorial(4));
    println!("The sum is {}",add(&2));
    let i = 10;
    println!("The sum is {}",add(&i));
    println!("The sum is {}",i);
    println!("The abs diff is {}",cosine());
}