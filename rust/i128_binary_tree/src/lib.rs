use std::{borrow::Borrow, cell::RefCell, rc::Rc};

#[derive(Debug, Clone)]
pub struct TreeNode {
    val: i128,
    left: Option<TreeNodeRef>,
    right: Option<TreeNodeRef>,
}

type TreeNodeRef = Rc<RefCell<TreeNode>>;

struct Tree (Option<TreeNodeRef>);

impl Tree {
    pub fn new() -> Tree {
        Tree(None)
    }

    pub fn insert(&mut self, d: i128) {
        let mut nn = Rc::new(RefCell::new(TreeNode{
            val: d,
            left: None,
            right: None,
        }));
        let curr = match &self.0 {
            None => {
                self.0 = Some(nn);
                return;
            }
            Some(node) => node
        };
        let mut stack: Vec<TreeNodeRef> = Vec::new();
        if curr.borrow().val > &nn.borrow().val {
            if let Some(left) = &curr.borrow().left {
                curr = left;
            }


            match curr.borrow().left {
                None => {
                    curr.borrow_mut().left = Some(nn);
                    return;
                }
                Some(n) => {
                    stack.push(n);
                    curr = n;
                }
            }
        }
    }

    pub fn delete(&mut self, d: i128) {

    }

    pub fn has(&self, d: i128) -> bool {
        false
    }
}






#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn new_tree() {
        let t = Tree::new();
        assert_eq!(t.has(0), false)        
    }

    #[test]
    fn insert_test() {
        let mut t = Tree::new();

        t.insert(1);
        assert_eq!(t.has(1), true)
    }

}
