-- generic_list.adb          -*- Ada -*-
--
-- This package defines a generic list and list iterator.
--
-- Author:  Eric Gustafson
-- Date:    25 August 1998
--

-- ------------------------------------------------------------
--
-- $Revision$
--

-- $Log$

-- ------------------------------------------------------------

package body Generic_List is

   -- ----- List_Type Methods ---------------------------------

   procedure List_Add(          List     : in out List_Type;
                                Element  : in     Element_Type ) is

   begin

      if List.Num_Elements = List.List'Last then
         declare
            New_List : Element_Array_Access
              :=  new Element_Array(1..List.List'Last+3);
         begin
            New_List(List.List'Range) := List.List.all;
            -- Deallocate list.list access
            List.List := New_List;
         end;
      end if;

      List.Num_Elements            := List.Num_Elements + 1;
      List.List(List.Num_Elements) := Element;

   end List_Add;


   -- ---------------------------------------------------------

   function List_New_Iterator( List : in List_Type )
     return List_Iterator_Type is

      List_Iterator : List_Iterator_Type;
   begin

      List_Iterator.List         := List.List;
      List_Iterator.Num_Elements := List.Num_Elements;

      return List_Iterator;

   end List_New_Iterator;



   -- ----- List_Iterator_Type Methods ------------------------

   function  Is_Next( List_Iterator  : in List_Iterator_Type )
     return Boolean is
   begin
      if List_Iterator.Index <= List_Iterator.Num_Elements then
         return True;
      else
         return False;
      end if;
   end Is_Next;


   -- ---------------------------------------------------------

   procedure Get_Next( List_Iterator  : in out List_Iterator_Type;
                       Next_Element   : out    Element_Type ) is
   begin

      if not Is_Next( List_Iterator ) then
         raise Iterator_Bound_Error;
      end if;

      Next_Element := List_Iterator.List(List_Iterator.Index);
      List_Iterator.Index := List_Iterator.Index + 1;

   end Get_Next;


end Generic_List;
