#include <stdio.h>
#include <string.h>

int main() {
    int hour, min, sec;
    char hours[3], mins[3], secs[3];
    char result[10];
    scanf("%d %d %d", &hour, &min, &sec);

    if (hour > 23 || hour < 0) {
        strcpy(hours, "--");
    } else {
        sprintf(hours, "%02d", hour);
    }
    if (min > 59 || min < 0) {
        strcpy(mins, "--");
    } else {
        sprintf(mins, "%02d", min);
    }
    if (sec > 59 || sec < 0) {
        strcpy(secs, "--");
    } else {
        sprintf(secs, "%02d", sec);
    }
    sprintf(result, "%s:%s:%s", hours, mins, secs);
    puts(result);
    return 0;
}

