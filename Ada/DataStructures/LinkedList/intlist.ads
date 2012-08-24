package IntList is

   type T is limited private;

   function  IsInList( List: in T;     Number: in Integer ) return Boolean;
   procedure Insert(   List: in out T; Number: in Integer );
   procedure Remove(   List: in out T; Number: in Integer );
   procedure Print(    List: in T );

private

   type IntList_Node;
   type IntList_Node_Ptr is access IntList_Node;

   type IntList_Node is
      record
         Datum : Integer;
         Next  : IntList_Node_Ptr;
      end record;

   type T is
      record
         Head : IntList_Node_Ptr := null;
         Tail : IntList_Node_Ptr := null;
      end record;

   function Find_Element( List: in T; Number: in Integer ) 
                         return IntList_Node_Ptr;

end IntList;
