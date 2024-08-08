use std::thread;
use std::time::Duration;

fn main() {

    #[allow(unused)]
    let v = vec![1, 2, 3];

    let handle = thread::spawn(move || {
        println!("vector v: {v:?}");
        for i in 1..10 {
            println!("hi number {i} from the spawned thread");
            thread::sleep(Duration::from_millis(1));
        }
    });

    for i in 1..5 {
        println!("hi number {i} from the main thread");
        thread::sleep(Duration::from_millis(1));
    }

    handle.join().unwrap()
}
