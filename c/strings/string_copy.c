
#include <stdio.h>
#include <string.h>
#define SIZE 50

int main() {
    char ps[7][SIZE];
    size_t ps_size = sizeof(ps)/sizeof(*ps);
    char ps_res[(SIZE+1) * 7] = "";
    
    for (size_t idx = 0; idx < ps_size; idx++) {
        scanf("%49s", ps[idx]);
    }
    for (size_t idx = 0; idx < ps_size; idx++) {
        if (idx != 0 && idx != ps_size) {
            strcat(ps_res, " ");
        }
        strcat(ps_res, ps[idx]);
    }

    puts(ps_res);
    return 0;
}