fn main(){
    for i in 0..5 {
        let even_odd = if i % 2 == 0 {"even"} else {"odd"};
        println!("{} {}", i, even_odd);
    }

    println!();
    
    let mut sum = 0;
    for i in 0..5 {
        sum += i;
    }
    println!("{}", sum);
}