/* intlist.cpp          -*-  C++ -*- */

#include "intlist.h"

#include <stdio.h>

// ------------------------------------------------------------

IntList::IntList()
    : Head(NULL), Tail(NULL)
{ }

// ------------------------------------------------------------

IntList_Node* IntList::find_Element(const int Number) const {

    IntList_Node* Next_Node = Head;

    while ( NULL != Next_Node ) {
        if ( Number == Next_Node->Datum ) {
            return Next_Node;
        }
        Next_Node = Next_Node->Next;
    }
    return NULL;
}

// ------------------------------------------------------------

bool IntList::isInList(const int Number) const {

    return ( find_Element(Number) != NULL );
}

// ------------------------------------------------------------

void IntList::insert(const int Number) {

    IntList_Node* New_Node;

    if ( isInList( Number ) ) {
        return;
    }

    New_Node = new IntList_Node;
    New_Node->Datum = Number;
    New_Node->Next  = NULL;

    if ( NULL == Head ) {
        Head = New_Node;
        Tail = New_Node;
    } else {
        Tail->Next = New_Node;
        Tail       = New_Node;
    }
}

// ------------------------------------------------------------

void IntList::remove(const int Number) {

    fprintf(stderr, "Not Implemented\n");
    exit(1);
}

// ------------------------------------------------------------

void IntList::print() const {

    IntList_Node* Next_Node = Head;

    while ( NULL != Next_Node ) {
        printf("%d\n", Next_Node->Datum);
        Next_Node = Next_Node->Next;
    }
}

// ------------------------------------------------------------
//
// End of intlist.cpp
//
