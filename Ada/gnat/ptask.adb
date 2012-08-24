with text_io, task_control;
use text_io, task_control;

procedure ptask is

-- y : duration := 0.0005;

  package int_io is new integer_io (integer);
  use int_io;

  task looper is
  end;

  task body looper is
  begin
     for j in 1 .. 100 loop
       put (j);
       new_line;
     end loop;
     new_line;
     put_line ("task integers complete");
     new_line;
  end;

  task scooper is
  end;

  task body scooper is
  begin
  task_control.set_time_slice (0.001);
   for k in 1..4 loop
     for j in 'A'..'Z' loop
       put (j);
       new_line;
     end loop;
     new_line;
  end loop;
  put_line ("task letters complete");
  new_line;
  end;

begin -- ptask
  task_control.pre_emption_on;

  for x in 1000 ..1020 loop
--    task_control.set_time_slice (y);
--    y := y + y;
    put(x);
    new_line;
  end loop;
  put_line ("proc complete");

end ptask;

