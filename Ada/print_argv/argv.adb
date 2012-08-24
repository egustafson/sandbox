with Ada.Command_Line, Ada.Text_IO;
use  Ada.Command_Line, Ada.Text_IO;

procedure Argv is
begin
   Put_Line( "Command   :  " & Command_Name );
   Put( "Args(" & Integer'Image(Argument_Count) & " ) :  " );
   Put( "[ " );
   for I in 1 .. Argument_Count loop
      Put( Argument(I) & " ");
   end loop;
   Put_Line( "]" );
end Argv;
