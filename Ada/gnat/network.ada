-- network.ada
--
--  Title: 	Network
--  Created:	Tue Apr 13 00:56:28 1993
--  Author: 	Eric Gustafson
--		<egustafs@play-doh>
--
--  Interface:
--    All nodes will be lettered sequentially from 'A' nodes are
--  case insensitive.  lambda and gama values are integer values 
--  in the range of 0..max_int.

with TEXT_IO;

procedure NETWORK is

  package INT_IO is new TEXT_IO.INTEGER_IO( INTEGER );

  use TEXT_IO;
  use INT_IO;

  MAX_NUMBER_NODES : constant INTEGER := 26;

  type NETWORK_MATRIX_TYPE is array(1..MAX_NUMBER_NODES, 
    	    	    	    	    1..MAX_NUMBER_NODES) of INTEGER;
  
  GAMA_IJ    : NETWORK_MATRIX_TYPE;	 -- End to end trafic
  GAMA       : INTEGER;			 -- Total end to end trafic
  LAMBDA_IJ  : NETWORK_MATRIX_TYPE;	 -- Link trafic
  LAMBDA     : INTEGER;			 -- Total link trafic
  CAPACITY_I : NETWORK_MATRIX_TYPE;	 -- Link Capacities
  C_BAR      : INTEGER;			 -- Average Link Capacity
  TIME_I     : NETWORK_MATRIX_TYPE;	 -- Link time delayes
  TIME_BAR   : INTEGER;			 -- Average time delay
  N_BAR      : INTEGER;			 -- Average number of packets in the net

  TEMPORARY  : INTEGER;

-- =====================================================================================

  procedure GET_INTEGER( X : out INTEGER ) is

    VALID_INPUT : BOOLEAN := false;
    
  begin -- GET_INTEGER

    while NOT( VALID_INPUT ) loop
    begin
      GET( X );
      VALID_INPUT := TRUE;
    exception
      when DATA_ERROR =>
        SKIP_LINE;
        PUT_LINE("Invalid format!");
      when others =>
        raise;
    end;    
    end loop;
  
  end GET_INTEGER;

-- =====================================================================================

  procedure GET_CHARACTER( X : out CHARACTER ) is
  
    VALID_INPUT : BOOLEAN := false;
    CH          : CHARACTER;
  
  begin -- GET_CHARACTER
  
    while NOT( VALID_INPUT ) loop
      GET( CH );
      if ( CH >= 'A' and CH <= 'Z' ) then 
        VALID_INPUT := true;
      elsif ( CH >= 'a' and CH <= 'z' ) then
        VALID_INPUT := true;
	CH := CHARACTER'VAL( CHARACTER'POS(CH) - 
	      (CHARACTER'POS('a') - CHARACTER'POS('A')));
      end if;    
      if NOT( VALID_INPUT ) then
        SKIP_LINE;
	PUT_LINE("Invalid format!");
      end if;
    end loop;
    X := CH;
  
  end GET_CHARACTER;

-- =====================================================================================

  procedure GET_GAMA( GAMA_IJ : out NETWORK_MATRIX_TYPE ) is
  
    CH          : CHARACTER;
    START_NODE  : INTEGER;
    END_NODE    : INTEGER;
    LINK_WEIGHT : INTEGER;
      
  begin -- GET_GAMA
    PUT_LINE("Please enter each gama value in the following form:");
    PUT_LINE("     <node><node> <weight>");
    NEW_LINE;
    PUT("1: ");
    GET_CHARACTER( CH );
    START_NODE := CHARACTER'POS( CH ) - CHARACTER'POS('A') + 1;
    GET_CHARACTER( CH );
    END_NODE := CHARACTER'POS( CH ) - CHARACTER'POS('A') + 1;
    GET_INTEGER( LINK_WEIGHT );
    
    NEW_LINE;
    PUT("From:   ");
    PUT( START_NODE );
    NEW_LINE;
    PUT("To:     ");
    PUT( END_NODE );
    NEW_LINE;
    PUT("Weight: ");
    PUT( LINK_WEIGHT );
    NEW_LINE;

  end GET_GAMA;

-- =====================================================================================

begin -- NETWORK

  GET_GAMA( GAMA_IJ );
    
end NETWORK;
