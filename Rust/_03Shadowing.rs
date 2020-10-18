fn main(){

    let mut mutable: bool = true;
    println!("{}", mutable);
    mutable = false; //Try commenting this to get unused_mut warning
    println!("{}", mutable);
    let mutable = 21;
    println!("{}", mutable);


}