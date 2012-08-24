/* commands.h */

#ifndef COMMANDS_H
#define COMMANDS_H


#define PING_REQUEST   1
#define PING_RESPONSE  2
#define ECHO           3
#define QUIET        252
#define VERBOSE      253
#define CLOSE_SOCK   254
#define SHUTDOWN     255

#define BUFF_SIZE 255

typedef unsigned char uchar;

typedef struct {
    uchar msgType;
    char  buff[BUFF_SIZE];
} message_s;

#endif /* COMMANDS_H */
