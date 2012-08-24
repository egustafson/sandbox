with TEXT_IO; use TEXT_IO;
with TASK_CONTROL;
use TASK_CONTROL;

procedure AB is

  task A;
  task B;  
  task C;  
  task D;

--                            --
-- End of type declairations. --
--                            --

  task body A is
  begin
    delay(0.01);
    loop
      PUT("A");
      NEW_LINE;
    end loop;
  end A;

  task body B is
  begin
    delay(0.01);
    loop
      PUT("  B");
      NEW_LINE;
    end loop;
  end B;

  task body C is
  begin
    delay(0.01);
    loop
      PUT("  C");
      NEW_LINE;
    end loop;
  end C;
  
  task body D is
  begin
    delay(0.01);
    loop
      PUT("    D");
      NEW_LINE;
    end loop;
  end D;

-- The tasks become active as soon as the procedure under which
-- their scope falls becomes active.  (i.e. when procedure HW
-- becomes active then it's tasks are started.

begin -- AB
  task_control.pre_emption_on;
  task_control.set_time_slice(0.001);  
end AB;






