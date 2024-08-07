use std::thread;

fn main() {
    thread_closure();
    tracing_closure();
    counting_closure();
}

fn thread_closure() {
    let list = vec![1, 2, 3];
    println!("before defining closure: {list:?}");

    thread::spawn(move || println!("from thread: {list:?}"))
        .join()
        .unwrap();

    // println!("after calling thread: {list:?}"); 
}

#[allow(unused)]
#[derive(Debug)]
struct Rect {
    width: u32,
    height: u32,
}

#[allow(unused)]
fn tracing_closure() {
    let mut list = [
        Rect{ width: 10, height: 1 },
        Rect{ width: 3, height: 5 },
        Rect{ width: 7, height: 12 },
    ] ;

    let mut sort_operations: Vec<String> = vec![];
    let value = String::from("closure called");

    list.sort_by_key(|r| {
        sort_operations.push(value.clone());
        r.width
    });
    println!("{list:#?}");
    println!("{sort_operations:#?}");
}

fn counting_closure() {
    let mut list = [
        Rect{ width: 10, height: 1 },
        Rect{ width: 3, height: 5 },
        Rect{ width: 7, height: 12 },
    ] ;

    let mut num_sort_operations = 0;

    list.sort_by_key(|r| {
        num_sort_operations += 1;
        r.width
    });
    println!("{list:#?}, sorted in {num_sort_operations} operations");
}