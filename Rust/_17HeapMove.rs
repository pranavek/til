// s1's value is moved to s2 and we are trying to access s1 from the heap, which is not possible as the s1 is dereferenced
fn main() {
    let s1 = String::from("Hi");
    let s2 = s1;
    println!("s1 is {} and s2 is {}", s1,s2);
}
