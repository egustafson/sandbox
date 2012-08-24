with Generic_AVL_Tree;
with Ada.Text_IO, Ada.Integer_Text_IO;
use  Ada.Text_IO, Ada.Integer_Text_IO;

procedure Main is

   procedure Integer_Put( Item: Integer ) is
   begin
      Put( Item, 2 );
   end Integer_Put;

   package Int_AVL_Tree is
     new Generic_AVL_Tree( Element_Type => Integer,
                           Less_Than    => "<",
                           Greater_Than => ">",
                           Put          => Integer_Put );

   List        : Int_AVL_Tree.T;

   Num_Elements : constant Integer := 20;
   Init_Values : array ( Integer range <> ) of Integer :=
     ( 20, 10, 40, 30, 50, 35, 25, 55 );

--      ( 10, 5, 15, 3, 7, 12, 18,  2,  1,  4 );
--      ( 10, 5, 15, 3, 7, 12, 18, 20, 21, 17 );
--      ( 10, 5, 15, 3, 7, 12, 18, 2, 1, 4, 6, 8, 9, 11, 13, 14, 16, 17, 20, 19 );


begin

--    for I in Init_Values'range loop
--       Int_AVL_Tree.Insert( List, Init_Values(I) );
--    end loop;

   for I in 1 .. 2**22 loop
      Int_Avl_Tree.Insert( List, I );
   end loop;

--    for I in 1 .. Num_Elements loop
--       if I mod 3 = 0 then
--          Int_AVL_Tree.Remove( List, I );
--       end if;
--    end loop;

--    Int_AVL_Tree.Debug_Print( List );

end Main;
