use std::{
    collections::HashMap,
    error::Error,
};

pub trait CommandRegistry {
    fn register(&mut self, handler: Box<dyn CommandHandler>) -> Result<(), Box<dyn Error>>;
    fn lookup(&self, cmd_line: &str) -> Option<Box<dyn CommandHandler>>;
}

pub trait CommandHandler {
    fn usage(&self) -> String;
}

struct CommandRepo {
    handlers: HashMap<String, Box<dyn CommandHandler>>
}

impl CommandRepo {
    pub fn new() -> CommandRepo {
        CommandRepo{
            handlers: HashMap::new()
        }
    }
}

impl CommandRegistry for CommandRepo {
    fn register(&mut self, handler: Box<dyn CommandHandler>) -> Result<(), Box<dyn Error>> {
        self.handlers.insert(handler.usage(), handler);
        Ok(())
    }

    fn lookup(&self, cmd_line: &str) -> Option<Box<dyn CommandHandler>> {
        let (cmd, args) = parse_cmd_line(cmd_line);

        None
    }
}

fn parse_cmd_line(cmd_line: &str) -> (Vec<String>, Vec<String>) {
    let mut cmd = Vec::new();
    let mut args: Vec<String> = cmd_line
        .split_whitespace()
        .map(|s| s.to_string())
        .collect();
    loop {
        if args.len() > 0 {
            let a = &args[0];
            match a.chars().next() {
                Some(c) => {
                    if c.is_alphabetic() {
                        cmd.push(args.remove(0));
                    } else {
                        break
                    }
                },
                None => break,
            }
        } else {
            break;
        }
    }
    (cmd, args)
}

#[cfg(test)]
mod tests {
    use super::*;

    struct MockHandler {
        usage: String
    }

    impl CommandHandler for MockHandler {
        fn usage(&self) -> String {
            self.usage.clone()
        }
    }

    #[test]
    fn create_registry() {
        let mut cr = CommandRepo::new();
        let h = Box::new(MockHandler{usage: String::from("placeholder")});
        assert!(cr.register(h).is_ok());
        assert!(cr.lookup("cmd sub -flag").is_none());
    }
}
