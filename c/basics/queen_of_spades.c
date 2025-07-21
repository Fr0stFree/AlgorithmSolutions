#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

struct Card
{
    char rank;
    char suit;
};

int check(struct Card hand[]);

int main()
{
    struct Card hand1[] = {{'Q', 's'}, {'A', 'h'}, {'9', 'd'}, {0, 0}};
    struct Card hand2[] = {{'2', 's'}, {'A', 'h'}, {0, 0}};
    assert(1 == check(hand1));
    assert(0 == check(hand2));

    return 0;
}

int check(struct Card hand[])
{
    struct Card looking_card = {.rank = 'Q', .suit = 's'};
    struct Card *p_current_card;

    for (p_current_card = hand; p_current_card->rank != 0; p_current_card++)
    {
        if (p_current_card->rank == looking_card.rank && p_current_card->suit == looking_card.suit)
        {
            return 1;
        }
    }
    return 0;
}