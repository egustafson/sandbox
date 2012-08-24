with Ada.Finalization;

package Object is

   type Object is new Ada.Finalization.Controlled with private;

   function New_Object( X : in Integer ) return Object;

   procedure Put( O : Object );

private

   type Object is new Ada.Finalization.Controlled with
      record
         X : Integer := 0;
      end record;
