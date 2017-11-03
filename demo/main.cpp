#include "hu.h"

void vTest1()
{
    long long a = 0;
    long long b = 1;
    long long c = 2;
    long long arr[3] = { 0, 1, 2 };
    printf("0 1 2\n");

    set<long long> set2;
    for (long long i = 0; i < 5; i++)
    {
        for (long long j = 0; j < 5; j++)
        {
            long long iNum = i * 10 + j;
            if (iNum >= 10)
            {
                set2.insert(iNum);
            }
        }
    }

    set<long long>::iterator iter2 = set2.begin();
    while (iter2 != set2.end())
    {
        printf("%d ", *iter2);
        ++iter2;
    }
    printf("\n");


    set<long long> set3;
    iter2 = set2.begin();
    while (iter2 != set2.end())
    {
        long long iNum2 = *iter2;

        for (long long i = 0; i < 3; i++)
        {
            long long iNum3 = i * 100 + iNum2;
            if (iNum3 >= 100)
            {
                set3.insert(iNum3);
            }

            iNum3 = iNum2 * 10 + i;
            if (iNum3 >= 100)
            {
                set3.insert(iNum3);
            }

            iNum3 = iNum2 / 10 * 100 + i * 10 + iNum2 % 10;
            if (iNum3 >= 100)
            {
                set3.insert(iNum3);
            }
        }

        ++iter2;
    }

    set<long long>::iterator iter3 = set3.begin();
    while (iter3 != set3.end())
    {
        printf("%d ", *iter3);
        ++iter3;
    }
    printf("\n");
}

int main()
{
    //vTest1();

    CLaiZiHu oLaiZiHu;

    long long llRet = oLaiZiHu.GetMinLaiZi(111031110);
    printf("vTest1 llRet = %d\n", llRet);

    Card aucHandCards[MAX_HANDCARD_NUM] = {
        0x01, 0x01, 0x01,

        0x04, 0x05, 0x06, 0x06, 0x06, 0x07, 0x08, 0x08,

        0x15, 0x17, 0x19,
    };
    bool bCanHu = oLaiZiHu.bHu(aucHandCards, MAX_HANDCARD_NUM, 0x01);
    printf("vTest1 bCanHu = %d\n", bCanHu);

    return 0;
}