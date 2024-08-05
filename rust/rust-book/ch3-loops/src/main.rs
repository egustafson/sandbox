fn main() {
    loop1();
    let y = loop2();
    println!("y = {y}");
    loop3();
    loop4();
    loop5();
}

fn loop1() {
    loop {
        println!{"again!"};
        break
    }
}

fn loop2() -> u32 {
    let mut x = 1;
    loop {
        x += 1;
        if x > 3 {
            return x // Exit fn, return value
        }
    }
}

fn loop3() {
    let mut counter = 0;

    let result = loop {
        counter += 1;
        if counter == 10 {
            break counter * 2;  // Exit loop statement, assign to `result`
        }
    };
    println!("result = {result}")
}

fn loop4() {
    let mut counter : i32 = 0;
    while counter < 10 {
        counter += 1;
    };
    println!("result = {counter}")
}

fn loop5() {
    let a = [10, 20, 30, 40, 50];
    for el in a {
        println!("a[?] = {el}")
    }
}