with Ada.Text_IO, Ada.Integer_Text_IO;
use  Ada.Text_IO, Ada.Integer_Text_IO;

procedure Multiply is
   Number_Of_Cases, A, B : Integer;
begin

   Put("Integer range: ");
   Put(Integer'First);
   Put(" .. ");
   Put(Integer'Last);
   New_Line;

   Put_Line("How many cases? ");
   Flush;
   Get(Number_Of_Cases);

   for I in 1 .. Number_Of_Cases loop

      Put_Line("Enter two numbers.");
      Flush;
      Get(A);
      Get(B);
      Put("The product is ");
      Put(A*B);
      Put_Line(".");

   end loop;

end Multiply;
