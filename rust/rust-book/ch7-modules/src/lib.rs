#![allow(dead_code, unused)]  // #! ==> all of crate, (all files)

mod front_of_house; // --> `src/front_of_house.rs` (or `src/front_of_house/mod.rs`)

mod back_of_house {
    fn fix_incorrect_order() {
        cook_order();
        super::deliver_order()
    }
    fn cook_order() {}
}

fn deliver_order() {}

mod customer {
    pub fn eat_at_resturant() {
        // Absolute path
        crate::front_of_house::hosting::add_to_waitlist();
        
        // Relative path
        super::front_of_house::hosting::add_to_waitlist();
    }
    
    use crate::front_of_house::hosting;
    
    pub fn drink_at_resturant() {
        hosting::add_to_waitlist();
    }
}