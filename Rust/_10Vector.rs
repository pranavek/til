fn vector(){
    let mut v = Vec::new();
    v.push(10);
    v.push(39);
    v.push(12);

    let first = v[0];
    let maybe_first = v.get(0);

    println!("The vector is {:?}", v);
    println!("first is {}", first);
    println!("maybe_first is {:?}", maybe_first);
}

fn main(){
    vector();
}