fn main() {
    let x = 5;
    
    let add = |i:i32| -> i32 { i + i };
    println!("{}", add(x));

    let mul = |i| { i * i };
    println!("{}", mul(x));

    let add1 = |i: i32| -> i32 { i + i }(x);
    println!("{}", add1);

    let mul1 = |i| -> i32 { i * i }(x);
    println!("{}", mul1);
}
