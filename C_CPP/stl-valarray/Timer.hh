// Timer.hh - A stopwatch for processes
// ------------------------------------------------------------
//
// Author: Eric Gustafson
// Date:   13 May 2003
//
// $Id$
//
// ////////////////////////////////////////////////////////////

#include <sys/times.h>
#include <ostream.h>

// ------------------------------------------------------------

class Timer {
  public:
    Timer();
    ~Timer();

    void start();
    void stop();
    void reset();

    double getRealTime();
    double getUserTime();
    double getSysTime();

    void print(ostream& ostr);

//     void lap();                 // future implementation
  private:    
    struct tms tmsstart;
    clock_t    starttime;
    struct tms tmsend;
    clock_t    endtime;

    static long sys_clktck;     // initialized in constructor.
};
