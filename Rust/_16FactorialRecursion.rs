fn factorial(n: u64) -> u64 {
    match n {
        1 => 1,
        n => n * factorial(n-1)
    }
}

fn main(){
    let n: u64 = 3;
    println!("{}",factorial(n));
}

