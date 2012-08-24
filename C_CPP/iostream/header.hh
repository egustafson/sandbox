#ifdef _DEBUG

#define LOG_DEBUG cout<<__FILE__<<"("<<__LINE__<<") - DBUG: "
#define LOG_INFO  cout<<__FILE__<<"("<<__LINE__<<") - INFO: "
#define LOG_WARN  cout<<__FILE__<<"("<<__LINE__<<") - WARN: "

#else

#define LOG_DEBUG cout<<"DBUG: "
#define LOG_INFO  cout<<"INFO: "
#define LOG_WARN  cout<<"WARN: "

#endif
