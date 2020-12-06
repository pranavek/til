fn wifi_status(power: bool) {
    if power == false { panic!("May day!"); }
    println!("Turned on!");
} 

fn main() {
    wifi_status(true);
    wifi_status(false);
}
