with Ada.Text_IO, Ada.Integer_Text_IO;
use  Ada.Text_IO, Ada.Integer_Text_IO;

procedure Error is

   -- Note: If Constrained_Integer is constrained to 0
   -- .. Integer'Last then Constraint_Error is raised when
   -- Result1 is calculated.  It appears to me that the
   -- boundry check is occurring after the calculation
   -- completed and doesn't check the signs of the terms
   -- vs. the results.  Similar results occur with addition.

   subtype Constrained_Integer is Integer
     range Integer'First .. Integer'Last;

   subtype Small_Integer is Integer range 0 .. 1000;

   Multiplicand1 : Integer;
   Multiplicand2 : Integer;

   Result1       : Constrained_Integer;
   Result2       : Small_Integer;

begin

   Multiplicand1 := 2;
   Multiplicand2 := Constrained_Integer'Last;

   -- Constraint_Error should get raised with this
   -- statement.
   Result1 := Multiplicand1 * Multiplicand2;

   Put("The result of multiplying ");
   Put(Multiplicand1);
   Put(" and ");
   Put(Multiplicand2);
   Put(" is ");
   Put(Result1);
   Put_Line(".");
   Flush;

   Multiplicand2 := Small_Integer'Last;

   -- Constraint_Error should get raised with this
   -- statement.
   Result2 := Multiplicand1 * Multiplicand2;

   Put("The result of multiplying ");
   Put(Multiplicand1);
   Put(" and ");
   Put(Multiplicand2);
   Put(" is ");
   Put(Result2);
   Put_Line(".");
   Flush;

end Error;
