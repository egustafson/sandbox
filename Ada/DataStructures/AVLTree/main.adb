with Generic_AVL_Tree, Ada.Numerics.Discrete_Random;
with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Command_Line;
use  Ada.Text_IO, Ada.Integer_Text_IO, Ada.Command_Line;

procedure Main is

   Num_Insertions : Integer := 0;
   Num_Queries    : Integer;
   Scratch        : Boolean;
   Rand           : Integer;

   procedure Integer_Put( Item: Integer ) is
   begin
      Put( Item );
   end Integer_Put;

   package Int_AVL_Tree is
     new Generic_AVL_Tree( Element_Type => Integer,
                           Less_Than    => "<",
                           Greater_Than => ">",
                           Put          => Integer_Put );

   package Random_Int is new Ada.Numerics.Discrete_Random(Integer);

   List           : Int_AVL_Tree.T;
   G              : Random_Int.Generator;

--    Sequence_File  : Ada.Text_IO.File_Type;

   procedure Print_Usage is
   begin
      New_Line;
      Put("Usage:  " & Command_Name & " <number of insertions> <number of queries>");
      New_Line; New_Line;
   end Print_Usage;

begin

   if Argument_Count /= 2 then
      Print_Usage;
   else
      begin
         Num_Insertions := Positive'Value( Argument(1) );
         Num_Queries    := Natural'Value( Argument(2) );
      exception
         when Constraint_Error =>
            Print_Usage;
            raise;
         when others =>
            raise;
      end;
   end if;

   if Num_Insertions > 0 then

--       Put("Inserting 2^(" & Integer'Image(Num_Insertions));
--       Put(") elements and then making 2^(" & Integer'Image(Num_Queries));
--       Put_Line(") queries.");

--       Open( Sequence_File, In_File, Name => "sequence.dat" );

      Random_Int.Reset(G);

      begin
         for I in 1 .. 2**Num_Insertions loop
            Rand := Random_Int.Random(G);
--             Put( Sequence_File, Rand ); New_Line( Sequence_File );
--             Get( Sequence_File, Rand );
            Int_AVL_Tree.Insert( List, Rand );
         end loop;
      exception
         when others =>
--             Close( Sequence_File );
            raise;
      end;

--       Close( Sequence_File );

      for I in 1 .. 2**Num_Queries loop
         Scratch := Int_AVL_Tree.Is_In_Tree( List, Random_Int.Random(G) );
      end loop;

--       Put_Line("Done.");

   end if;

end Main;
