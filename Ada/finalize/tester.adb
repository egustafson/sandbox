with Obj;          use Obj;


procedure Tester is

   A, B, C  : Obj_T;

begin

   A := New_Obj( 1 );
   B := New_Obj( 2 );
   C := New_Obj( 3 );

   A := C;

end Tester;
