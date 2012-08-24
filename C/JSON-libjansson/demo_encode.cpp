#include <iostream>

#include <stdlib.h>
#include <string.h>

#include <jansson.h>

int main(int argc, char* argv[]) {

    std::cout<<"Jansson version: "<<JANSSON_VERSION<<std::endl;

    json_t* json_obj = json_object();
    json_object_set( json_obj, "key", json_string("value") );

    char* json_enc = json_dumps(json_obj, JSON_ENSURE_ASCII);

    std::cout<<json_enc<<std::endl;

    free(json_enc);

    std::cout<<"done."<<std::endl;

}
