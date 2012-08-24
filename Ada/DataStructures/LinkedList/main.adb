with IntList, Ada.Text_IO;
use Ada.Text_IO;

procedure Main is
   List : IntList.T;
begin

   for I in 1 .. 20 loop
      IntList.Insert( List, I );
   end loop;

   for I in 1 .. 20 loop
      if ( I mod 3 = 0 ) then
         IntList.Remove( List, I );
      end if;
   end loop;

   IntList.Print( List );

   Put_Line("Done.");

end Main;
