#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void get_ln(char* buffer, size_t max_len)
{
    fgets(buffer, max_len-1, stdin);
    char* ptr_n = strrchr(buffer, '\n');
    if(ptr_n != NULL)
        *ptr_n = '\0';
}

char * merge_string(const char* str1, const char* str2)
{
    
    size_t length1 = strlen(str1);
    size_t length2 = strlen(str2);
    char *result = malloc(sizeof(char) * (length1 + length2 + 1));
    memcpy(result, str1, length1);
    memcpy(result + length1, str2, length2 + 1);
    return result;
}

int main(void)
{
    char str_1[200] = {0};
    char str_2[200] = {0};

    get_ln(str_1, sizeof(str_1));
    get_ln(str_2, sizeof(str_2));

    char *sentence = merge_string(str_1, str_2);
    puts(sentence);
    free(sentence);

    return 0;
}