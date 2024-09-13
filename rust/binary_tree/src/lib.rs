use std::iter::Iterator;

struct Tree<T: PartialOrd> (Option<Box<Node<T>>>);

pub struct Node<T: PartialOrd> {
    data: T,
    left: Option<Box<Node<T>>>,
    right: Option<Box<Node<T>>>,
}

impl<T: PartialOrd> Tree<T> {
    pub fn new() -> Tree<T> {
        Tree(None)
    }

    pub fn insert(&mut self, d: T) {
        let mut nn = Box::new(Node{
            data: d,
            left: None,
            right: None,
        });
        let mut curr = match &self.0 {
            None => {
                self.0 = Some(nn);
                return;
            }
            Some(t) => t
        };
        let mut stack: Vec<&Box<Node<T>>> = Vec::new();
        if curr.data > nn.data {
            match &curr.left {
                None => {
                    curr.left = Some(nn);
                    return;
                }
                Some(n) => {
                    stack.push(curr);
                    curr = &n;
                }
            }
        }        

    }

    pub fn delete(&mut self, d: T) {

    }

    pub fn has(&self, d: T) -> bool {
        false
    }

    pub fn iter(&self) -> TreeIterator<T> {
        TreeIterator(None)
    }
}

struct TreeIterator<'a, T: PartialOrd> (Option<&'a Node<T>>);

impl<'a, T: PartialOrd> Iterator for TreeIterator<'a, T> {
    type Item = &'a T;

    fn next(&mut self) -> Option<Self::Item> {
        //
        // TODO
        //
        None
    }
}

// ---- boiler plate -- remove below
pub fn add(left: u64, right: u64) -> u64 {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn new_tree() {
        let _t: Tree<u32> = Tree::new();
    }

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
