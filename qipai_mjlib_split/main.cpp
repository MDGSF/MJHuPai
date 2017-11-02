#include <stdio.h>
#include "split.h"

void test_split()
{
    char cards[] = {
        0,1,1,0,0,0,0,1,0,
        1,1,0,1,0,4,1,1,1,
        0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,2
    };

    bool bCanHu = split::get_hu_info(cards, 34, 33);
    printf("test_split bCanHu = %d\n", bCanHu);
}

void vTestError1()
{
    Card aucHandCards[MAX_HANDCARD_NUM] = {
        0x01, 0x01, 0x01, 0x05, 0x06, 0x07, 0x11, 0x12, 0x14, 0x14, 0x15, 0x16, 0x16, 0x34,
    };
    bool bCanHu = split::bHuHasLaizi(aucHandCards, 0x01);
    printf("vTest1 bCanHu = %d\n", bCanHu);
}

void vTest1()
{
    Card aucHandCards[MAX_HANDCARD_NUM] = {
        0x02, 0x02, 0x03, 0x04, 0x05, 0x09, 0x09, 0x12, 0x12, 0x14, 0x15, 0x28, 0x28, 0x28,
    };
    bool bCanHu = split::bHuHasLaizi(aucHandCards, 0x28);
    printf("vTest1 bCanHu = %d\n", bCanHu);
}

int main() 
{
    //test_split();
    vTest1();

    return 0;
}