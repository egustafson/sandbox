with Ada.Text_IO, Ada.Integer_Text_IO;
use  Ada.Text_IO, Ada.Integer_Text_IO;

procedure Catch_Error is

   subtype Small_Integer is Integer range 0 .. 1000;

   Result       : Small_Integer;
   Multiplicand : Small_Integer;
begin


   Multiplicand := 2;

   -- The following should evaluate to an out of range value
   -- and thus should raise the Constraint_Error exception.

   Result := Small_Integer'Last * Multiplicand;

   Put("The result of multiplying ");
   Put(Small_Integer'Last);
   Put(" and ");
   Put(Multiplicand);
   Put(" is ");
   Put(Result);
   Put_Line(".");

exception
   when Constraint_Error =>
      Put("The result was out of the range ");
      Put(Small_Integer'First);
      Put(" .. ");
      Put(Small_Integer'Last);
      Put_Line(".");

end Catch_Error;
