fn main() {
    let m = method;
    let x = m(3);
    println!("{}", x);

    let m1: fn(i32) -> i32 = method;
    let y = m1(4);
    println!("{}", y);
}

fn method(x: i32) -> i32 {
    x + 1
}



