%token NAME NUMBER
%%
statement:  NAME '=' expression
    |       expression          { printf("= %d\n"); }
    ;

expression: expression '+' NUMBER { $$ = $1 + $3; }
    |       expression '-' NUMBER { $$ = $1 - $3; }
    |       NUMBER                { $$ = $1; }
    ;
%%
int yyerror(char* s) {

    fprintf(stderr, "%s\n", s);
}