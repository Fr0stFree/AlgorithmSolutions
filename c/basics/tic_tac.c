#include <stdio.h>

typedef struct {
    int h;
    int min;
} TicTac;

TicTac after(TicTac a, int min) {
    int total_min = a.h * 60 + a.min + min;
    total_min %= (12 * 60);
    if (total_min < 0) {
        total_min += 12 * 60;
    }
    TicTac result = {
        .h = total_min / 60,
        .min = total_min % 60,
    };
    return result;
}

void forward(TicTac * me, TicTac a) {
    int total_min = me -> h * 60 + me -> min + a.h * 60 + a.min;
    total_min %= (12 * 60);
    if (total_min < 0) {
        total_min += 12 * 60;
    }
    me -> h = total_min / 60;
    me -> min = total_min % 60;
};

void backward(TicTac * me, TicTac a) {
    int total_min = me -> h * 60 + me -> min - (a.h * 60 + a.min);
    total_min %= (12 * 60);
    if (total_min < 0) {
        total_min += 12 * 60;
    }
    me -> h = total_min / 60;
    me -> min = total_min % 60;
}

int isEqualTime(TicTac a, TicTac b) {
    if (a.h == b.h && a.min == b.min) {
        return 1;
    }
    return 0;
};

void printTic(TicTac a) {
    printf("%02d:%02d\n", a.h, a.min);
};

int main(){
    TicTac a,b,c;
    int mk;

    scanf("%d:%d", &(a.h), &(a.min));
    scanf("%d", &mk);
    scanf("%d:%d", &(b.h), &(b.min));

    printf("equal: %d\n",isEqualTime(a,b));
    c = after(a, mk);
    printf("after: ");
    printTic(c);

    c = a;
    printf("forward: ");
    forward(&a, b);
    printTic(a);

    printf("backward: ");
    backward(&c, b);
    printTic(c);
    return 0;
}

