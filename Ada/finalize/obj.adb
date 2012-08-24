with Ada.Text_IO;          use Ada.Text_IO;
-- with Ada.Integer_Text_IO;  use Ada.Integer_Text_IO;

with Ada.Finalization;     use Ada.Finalization;


package body Obj is


   Master_Serial : Integer := 1;


   function New_Obj( I : in Integer ) return Obj_T is
      Obj : Obj_T;
   begin
      Obj.X := I;
      return Obj;
   end New_Obj;


   procedure Put( O : Obj_T ) is
   begin
      Put("( " & Integer'Image(O.X) &
          "[" & Integer'Image(O.Serial) & "] )");
   end Put;


   procedure Initialize( Object: in out Obj_T ) is
   begin
      Object.Serial := Master_Serial;
      Master_Serial := Master_Serial + 1;
      Put("Initializing: ");
      Put( Object );
      New_Line;
   end Initialize;


   procedure Adjust( Object: in out Obj_T ) is
   begin
      Put("Adjusting:    ");
      Put( Object );
      New_Line;
   end Adjust;


   procedure Finalize( Object: in out Obj_T ) is
   begin
      Put("Finalizing:   ");
      Put( Object );
      New_Line;
   end Finalize;

end Obj;
