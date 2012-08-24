with Generic_Binary_Tree, Ada.Numerics.Discrete_Random;
with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Command_Line;
use  Ada.Text_IO, Ada.Integer_Text_IO, Ada.Command_Line;

procedure Main is

   Num_Insertions : Integer := 0;
   Num_Queries    : Integer;
   Scratch        : Boolean;

   procedure Integer_Put( Item: Integer ) is
   begin
      Put( Item );
   end Integer_Put;

   package Int_Binary_Tree is
     new Generic_Binary_Tree( Element_Type => Integer,
                              Less_Than    => "<",
                              Greater_Than => ">",
                              Put          => Integer_Put );
   package Random_Int is new Ada.Numerics.Discrete_Random(Integer);

   List           : Int_Binary_Tree.T;
   G              : Random_Int.Generator;

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

      Random_Int.Reset(G);
      for I in 1 .. 2**Num_Insertions loop
         Int_Binary_Tree.Insert( List, Random_Int.Random(G) );
      end loop;
      for I in 1 .. 2**Num_Queries loop
         Scratch := Int_Binary_Tree.Is_In_Tree( List, Random_Int.Random(G) );
      end loop;

--       Put_Line("Done.");

   end if;

end Main;
