use adder::add;

mod common;

#[test]
fn it_adds() {
    assert_eq!(4, add(2, 2))
}

#[test]
fn test_adds_with_setup() {
    common::setup();
    assert_eq!(10, add(5, 5))
}