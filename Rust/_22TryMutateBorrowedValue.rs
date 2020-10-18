fn main(){
    let s = String::from("Hi");
    change(&s);
}

fn change(s: &String){
    s.push_str("Pranav");
}

