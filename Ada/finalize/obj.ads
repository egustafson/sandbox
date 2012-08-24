with Ada.Finalization;

package Obj is

   type Obj_T is new Ada.Finalization.Controlled with private;

   function New_Obj( I : in Integer ) return Obj_T;

   procedure Put( O : Obj_T );

private

   type Obj_T is new Ada.Finalization.Controlled with
      record
         X      : Integer := 0;
         Serial : Integer := 0;
      end record;

   procedure Initialize( Object: in out Obj_T );

   procedure Adjust( Object: in out Obj_T );

   procedure Finalize( Object: in out Obj_T );

end Obj;
