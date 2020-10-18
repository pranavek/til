fn main(){
    let s = String::from("Hi");
    let len = cal_len(&s);
    println!("Size of {} is {}", s,len);
}

fn cal_len(s: &String)-> usize{
    s.len()
}
