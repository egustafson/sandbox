with Ada.Text_IO;
use  Ada.Text_IO;

generic
   type Element_Type is private;

   with function  Less_Than( Left, Right : in Element_Type ) return Boolean;
   with function  Greater_Than( Left, Right : in Element_Type ) return Boolean;
   with procedure Put( Item: in Element_Type );

package Generic_Binary_Tree is

   type T is limited private;

   function Is_In_Tree( Tree: in T; Element: in Element_Type ) return Boolean;

   procedure Insert( Tree: in out T; Element: in Element_Type );
   procedure Remove( Tree: in out T; Element: in Element_Type );

   procedure Print( Tree: in T );
   procedure Debug_Print( Tree: in T );

private

   type Tree_Node;
   type Tree_Node_Ptr is access Tree_Node;

   type Tree_Node is
      record
         Data  : Element_Type;
         Left  : Tree_Node_Ptr;
         Right : Tree_Node_Ptr;
      end record;

   type T is
      record
         Root : Tree_Node_Ptr;
      end record;

end Generic_Binary_Tree;
