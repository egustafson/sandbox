fn main() {
    lister();
    raw_deref();
    my_box1();
    my_box2();
    custom_drop();
    rc_lister();
}

// ----------------------------------------

use std::rc::Rc;
#[allow(dead_code)]
enum RcList {
    Cons(i32, Rc<RcList>),
    Nil,
}

#[allow(unused)]
fn rc_lister() {
    use crate::RcList::{Cons, Nil};
    let a = Rc::new(Cons(5, Rc::new(Cons(10, Rc::new(Nil)))));
    println!("a's ref-count after creating a is {}", Rc::strong_count(&a));
    let b = Cons(3, Rc::clone(&a));
    println!("a's ref-count after creating b is {}", Rc::strong_count(&a));
    {
        let c = Cons(4, Rc::clone(&a));
        println!("a's ref-count after creating c is {}", Rc::strong_count(&a));
    }
    println!("a's ref-count after c goes out of scope is {}", Rc::strong_count(&a));
}

// ----------------------------------------
#[allow(dead_code)]
enum List {
    Cons(i32, Box<List>),
    Nil,
}

fn lister() {
    use crate::List::{Cons, Nil};

    #[allow(unused)]
    let list = Cons(1, Box::new(Cons(2, Box::new(Cons(3, Box::new(Nil))))));
}

// ----------------------------------------
fn raw_deref() {
    let mut x = 5;
    let y = Box::new(x);

    x = 6;

    assert_eq!(6, x);
    assert_eq!(5, *y);
    println!("x = {x}, y = {y}")
}

// ----------------------------------------
struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

use std::ops::Deref;

impl<T> Deref for MyBox<T> {
    type Target = T;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

fn my_box1() {
    let mut x = 5;
    let y = MyBox::new(x);

    x = 6;

    assert_eq!(6, x);
    assert_eq!(5, *y);
    println!("x = {x}, y = {}", *y)
}

// ----------------------------------------
fn hello(name: &str) {
    println!("Hello, {name}.");
}

fn my_box2() {
    let m = MyBox::new(String::from("Rust"));
    hello(&m);
}

// ----------------------------------------
struct CustomSmartPointer {
    data: String,
}

// use core::ops::Drop;  // <-- unnecessary, included in prelude

impl Drop for CustomSmartPointer {
    fn drop(&mut self) {
        println!("Dropping CustomSmartPointer with data '{}'", self.data);
    }
}

#[allow(unused)]
fn custom_drop() {
    let c = CustomSmartPointer{
        data: String::from("my stuff"),
    };
    let d = CustomSmartPointer{
        data: String::from("other stuff"),
    };
    println!("CustomSmartPointers created.");

    //use core::mem::drop;  // <-- unnecessary, included in prelude
    drop(c);              // <-- `c` does not have to be mutable
    println!("CustomSmartPointer dropped before the end of main.");
}