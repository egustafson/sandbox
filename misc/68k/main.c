/* ----- */


void  set_sr(short new_sr);
short get_sr();

/* **************************************** */

int main() {

    short sr;

    sr = get_sr();
    sr = sr - 0x700;
    set_sr(sr);
    return 0;
}

/* **************************************** */

void set_sr(short new_sr) {
    asm("move sp@(6),sr");
}

short get_sr() {
    
    asm("clrl d0");
    asm("move sr,d0");
}
