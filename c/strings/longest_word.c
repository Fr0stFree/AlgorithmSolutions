#include <stdio.h>
#include <string.h>

const char *delimiters = " \n";

int main()
{
    char words[1000];
    char longest_word[100] = "\0";

    while (scanf("%999s", words) == 1)
    {   
        for (char *token = strtok(words, delimiters); token != NULL; token = strtok(NULL, delimiters))
        {
            if (strlen(longest_word) < strlen(token))
            {
                strcpy(longest_word, token);
            }
        }
    }

    printf("%s %lu", longest_word, strlen(longest_word));

    return 0;
}
