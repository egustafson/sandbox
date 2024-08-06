fn main() {
    // do_panic();  // <-- panics!

    err_result_expect();  // <-- panics!
    err_result_unwrap();  // <-- panics!
    err_result();  // <-- panics!
}

#[allow(dead_code)]
fn do_panic() {
    let v = vec![1, 2, 3];
    v[99]; // <-- panic!
}

use std::fs::File;

#[allow(dead_code)]
fn err_result() {
    let file_result = File::open("bogus-filename.txt");

    #[allow(unused)]
    let file = match file_result {
        Ok(file) => file,
        Err(error) => panic!("Problem opening the file: {error:?}")
    };
}

#[allow(dead_code)]
fn err_result_unwrap() {
    #[allow(unused)]
    let file = File::open("bogus-filename.txt").unwrap();
}

#[allow(dead_code)]
fn err_result_expect() {
    #[allow(unused)]
    let file = File::open("bogus-filename.txt")
        .expect("file missing");
}