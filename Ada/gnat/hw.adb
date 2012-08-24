with TEXT_IO; use TEXT_IO;

procedure HW is

  task HELLO;

  task WORLD is
    entry PRINT_NOW;
  end WORLD;

--                            --
-- End of type declairations. --
--                            --

  task body HELLO is
  begin
    PUT("Hello, ");
    WORLD.PRINT_NOW;
  end HELLO;

  task body WORLD is
  begin
    accept PRINT_NOW;
    PUT("World");
    NEW_LINE;
  end WORLD;

-- The tasks become active as soon as the procedure under which
-- their scope falls becomes active.  (i.e. when procedure HW
-- becomes active then it's tasks are started.

begin -- HW (Hello World)
  null;
end HW;
