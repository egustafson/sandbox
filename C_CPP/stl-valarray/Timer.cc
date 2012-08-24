// Timer.cc - A stopwatch for processes
// ------------------------------------------------------------
//
// Author: Eric Gustafson
// Date:   13 May 2003
//
// $Id$
//
// ////////////////////////////////////////////////////////////

#include "Timer.hh"

#include <unistd.h>
#include <iostream>
#include <iomanip>

// ------------------------------------------------------------

long Timer::sys_clktck = 0;

// --------------------

Timer::Timer()
{
    reset();

    if ( 0 == sys_clktck ) {
        sys_clktck = sysconf(_SC_CLK_TCK);
    }
}

// --------------------

Timer::~Timer()
{ /* nothing to do */ }

// --------------------

void Timer::start() 
{
    reset();
    starttime = times(&tmsstart);
}

// --------------------

void Timer::stop() 
{
    endtime = times(&tmsend);
}

// --------------------

void Timer::reset() 
{
    starttime = 0;
    endtime   = 0;

    tmsstart.tms_utime = 0;
    tmsstart.tms_stime = 0;

    tmsend.tms_utime = 0;
    tmsend.tms_stime = 0;
}

// --------------------

double Timer::getRealTime()
{
    double rt = (endtime-starttime) / (double) sys_clktck;
    return ( rt > 0.0 ? rt : 0.0 );
}

// --------------------

double Timer::getUserTime()
{
    double ut = (tmsend.tms_utime-tmsstart.tms_utime) / (double) sys_clktck;
    return ( ut > 0.0 ? ut : 0.0 );
}

// --------------------

double Timer::getSysTime()
{
    double st = (tmsend.tms_stime-tmsstart.tms_stime) / (double) sys_clktck;
    return ( st > 0.0 ? st : 0.0 );
}

// --------------------

void Timer::print(ostream& ostr) 
{
    ostr<<  "  real:  "<<std::setprecision(5)<<getRealTime()
        <<"\n  user:  "<<getUserTime()
        <<"\n   sys:  "<<getSysTime()<<endl;
}

// --------------------
