with Generic_Stack;
with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Unchecked_Deallocation;
use  Ada.Text_IO, Ada.Integer_Text_IO;

package body Generic_AVL_Tree is

   type Compare_Result_Type is ( Less_Than, Equal, Greater_Than );

   ----------------------------------------------------------------------

   procedure Free is new Ada.Unchecked_Deallocation( Tree_Node,
                                                     Tree_Node_Ptr );

   ----------------------------------------------------------------------

   procedure Put(Ballance: in Pivot_Type) is
   begin
      case Ballance is
         when Pivot_Left   =>  Put("Pivot Left");
         when Left         =>  Put(" Left");
         when Equal        =>  Put("Equal");
         when Right        =>  Put("Right");
         when Pivot_Right  =>  Put("Pivot Right");
      end case;
   end Put;

   ----------------------------------------------------------------------

   function Compare( Left, Right : in Element_Type ) return
     Compare_Result_Type is
   begin
      if Less_Than( Left, Right ) then
         return Less_Than;
      elsif Greater_Than( Left, Right ) then
         return Greater_Than;
      else
         return Equal;
      end if;

   end Compare;

   ----------------------------------------------------------------------

   function Find_Node( Tree: in T; Element: in Element_Type ) return
     Tree_Node_Ptr is

      Node   : Tree_Node_Ptr;

   begin

      while Node /= null loop

         case Compare ( Element, Node.Data ) is
            when Equal =>
               return Node;
            when Less_Than =>
               Node   := Node.Left;
            when Greater_Than =>
               Node   := Node.Right;
         end case;

      end loop;

      return Node;

   end Find_Node;

   ----------------------------------------------------------------------

   function Is_In_Tree( Tree: in T; Element: in Element_Type ) return Boolean is

      Node : Tree_Node_Ptr := Find_Node( Tree, Element );

   begin

      if Node /= null then
         return True;
      else
         return False;
      end if;

   end Is_In_Tree;

   ----------------------------------------------------------------------

   procedure Insert( Tree: in out T; Element: in Element_Type ) is

      New_Node : Tree_Node_Ptr    := new Tree_Node'( Data     => Element,
                                                     Left     => null,
                                                     Right    => null,
                                                     Parent   => null,
                                                     Ballance => Equal,
                                                     Branch   => Root );
      -------------------------------------------------------------------
      procedure Inject_Node( Tree:     in out  T;
                             New_Node: in out  Tree_Node_Ptr ) is
         Child    : Tree_Node_Ptr    := Tree.Root;
         Node     : Tree_Node_Ptr    := Tree.Root;
      begin
         if Node = null then
            Tree.Root := New_Node;
         else
            while Child /= null loop
               Node := Child;
               case Compare ( Element, Node.Data ) is
                  when Equal =>
                     return;
                  when Less_Than =>
                     Child           := Child.Left;
                     New_Node.Branch := Left;
                  when Greater_Than =>
                     Child           := Child.Right;
                     New_Node.Branch := Right;
               end case;
            end loop;
            New_Node.Parent := Node;
            case Compare( Element, Node.Data ) is
               when Less_Than =>
                  Node.Left  := New_Node;
               when Greater_Than =>
                  Node.Right := New_Node;
               when Equal =>
                  null;
            end case;
         end if;
      end Inject_Node;
      -------------------------------------------------------------------
      procedure Reballance( Tree: in out T; Leaf: in Tree_Node_Ptr ) is
         Node : Tree_Node_Ptr := Leaf;
         New_Ballance : Pivot_Type;

         Ballance_State_Table : constant
           array(Branch_Type range Left..Right, Ballance_Type)
           of Pivot_Type := (( Pivot_Left,  Left,  Equal       ),
                             ( Equal,       Right, Pivot_Right ));
         ----------------------------------------------------------------
         procedure Swap_Child( Tree   : in out T;
                               Parent : in Tree_Node_Ptr;
                               Branch : in Branch_Type;
                               Child  : in Tree_Node_Ptr ) is
         begin
            case Branch is
               when Left  => Parent.Left  := Child;
               when Right => Parent.Right := Child;
               when Root  => Tree.Root    := Child;
            end case;
            if Child /= null then
               Child.Parent := Parent;
               Child.Branch := Branch;
            end if;
         end Swap_Child;
         ----------------------------------------------------------------
         procedure Rotate_Right(Tree  : in out T;
                                Pivot : in Tree_Node_Ptr ) is
            Child  : Tree_Node_Ptr := Pivot.Left;
         begin
--             Put_Line("Rotate Right - Begin.");
            Swap_Child( Tree, Pivot.Parent, Pivot.Branch, Child );
            Swap_Child( Tree, Pivot,        Left,         Child.Right );
            Swap_Child( Tree, Child,        Right,        Pivot );
            Pivot.Ballance := Equal;
            Child.Ballance := Equal;
--             Put_Line("Rotate Right - End.");
         end Rotate_Right;
         ----------------------------------------------------------------
         procedure Double_Right(Tree  : in out T;
                                Pivot : in Tree_Node_Ptr ) is
            Child1      : Tree_Node_Ptr := Pivot.Left;
            Child2      : Tree_Node_Ptr := Pivot.Left.Right;
         begin
--             Put_Line("Double Right - Begin.");
            Swap_Child( Tree, Pivot.Parent, Pivot.Branch, Child2 );
            Swap_Child( Tree, Pivot,        Left,         Child2.Right );
            Swap_Child( Tree, Child1,       Right,        Child2.Left );
            Swap_Child( Tree, Child2,       Left,         Child1);
            Swap_Child( Tree, Child2,       Right,        Pivot);
            Pivot.Ballance  := Equal;
            Child1.Ballance := Equal;
            case Child2.Ballance is
               when Left   =>
                  Pivot.Ballance  := Right;
                  if Child1.Left = null then
                     Child1.Ballance := Right;
                  end if;
               when Right  =>
                  Child1.Ballance := Left;
               when Equal  =>
                  null;
            end case;
            Child2.Ballance := Equal;
--             Put_Line("Double Right - End.");
         end Double_Right;
         ----------------------------------------------------------------
         procedure Rotate_Left(Tree  : in out T;
                               Pivot : in Tree_Node_Ptr ) is
            Child  : Tree_Node_Ptr := Pivot.Right;
         begin
--             Put("Rotate Left - Begin.");
            Swap_Child( Tree, Pivot.Parent, Pivot.Branch, Child );
            Swap_Child( Tree, Pivot,        Right,        Child.Left );
            Swap_Child( Tree, Child,        Left,         Pivot );
            Pivot.Ballance := Equal;
            Child.Ballance := Equal;
--             Put_Line("Rotate Left - End.");
         end Rotate_Left;
         ----------------------------------------------------------------
         procedure Double_Left(Tree  : in out T;
                               Pivot : in Tree_Node_Ptr ) is
            Child1  : Tree_Node_Ptr := Pivot.Right;
            Child2  : Tree_Node_Ptr := Pivot.Right.Left;
         begin
--             Put_Line("Double Left - Begin.");
            Swap_Child( Tree, Pivot.Parent, Pivot.Branch, Child2 );
            Swap_Child( Tree, Pivot,        Right,        Child2.Left );
            Swap_Child( Tree, Child1,       Left,         Child2.Right );
            Swap_Child( Tree, Child2,       Left,         Pivot );
            Swap_Child( Tree, Child2,       Right,        Child1 );
            Pivot.Ballance  := Equal;
            Child1.Ballance := Equal;
            case Child2.Ballance is
               when Left   =>
                  Child1.Ballance := Right;
               when Right  =>
                  Pivot.Ballance  := Left;
                  if Child1.Right = null then
                     Child1.Ballance := Left;
                  end if;
               when Equal  =>
                  null;
            end case;
            Child2.Ballance := Equal;
--             Put_Line("Double Left - End.");
         end Double_Left;
         ----------------------------------------------------------------
      begin
         while Node.Parent /= null loop -- top and _bottom_ tested

            New_Ballance := Ballance_State_Table( Node.Branch,
                                                  Node.Parent.Ballance );

--             Put("node "); Put(Node.Parent.Data); Put(" is ballanced ");
--             Put(Node.Parent.Ballance);Put(" -> "); Put(New_Ballance);
--             Put(" from child "); Put(Node.Data); New_Line;

            if New_Ballance = Pivot_Left then
               if Node.Ballance = Left then
                  Node := Node.Parent;
                  Rotate_Right( Tree, Node );
--                   if Node.Ballance /= Equal then
--                      Put_Line("PROBLEM:  Pivot.Ballance is not Equal");
--                   end if;
               else
                  Node := Node.Parent;
                  Double_Right( Tree, Node );
               end if;
               exit;
            elsif New_Ballance = Pivot_Right then
               if Node.Ballance = Right then
                  Node := Node.Parent;
                  Rotate_Left( Tree, Node );
--                   if Node.Ballance /= Equal then
--                      Put_Line("PROBLEM:  Pivot.Ballance is not Equal");
--                   end if;
               else
                  Node := Node.Parent;
                  Double_Left( Tree, Node );
               end if;
               exit;
            else
               Node.Parent.Ballance := New_Ballance;
            end if;

            Node := Node.Parent;
            exit when New_Ballance = Equal;
         end loop;
      end Reballance;
      -------------------------------------------------------------------
   begin -- Insert

--       Put_Line("----------------------------------------");
--       Put("Inserting :  "); Put(Element); New_Line;

      Inject_Node( Tree, New_Node );
--       Put("Injected successfully; parent : ");

--       if New_Node.Parent /= null then
--          Put(New_Node.Parent.Data); New_Line;
--       else
--          Put("Root"); New_Line;
--          if Tree.Root /= New_Node then
--             Put_Line("Error, this node is not the root.");
--             raise Program_Error;
--          end if;
--       end if;
--
--       Check_Tree( Tree );

      Reballance( Tree, New_Node );
--       Put_Line("Reballanced Successfully.");

--       Check_Tree( Tree );

   end Insert;

   ----------------------------------------------------------------------

   procedure Remove( Tree: in out T; Element: in Element_Type ) is

      Remove_Node       : Tree_Node_Ptr := Find_Node( Tree, Element );
      Pivot_Node        : Tree_Node_Ptr;
      Pivot_Node_Parent : Tree_Node_Ptr;

      procedure Graft( Tree: in out T;
                       Old_Node, New_Node: in out Tree_Node_Ptr ) is

         Parent : Tree_Node_Ptr := Old_Node.Parent;

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

      if Remove_Node = null then
         return;
      end if;

      if Remove_Node.Left = null then
         Graft( Tree, Remove_Node, Remove_Node.Right );
         Free( Remove_Node );
      elsif Remove_Node.Right = null then
         Graft( Tree, Remove_Node, Remove_Node.Left );
         Free( Remove_Node );
      else
         Pivot_Node        := Remove_Node.Left;
         while Pivot_Node.Right /= null loop
            Pivot_Node        := Pivot_Node.Right;
         end loop;
         if Pivot_Node.Parent = Remove_Node then
            Pivot_Node.Right := Remove_Node.Right;
         else
            Pivot_Node.Parent.Right := Pivot_Node.Left;
            Pivot_Node.Left  := Remove_Node.Left;
            Pivot_Node.Right := Remove_Node.Right;
         end if;
         Graft( Tree, Remove_Node, Pivot_Node );
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

      procedure Print_Node( Node: in Tree_Node_Ptr;
                            Node_ID: in out Integer ) is

         procedure Print_Edge( Source, Target : in Integer;
                               Branch         : in Branch_Type ) is
         begin
            Put("edge [ source ");
            Put( Source );
            Put(" target ");
            Put( Target );
            if Branch = Left then
               Put_Line(" label ""L"" ]");
            else
               Put_Line(" label ""R"" ]");
            end if;
         end Print_Edge;

         My_ID : Integer := Node_ID;

      begin
         if Node.Left /= null then
            Node_ID := Node_ID + 1;
            Print_Edge( My_ID, Node_ID, Left );
            Print_Node( Node.Left, Node_ID );
         end if;
         Put("node [ id ");
         Put( My_ID );
         Put(" label """);
         Put(Node.Data);
         Put_Line(""" ]");
         if Node.Right /= null then
            Node_ID := Node_ID + 1;
            Print_Edge( My_ID, Node_ID, Right );
            Print_Node( Node.Right, Node_ID );
         end if;
      end Print_Node;

      Node_ID : Integer := 0;

   begin
      Put_Line("graph [ directed 1");
      Print_Node( Tree.Root, Node_ID );
      Put_Line("]");
   end Debug_Print;

   ----------------------------------------------------------------------

   procedure Check_Tree( Tree: in T ) is

      package Tree_Node_Stack is new Generic_Stack(Tree_Node_Ptr);
      use Tree_Node_Stack;

      Node_Count : Integer := 0;

      Node_Stack : Tree_Node_Stack.T;
      Node       : Tree_Node_Ptr;

      procedure Put( Branch : Branch_Type ) is
      begin
         case Branch is
            when Left  => Put("left");
            when Right => Put("right");
            when Root  => Put("root");
         end case;
      end Put;

   begin

      if Tree.Root /= null then

         if Tree.Root.Parent /= null then
            Put_Line("Tree.Root.Parent /= null.");
            raise AVL_Tree_Error;
         end if;

         if Tree.Root.Branch /= Root then
            Put_Line("Tree.Root.Branch /= root.");
            raise AVL_Tree_Error;
         end if;

         Push( Node_Stack, Tree.Root );
         while not Empty( Node_Stack ) loop

            Pop( Node_Stack, Node );
            Node_Count := Node_Count + 1;

            if Node.Left /= null then
               Push( Node_Stack, Node.Left );
               if Node.Left.Parent /= Node or Node.Left.Branch /= Left then
                  Put("Node "); Put(Node.Data);
                  Put(" has left child "); Put(Node.Left.Data);
                  Put(" which thinks parent is "); Put(Node.Left.Parent.Data);
                  Put(" and thinks it is a "); Put(Node.Left.Branch);
                  Put(" branch."); New_Line;
                  raise AVL_Tree_Error;
               end if;
            end if;

            if Node.Right /= null then
               Push( Node_Stack, Node.Right );
               if Node.Right.Parent /= Node or Node.Right.Branch /= Right then
                  Put("Node "); Put(Node.Data);
                  Put(" has right child "); Put(Node.Right.Data);
                  Put(" which thinks parent is "); Put(Node.Right.Parent.Data);
                  Put(" and thinks it is a "); Put(Node.Right.Branch);
                  Put(" branch."); New_Line;
                  raise AVL_Tree_Error;
               end if;
            end if;

         end loop;
      end if;
      Put_Line(" Tree Check - " & Integer'Image(Node_Count) & " nodes ok.");
   end Check_Tree;

end Generic_AVL_Tree;
