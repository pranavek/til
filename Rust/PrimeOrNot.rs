//use std::process;

fn main(){

    let x : u32 = 13;

    //let mut result : bool = for i in 2..x {}
    //let result : bool = is_prime(x);
    println!("Is {} prime? {}",x, is_prime(x));
}

fn is_prime(x : u32) -> bool {
    if x <= 1 {
        return false;
    }
    let limit = (x as f64).sqrt() as u32;
    for i in 2..=limit {
        if x % i == 0 {
            return false;
        }
    }
    return true;
}