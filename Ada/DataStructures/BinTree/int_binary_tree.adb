with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Unchecked_Deallocation;
use  Ada.Text_IO, Ada.Integer_Text_IO;

package body Int_Binary_Tree is

   type Compare_Result_Type is ( Less_Than, Equal, Greater_Than );

   type Traced_Node_Type is
      record
         Node   : Tree_Node_Ptr;
         Parent : Tree_Node_Ptr;
      end record;

   ----------------------------------------------------------------------

   procedure Free is new Ada.Unchecked_Deallocation( Tree_Node,
                                                     Tree_Node_Ptr );

   ----------------------------------------------------------------------

   function Compare( Left, Right : in Integer ) return
     Compare_Result_Type is
   begin
      if Left < Right then
         return Less_Than;
      elsif Left > Right then
         return Greater_Than;
      else
         return Equal;
      end if;

   end Compare;

   ----------------------------------------------------------------------

   function Find_Node( Tree: in T; Number: in Integer ) return
     Traced_Node_Type is

      T_Node : Traced_Node_Type := Traced_Node_Type'( Node   => Tree.Root,
                                                      Parent => null );
      Node   : Tree_Node_Ptr renames T_Node.Node;
      Parent : Tree_Node_Ptr renames T_Node.Parent;
   begin

      while Node /= null loop

         case Compare ( Number, Node.Data ) is
            when Equal =>
               return T_Node;
            when Less_Than =>
               Parent := Node;
               Node   := Node.Left;
            when Greater_Than =>
               Parent := Node;
               Node   := Node.Right;
         end case;

      end loop;

      return T_Node;

   end Find_Node;

   ----------------------------------------------------------------------

   function Is_In_Tree( Tree: in T; Number: in Integer ) return Boolean is

      T_Node : Traced_Node_Type := Find_Node( Tree, Number );

   begin

      if T_Node.Node /= null then
         return True;
      else
         return False;
      end if;

   end Is_In_Tree;

   ----------------------------------------------------------------------

   procedure Insert( Tree: in out T; Number: in Integer ) is

      T_Node   : Traced_Node_Type := Find_Node( Tree, Number );
      New_Node : Tree_Node_Ptr    := new Tree_Node'( Data  => Number,
                                                     Left  => null,
                                                     Right => null );
   begin

      if T_Node.Node /= null then
         return;
      elsif T_Node.Parent = null then
         Tree.Root := New_Node;
      else
         case Compare( Number, T_Node.Parent.Data ) is
            when Less_Than =>
               T_Node.Parent.Left  := New_Node;
            when Greater_Than =>
               T_Node.Parent.Right := New_Node;
            when Equal =>
               null;
         end case;
      end if;

   end Insert;

   ----------------------------------------------------------------------

   procedure Remove( Tree: in out T; Number: in Integer ) is

      Remove_Node       : Traced_Node_Type := Find_Node( Tree, Number );
      Pivot_Node        : Tree_Node_Ptr;
      Pivot_Node_Parent : Tree_Node_Ptr;

      procedure Graft( Tree: in out T;
                       Old_Node,
                       New_Node,
                       Parent: in out Tree_Node_Ptr ) is
      begin
         if Parent /= null then
            if Parent.Left = Old_Node then
               Parent.Left := New_Node;
            else
               Parent.Right := New_Node;
            end if;
         else
            Tree.Root := New_Node;
         end if;
      end Graft;

   begin

      if Remove_Node.Node = null then
         return;
      end if;

      if Remove_Node.Node.Left = null then
         Graft( Tree, Remove_Node.Node, Remove_Node.Node.Right, Remove_Node.Parent );
         Free( Remove_Node.Node );
      elsif Remove_Node.Node.Right = null then
         Graft( Tree, Remove_Node.Node, Remove_Node.Node.Left, Remove_Node.Parent );
         Free( Remove_Node.Node );
      else
         Pivot_Node_Parent := Remove_Node.Node;
         Pivot_Node        := Remove_Node.Node.Left;
         while Pivot_Node.Right /= null loop
            Pivot_Node_Parent := Pivot_Node;
            Pivot_Node        := Pivot_Node.Right;
         end loop;
         if Pivot_Node_Parent = Remove_Node.Node then
            Pivot_Node.Right := Remove_Node.Node.Right;
         else
            Pivot_Node_Parent.Right := Pivot_Node.Left;
            Pivot_Node.Left  := Remove_Node.Node.Left;
            Pivot_Node.Right := Remove_Node.Node.Right;
         end if;
         Graft( Tree, Remove_Node.Node, Pivot_Node, Remove_Node.Parent );
      end if;

   end Remove;

   ----------------------------------------------------------------------

   procedure Print( Tree: in T ) is

      procedure Print_Node( Node: in Tree_Node_Ptr ) is
      begin
         if Node.Left /= null then
            Print_Node( Node.Left );
         end if;
         Put(Node.Data);
         New_Line;
         if Node.Right /= null then
            Print_Node( Node.Right );
         end if;
      end;

   begin
      Print_Node( Tree.Root );
   end Print;

   ----------------------------------------------------------------------

   procedure Debug_Print( Tree: in T ) is

      procedure Print_Node( Node: in Tree_Node_Ptr ) is

         procedure Print_Branch( Node: in Tree_Node_Ptr ) is
         begin
            if Node = null then
               Put("       null");
            else
               Put(Node.Data);
            end if;
         end Print_Branch;

      begin
         if Node.Left /= null then
            Print_Node( Node.Left );
         end if;
         Put(Node.Data);
         Put("   Left: "); Print_Branch(Node.Left);
         Put("   Right:"); Print_Branch(Node.Right);
         New_Line;
         if Node.Right /= null then
            Print_Node( Node.Right );
         end if;
      end Print_Node;

   begin
      Print_Node( Tree.Root );
   end Debug_Print;

end Int_Binary_Tree;



