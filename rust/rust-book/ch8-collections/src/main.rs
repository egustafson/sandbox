fn main() {
    //vec_panic();  // <- panic's
    vec_iter();
    vec_hetero();

    hash_entry();
    hash_update();
}

#[allow(unused)]
fn vec_panic() {
    let v = vec![1, 2, 3, 4, 5];

    let does_not_exist = &v[100];  // <-- panic's
    let does_not_exist = v.get(100);
}

fn vec_iter() {
    let mut v = vec![100, 32, 57];
    for i in &v {
        println!("{i}");
    }

    // change element values
    for i in &mut v {
        *i += 50;
    }

    println!("---");
    for i in &v {
        println!("{i}");
    }
}

#[allow(unused, dead_code)]
fn vec_hetero() {
    enum SSCell {
        Int(i32),
        Float(f64),
        Text(String),
    }

    let row = vec![
        SSCell::Int(3),
        SSCell::Float(23.71),
        SSCell::Text(String::from("some text")),
    ];
}

fn hash_entry() {
    use std::collections::HashMap;

    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);

    scores.entry(String::from("Yellow")).or_insert(50);
    scores.entry(String::from("Blue")).or_insert(50);

    println!("{scores:?}");
}

fn hash_update() {
    use std::collections::HashMap;

    let text = "hello world wonderful world";

    let mut map = HashMap::new();

    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;
    }

    println!("{map:?}");
}