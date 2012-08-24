--  int_in.ada
--
--
--  Title: 	Integer Input
--  Created:	Wed Apr 14 12:53:06 1993
--  Author: 	Eric Gustafson
--		<egustafs@play-doh>
--


with TEXT_IO;

procedure INT_IN is

  package INT_IO is new TEXT_IO.INTEGER_IO( INTEGER );

  use TEXT_IO;
  use INT_IO;

  I : INTEGER;
  VALID_INPUT : BOOLEAN;

begin -- INT_IN

  VALID_INPUT := FALSE;
  WHILE NOT( VALID_INPUT ) loop
  begin
    PUT("Enter Number: ");
    GET( I );
    VALID_INPUT := TRUE;
    NEW_LINE(2);
    
  exception
    when DATA_ERROR =>
      SKIP_LINE;
      NEW_LINE;
      PUT_LINE("Invalid format!");
    when others =>
      raise;
  end;    
  end loop;
  
  PUT("You entered:");
  PUT( I );
  NEW_LINE;
  
end INT_IN;
