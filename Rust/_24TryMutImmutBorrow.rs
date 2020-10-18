fn main(){
    let mut a = String::from("Hello");

    let a1 = &a;
    let a2 = &a;
    let a3 = &mut a;

   println!("{} {} {}", a1, a2, a3); 
}
