struct Door{
    is_open: bool
}

impl Door{
    fn new(is_open: bool) -> Door {
        Door {is_open: is_open}
    }
}

trait Openable{
    fn open(&mut self);
}

impl Openable for Door {
    fn open(&mut is_open){
        self.is_open = true;
    }
}

#[cfg(test)]
mod tests{
    use super::*;

    #[test]
    fn door_open(){
        let mut door = Door::new(false);
        door.open();
        assert!(door.is_open);
    }
}