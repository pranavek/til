fn main(){
    let string = String::from("Hi there");
    //let len = calculate_len(string);
    let (string,len) = calculate_len1(string);
    println!("Length of {} is  {}",string, len);
}

fn calculate_len(s: String) -> usize {
    let l = s.len();
    l
}

fn calculate_len1(s: String) -> (String,usize) {
    let l = s.len();
    (s,l)
}
