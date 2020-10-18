use std::io;

fn main(){
    let mut count = String::new();
    match io::stdin().read_line(&mut count){
        Ok(n) => {
            println!("The size is {}",n);
            println!("The count is {}",count);
        }
        Err(e) => {
            println!("The error is {}",e);
        }
    }
}
     
