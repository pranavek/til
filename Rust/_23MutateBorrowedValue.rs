fn main(){
    let mut s = String::from("Hi");
    change(&mut s);
}

fn change(s: &mut String){
    s.push_str("Pranav");
}

