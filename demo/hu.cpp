#include "hu.h"

CLaiZiHu::CLaiZiHu()
{
    vInit();
}

CLaiZiHu::~CLaiZiHu()
{

}

void CLaiZiHu::vInit()
{
    //<handcards, need LaiZi Num>

    handCardsMap[0] = 0;
    handCardsMap[1] = 2;
    handCardsMap[2] = 1;
    handCardsMap[3] = 0;
    handCardsMap[4] = 2;

    // 23 24 30 31 32 33 34 40 41 42 43 44
    handCardsMap[10] = 2;
    handCardsMap[11] = 1;
    handCardsMap[12] = 3;
    handCardsMap[13] = 2;
    handCardsMap[14] = 1;
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
    set1.insert(3);
    set1.insert(4);
    vec.push_back(set1);

    set<long long> set2;
    //10 11 12 13 14 20 21 22 23 24 30 31 32 33 34 40 41 42 43 44
    set2.insert(10);
    set2.insert(11);
    set2.insert(12);
    set2.insert(20);
    set2.insert(21);
    set2.insert(22);
    //vec.push_back(set2);

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
    //vec.push_back(set3);

    long long iStart = 2;
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

            for (long long i = 0; i < 5; i++)
            {
                vGetNext(iPreNum, iStart - 1, i, curset);
            }

            ++iterPre;
        }

        vec.push_back(curset);
        iStart++;
    }

    vShowVec(vec);

    //vector< set<long long> > vec;
}

void CLaiZiHu::vShowVec(vector< set<long long> > & vec)
{
    printf("### vShowVec\n");
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
    printf("### vShowVec End\n");
}


bool CLaiZiHu::bHu(Card aucHandCards[MAX_HANDCARD_NUM], int iHandCardsLen, Card ucLaiZi)
{
    /*计算出每张牌的张数。iCardsNum的下标代表每张牌，value就是这张牌的数量。*/
    int iCardsNum[MAX_CARD_ARRAY_SIZE] = { 0 };
    for (int i = 0; i < iHandCardsLen; i++)
    {
        iCardsNum[aucHandCards[i]]++;
    }

    /*计算赖子的数量，同时先把赖子移除。*/
    int iLaiZiNum = iCardsNum[ucLaiZi];
    iCardsNum[ucLaiZi] = 0;

    /*遍历所有将牌的情况。*/
    for (int i = 0; i < MAX_CARD_ARRAY_SIZE; i++)
    {
        Card ucJiang = 0;
        int iCurLaiZiNum = iLaiZiNum;
        int iCardsNumNoJiang[MAX_CARD_ARRAY_SIZE] = { 0 };
        memcpy(iCardsNumNoJiang, iCardsNum, MAX_CARD_ARRAY_SIZE * sizeof(int));

        if (iCardsNumNoJiang[i] == 1 && iCurLaiZiNum > 0)
        {
            ucJiang = i;
            iCardsNumNoJiang[i] = 0;
            iCurLaiZiNum = iCurLaiZiNum - 1;
        }
        else if (iCardsNumNoJiang[i] > 1)
        {
            ucJiang = i;
            iCardsNumNoJiang[i] -= 2;
        }
        else
        {
            continue;
        }

        /*判断扣除了将牌之后，剩余的牌能不能构成顺子(1万、2万、3万)或者刻子(1万、1万、1万)。*/
        if (bCanBePu(iCardsNumNoJiang, ucLaiZi, iCurLaiZiNum))
        {
            return true;
        }
    }

    return false;
}

void CLaiZiHu::vGetJiang(int iCardsNum[MAX_CARD_ARRAY_SIZE], Card aucJiang[MAX_HANDCARD_NUM], int & riJiangNum)
{
    memset(aucJiang, 0, MAX_HANDCARD_NUM);
    riJiangNum = 0;

    for (int i = 0; i < MAX_CARD_ARRAY_SIZE; i++)
    {
        if (iCardsNum[i] >= 2)
        {
            aucJiang[riJiangNum++] = i;
        }
    }
}


long long CLaiZiHu::ToNum(int iCardsNum[MAX_CARD_ARRAY_SIZE], int iStartIndex, int iEndIndex)
{
    long long num = 0;
    long long base = 1;
    for (int i = iStartIndex; i <= iEndIndex; i++)
    {
        num = num + iCardsNum[i] * base;
        base = base * 10;
    }
    return num;
}

bool CLaiZiHu::bCanBePu(int iCardsNum[MAX_CARD_ARRAY_SIZE], Card ucLaiZi, int iLaiZiNum)
{
    //判断万牌能不能构成顺子或者刻子
    long long llWan = ToNum(iCardsNum, CardYiWan, CardJiuWan);
    int iNeedLaiZiNum = GetMinLaiZi(llWan);
    if (iNeedLaiZiNum > iLaiZiNum)
    {
        return false;
    }
    iLaiZiNum -= iNeedLaiZiNum;

    //判断条牌能不能构成顺子或者刻子
    long long llTiao = ToNum(iCardsNum, CardYaoJi, CardJiuTiao);
    iNeedLaiZiNum = GetMinLaiZi(llTiao);
    if (iNeedLaiZiNum > iLaiZiNum)
    {
        return false;
    }
    iLaiZiNum -= iNeedLaiZiNum;

    //判断筒牌能不能构成顺子或者刻子
    long long llTong = ToNum(iCardsNum, CardYiTong, CardJiuTong);
    iNeedLaiZiNum = GetMinLaiZi(llTong);
    if (iNeedLaiZiNum > iLaiZiNum)
    {
        return false;
    }
    iLaiZiNum -= iNeedLaiZiNum;

    //判断风牌能不能构成刻子
    if (!bTongHuaSeCanBeKeZi(iCardsNum, CardDONG, CardBEI, iNeedLaiZiNum))
    {
        if (iNeedLaiZiNum > iLaiZiNum)
        {
            return false;
        }
        else
        {
            iLaiZiNum -= iNeedLaiZiNum;
        }
    }

    //判断箭牌能不能构成刻子
    if (!bTongHuaSeCanBeKeZi(iCardsNum, CardZHONG, CardBAI, iNeedLaiZiNum))
    {
        if (iNeedLaiZiNum > iLaiZiNum)
        {
            return false;
        }
        else
        {
            iLaiZiNum -= iNeedLaiZiNum;
        }
    }

    return true;
}

/*
@brief: 能否构成刻子。
*/
bool CLaiZiHu::bTongHuaSeCanBeKeZi(int iCardsNum[MAX_CARD_ARRAY_SIZE], int iStartIndex, int iEndIndex, int & iNeedLaiZiNum)
{
    iNeedLaiZiNum = 0;

    for (int i = iStartIndex; i <= iEndIndex; i++)
    {
        int iNum = iCardsNum[i];
        if (iNum == 1)
        {
            iNeedLaiZiNum += 2;
        }
        else if (iNum == 2)
        {
            iNeedLaiZiNum += 1;
        }
        else if (iNum == 4)
        {
            iNeedLaiZiNum += 2;
        }
    }

    return iNeedLaiZiNum > 0 ? false : true;
}


bool CLaiZiHu::bIsNumBitMoreThan14(long long iNum)
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

void CLaiZiHu::vGetNext(long long iNum, long long iNumBits, long long iInsertNum, set<long long> & curset)
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






int CLaiZiHu::iGetNumBits(long long iNum)
{
    int iBitsNum = 0;
    while (iNum > 0)
    {
        iBitsNum++;
        iNum /= 10;
    }
    return iBitsNum;
}

int CLaiZiHu::iGetNumRightBit(long long iNum)
{
    return iNum % 10;
}

int CLaiZiHu::iGetNumLeftBit(long long iNum)
{
    int iBitsNum = iGetNumBits(iNum);
    return iNum / (pow(10, iBitsNum - 1));
}

void CLaiZiHu::vReverseNumArray(long long arr[], int iArrLen)
{
    int iStart = 0;
    int iEnd = iArrLen - 1;
    while (iStart < iEnd)
    {
        long long tmp = arr[iStart];
        arr[iStart] = arr[iEnd];
        arr[iEnd] = tmp;

        ++iStart;
        --iEnd;
    }
}

Sub CLaiZiHu::SplitNum(long long iNum, int i)
{
    int iSplitLine = pow(10, i);
    Sub sub;
    sub.a = iNum / iSplitLine;
    sub.b = iNum % iSplitLine;

    long long numArr[14] = { 0 };
    int iNumArrLen = 0;
    long long sourceNum = iNum;
    while (sourceNum > 0)
    {
        numArr[iNumArrLen++] = sourceNum % 10;
        sourceNum /= 10;
    }

    int k = 0;
    while (k < i)
    {
        sub.brr[sub.brrLen++] = numArr[k++];
    }
    //vReverseNumArray(sub.brr, sub.brrLen);

    while (k < iNumArrLen)
    {
        sub.arr[sub.arrLen++] = numArr[k++];
    }

    return sub;
}

void CLaiZiHu::CalcMainSub(long long iNum, Sub MainSub, vector<Sub> & subVec)
{
    int iLeftNum = MainSub.arrLen;
    int iRightNum = MainSub.brrLen;

    if (iLeftNum == 1)
    {
        int l1 = MainSub.arr[0];
        for (int i = 0; i <= l1; i++)
        {
            Sub sub;
            sub.a = i;

            sub.b = MainSub.b + (l1 - i) * pow(10, MainSub.brrLen);

            if (sub.a > 0 && sub.a < iNum && sub.b > 0 && sub.b < iNum)
            {
                subVec.push_back(sub);
            }
        }
    }
    else if (iLeftNum >= 2)
    {
        int l1 = MainSub.arr[0];
        int l2 = MainSub.arr[1];
        for (int i = 0; i <= l1; i++)
        {
            for (int j = 0; j <= l2; j++)
            {
                Sub sub;
                sub.a = MainSub.a / 100 * 100 + j * 10 + i;
                sub.b = MainSub.b + (l1 - i) * pow(10, MainSub.brrLen) + (l2 - j) * pow(10, MainSub.brrLen + 1);

                if (sub.a > 0 && sub.a < iNum && sub.b > 0 && sub.b < iNum)
                {
                    subVec.push_back(sub);
                }
            }
        }
    }

    if (iRightNum == 1)
    {
        int r1 = MainSub.brr[0];
        for (int i = 0; i <= r1; i++)
        {
            Sub sub;
            sub.a = MainSub.a * 10 + (r1 - i);

            sub.b = i;

            if (sub.a > 0 && sub.a < iNum && sub.b > 0 && sub.b < iNum)
            {
                subVec.push_back(sub);
            }
        }
    }
    else if (iRightNum >= 2)
    {
        int r1 = MainSub.brr[MainSub.brrLen - 1];
        int r2 = MainSub.brr[MainSub.brrLen - 2];
        for (int i = 0; i <= r1; i++)
        {
            for (int j = 0; j <= r2; j++)
            {
                Sub sub;
                sub.a = MainSub.a * 100 + (r1 - i) * 10 + (r2 - j);

                sub.b = MainSub.b % ((long long)pow(10, MainSub.brrLen - 2)) + i * pow(10, MainSub.brrLen - 1) + j * pow(10, MainSub.brrLen - 2);

                if (sub.a > 0 && sub.a < iNum && sub.b > 0 && sub.b < iNum)
                {
                    subVec.push_back(sub);
                }
            }
        }
    }
}

void CLaiZiHu::GetNumSub(long long iNum, vector<Sub> & subVec)
{
    int iBitsNum = iGetNumBits(iNum);
    for (int i = 1; i < iBitsNum; i++)
    {
        Sub MainSub = SplitNum(iNum, i);
        if (MainSub.a > 0 && MainSub.a < iNum && MainSub.b > 0 && MainSub.b < iNum)
        {
            subVec.push_back(MainSub);
            CalcMainSub(iNum, MainSub, subVec);
        }
    }
}

bool CLaiZiHu::bSplitWithTwoBlank(long long iNum, vector<long long> & setNum)
{
    long long sourceNum = iNum;
    long long iBase = 10;
    long long iCurNum = sourceNum % 10;
    sourceNum /= 10;
    bool bPreIsZero = iCurNum == 0 ? true : false;
    while (sourceNum > 0)
    {
        long long iCurBit = sourceNum % 10;
        sourceNum /= 10;

        if (iCurBit == 0 && bPreIsZero)
        {
            if (iCurNum > 0 && iCurNum < iNum)
            {
                setNum.push_back(iCurNum);
            }
            iCurNum = 0;
            iBase = 1;
            bPreIsZero = (iCurBit == 0) ? true : false;
        }
        else
        {
            iCurNum = iCurNum + iCurBit * iBase;
            iBase *= 10;
            bPreIsZero = (iCurBit == 0) ? true : false;
        }
    }

    if (iCurNum > 0 && iCurNum < iNum)
    {
        setNum.push_back(iCurNum);
    }

    return setNum.size() > 0;
}

void CLaiZiHu::vTrimNum(long long & iNum)
{
    while (iNum % 10 == 0)
    {
        iNum /= 10;
    }
}

long long CLaiZiHu::GetMinLaiZi(long long iNum)
{
    long long sourceNum = iNum;

    if (iNum <= 0)
    {
        return 0;
    }

    vTrimNum(iNum);

    map<long long, long long>::iterator iterMap = handCardsMap.find(iNum);
    if (iterMap != handCardsMap.end())
    {
        return iterMap->second;
    }

    vector<long long> setNum;
    bool bRet = bSplitWithTwoBlank(iNum, setNum);
    if (bRet)
    {
        int iMin = 0;
        vector<long long>::iterator iter = setNum.begin();
        while (iter != setNum.end())
        {
            iMin += GetMinLaiZi(*iter);
            ++iter;
        }
        return iMin;
    }

    int iNumBits = iGetNumBits(iNum);
    if (iNumBits <= 3)
    {
        long long iNewNum = 0;
        long long base = 1;
        while (iNum > 0)
        {
            int iCurBit = iNum % 10;
            iNum /= 10;

            if (iCurBit > 2)
            {
                iCurBit -= 3;
            }

            iNewNum = iNewNum + iCurBit * base;
            base *= 10;
        }

        iNum = iNewNum;

        return GetMinLaiZi(iNum);
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
    handCardsMap[sourceNum] = iMin;

    return iMin;
}

void CLaiZiHu::vTest2()
{
    long long lRet = GetMinLaiZi(10121);
    printf("lRet = %lld\n", lRet);


    vector< set<long long> >::iterator iterVec = vec.begin();
    while (iterVec != vec.end())
    {
        set<long long> & curset = *iterVec;
        set<long long>::iterator iterSet = curset.begin();
        while (iterSet != curset.end())
        {
            printf("m(%lld) = %lld\n", *iterSet, GetMinLaiZi(*iterSet));

            ++iterSet;
        }
        printf("\n\n");
        ++iterVec;
    }
}
