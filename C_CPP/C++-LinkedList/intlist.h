/* intlist.h          -*- C++ -*- */

class IntList_Node {
public:
    int           Datum;
    IntList_Node* Next;
};

class IntList {
private:
    IntList_Node* Head;
    IntList_Node* Tail;

    IntList_Node* find_Element(const int Number) const;
public:
    IntList();

    bool isInList(const int Number) const;
    void insert(  const int Number);
    void remove(  const int Number);
    void print() const;
};

// End of intlist.h
