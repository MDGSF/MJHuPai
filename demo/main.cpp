#include <stdio.h>
#include <iostream>
#include <set>
#include <map>
#include <vector>
#include <queue>
using namespace std;

void vTest1()
{
    long long a = 0;
    long long b = 1;
    long long c = 2;
    long long arr[3] = { 0, 1, 2 };
    printf("0 1 2\n");

    set<long long> set2;
    for (long long i = 0; i < 3; i++)
    {
        for (long long j = 0; j < 3; j++)
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

void vShowVec(vector< set<long long> > & vec)
{
    vector< set<long long> >::iterator iterVec = vec.begin();
    while (iterVec != vec.end())
    {
        set<long long> & curset = *iterVec;
        set<long long>::iterator iterSet = curset.begin();
        while (iterSet != curset.end())
        {
            printf("%lld ", *iterSet);
            ++iterSet;
        }
        printf("\n\n");
        ++iterVec;
    }
}

bool bIsNumBitMoreThan14(long long iNum)
{
    int iSum = 0;
    while (iNum > 0)
    {
        iSum = iSum + iNum % 10;
        if (iSum > 14)
        {
            return true;
        }
        iNum = iNum / 10;
    }

    if (iSum > 14)
    {
        return true;
    }
    else
    {
        return false;
    }
}

void vGetNext(long long iNum, long long iNumBits, long long iInsertNum, set<long long> & curset)
{
    long long iStart = 1;
    long long iEnd = pow(10, iNumBits);
    for (long long i = iStart; i <= iEnd; i *= 10)
    {
        long long iNewNum = iNum / i*i * 10 + iInsertNum*i + iNum%i;
        //printf("%d ", iNewNum);
        if (iNewNum > iEnd && !bIsNumBitMoreThan14(iNewNum))
        {
            curset.insert(iNewNum);
        }
    }
}


map<long long, long long> handCardsMap;
vector< set<long long> > vec;

void vInit()
{
    //<handcards, need LaiZi Num>
    
    handCardsMap[0] = 0;
    handCardsMap[1] = 2;
    handCardsMap[2] = 1;

    handCardsMap[10] = 2;
    handCardsMap[11] = 1;
    handCardsMap[12] = 3;
    handCardsMap[20] = 1;
    handCardsMap[21] = 3;
    handCardsMap[22] = 2;

    handCardsMap[100] = 2;
    handCardsMap[101] = 1;
    handCardsMap[102] = 3;
    handCardsMap[110] = 1;
    handCardsMap[111] = 0;
    handCardsMap[112] = 2;
    handCardsMap[120] = 3;
    handCardsMap[121] = 2;
    handCardsMap[122] = 1;
    handCardsMap[200] = 1;
    handCardsMap[201] = 3;
    handCardsMap[202] = 2;
    handCardsMap[210] = 3;
    handCardsMap[211] = 2;
    handCardsMap[212] = 1;
    handCardsMap[220] = 2;
    handCardsMap[221] = 1;
    handCardsMap[222] = 0;

    
    set<long long> set1;
    set1.insert(0);
    set1.insert(1);
    set1.insert(2);
    vec.push_back(set1);

    set<long long> set2;
    set2.insert(10);
    set2.insert(11);
    set2.insert(12);
    set2.insert(20);
    set2.insert(21);
    set2.insert(22);
    vec.push_back(set2);

    set<long long> set3;
    set3.insert(100);
    set3.insert(101);
    set3.insert(102);
    set3.insert(110);
    set3.insert(111);
    set3.insert(112);
    set3.insert(120);
    set3.insert(121);
    set3.insert(122);
    set3.insert(200);
    set3.insert(201);
    set3.insert(202);
    set3.insert(210);
    set3.insert(211);
    set3.insert(212);
    set3.insert(220);
    set3.insert(221);
    set3.insert(222);
    vec.push_back(set3);

    long long iStart = 4;
    long long iStartNum = pow(10, iStart - 1);
    while (iStart < 6)
    {
        printf("iStart = %d\n", iStart);

        set<long long> curset;
        set<long long> & preset = vec.back();
        set<long long>::iterator iterPre = preset.begin();
        while (iterPre != preset.end())
        {
            long long iPreNum = *iterPre;

            for (long long i = 0; i < 3; i++)
            {
                vGetNext(iPreNum, iStart - 1, i, curset);
            }

            ++iterPre;
        }

        vec.push_back(curset);
        iStart++;
    }
    
    vShowVec(vec);
}

struct Sub {
    long long a;
    long long b;
};

int iGetNumBits(long long iNum)
{
    int iBitsNum = 0;
    while (iNum > 0)
    {
        iBitsNum++;
        iNum /= 10;
    }
    return iBitsNum;
}

int iGetNumRightBit(long long iNum)
{
    return iNum % 10;
}

int iGetNumLeftBit(long long iNum)
{
    int iBitsNum = iGetNumBits(iNum);
    return iNum / (pow(10, iBitsNum - 1));
}

Sub SplitNum(long long iNum, int i)
{
    int iSplitLine = pow(10, i);
    Sub sub;
    sub.a = iNum / iSplitLine;
    sub.b = iNum % iSplitLine;
    return sub;
}

void CalcMainSub(Sub MainSub, vector<Sub> & subVec)
{

}

void GetNumSub(long long iNum, vector<Sub> & subVec)
{
    int iBitsNum = iGetNumBits(iNum);
    for (int i = 1; i < iBitsNum; i++)
    {
        Sub MainSub = SplitNum(iNum, i);
        subVec.push_back(MainSub);

        CalcMainSub(MainSub, subVec);
    }
}

long long GetMinLaiZi(long long iNum)
{
    map<long long, long long>::iterator iterMap = handCardsMap.find(iNum);
    if (iterMap != handCardsMap.end())
    {
        return iterMap->second;
    }

    vector<Sub> subVec;
    GetNumSub(iNum, subVec);
    long long iMin = 999999;
    vector<Sub>::iterator iterSubVec = subVec.begin();
    while (iterSubVec != subVec.end())
    {
        Sub & sub = *iterSubVec;

        long long iLaiZi = GetMinLaiZi(sub.a) + GetMinLaiZi(sub.b);
        if (iLaiZi < iMin)
        {
            iMin = iLaiZi;
        }

        ++iterSubVec;
    }

    handCardsMap[iNum] = iMin;

    return iMin;
}

void vTest2()
{

}


int main()
{
    //100, 3, 2 --> result: 1002, 1020, 1200, 2100
    //vGetNext(1030, 4, 2);

    vInit();

    return 0;
}