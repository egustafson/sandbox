pub fn add(left: u64, right: u64) -> u64 {
    left + right
}

#[allow(dead_code)]
fn priv_add2x(left: u64, right: u64) -> u64 {
    2 * (left + right)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
        println!("output from inside a successful test");
    }

    #[test]
    fn test_priv_add2x() {
        let result = priv_add2x(4, 4);
        assert_eq!(result, 16)
    }

    #[test]
    #[ignore]
    fn test_result_failure() -> Result<(), String> {
        println!("output from inside the failing test");
        Err(String::from("always fails"))
    }
}
