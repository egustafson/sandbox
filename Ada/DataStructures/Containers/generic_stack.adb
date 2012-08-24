with Ada.Unchecked_Deallocation;

package body Generic_Stack is

   procedure Free is new Ada.Unchecked_Deallocation( Stack_Node,
                                                     Stack_Node_Ptr );

   ------------------------------------------------------------

   procedure Push(  Stack: in out T; Element: in Element_Type ) is
      New_Node : Stack_Node_Ptr := new Stack_Node'( Data => Element,
                                                    Next => Stack.Head );
   begin
      Stack.Head := New_Node;
   end Push;

   ------------------------------------------------------------

   procedure Pop( Stack: in out T; Element: out Element_Type ) is
      Old_Head  : Stack_Node_Ptr := Stack.Head;
   begin

      if Stack.Head = null then
         raise Underflow;
      end if;

      Stack.Head := Old_Head.Next;

      Element := Old_Head.Data;
      Free( Old_Head );

   end Pop;

   ------------------------------------------------------------

   procedure Peek( Stack: in out T; Element: out Element_Type ) is
   begin

      if Stack.Head = null then
         raise Underflow;
      end if;

      Element := Stack.Head.Data;

   end Peek;

   ------------------------------------------------------------

   function Empty( Stack: in T ) return Boolean is
   begin
      return Stack.Head = null;
   end Empty;

   ------------------------------------------------------------

   procedure Reset( Stack: in out T ) is
      Old_Head : Stack_Node_Ptr := Stack.Head;
   begin

      while Old_Head /= null loop
         Stack.Head := Old_Head.Next;
         Free( Old_Head );
         Old_Head := Stack.Head;
      end loop;

   end Reset;

end Generic_Stack;
