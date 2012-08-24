with Generic_Stack, Generic_Queue;

with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Strings.Bounded;
use  Ada.Text_IO, Ada.Integer_Text_IO;

procedure Test_Containers is

   package String80 is new Ada.Strings.Bounded.Generic_Bounded_Length(80);

   package Int_Stack is new Generic_Stack(Element_Type => Integer);
   package Str_Stack is new Generic_Stack(Element_Type => String80.Bounded_String);

   package Int_Queue is new Generic_Queue(Element_Type => Integer);

   Stack : Int_Stack.T;
   Queue : Int_Queue.T;
   Data  : Integer;

   Stack_Str : Str_Stack.T;
   Str       : String80.Bounded_String;
   Str_Array : array (1..3) of String80.Bounded_String
     := (String80.To_Bounded_String("One"),
         String80.To_Bounded_String("Two"),
         String80.To_Bounded_String("Three") );
begin

   for I in 1 .. 20 loop
      Int_Stack.Push( Stack, I );
      Int_Queue.Enqueue( Queue, I );
   end loop;

   Put("The top of the stack is ");
   Int_Stack.Peek( Stack, Data );
   Put( Data );
   New_Line;

   Put("The head of the queue is");
   Int_Queue.Peek( Queue, Data );
   Put( Data );
   New_Line;

   Put_Line("      Stack :    Queue ");

   while not Int_Stack.Empty( Stack ) loop
      Int_Stack.Pop( Stack, Data );
      Put( Data );
      Int_Queue.Dequeue( Queue, Data );
      Put( Data );
      New_Line;
   end loop;
   New_Line;



   for J in Str_Array'Range loop
      Str_Stack.Push(Stack_Str, Str_Array(J));
   end loop;

   Put("The top of the Stack is ");
   Str_Stack.Peek( Stack_Str, Str );
   Put( String80.To_String(Str) );
   New_Line;

   Put_Line("The Stack is:");

   while not Str_Stack.Empty( Stack_Str ) loop
      Str_Stack.Pop( Stack_Str, Str );
      Put( String80.To_String(Str) );
      New_Line;
   end loop;
   New_Line;

   Put_Line("Done.");

end Test_Containers;
