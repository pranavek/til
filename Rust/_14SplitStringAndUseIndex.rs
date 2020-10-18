use std::io;

fn main(){
    let mut input_array = String::new();
    io::stdin().read_line(&mut input_array).expect("Error");

    let list: Vec<_> = input_array.split(' ').collect();
    for i in &list[1..]{
        println!("{}",&list[2]);
    }
}
