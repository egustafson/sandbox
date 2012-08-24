#include <list>
#include <string>

class Registry {
public:
    ~Registry();
    static Registry* get_instance();
    void log(const string& name);
    void print_entries() const;
private:
    Registry();

    static Registry* instance;
    list<string> myEntries;
};
