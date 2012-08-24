generic
   type Element_Type is private;

package Generic_Queue is

   type T is limited private;

   procedure Enqueue( Queue: in out T; Element: in  Element_Type );
   procedure Dequeue( Queue: in out T; Element: out Element_Type );
   procedure Peek   ( Queue: in out T; Element: out Element_Type );
   function  Empty  ( Queue: in T ) return Boolean;
   procedure Reset  ( Queue: in out T );

   Underflow : exception;

private

   type Queue_Node;
   type Queue_Node_Ptr is access Queue_Node;

   type Queue_Node is
      record
         Data : Element_Type;
         Next : Queue_Node_Ptr;
      end record;

   type T is
      record
         Head : Queue_Node_Ptr := null;
         Tail : Queue_Node_Ptr := null;
      end record;

end Generic_Queue;
