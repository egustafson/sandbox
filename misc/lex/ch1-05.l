%{
/*
 * We now build a lexical analyzer to be used by a higher-level parser.
 */

#include "y.tab.h"

#define   LOOKUP 0  /* default - not a defined word type */

int state;

%}

%%

\n      { state = LOOKUP; }

\.\n    { state = LOOKUP;
          return 0;  /* end of sentence */
        }

^verb   { state = VERB; }
^adj    { state = ADJECTIVE; }
^adv    { state = ADVERB; }
^noun   { state = NOUN; }
^prep   { state = PREPOSITION; }
^pron   { state = PRONOUN; }
^conj   { state = CONJUNCTION; }

[a-zA-Z]+ {
            if ( state != LOOKUP ) {
              add_word(state, yytext);
            } else {
              switch( lookup_word(yytext) ) {
                case VERB:
                  printf("VERB:  %s\n", yytext);
                  return(VERB);
                case ADJECTIVE:
                  printf("ADJECTIVE:  %s\n", yytext);
                  return(ADJECTIVE);
                case ADVERB:
                  printf("ADVERB:  %s\n", yytext);
                  return(ADVERB);
                case NOUN:
                  printf("NOUN:  %s\n", yytext);
                  return(NOUN);
                case PREPOSITION:
                  printf("PREPOSITION:  %s\n", yytext);
                  return(PREPOSITION);
                case PRONOUN:
                  printf("PRONOUN:  %s\n", yytext);
                  return(PRONOUN);
                case CONJUNCTION:
                  printf("CONJUNCTION:  %s\n", yytext);
                  return(CONJUNCTION);
                default:
                  printf("%s:  don't recognize.\n", yytext);
                  /* don't return, just ignore it. */
              }
            }
          }

.       ;

%%

/* define a linked list of words and types */

struct word {
    char *word_name;
    int   word_type;
    struct word *next;
};

struct word *word_list; /* first element in word list */

extern void *malloc();

int add_word(int type, char *word) {

    struct word *wp;

    printf("Adding word:  %s\n", word);

    if ( lookup_word(word) != LOOKUP ) {
        printf("!!!warning:  word %s already defined \n", word);
        return 0;
    }

    wp = (struct word *)malloc(sizeof(struct word));

    wp->next = word_list;

    wp->word_name = (char *)malloc(strlen(word)+1);
    strcpy(wp->word_name, word);
    wp->word_type = type;
    word_list = wp;
    return 1;
}

int lookup_word(char *word) {

    struct word *wp = word_list;

    for (; wp; wp = wp->next) {
    
        if ( strcmp(wp->word_name, word) == 0) {
            return wp->word_type;
        }
    }
    return LOOKUP;  /* not found */
}