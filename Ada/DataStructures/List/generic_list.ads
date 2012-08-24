-- generic_list.ads          -*- Ada -*-
--
-- This package defines a generic list and list iterator.
--
-- Author:  Eric Gustafson
-- Date:    11 August 1998
--

-- ------------------------------------------------------------
--
-- $Revision$
--

-- $Log$

-- ------------------------------------------------------------

generic

   type Element_Type is private;

package Generic_List is


   -- ----- Fundamental Object Types --------------------------

   type List_Type is private;
   type List_Iterator_Type is private;

   -- Raised by Iterator methods
   Iterator_Bound_Error  : exception;

   -- ----- List_Type Methods ---------------------------------

   procedure List_Add(          List     : in out List_Type;
                                Element  : in     Element_Type );

   function  List_New_Iterator( List     : in     List_Type )
     return List_Iterator_Type;


   -- ----- List_Iterator_Type Methods ------------------------

   function  Is_Next( List_Iterator  : in List_Iterator_Type )
     return Boolean;

   procedure Get_Next( List_Iterator  : in out List_Iterator_Type;
                       Next_Element   : out    Element_Type );

-- ------------------------------------------------------------
-- ----------          Private Section               ----------
-- ------------------------------------------------------------
private

   type Element_Array is array (Positive range <>) of Element_Type;
   type Element_Array_Access is access Element_Array;

   type List_Type is
      record
         List          : aliased Element_Array_Access := new Element_Array(1..3);
         Num_Elements  : Natural := 0;
      end record;

   type List_Iterator_Type is
      record
         List          : Element_Array_Access;
         Num_Elements  : Natural;
         Index         : Positive := 1;
      end record;

end Generic_List;
