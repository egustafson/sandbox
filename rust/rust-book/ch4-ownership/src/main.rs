fn main() {
    ex1();

    let s = String::from("hola");
    let s2 = ex2(s);
    println!("s2 = {s2}");

    ex3_byref(&s2);
    println!("outer func: {s2}");
}

fn ex1() {
    let mut s = String::from("hello");
    s.push_str(", world!");
    println!("{s}");
}

fn ex2(str : String) -> String {
    str
}

fn ex3_byref(str: &String) {
    println!("nested func: {str}")
}