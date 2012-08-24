with Ada.Integer_Text_IO, Ada.Text_IO;
use  Ada.Integer_Text_IO, Ada.Text_IO;

package body IntList is

   function IsInList( List: in T; Number: in Integer ) return Boolean is
   begin

      return Find_Element( List, Number ) /= null;

   end IsInList;

   ------------------------------------------------------------

   procedure Insert( List: in out T; Number: in Integer ) is
      New_Node : IntList_Node_Ptr;
   begin
      
      if IsInList( List, Number ) then
         return;
      end if;

      New_Node := new IntList_Node'(Datum => Number, Next => null);

      if List.Head = null then
         List.Head := New_Node;
         List.Tail := New_Node;
      else
         List.Tail.Next := New_Node;
         List.Tail      := New_Node;
      end if;
      
   end Insert;

   ------------------------------------------------------------

   procedure Remove( List: in out T; Number: in Integer ) is

      Node : IntList_Node_Ptr := Find_Element( List, Number );
      Prev : IntList_Node_Ptr := List.Head;

   begin
      
      if Node = null then
         return;
      end if;

      if Node = List.Head then
         List.Head := Node.Next;
      else 

         while Prev.Next /= Node loop
            Prev := Prev.Next;
         end loop;
         
         Prev.Next := Node.Next;
         
         if Node = List.Tail then
            List.Tail := Prev;
         end if;

      end if;

      -- Free( Node );

   end Remove;

   ------------------------------------------------------------

   procedure Print( List: in T ) is
      Next_Node : IntList_Node_Ptr := List.Head;
   begin

      while Next_Node /= null loop
         Put(Next_Node.Datum);
         New_Line;
         Next_Node := Next_Node.Next;
      end loop;

   end Print;

   ------------------------------------------------------------

   function Find_Element( List: in T; Number: in Integer ) 
                         return IntList_Node_Ptr is

      Next_Node : IntList_Node_Ptr := List.Head;

   begin

      while Next_Node /= null loop
         if Next_Node.Datum = Number then
            return Next_Node;
         end if;
         Next_Node := Next_Node.Next;
      end loop;

      return null;

   end Find_Element;

end IntList;
