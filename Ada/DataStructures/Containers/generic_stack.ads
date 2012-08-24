generic
   type Element_Type is private;

package Generic_Stack is

   type T is limited private;

   procedure Push(  Stack: in out T; Element: in  Element_Type );
   procedure Pop(   Stack: in out T; Element: out Element_Type );
   procedure Peek(  Stack: in out T; Element: out Element_Type );
   function  Empty( Stack: in T ) return Boolean;
   procedure Reset( Stack: in out T );

   Underflow : exception;

private

   type Stack_Node;
   type Stack_Node_Ptr is access Stack_Node;

   type Stack_Node is
      record
         Data : Element_Type;
         Next : Stack_Node_Ptr;
      end record;

   type T is
      record
         Head : Stack_Node_Ptr;
      end record;

end Generic_Stack;
