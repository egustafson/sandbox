with Ada.Unchecked_Deallocation;

package body Generic_Queue is

   procedure Free is new Ada.Unchecked_Deallocation( Queue_Node,
                                                     Queue_Node_Ptr );

   ------------------------------------------------------------

   procedure Enqueue( Queue: in out T; Element: in Element_Type ) is

      New_Node : Queue_Node_Ptr := new Queue_Node'( Data => Element,
                                                    Next => null );
   begin

      if Queue.Head = null then
         Queue.Head := New_Node;
         Queue.Tail := New_Node;
      else
         Queue.Tail.Next := New_Node;
         Queue.Tail      := New_Node;
      end if;

   end Enqueue;

   ------------------------------------------------------------

   procedure Dequeue( Queue: in out T; Element: out Element_Type ) is
      Old_Head : Queue_Node_Ptr := Queue.Head;
   begin

      if Queue.Head = null then
         raise Underflow;
      end if;

      Element := Queue.Head.Data;

      if Queue.Tail = Queue.Head then
         Queue.Tail := null;
      end if;

      Queue.Head := Queue.Head.Next;
      Free( Old_Head );

   end Dequeue;

   ------------------------------------------------------------

   procedure Peek( Queue: in out T; Element: out Element_Type ) is
   begin

      if Queue.Head = null then
         raise Underflow;
      end if;

      Element := Queue.Head.Data;

   end Peek;

   ------------------------------------------------------------

   function Empty( Queue: in T ) return Boolean is
   begin

      return (Queue.Head = null);

   end Empty;

   ------------------------------------------------------------

   procedure Reset( Queue: in out T ) is
      Old_Head : Queue_Node_Ptr;
   begin

      while Queue.Head /= null loop
         Old_Head   := Queue.Head;
         Queue.Head := Queue.Head.Next;
         Free( Old_Head );
      end loop;

      Queue.Tail := null;

   end Reset;

end Generic_Queue;
