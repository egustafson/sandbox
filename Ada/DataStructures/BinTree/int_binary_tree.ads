package Int_Binary_Tree is

   type T is limited private;

   function Is_In_Tree( Tree: in T; Number: in Integer ) return Boolean;

   procedure Insert( Tree: in out T; Number: in Integer );
   procedure Remove( Tree: in out T; Number: in Integer );

   procedure Print( Tree: in T );
   procedure Debug_Print( Tree: in T );

private

   type Tree_Node;
   type Tree_Node_Ptr is access Tree_Node;

   type Tree_Node is
      record
         Data  : Integer;
         Left  : Tree_Node_Ptr;
         Right : Tree_Node_Ptr;
      end record;

   type T is
      record
         Root : Tree_Node_Ptr;
      end record;

end Int_Binary_Tree;
