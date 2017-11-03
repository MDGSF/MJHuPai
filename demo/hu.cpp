#include "hu.h"

struct Sub {
    long long a;
    long long b;

    long long arr[MAX_HANDCARD_NUM];
    int arrLen;

    long long brr[MAX_HANDCARD_NUM];
    int brrLen;

    Sub()
    {
        a = 0;
        memset(arr, 0, sizeof(arr));
        arrLen = 0;

        b = 0;
        memset(brr, 0, sizeof(brr));
        brrLen = 0;
    }
};

CLaiZiHu::CLaiZiHu()
{
    vInit();
}

CLaiZiHu::~CLaiZiHu()
{

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

void CLaiZiHu::vTrimNumZero(long long & iNum)
{
    while (iNum % 10 == 0)
    {
        iNum /= 10;
    }
}

long long CLaiZiHu::GetMinLaiZi(long long iNum)
{
    if (iNum <= 0)
    {
        return 0;
    }

    vector<long long> setNum;
    bool bRet = bSplitWithTwoBlank(iNum, setNum);
    if (bRet)
    {
        int iMin = 0;
        vector<long long>::iterator iter = setNum.begin();
        while (iter != setNum.end())
        {
            iMin += GetMinLaiZiNoTwoBlank(*iter);
            ++iter;
        }
        return iMin;
    }

    return GetMinLaiZiNoTwoBlank(iNum);
}

long long CLaiZiHu::GetMinLaiZiNoTwoBlank(long long iNum)
{
    long long llSourceNum = iNum;

    vTrimNumZero(iNum);

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

        long long iLaiZi = GetMinLaiZiNoTwoBlank(sub.a) + GetMinLaiZiNoTwoBlank(sub.b);
        if (iLaiZi < iMin)
        {
            iMin = iLaiZi;
        }

        ++iterSubVec;
    }

    handCardsMap[iNum] = iMin;
    handCardsMap[llSourceNum] = iMin;

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


void CLaiZiHu::vInit()
{
    handCardsMap[0] = 0;
    handCardsMap[1] = 2;
    handCardsMap[2] = 1;
    handCardsMap[3] = 0;
    handCardsMap[4] = 2;

    handCardsMap[10] = 2;
    handCardsMap[11] = 1;
    handCardsMap[12] = 3;
    handCardsMap[13] = 2;
    handCardsMap[14] = 1;
    handCardsMap[20] = 1;
    handCardsMap[21] = 3;
    handCardsMap[22] = 2;
    handCardsMap[23] = 1;
    handCardsMap[24] = 3;
    handCardsMap[30] = 0;
    handCardsMap[31] = 2;
    handCardsMap[32] = 1;
    handCardsMap[33] = 0;
    handCardsMap[34] = 2;
    handCardsMap[40] = 2;
    handCardsMap[41] = 1;
    handCardsMap[42] = 3;
    handCardsMap[43] = 2;
    handCardsMap[44] = 1;

    handCardsMap[100] = 2;
    handCardsMap[101] = 1;
    handCardsMap[102] = 3;
    handCardsMap[103] = 2;
    handCardsMap[104] = 1;
    handCardsMap[110] = 1;
    handCardsMap[111] = 0;
    handCardsMap[112] = 2;
    handCardsMap[113] = 1;
    handCardsMap[114] = 0;
    handCardsMap[120] = 3;
    handCardsMap[121] = 2;
    handCardsMap[122] = 1;
    handCardsMap[123] = 3;
    handCardsMap[124] = 2;
    handCardsMap[130] = 2;
    handCardsMap[131] = 1;
    handCardsMap[132] = 3;
    handCardsMap[133] = 2;
    handCardsMap[134] = 1;
    handCardsMap[140] = 1;
    handCardsMap[141] = 0;
    handCardsMap[142] = 2;
    handCardsMap[143] = 1;
    handCardsMap[144] = 0;
    handCardsMap[200] = 1;
    handCardsMap[201] = 3;
    handCardsMap[202] = 2;
    handCardsMap[203] = 1;
    handCardsMap[204] = 3;
    handCardsMap[210] = 3;
    handCardsMap[211] = 2;
    handCardsMap[212] = 1;
    handCardsMap[213] = 3;
    handCardsMap[214] = 2;
    handCardsMap[220] = 2;
    handCardsMap[221] = 1;
    handCardsMap[222] = 0;
    handCardsMap[223] = 2;
    handCardsMap[224] = 1;
    handCardsMap[230] = 1;
    handCardsMap[231] = 3;
    handCardsMap[232] = 2;
    handCardsMap[233] = 1;
    handCardsMap[234] = 3;
    handCardsMap[240] = 3;
    handCardsMap[241] = 2;
    handCardsMap[242] = 1;
    handCardsMap[243] = 3;
    handCardsMap[244] = 2;
    handCardsMap[300] = 0;
    handCardsMap[301] = 2;
    handCardsMap[302] = 1;
    handCardsMap[303] = 0;
    handCardsMap[304] = 2;
    handCardsMap[310] = 2;
    handCardsMap[311] = 1;
    handCardsMap[312] = 3;
    handCardsMap[313] = 2;
    handCardsMap[314] = 1;
    handCardsMap[320] = 1;
    handCardsMap[321] = 3;
    handCardsMap[322] = 2;
    handCardsMap[323] = 1;
    handCardsMap[324] = 3;
    handCardsMap[330] = 0;
    handCardsMap[331] = 2;
    handCardsMap[332] = 1;
    handCardsMap[333] = 0;
    handCardsMap[334] = 2;
    handCardsMap[340] = 2;
    handCardsMap[341] = 1;
    handCardsMap[342] = 3;
    handCardsMap[343] = 2;
    handCardsMap[344] = 1;
    handCardsMap[400] = 2;
    handCardsMap[401] = 1;
    handCardsMap[402] = 3;
    handCardsMap[403] = 2;
    handCardsMap[404] = 1;
    handCardsMap[410] = 1;
    handCardsMap[411] = 0;
    handCardsMap[412] = 2;
    handCardsMap[413] = 1;
    handCardsMap[414] = 0;
    handCardsMap[420] = 3;
    handCardsMap[421] = 2;
    handCardsMap[422] = 1;
    handCardsMap[423] = 3;
    handCardsMap[424] = 2;
    handCardsMap[430] = 2;
    handCardsMap[431] = 1;
    handCardsMap[432] = 3;
    handCardsMap[433] = 2;
    handCardsMap[434] = 1;
    handCardsMap[440] = 1;
    handCardsMap[441] = 0;
    handCardsMap[442] = 2;
    handCardsMap[443] = 1;
    handCardsMap[444] = 0;


    handCardsMap[1001] = 4;
    handCardsMap[1002] = 3;
    handCardsMap[1003] = 2;
    handCardsMap[1004] = 4;
    handCardsMap[1010] = 1;
    handCardsMap[1011] = 3;
    handCardsMap[1012] = 2;
    handCardsMap[1013] = 1;
    handCardsMap[1014] = 3;
    handCardsMap[1020] = 3;
    handCardsMap[1021] = 2;
    handCardsMap[1022] = 4;
    handCardsMap[1023] = 3;
    handCardsMap[1024] = 2;
    handCardsMap[1030] = 2;
    handCardsMap[1031] = 4;
    handCardsMap[1032] = 3;
    handCardsMap[1033] = 2;
    handCardsMap[1034] = 4;
    handCardsMap[1040] = 1;
    handCardsMap[1041] = 3;
    handCardsMap[1042] = 2;
    handCardsMap[1043] = 1;
    handCardsMap[1044] = 3;
    handCardsMap[1100] = 1;
    handCardsMap[1101] = 3;
    handCardsMap[1102] = 2;
    handCardsMap[1103] = 1;
    handCardsMap[1104] = 3;
    handCardsMap[1110] = 0;
    handCardsMap[1111] = 2;
    handCardsMap[1112] = 1;
    handCardsMap[1113] = 0;
    handCardsMap[1114] = 2;
    handCardsMap[1120] = 2;
    handCardsMap[1121] = 1;
    handCardsMap[1122] = 3;
    handCardsMap[1123] = 2;
    handCardsMap[1124] = 1;
    handCardsMap[1130] = 1;
    handCardsMap[1131] = 3;
    handCardsMap[1132] = 2;
    handCardsMap[1133] = 1;
    handCardsMap[1134] = 3;
    handCardsMap[1140] = 0;
    handCardsMap[1141] = 2;
    handCardsMap[1142] = 1;
    handCardsMap[1143] = 0;
    handCardsMap[1144] = 2;
    handCardsMap[1200] = 3;
    handCardsMap[1201] = 2;
    handCardsMap[1202] = 4;
    handCardsMap[1203] = 3;
    handCardsMap[1204] = 2;
    handCardsMap[1210] = 2;
    handCardsMap[1211] = 1;
    handCardsMap[1212] = 3;
    handCardsMap[1213] = 2;
    handCardsMap[1214] = 1;
    handCardsMap[1220] = 1;
    handCardsMap[1221] = 0;
    handCardsMap[1222] = 2;
    handCardsMap[1223] = 1;
    handCardsMap[1224] = 0;
    handCardsMap[1230] = 3;
    handCardsMap[1231] = 2;
    handCardsMap[1232] = 1;
    handCardsMap[1233] = 3;
    handCardsMap[1234] = 2;
    handCardsMap[1240] = 2;
    handCardsMap[1241] = 1;
    handCardsMap[1242] = 3;
    handCardsMap[1243] = 2;
    handCardsMap[1244] = 1;
    handCardsMap[1300] = 2;
    handCardsMap[1301] = 4;
    handCardsMap[1302] = 3;
    handCardsMap[1303] = 2;
    handCardsMap[1304] = 4;
    handCardsMap[1310] = 1;
    handCardsMap[1311] = 3;
    handCardsMap[1312] = 2;
    handCardsMap[1313] = 1;
    handCardsMap[1314] = 3;
    handCardsMap[1320] = 3;
    handCardsMap[1321] = 2;
    handCardsMap[1322] = 1;
    handCardsMap[1323] = 3;
    handCardsMap[1324] = 2;
    handCardsMap[1330] = 2;
    handCardsMap[1331] = 1;
    handCardsMap[1332] = 0;
    handCardsMap[1333] = 2;
    handCardsMap[1334] = 1;
    handCardsMap[1340] = 1;
    handCardsMap[1341] = 3;
    handCardsMap[1342] = 2;
    handCardsMap[1343] = 1;
    handCardsMap[1344] = 3;
    handCardsMap[1400] = 1;
    handCardsMap[1401] = 3;
    handCardsMap[1402] = 2;
    handCardsMap[1403] = 1;
    handCardsMap[1404] = 3;
    handCardsMap[1410] = 0;
    handCardsMap[1411] = 2;
    handCardsMap[1412] = 1;
    handCardsMap[1413] = 0;
    handCardsMap[1414] = 2;
    handCardsMap[1420] = 2;
    handCardsMap[1421] = 1;
    handCardsMap[1422] = 3;
    handCardsMap[1423] = 2;
    handCardsMap[1424] = 1;
    handCardsMap[1430] = 1;
    handCardsMap[1431] = 3;
    handCardsMap[1432] = 2;
    handCardsMap[1433] = 1;
    handCardsMap[1434] = 3;
    handCardsMap[1440] = 0;
    handCardsMap[1441] = 2;
    handCardsMap[1442] = 1;
    handCardsMap[1443] = 0;
    handCardsMap[1444] = 2;
    handCardsMap[2000] = 1;
    handCardsMap[2001] = 3;
    handCardsMap[2002] = 2;
    handCardsMap[2003] = 1;
    handCardsMap[2004] = 3;
    handCardsMap[2010] = 3;
    handCardsMap[2011] = 2;
    handCardsMap[2012] = 4;
    handCardsMap[2013] = 3;
    handCardsMap[2014] = 2;
    handCardsMap[2020] = 2;
    handCardsMap[2021] = 4;
    handCardsMap[2022] = 3;
    handCardsMap[2023] = 2;
    handCardsMap[2024] = 4;
    handCardsMap[2030] = 1;
    handCardsMap[2031] = 3;
    handCardsMap[2032] = 2;
    handCardsMap[2033] = 1;
    handCardsMap[2034] = 3;
    handCardsMap[2040] = 3;
    handCardsMap[2041] = 2;
    handCardsMap[2042] = 4;
    handCardsMap[2043] = 3;
    handCardsMap[2044] = 2;
    handCardsMap[2100] = 3;
    handCardsMap[2101] = 2;
    handCardsMap[2102] = 4;
    handCardsMap[2103] = 3;
    handCardsMap[2104] = 2;
    handCardsMap[2110] = 2;
    handCardsMap[2111] = 1;
    handCardsMap[2112] = 3;
    handCardsMap[2113] = 2;
    handCardsMap[2114] = 1;
    handCardsMap[2120] = 1;
    handCardsMap[2121] = 3;
    handCardsMap[2122] = 2;
    handCardsMap[2123] = 1;
    handCardsMap[2124] = 3;
    handCardsMap[2130] = 3;
    handCardsMap[2131] = 2;
    handCardsMap[2132] = 4;
    handCardsMap[2133] = 3;
    handCardsMap[2134] = 2;
    handCardsMap[2140] = 2;
    handCardsMap[2141] = 1;
    handCardsMap[2142] = 3;
    handCardsMap[2143] = 2;
    handCardsMap[2144] = 1;
    handCardsMap[2200] = 2;
    handCardsMap[2201] = 4;
    handCardsMap[2202] = 3;
    handCardsMap[2203] = 2;
    handCardsMap[2204] = 4;
    handCardsMap[2210] = 1;
    handCardsMap[2211] = 3;
    handCardsMap[2212] = 2;
    handCardsMap[2213] = 1;
    handCardsMap[2214] = 3;
    handCardsMap[2220] = 0;
    handCardsMap[2221] = 2;
    handCardsMap[2222] = 1;
    handCardsMap[2223] = 0;
    handCardsMap[2224] = 2;
    handCardsMap[2230] = 2;
    handCardsMap[2231] = 1;
    handCardsMap[2232] = 3;
    handCardsMap[2233] = 2;
    handCardsMap[2234] = 1;
    handCardsMap[2240] = 1;
    handCardsMap[2241] = 3;
    handCardsMap[2242] = 2;
    handCardsMap[2243] = 1;
    handCardsMap[2244] = 3;
    handCardsMap[2300] = 1;
    handCardsMap[2301] = 3;
    handCardsMap[2302] = 2;
    handCardsMap[2303] = 1;
    handCardsMap[2304] = 3;
    handCardsMap[2310] = 3;
    handCardsMap[2311] = 2;
    handCardsMap[2312] = 4;
    handCardsMap[2313] = 3;
    handCardsMap[2314] = 2;
    handCardsMap[2320] = 2;
    handCardsMap[2321] = 1;
    handCardsMap[2322] = 3;
    handCardsMap[2323] = 2;
    handCardsMap[2324] = 1;
    handCardsMap[2330] = 1;
    handCardsMap[2331] = 0;
    handCardsMap[2332] = 2;
    handCardsMap[2333] = 1;
    handCardsMap[2334] = 0;
    handCardsMap[2340] = 3;
    handCardsMap[2341] = 2;
    handCardsMap[2342] = 1;
    handCardsMap[2343] = 3;
    handCardsMap[2344] = 2;
    handCardsMap[2400] = 3;
    handCardsMap[2401] = 2;
    handCardsMap[2402] = 4;
    handCardsMap[2403] = 3;
    handCardsMap[2404] = 2;
    handCardsMap[2410] = 2;
    handCardsMap[2411] = 1;
    handCardsMap[2412] = 3;
    handCardsMap[2413] = 2;
    handCardsMap[2414] = 1;
    handCardsMap[2420] = 1;
    handCardsMap[2421] = 3;
    handCardsMap[2422] = 2;
    handCardsMap[2423] = 1;
    handCardsMap[2424] = 3;
    handCardsMap[2430] = 3;
    handCardsMap[2431] = 2;
    handCardsMap[2432] = 1;
    handCardsMap[2433] = 3;
    handCardsMap[2434] = 2;
    handCardsMap[2440] = 2;
    handCardsMap[2441] = 1;
    handCardsMap[2442] = 0;
    handCardsMap[2443] = 2;
    handCardsMap[2444] = 1;
    handCardsMap[3000] = 0;
    handCardsMap[3001] = 2;
    handCardsMap[3002] = 1;
    handCardsMap[3003] = 0;
    handCardsMap[3004] = 2;
    handCardsMap[3010] = 2;
    handCardsMap[3011] = 1;
    handCardsMap[3012] = 3;
    handCardsMap[3013] = 2;
    handCardsMap[3014] = 1;
    handCardsMap[3020] = 1;
    handCardsMap[3021] = 3;
    handCardsMap[3022] = 2;
    handCardsMap[3023] = 1;
    handCardsMap[3024] = 3;
    handCardsMap[3030] = 0;
    handCardsMap[3031] = 2;
    handCardsMap[3032] = 1;
    handCardsMap[3033] = 0;
    handCardsMap[3034] = 2;
    handCardsMap[3040] = 2;
    handCardsMap[3041] = 1;
    handCardsMap[3042] = 3;
    handCardsMap[3043] = 2;
    handCardsMap[3044] = 1;
    handCardsMap[3100] = 2;
    handCardsMap[3101] = 1;
    handCardsMap[3102] = 3;
    handCardsMap[3103] = 2;
    handCardsMap[3104] = 1;
    handCardsMap[3110] = 1;
    handCardsMap[3111] = 0;
    handCardsMap[3112] = 2;
    handCardsMap[3113] = 1;
    handCardsMap[3114] = 0;
    handCardsMap[3120] = 3;
    handCardsMap[3121] = 2;
    handCardsMap[3122] = 1;
    handCardsMap[3123] = 3;
    handCardsMap[3124] = 2;
    handCardsMap[3130] = 2;
    handCardsMap[3131] = 1;
    handCardsMap[3132] = 3;
    handCardsMap[3133] = 2;
    handCardsMap[3134] = 1;
    handCardsMap[3140] = 1;
    handCardsMap[3141] = 0;
    handCardsMap[3142] = 2;
    handCardsMap[3143] = 1;
    handCardsMap[3144] = 0;
    handCardsMap[3200] = 1;
    handCardsMap[3201] = 3;
    handCardsMap[3202] = 2;
    handCardsMap[3203] = 1;
    handCardsMap[3204] = 3;
    handCardsMap[3210] = 3;
    handCardsMap[3211] = 2;
    handCardsMap[3212] = 1;
    handCardsMap[3213] = 3;
    handCardsMap[3214] = 2;
    handCardsMap[3220] = 2;
    handCardsMap[3221] = 1;
    handCardsMap[3222] = 0;
    handCardsMap[3223] = 2;
    handCardsMap[3224] = 1;
    handCardsMap[3230] = 1;
    handCardsMap[3231] = 3;
    handCardsMap[3232] = 2;
    handCardsMap[3233] = 1;
    handCardsMap[3234] = 3;
    handCardsMap[3240] = 3;
    handCardsMap[3241] = 2;
    handCardsMap[3242] = 1;
    handCardsMap[3243] = 3;
    handCardsMap[3244] = 2;
    handCardsMap[3300] = 0;
    handCardsMap[3301] = 2;
    handCardsMap[3302] = 1;
    handCardsMap[3303] = 0;
    handCardsMap[3304] = 2;
    handCardsMap[3310] = 2;
    handCardsMap[3311] = 1;
    handCardsMap[3312] = 3;
    handCardsMap[3313] = 2;
    handCardsMap[3314] = 1;
    handCardsMap[3320] = 1;
    handCardsMap[3321] = 3;
    handCardsMap[3322] = 2;
    handCardsMap[3323] = 1;
    handCardsMap[3324] = 3;
    handCardsMap[3330] = 0;
    handCardsMap[3331] = 2;
    handCardsMap[3332] = 1;
    handCardsMap[3333] = 0;
    handCardsMap[3334] = 2;
    handCardsMap[3340] = 2;
    handCardsMap[3341] = 1;
    handCardsMap[3342] = 3;
    handCardsMap[3343] = 2;
    handCardsMap[3344] = 1;
    handCardsMap[3400] = 2;
    handCardsMap[3401] = 1;
    handCardsMap[3402] = 3;
    handCardsMap[3403] = 2;
    handCardsMap[3404] = 1;
    handCardsMap[3410] = 1;
    handCardsMap[3411] = 0;
    handCardsMap[3412] = 2;
    handCardsMap[3413] = 1;
    handCardsMap[3414] = 0;
    handCardsMap[3420] = 3;
    handCardsMap[3421] = 2;
    handCardsMap[3422] = 1;
    handCardsMap[3423] = 3;
    handCardsMap[3424] = 2;
    handCardsMap[3430] = 2;
    handCardsMap[3431] = 1;
    handCardsMap[3432] = 3;
    handCardsMap[3433] = 2;
    handCardsMap[3434] = 1;
    handCardsMap[3440] = 1;
    handCardsMap[3441] = 0;
    handCardsMap[3442] = 2;
    handCardsMap[3443] = 1;
    handCardsMap[4000] = 2;
    handCardsMap[4001] = 4;
    handCardsMap[4002] = 3;
    handCardsMap[4003] = 2;
    handCardsMap[4004] = 4;
    handCardsMap[4010] = 1;
    handCardsMap[4011] = 3;
    handCardsMap[4012] = 2;
    handCardsMap[4013] = 1;
    handCardsMap[4014] = 3;
    handCardsMap[4020] = 3;
    handCardsMap[4021] = 2;
    handCardsMap[4022] = 4;
    handCardsMap[4023] = 3;
    handCardsMap[4024] = 2;
    handCardsMap[4030] = 2;
    handCardsMap[4031] = 4;
    handCardsMap[4032] = 3;
    handCardsMap[4033] = 2;
    handCardsMap[4034] = 4;
    handCardsMap[4040] = 1;
    handCardsMap[4041] = 3;
    handCardsMap[4042] = 2;
    handCardsMap[4043] = 1;
    handCardsMap[4044] = 3;
    handCardsMap[4100] = 1;
    handCardsMap[4101] = 3;
    handCardsMap[4102] = 2;
    handCardsMap[4103] = 1;
    handCardsMap[4104] = 3;
    handCardsMap[4110] = 0;
    handCardsMap[4111] = 2;
    handCardsMap[4112] = 1;
    handCardsMap[4113] = 0;
    handCardsMap[4114] = 2;
    handCardsMap[4120] = 2;
    handCardsMap[4121] = 1;
    handCardsMap[4122] = 3;
    handCardsMap[4123] = 2;
    handCardsMap[4124] = 1;
    handCardsMap[4130] = 1;
    handCardsMap[4131] = 3;
    handCardsMap[4132] = 2;
    handCardsMap[4133] = 1;
    handCardsMap[4134] = 3;
    handCardsMap[4140] = 0;
    handCardsMap[4141] = 2;
    handCardsMap[4142] = 1;
    handCardsMap[4143] = 0;
    handCardsMap[4144] = 2;
    handCardsMap[4200] = 3;
    handCardsMap[4201] = 2;
    handCardsMap[4202] = 4;
    handCardsMap[4203] = 3;
    handCardsMap[4204] = 2;
    handCardsMap[4210] = 2;
    handCardsMap[4211] = 1;
    handCardsMap[4212] = 3;
    handCardsMap[4213] = 2;
    handCardsMap[4214] = 1;
    handCardsMap[4220] = 1;
    handCardsMap[4221] = 0;
    handCardsMap[4222] = 2;
    handCardsMap[4223] = 1;
    handCardsMap[4224] = 0;
    handCardsMap[4230] = 3;
    handCardsMap[4231] = 2;
    handCardsMap[4232] = 1;
    handCardsMap[4233] = 3;
    handCardsMap[4234] = 2;
    handCardsMap[4240] = 2;
    handCardsMap[4241] = 1;
    handCardsMap[4242] = 3;
    handCardsMap[4243] = 2;
    handCardsMap[4244] = 1;
    handCardsMap[4300] = 2;
    handCardsMap[4301] = 4;
    handCardsMap[4302] = 3;
    handCardsMap[4303] = 2;
    handCardsMap[4304] = 4;
    handCardsMap[4310] = 1;
    handCardsMap[4311] = 3;
    handCardsMap[4312] = 2;
    handCardsMap[4313] = 1;
    handCardsMap[4314] = 3;
    handCardsMap[4320] = 3;
    handCardsMap[4321] = 2;
    handCardsMap[4322] = 1;
    handCardsMap[4323] = 3;
    handCardsMap[4324] = 2;
    handCardsMap[4330] = 2;
    handCardsMap[4331] = 1;
    handCardsMap[4332] = 0;
    handCardsMap[4333] = 2;
    handCardsMap[4334] = 1;
    handCardsMap[4340] = 1;
    handCardsMap[4341] = 3;
    handCardsMap[4342] = 2;
    handCardsMap[4343] = 1;
    handCardsMap[4400] = 1;
    handCardsMap[4401] = 3;
    handCardsMap[4402] = 2;
    handCardsMap[4403] = 1;
    handCardsMap[4404] = 3;
    handCardsMap[4410] = 0;
    handCardsMap[4411] = 2;
    handCardsMap[4412] = 1;
    handCardsMap[4413] = 0;
    handCardsMap[4414] = 2;
    handCardsMap[4420] = 2;
    handCardsMap[4421] = 1;
    handCardsMap[4422] = 3;
    handCardsMap[4423] = 2;
    handCardsMap[4424] = 1;
    handCardsMap[4430] = 1;
    handCardsMap[4431] = 3;
    handCardsMap[4432] = 2;
    handCardsMap[4433] = 1;
    handCardsMap[4440] = 0;
    handCardsMap[4441] = 2;
    handCardsMap[4442] = 1;


    handCardsMap[10001] = 4;
    handCardsMap[10002] = 3;
    handCardsMap[10003] = 2;
    handCardsMap[10004] = 4;
    handCardsMap[10010] = 4;
    handCardsMap[10011] = 3;
    handCardsMap[10012] = 5;
    handCardsMap[10013] = 4;
    handCardsMap[10014] = 3;
    handCardsMap[10020] = 3;
    handCardsMap[10021] = 5;
    handCardsMap[10022] = 4;
    handCardsMap[10023] = 3;
    handCardsMap[10024] = 5;
    handCardsMap[10030] = 2;
    handCardsMap[10031] = 4;
    handCardsMap[10032] = 3;
    handCardsMap[10033] = 2;
    handCardsMap[10034] = 4;
    handCardsMap[10040] = 4;
    handCardsMap[10041] = 3;
    handCardsMap[10042] = 5;
    handCardsMap[10043] = 4;
    handCardsMap[10044] = 3;
    handCardsMap[10100] = 1;
    handCardsMap[10101] = 3;
    handCardsMap[10102] = 2;
    handCardsMap[10103] = 1;
    handCardsMap[10104] = 3;
    handCardsMap[10110] = 3;
    handCardsMap[10111] = 2;
    handCardsMap[10112] = 4;
    handCardsMap[10113] = 3;
    handCardsMap[10114] = 2;
    handCardsMap[10120] = 2;
    handCardsMap[10121] = 4;
    handCardsMap[10122] = 3;
    handCardsMap[10123] = 2;
    handCardsMap[10124] = 4;
    handCardsMap[10130] = 1;
    handCardsMap[10131] = 3;
    handCardsMap[10132] = 2;
    handCardsMap[10133] = 1;
    handCardsMap[10134] = 3;
    handCardsMap[10140] = 3;
    handCardsMap[10141] = 2;
    handCardsMap[10142] = 4;
    handCardsMap[10143] = 3;
    handCardsMap[10144] = 2;
    handCardsMap[10200] = 3;
    handCardsMap[10201] = 2;
    handCardsMap[10202] = 4;
    handCardsMap[10203] = 3;
    handCardsMap[10204] = 2;
    handCardsMap[10210] = 2;
    handCardsMap[10211] = 1;
    handCardsMap[10212] = 3;
    handCardsMap[10213] = 2;
    handCardsMap[10214] = 1;
    handCardsMap[10220] = 4;
    handCardsMap[10221] = 3;
    handCardsMap[10222] = 2;
    handCardsMap[10223] = 4;
    handCardsMap[10224] = 3;
    handCardsMap[10230] = 3;
    handCardsMap[10231] = 2;
    handCardsMap[10232] = 4;
    handCardsMap[10233] = 3;
    handCardsMap[10234] = 2;
    handCardsMap[10240] = 2;
    handCardsMap[10241] = 1;
    handCardsMap[10242] = 3;
    handCardsMap[10243] = 2;
    handCardsMap[10244] = 1;
    handCardsMap[10300] = 2;
    handCardsMap[10301] = 4;
    handCardsMap[10302] = 3;
    handCardsMap[10303] = 2;
    handCardsMap[10304] = 4;
    handCardsMap[10310] = 4;
    handCardsMap[10311] = 3;
    handCardsMap[10312] = 2;
    handCardsMap[10313] = 4;
    handCardsMap[10314] = 3;
    handCardsMap[10320] = 3;
    handCardsMap[10321] = 2;
    handCardsMap[10322] = 1;
    handCardsMap[10323] = 3;
    handCardsMap[10324] = 2;
    handCardsMap[10330] = 2;
    handCardsMap[10331] = 4;
    handCardsMap[10332] = 3;
    handCardsMap[10333] = 2;
    handCardsMap[10334] = 4;
    handCardsMap[10340] = 4;
    handCardsMap[10341] = 3;
    handCardsMap[10342] = 2;
    handCardsMap[10343] = 4;
    handCardsMap[10344] = 3;
    handCardsMap[10400] = 1;
    handCardsMap[10401] = 3;
    handCardsMap[10402] = 2;
    handCardsMap[10403] = 1;
    handCardsMap[10404] = 3;
    handCardsMap[10410] = 3;
    handCardsMap[10411] = 2;
    handCardsMap[10412] = 4;
    handCardsMap[10413] = 3;
    handCardsMap[10414] = 2;
    handCardsMap[10420] = 2;
    handCardsMap[10421] = 4;
    handCardsMap[10422] = 3;
    handCardsMap[10423] = 2;
    handCardsMap[10424] = 4;
    handCardsMap[10430] = 1;
    handCardsMap[10431] = 3;
    handCardsMap[10432] = 2;
    handCardsMap[10433] = 1;
    handCardsMap[10434] = 3;
    handCardsMap[10440] = 3;
    handCardsMap[10441] = 2;
    handCardsMap[10442] = 4;
    handCardsMap[10443] = 3;
    handCardsMap[10444] = 2;
    handCardsMap[11000] = 1;
    handCardsMap[11001] = 3;
    handCardsMap[11002] = 2;
    handCardsMap[11003] = 1;
    handCardsMap[11004] = 3;
    handCardsMap[11010] = 3;
    handCardsMap[11011] = 2;
    handCardsMap[11012] = 4;
    handCardsMap[11013] = 3;
    handCardsMap[11014] = 2;
    handCardsMap[11020] = 2;
    handCardsMap[11021] = 4;
    handCardsMap[11022] = 3;
    handCardsMap[11023] = 2;
    handCardsMap[11024] = 4;
    handCardsMap[11030] = 1;
    handCardsMap[11031] = 3;
    handCardsMap[11032] = 2;
    handCardsMap[11033] = 1;
    handCardsMap[11034] = 3;
    handCardsMap[11040] = 3;
    handCardsMap[11041] = 2;
    handCardsMap[11042] = 4;
    handCardsMap[11043] = 3;
    handCardsMap[11044] = 2;
    handCardsMap[11100] = 0;
    handCardsMap[11101] = 2;
    handCardsMap[11102] = 1;
    handCardsMap[11103] = 0;
    handCardsMap[11104] = 2;
    handCardsMap[11110] = 2;
    handCardsMap[11111] = 1;
    handCardsMap[11112] = 3;
    handCardsMap[11113] = 2;
    handCardsMap[11114] = 1;
    handCardsMap[11120] = 1;
    handCardsMap[11121] = 3;
    handCardsMap[11122] = 2;
    handCardsMap[11123] = 1;
    handCardsMap[11124] = 3;
    handCardsMap[11130] = 0;
    handCardsMap[11131] = 2;
    handCardsMap[11132] = 1;
    handCardsMap[11133] = 0;
    handCardsMap[11134] = 2;
    handCardsMap[11140] = 2;
    handCardsMap[11141] = 1;
    handCardsMap[11142] = 3;
    handCardsMap[11143] = 2;
    handCardsMap[11144] = 1;
    handCardsMap[11200] = 2;
    handCardsMap[11201] = 1;
    handCardsMap[11202] = 3;
    handCardsMap[11203] = 2;
    handCardsMap[11204] = 1;
    handCardsMap[11210] = 1;
    handCardsMap[11211] = 0;
    handCardsMap[11212] = 2;
    handCardsMap[11213] = 1;
    handCardsMap[11214] = 0;
    handCardsMap[11220] = 3;
    handCardsMap[11221] = 2;
    handCardsMap[11222] = 1;
    handCardsMap[11223] = 3;
    handCardsMap[11224] = 2;
    handCardsMap[11230] = 2;
    handCardsMap[11231] = 1;
    handCardsMap[11232] = 3;
    handCardsMap[11233] = 2;
    handCardsMap[11234] = 1;
    handCardsMap[11240] = 1;
    handCardsMap[11241] = 0;
    handCardsMap[11242] = 2;
    handCardsMap[11243] = 1;
    handCardsMap[11244] = 0;
    handCardsMap[11300] = 1;
    handCardsMap[11301] = 3;
    handCardsMap[11302] = 2;
    handCardsMap[11303] = 1;
    handCardsMap[11304] = 3;
    handCardsMap[11310] = 3;
    handCardsMap[11311] = 2;
    handCardsMap[11312] = 1;
    handCardsMap[11313] = 3;
    handCardsMap[11314] = 2;
    handCardsMap[11320] = 2;
    handCardsMap[11321] = 1;
    handCardsMap[11322] = 0;
    handCardsMap[11323] = 2;
    handCardsMap[11324] = 1;
    handCardsMap[11330] = 1;
    handCardsMap[11331] = 3;
    handCardsMap[11332] = 2;
    handCardsMap[11333] = 1;
    handCardsMap[11334] = 3;
    handCardsMap[11340] = 3;
    handCardsMap[11341] = 2;
    handCardsMap[11342] = 1;
    handCardsMap[11343] = 3;
    handCardsMap[11344] = 2;
    handCardsMap[11400] = 0;
    handCardsMap[11401] = 2;
    handCardsMap[11402] = 1;
    handCardsMap[11403] = 0;
    handCardsMap[11404] = 2;
    handCardsMap[11410] = 2;
    handCardsMap[11411] = 1;
    handCardsMap[11412] = 3;
    handCardsMap[11413] = 2;
    handCardsMap[11414] = 1;
    handCardsMap[11420] = 1;
    handCardsMap[11421] = 3;
    handCardsMap[11422] = 2;
    handCardsMap[11423] = 1;
    handCardsMap[11424] = 3;
    handCardsMap[11430] = 0;
    handCardsMap[11431] = 2;
    handCardsMap[11432] = 1;
    handCardsMap[11433] = 0;
    handCardsMap[11434] = 2;
    handCardsMap[11440] = 2;
    handCardsMap[11441] = 1;
    handCardsMap[11442] = 3;
    handCardsMap[11443] = 2;
    handCardsMap[11444] = 1;
    handCardsMap[12000] = 3;
    handCardsMap[12001] = 5;
    handCardsMap[12002] = 4;
    handCardsMap[12003] = 3;
    handCardsMap[12004] = 5;
    handCardsMap[12010] = 2;
    handCardsMap[12011] = 4;
    handCardsMap[12012] = 3;
    handCardsMap[12013] = 2;
    handCardsMap[12014] = 4;
    handCardsMap[12020] = 4;
    handCardsMap[12021] = 3;
    handCardsMap[12022] = 5;
    handCardsMap[12023] = 4;
    handCardsMap[12024] = 3;
    handCardsMap[12030] = 3;
    handCardsMap[12031] = 5;
    handCardsMap[12032] = 4;
    handCardsMap[12033] = 3;
    handCardsMap[12034] = 5;
    handCardsMap[12040] = 2;
    handCardsMap[12041] = 4;
    handCardsMap[12042] = 3;
    handCardsMap[12043] = 2;
    handCardsMap[12044] = 4;
    handCardsMap[12100] = 2;
    handCardsMap[12101] = 4;
    handCardsMap[12102] = 3;
    handCardsMap[12103] = 2;
    handCardsMap[12104] = 4;
    handCardsMap[12110] = 1;
    handCardsMap[12111] = 3;
    handCardsMap[12112] = 2;
    handCardsMap[12113] = 1;
    handCardsMap[12114] = 3;
    handCardsMap[12120] = 3;
    handCardsMap[12121] = 2;
    handCardsMap[12122] = 4;
    handCardsMap[12123] = 3;
    handCardsMap[12124] = 2;
    handCardsMap[12130] = 2;
    handCardsMap[12131] = 4;
    handCardsMap[12132] = 3;
    handCardsMap[12133] = 2;
    handCardsMap[12134] = 4;
    handCardsMap[12140] = 1;
    handCardsMap[12141] = 3;
    handCardsMap[12142] = 2;
    handCardsMap[12143] = 1;
    handCardsMap[12144] = 3;
    handCardsMap[12200] = 1;
    handCardsMap[12201] = 3;
    handCardsMap[12202] = 2;
    handCardsMap[12203] = 1;
    handCardsMap[12204] = 3;
    handCardsMap[12210] = 0;
    handCardsMap[12211] = 2;
    handCardsMap[12212] = 1;
    handCardsMap[12213] = 0;
    handCardsMap[12214] = 2;
    handCardsMap[12220] = 2;
    handCardsMap[12221] = 1;
    handCardsMap[12222] = 3;
    handCardsMap[12223] = 2;
    handCardsMap[12224] = 1;
    handCardsMap[12230] = 1;
    handCardsMap[12231] = 3;
    handCardsMap[12232] = 2;
    handCardsMap[12233] = 1;
    handCardsMap[12234] = 3;
    handCardsMap[12240] = 0;
    handCardsMap[12241] = 2;
    handCardsMap[12242] = 1;
    handCardsMap[12243] = 0;
    handCardsMap[12244] = 2;
    handCardsMap[12300] = 3;
    handCardsMap[12301] = 2;
    handCardsMap[12302] = 4;
    handCardsMap[12303] = 3;
    handCardsMap[12304] = 2;
    handCardsMap[12310] = 2;
    handCardsMap[12311] = 1;
    handCardsMap[12312] = 3;
    handCardsMap[12313] = 2;
    handCardsMap[12314] = 1;
    handCardsMap[12320] = 1;
    handCardsMap[12321] = 0;
    handCardsMap[12322] = 2;
    handCardsMap[12323] = 1;
    handCardsMap[12324] = 0;
    handCardsMap[12330] = 3;
    handCardsMap[12331] = 2;
    handCardsMap[12332] = 1;
    handCardsMap[12333] = 3;
    handCardsMap[12334] = 2;
    handCardsMap[12340] = 2;
    handCardsMap[12341] = 1;
    handCardsMap[12342] = 3;
    handCardsMap[12343] = 2;
    handCardsMap[12344] = 1;
    handCardsMap[12400] = 2;
    handCardsMap[12401] = 4;
    handCardsMap[12402] = 3;
    handCardsMap[12403] = 2;
    handCardsMap[12404] = 4;
    handCardsMap[12410] = 1;
    handCardsMap[12411] = 3;
    handCardsMap[12412] = 2;
    handCardsMap[12413] = 1;
    handCardsMap[12414] = 3;
    handCardsMap[12420] = 3;
    handCardsMap[12421] = 2;
    handCardsMap[12422] = 1;
    handCardsMap[12423] = 3;
    handCardsMap[12424] = 2;
    handCardsMap[12430] = 2;
    handCardsMap[12431] = 1;
    handCardsMap[12432] = 0;
    handCardsMap[12433] = 2;
    handCardsMap[12434] = 1;
    handCardsMap[12440] = 1;
    handCardsMap[12441] = 3;
    handCardsMap[12442] = 2;
    handCardsMap[12443] = 1;
    handCardsMap[13000] = 2;
    handCardsMap[13001] = 4;
    handCardsMap[13002] = 3;
    handCardsMap[13003] = 2;
    handCardsMap[13004] = 4;
    handCardsMap[13010] = 4;
    handCardsMap[13011] = 3;
    handCardsMap[13012] = 5;
    handCardsMap[13013] = 4;
    handCardsMap[13014] = 3;
    handCardsMap[13020] = 3;
    handCardsMap[13021] = 5;
    handCardsMap[13022] = 4;
    handCardsMap[13023] = 3;
    handCardsMap[13024] = 5;
    handCardsMap[13030] = 2;
    handCardsMap[13031] = 4;
    handCardsMap[13032] = 3;
    handCardsMap[13033] = 2;
    handCardsMap[13034] = 4;
    handCardsMap[13040] = 4;
    handCardsMap[13041] = 3;
    handCardsMap[13042] = 5;
    handCardsMap[13043] = 4;
    handCardsMap[13044] = 3;
    handCardsMap[13100] = 1;
    handCardsMap[13101] = 3;
    handCardsMap[13102] = 2;
    handCardsMap[13103] = 1;
    handCardsMap[13104] = 3;
    handCardsMap[13110] = 3;
    handCardsMap[13111] = 2;
    handCardsMap[13112] = 4;
    handCardsMap[13113] = 3;
    handCardsMap[13114] = 2;
    handCardsMap[13120] = 2;
    handCardsMap[13121] = 4;
    handCardsMap[13122] = 3;
    handCardsMap[13123] = 2;
    handCardsMap[13124] = 4;
    handCardsMap[13130] = 1;
    handCardsMap[13131] = 3;
    handCardsMap[13132] = 2;
    handCardsMap[13133] = 1;
    handCardsMap[13134] = 3;
    handCardsMap[13140] = 3;
    handCardsMap[13141] = 2;
    handCardsMap[13142] = 4;
    handCardsMap[13143] = 3;
    handCardsMap[13144] = 2;
    handCardsMap[13200] = 3;
    handCardsMap[13201] = 2;
    handCardsMap[13202] = 4;
    handCardsMap[13203] = 3;
    handCardsMap[13204] = 2;
    handCardsMap[13210] = 2;
    handCardsMap[13211] = 1;
    handCardsMap[13212] = 3;
    handCardsMap[13213] = 2;
    handCardsMap[13214] = 1;
    handCardsMap[13220] = 1;
    handCardsMap[13221] = 3;
    handCardsMap[13222] = 2;
    handCardsMap[13223] = 1;
    handCardsMap[13224] = 3;
    handCardsMap[13230] = 3;
    handCardsMap[13231] = 2;
    handCardsMap[13232] = 4;
    handCardsMap[13233] = 3;
    handCardsMap[13234] = 2;
    handCardsMap[13240] = 2;
    handCardsMap[13241] = 1;
    handCardsMap[13242] = 3;
    handCardsMap[13243] = 2;
    handCardsMap[13244] = 1;
    handCardsMap[13300] = 2;
    handCardsMap[13301] = 4;
    handCardsMap[13302] = 3;
    handCardsMap[13303] = 2;
    handCardsMap[13304] = 4;
    handCardsMap[13310] = 1;
    handCardsMap[13311] = 3;
    handCardsMap[13312] = 2;
    handCardsMap[13313] = 1;
    handCardsMap[13314] = 3;
    handCardsMap[13320] = 0;
    handCardsMap[13321] = 2;
    handCardsMap[13322] = 1;
    handCardsMap[13323] = 0;
    handCardsMap[13324] = 2;
    handCardsMap[13330] = 2;
    handCardsMap[13331] = 1;
    handCardsMap[13332] = 3;
    handCardsMap[13333] = 2;
    handCardsMap[13334] = 1;
    handCardsMap[13340] = 1;
    handCardsMap[13341] = 3;
    handCardsMap[13342] = 2;
    handCardsMap[13343] = 1;
    handCardsMap[13400] = 1;
    handCardsMap[13401] = 3;
    handCardsMap[13402] = 2;
    handCardsMap[13403] = 1;
    handCardsMap[13404] = 3;
    handCardsMap[13410] = 3;
    handCardsMap[13411] = 2;
    handCardsMap[13412] = 4;
    handCardsMap[13413] = 3;
    handCardsMap[13414] = 2;
    handCardsMap[13420] = 2;
    handCardsMap[13421] = 1;
    handCardsMap[13422] = 3;
    handCardsMap[13423] = 2;
    handCardsMap[13424] = 1;
    handCardsMap[13430] = 1;
    handCardsMap[13431] = 0;
    handCardsMap[13432] = 2;
    handCardsMap[13433] = 1;
    handCardsMap[13440] = 3;
    handCardsMap[13441] = 2;
    handCardsMap[13442] = 1;
    handCardsMap[14000] = 1;
    handCardsMap[14001] = 3;
    handCardsMap[14002] = 2;
    handCardsMap[14003] = 1;
    handCardsMap[14004] = 3;
    handCardsMap[14010] = 3;
    handCardsMap[14011] = 2;
    handCardsMap[14012] = 4;
    handCardsMap[14013] = 3;
    handCardsMap[14014] = 2;
    handCardsMap[14020] = 2;
    handCardsMap[14021] = 4;
    handCardsMap[14022] = 3;
    handCardsMap[14023] = 2;
    handCardsMap[14024] = 4;
    handCardsMap[14030] = 1;
    handCardsMap[14031] = 3;
    handCardsMap[14032] = 2;
    handCardsMap[14033] = 1;
    handCardsMap[14034] = 3;
    handCardsMap[14040] = 3;
    handCardsMap[14041] = 2;
    handCardsMap[14042] = 4;
    handCardsMap[14043] = 3;
    handCardsMap[14044] = 2;
    handCardsMap[14100] = 0;
    handCardsMap[14101] = 2;
    handCardsMap[14102] = 1;
    handCardsMap[14103] = 0;
    handCardsMap[14104] = 2;
    handCardsMap[14110] = 2;
    handCardsMap[14111] = 1;
    handCardsMap[14112] = 3;
    handCardsMap[14113] = 2;
    handCardsMap[14114] = 1;
    handCardsMap[14120] = 1;
    handCardsMap[14121] = 3;
    handCardsMap[14122] = 2;
    handCardsMap[14123] = 1;
    handCardsMap[14124] = 3;
    handCardsMap[14130] = 0;
    handCardsMap[14131] = 2;
    handCardsMap[14132] = 1;
    handCardsMap[14133] = 0;
    handCardsMap[14134] = 2;
    handCardsMap[14140] = 2;
    handCardsMap[14141] = 1;
    handCardsMap[14142] = 3;
    handCardsMap[14143] = 2;
    handCardsMap[14144] = 1;
    handCardsMap[14200] = 2;
    handCardsMap[14201] = 1;
    handCardsMap[14202] = 3;
    handCardsMap[14203] = 2;
    handCardsMap[14204] = 1;
    handCardsMap[14210] = 1;
    handCardsMap[14211] = 0;
    handCardsMap[14212] = 2;
    handCardsMap[14213] = 1;
    handCardsMap[14214] = 0;
    handCardsMap[14220] = 3;
    handCardsMap[14221] = 2;
    handCardsMap[14222] = 1;
    handCardsMap[14223] = 3;
    handCardsMap[14224] = 2;
    handCardsMap[14230] = 2;
    handCardsMap[14231] = 1;
    handCardsMap[14232] = 3;
    handCardsMap[14233] = 2;
    handCardsMap[14234] = 1;
    handCardsMap[14240] = 1;
    handCardsMap[14241] = 0;
    handCardsMap[14242] = 2;
    handCardsMap[14243] = 1;
    handCardsMap[14300] = 1;
    handCardsMap[14301] = 3;
    handCardsMap[14302] = 2;
    handCardsMap[14303] = 1;
    handCardsMap[14304] = 3;
    handCardsMap[14310] = 3;
    handCardsMap[14311] = 2;
    handCardsMap[14312] = 1;
    handCardsMap[14313] = 3;
    handCardsMap[14314] = 2;
    handCardsMap[14320] = 2;
    handCardsMap[14321] = 1;
    handCardsMap[14322] = 0;
    handCardsMap[14323] = 2;
    handCardsMap[14324] = 1;
    handCardsMap[14330] = 1;
    handCardsMap[14331] = 3;
    handCardsMap[14332] = 2;
    handCardsMap[14333] = 1;
    handCardsMap[14340] = 3;
    handCardsMap[14341] = 2;
    handCardsMap[14342] = 1;
    handCardsMap[14400] = 0;
    handCardsMap[14401] = 2;
    handCardsMap[14402] = 1;
    handCardsMap[14403] = 0;
    handCardsMap[14404] = 2;
    handCardsMap[14410] = 2;
    handCardsMap[14411] = 1;
    handCardsMap[14412] = 3;
    handCardsMap[14413] = 2;
    handCardsMap[14414] = 1;
    handCardsMap[14420] = 1;
    handCardsMap[14421] = 3;
    handCardsMap[14422] = 2;
    handCardsMap[14423] = 1;
    handCardsMap[14430] = 0;
    handCardsMap[14431] = 2;
    handCardsMap[14432] = 1;
    handCardsMap[14440] = 2;
    handCardsMap[14441] = 1;
    handCardsMap[20000] = 1;
    handCardsMap[20001] = 3;
    handCardsMap[20002] = 2;
    handCardsMap[20003] = 1;
    handCardsMap[20004] = 3;
    handCardsMap[20010] = 3;
    handCardsMap[20011] = 2;
    handCardsMap[20012] = 4;
    handCardsMap[20013] = 3;
    handCardsMap[20014] = 2;
    handCardsMap[20020] = 2;
    handCardsMap[20021] = 4;
    handCardsMap[20022] = 3;
    handCardsMap[20023] = 2;
    handCardsMap[20024] = 4;
    handCardsMap[20030] = 1;
    handCardsMap[20031] = 3;
    handCardsMap[20032] = 2;
    handCardsMap[20033] = 1;
    handCardsMap[20034] = 3;
    handCardsMap[20040] = 3;
    handCardsMap[20041] = 2;
    handCardsMap[20042] = 4;
    handCardsMap[20043] = 3;
    handCardsMap[20044] = 2;
    handCardsMap[20100] = 3;
    handCardsMap[20101] = 2;
    handCardsMap[20102] = 4;
    handCardsMap[20103] = 3;
    handCardsMap[20104] = 2;
    handCardsMap[20110] = 2;
    handCardsMap[20111] = 1;
    handCardsMap[20112] = 3;
    handCardsMap[20113] = 2;
    handCardsMap[20114] = 1;
    handCardsMap[20120] = 4;
    handCardsMap[20121] = 3;
    handCardsMap[20122] = 2;
    handCardsMap[20123] = 4;
    handCardsMap[20124] = 3;
    handCardsMap[20130] = 3;
    handCardsMap[20131] = 2;
    handCardsMap[20132] = 4;
    handCardsMap[20133] = 3;
    handCardsMap[20134] = 2;
    handCardsMap[20140] = 2;
    handCardsMap[20141] = 1;
    handCardsMap[20142] = 3;
    handCardsMap[20143] = 2;
    handCardsMap[20144] = 1;
    handCardsMap[20200] = 2;
    handCardsMap[20201] = 4;
    handCardsMap[20202] = 3;
    handCardsMap[20203] = 2;
    handCardsMap[20204] = 4;
    handCardsMap[20210] = 4;
    handCardsMap[20211] = 3;
    handCardsMap[20212] = 2;
    handCardsMap[20213] = 4;
    handCardsMap[20214] = 3;
    handCardsMap[20220] = 3;
    handCardsMap[20221] = 2;
    handCardsMap[20222] = 1;
    handCardsMap[20223] = 3;
    handCardsMap[20224] = 2;
    handCardsMap[20230] = 2;
    handCardsMap[20231] = 4;
    handCardsMap[20232] = 3;
    handCardsMap[20233] = 2;
    handCardsMap[20234] = 4;
    handCardsMap[20240] = 4;
    handCardsMap[20241] = 3;
    handCardsMap[20242] = 2;
    handCardsMap[20243] = 4;
    handCardsMap[20244] = 3;
    handCardsMap[20300] = 1;
    handCardsMap[20301] = 3;
    handCardsMap[20302] = 2;
    handCardsMap[20303] = 1;
    handCardsMap[20304] = 3;
    handCardsMap[20310] = 3;
    handCardsMap[20311] = 2;
    handCardsMap[20312] = 4;
    handCardsMap[20313] = 3;
    handCardsMap[20314] = 2;
    handCardsMap[20320] = 2;
    handCardsMap[20321] = 4;
    handCardsMap[20322] = 3;
    handCardsMap[20323] = 2;
    handCardsMap[20324] = 4;
    handCardsMap[20330] = 1;
    handCardsMap[20331] = 3;
    handCardsMap[20332] = 2;
    handCardsMap[20333] = 1;
    handCardsMap[20334] = 3;
    handCardsMap[20340] = 3;
    handCardsMap[20341] = 2;
    handCardsMap[20342] = 4;
    handCardsMap[20343] = 3;
    handCardsMap[20344] = 2;
    handCardsMap[20400] = 3;
    handCardsMap[20401] = 2;
    handCardsMap[20402] = 4;
    handCardsMap[20403] = 3;
    handCardsMap[20404] = 2;
    handCardsMap[20410] = 2;
    handCardsMap[20411] = 1;
    handCardsMap[20412] = 3;
    handCardsMap[20413] = 2;
    handCardsMap[20414] = 1;
    handCardsMap[20420] = 4;
    handCardsMap[20421] = 3;
    handCardsMap[20422] = 2;
    handCardsMap[20423] = 4;
    handCardsMap[20424] = 3;
    handCardsMap[20430] = 3;
    handCardsMap[20431] = 2;
    handCardsMap[20432] = 4;
    handCardsMap[20433] = 3;
    handCardsMap[20434] = 2;
    handCardsMap[20440] = 2;
    handCardsMap[20441] = 1;
    handCardsMap[20442] = 3;
    handCardsMap[20443] = 2;
    handCardsMap[20444] = 1;
    handCardsMap[21000] = 3;
    handCardsMap[21001] = 5;
    handCardsMap[21002] = 4;
    handCardsMap[21003] = 3;
    handCardsMap[21004] = 5;
    handCardsMap[21010] = 2;
    handCardsMap[21011] = 4;
    handCardsMap[21012] = 3;
    handCardsMap[21013] = 2;
    handCardsMap[21014] = 4;
    handCardsMap[21020] = 4;
    handCardsMap[21021] = 3;
    handCardsMap[21022] = 5;
    handCardsMap[21023] = 4;
    handCardsMap[21024] = 3;
    handCardsMap[21030] = 3;
    handCardsMap[21031] = 5;
    handCardsMap[21032] = 4;
    handCardsMap[21033] = 3;
    handCardsMap[21034] = 5;
    handCardsMap[21040] = 2;
    handCardsMap[21041] = 4;
    handCardsMap[21042] = 3;
    handCardsMap[21043] = 2;
    handCardsMap[21044] = 4;
    handCardsMap[21100] = 2;
    handCardsMap[21101] = 4;
    handCardsMap[21102] = 3;
    handCardsMap[21103] = 2;
    handCardsMap[21104] = 4;
    handCardsMap[21110] = 1;
    handCardsMap[21111] = 3;
    handCardsMap[21112] = 2;
    handCardsMap[21113] = 1;
    handCardsMap[21114] = 3;
    handCardsMap[21120] = 3;
    handCardsMap[21121] = 2;
    handCardsMap[21122] = 4;
    handCardsMap[21123] = 3;
    handCardsMap[21124] = 2;
    handCardsMap[21130] = 2;
    handCardsMap[21131] = 4;
    handCardsMap[21132] = 3;
    handCardsMap[21133] = 2;
    handCardsMap[21134] = 4;
    handCardsMap[21140] = 1;
    handCardsMap[21141] = 3;
    handCardsMap[21142] = 2;
    handCardsMap[21143] = 1;
    handCardsMap[21144] = 3;
    handCardsMap[21200] = 1;
    handCardsMap[21201] = 3;
    handCardsMap[21202] = 2;
    handCardsMap[21203] = 1;
    handCardsMap[21204] = 3;
    handCardsMap[21210] = 3;
    handCardsMap[21211] = 2;
    handCardsMap[21212] = 4;
    handCardsMap[21213] = 3;
    handCardsMap[21214] = 2;
    handCardsMap[21220] = 2;
    handCardsMap[21221] = 1;
    handCardsMap[21222] = 3;
    handCardsMap[21223] = 2;
    handCardsMap[21224] = 1;
    handCardsMap[21230] = 1;
    handCardsMap[21231] = 3;
    handCardsMap[21232] = 2;
    handCardsMap[21233] = 1;
    handCardsMap[21234] = 3;
    handCardsMap[21240] = 3;
    handCardsMap[21241] = 2;
    handCardsMap[21242] = 4;
    handCardsMap[21243] = 3;
    handCardsMap[21244] = 2;
    handCardsMap[21300] = 3;
    handCardsMap[21301] = 2;
    handCardsMap[21302] = 4;
    handCardsMap[21303] = 3;
    handCardsMap[21304] = 2;
    handCardsMap[21310] = 2;
    handCardsMap[21311] = 1;
    handCardsMap[21312] = 3;
    handCardsMap[21313] = 2;
    handCardsMap[21314] = 1;
    handCardsMap[21320] = 4;
    handCardsMap[21321] = 3;
    handCardsMap[21322] = 2;
    handCardsMap[21323] = 4;
    handCardsMap[21324] = 3;
    handCardsMap[21330] = 3;
    handCardsMap[21331] = 2;
    handCardsMap[21332] = 1;
    handCardsMap[21333] = 3;
    handCardsMap[21334] = 2;
    handCardsMap[21340] = 2;
    handCardsMap[21341] = 1;
    handCardsMap[21342] = 3;
    handCardsMap[21343] = 2;
    handCardsMap[21344] = 1;
    handCardsMap[21400] = 2;
    handCardsMap[21401] = 4;
    handCardsMap[21402] = 3;
    handCardsMap[21403] = 2;
    handCardsMap[21404] = 4;
    handCardsMap[21410] = 1;
    handCardsMap[21411] = 3;
    handCardsMap[21412] = 2;
    handCardsMap[21413] = 1;
    handCardsMap[21414] = 3;
    handCardsMap[21420] = 3;
    handCardsMap[21421] = 2;
    handCardsMap[21422] = 1;
    handCardsMap[21423] = 3;
    handCardsMap[21424] = 2;
    handCardsMap[21430] = 2;
    handCardsMap[21431] = 4;
    handCardsMap[21432] = 3;
    handCardsMap[21433] = 2;
    handCardsMap[21434] = 4;
    handCardsMap[21440] = 1;
    handCardsMap[21441] = 3;
    handCardsMap[21442] = 2;
    handCardsMap[21443] = 1;
    handCardsMap[22000] = 2;
    handCardsMap[22001] = 4;
    handCardsMap[22002] = 3;
    handCardsMap[22003] = 2;
    handCardsMap[22004] = 4;
    handCardsMap[22010] = 4;
    handCardsMap[22011] = 3;
    handCardsMap[22012] = 5;
    handCardsMap[22013] = 4;
    handCardsMap[22014] = 3;
    handCardsMap[22020] = 3;
    handCardsMap[22021] = 5;
    handCardsMap[22022] = 4;
    handCardsMap[22023] = 3;
    handCardsMap[22024] = 5;
    handCardsMap[22030] = 2;
    handCardsMap[22031] = 4;
    handCardsMap[22032] = 3;
    handCardsMap[22033] = 2;
    handCardsMap[22034] = 4;
    handCardsMap[22040] = 4;
    handCardsMap[22041] = 3;
    handCardsMap[22042] = 5;
    handCardsMap[22043] = 4;
    handCardsMap[22044] = 3;
    handCardsMap[22100] = 1;
    handCardsMap[22101] = 3;
    handCardsMap[22102] = 2;
    handCardsMap[22103] = 1;
    handCardsMap[22104] = 3;
    handCardsMap[22110] = 3;
    handCardsMap[22111] = 2;
    handCardsMap[22112] = 4;
    handCardsMap[22113] = 3;
    handCardsMap[22114] = 2;
    handCardsMap[22120] = 2;
    handCardsMap[22121] = 4;
    handCardsMap[22122] = 3;
    handCardsMap[22123] = 2;
    handCardsMap[22124] = 4;
    handCardsMap[22130] = 1;
    handCardsMap[22131] = 3;
    handCardsMap[22132] = 2;
    handCardsMap[22133] = 1;
    handCardsMap[22134] = 3;
    handCardsMap[22140] = 3;
    handCardsMap[22141] = 2;
    handCardsMap[22142] = 4;
    handCardsMap[22143] = 3;
    handCardsMap[22144] = 2;
    handCardsMap[22200] = 0;
    handCardsMap[22201] = 2;
    handCardsMap[22202] = 1;
    handCardsMap[22203] = 0;
    handCardsMap[22204] = 2;
    handCardsMap[22210] = 2;
    handCardsMap[22211] = 1;
    handCardsMap[22212] = 3;
    handCardsMap[22213] = 2;
    handCardsMap[22214] = 1;
    handCardsMap[22220] = 1;
    handCardsMap[22221] = 3;
    handCardsMap[22222] = 2;
    handCardsMap[22223] = 1;
    handCardsMap[22224] = 3;
    handCardsMap[22230] = 0;
    handCardsMap[22231] = 2;
    handCardsMap[22232] = 1;
    handCardsMap[22233] = 0;
    handCardsMap[22234] = 2;
    handCardsMap[22240] = 2;
    handCardsMap[22241] = 1;
    handCardsMap[22242] = 3;
    handCardsMap[22243] = 2;
    handCardsMap[22244] = 1;
    handCardsMap[22300] = 2;
    handCardsMap[22301] = 1;
    handCardsMap[22302] = 3;
    handCardsMap[22303] = 2;
    handCardsMap[22304] = 1;
    handCardsMap[22310] = 1;
    handCardsMap[22311] = 0;
    handCardsMap[22312] = 2;
    handCardsMap[22313] = 1;
    handCardsMap[22314] = 0;
    handCardsMap[22320] = 3;
    handCardsMap[22321] = 2;
    handCardsMap[22322] = 1;
    handCardsMap[22323] = 3;
    handCardsMap[22324] = 2;
    handCardsMap[22330] = 2;
    handCardsMap[22331] = 1;
    handCardsMap[22332] = 3;
    handCardsMap[22333] = 2;
    handCardsMap[22334] = 1;
    handCardsMap[22340] = 1;
    handCardsMap[22341] = 0;
    handCardsMap[22342] = 2;
    handCardsMap[22343] = 1;
    handCardsMap[22400] = 1;
    handCardsMap[22401] = 3;
    handCardsMap[22402] = 2;
    handCardsMap[22403] = 1;
    handCardsMap[22404] = 3;
    handCardsMap[22410] = 3;
    handCardsMap[22411] = 2;
    handCardsMap[22412] = 1;
    handCardsMap[22413] = 3;
    handCardsMap[22414] = 2;
    handCardsMap[22420] = 2;
    handCardsMap[22421] = 1;
    handCardsMap[22422] = 0;
    handCardsMap[22423] = 2;
    handCardsMap[22424] = 1;
    handCardsMap[22430] = 1;
    handCardsMap[22431] = 3;
    handCardsMap[22432] = 2;
    handCardsMap[22433] = 1;
    handCardsMap[22440] = 3;
    handCardsMap[22441] = 2;
    handCardsMap[22442] = 1;
    handCardsMap[23000] = 1;
    handCardsMap[23001] = 3;
    handCardsMap[23002] = 2;
    handCardsMap[23003] = 1;
    handCardsMap[23004] = 3;
    handCardsMap[23010] = 3;
    handCardsMap[23011] = 2;
    handCardsMap[23012] = 4;
    handCardsMap[23013] = 3;
    handCardsMap[23014] = 2;
    handCardsMap[23020] = 2;
    handCardsMap[23021] = 4;
    handCardsMap[23022] = 3;
    handCardsMap[23023] = 2;
    handCardsMap[23024] = 4;
    handCardsMap[23030] = 1;
    handCardsMap[23031] = 3;
    handCardsMap[23032] = 2;
    handCardsMap[23033] = 1;
    handCardsMap[23034] = 3;
    handCardsMap[23040] = 3;
    handCardsMap[23041] = 2;
    handCardsMap[23042] = 4;
    handCardsMap[23043] = 3;
    handCardsMap[23044] = 2;
    handCardsMap[23100] = 3;
    handCardsMap[23101] = 2;
    handCardsMap[23102] = 4;
    handCardsMap[23103] = 3;
    handCardsMap[23104] = 2;
    handCardsMap[23110] = 2;
    handCardsMap[23111] = 1;
    handCardsMap[23112] = 3;
    handCardsMap[23113] = 2;
    handCardsMap[23114] = 1;
    handCardsMap[23120] = 4;
    handCardsMap[23121] = 3;
    handCardsMap[23122] = 2;
    handCardsMap[23123] = 4;
    handCardsMap[23124] = 3;
    handCardsMap[23130] = 3;
    handCardsMap[23131] = 2;
    handCardsMap[23132] = 4;
    handCardsMap[23133] = 3;
    handCardsMap[23134] = 2;
    handCardsMap[23140] = 2;
    handCardsMap[23141] = 1;
    handCardsMap[23142] = 3;
    handCardsMap[23143] = 2;
    handCardsMap[23144] = 1;
    handCardsMap[23200] = 2;
    handCardsMap[23201] = 4;
    handCardsMap[23202] = 3;
    handCardsMap[23203] = 2;
    handCardsMap[23204] = 4;
    handCardsMap[23210] = 1;
    handCardsMap[23211] = 3;
    handCardsMap[23212] = 2;
    handCardsMap[23213] = 1;
    handCardsMap[23214] = 3;
    handCardsMap[23220] = 3;
    handCardsMap[23221] = 2;
    handCardsMap[23222] = 1;
    handCardsMap[23223] = 3;
    handCardsMap[23224] = 2;
    handCardsMap[23230] = 2;
    handCardsMap[23231] = 4;
    handCardsMap[23232] = 3;
    handCardsMap[23233] = 2;
    handCardsMap[23234] = 4;
    handCardsMap[23240] = 1;
    handCardsMap[23241] = 3;
    handCardsMap[23242] = 2;
    handCardsMap[23243] = 1;
    handCardsMap[23300] = 1;
    handCardsMap[23301] = 3;
    handCardsMap[23302] = 2;
    handCardsMap[23303] = 1;
    handCardsMap[23304] = 3;
    handCardsMap[23310] = 0;
    handCardsMap[23311] = 2;
    handCardsMap[23312] = 1;
    handCardsMap[23313] = 0;
    handCardsMap[23314] = 2;
    handCardsMap[23320] = 2;
    handCardsMap[23321] = 1;
    handCardsMap[23322] = 3;
    handCardsMap[23323] = 2;
    handCardsMap[23324] = 1;
    handCardsMap[23330] = 1;
    handCardsMap[23331] = 3;
    handCardsMap[23332] = 2;
    handCardsMap[23333] = 1;
    handCardsMap[23340] = 0;
    handCardsMap[23341] = 2;
    handCardsMap[23342] = 1;
    handCardsMap[23400] = 3;
    handCardsMap[23401] = 2;
    handCardsMap[23402] = 4;
    handCardsMap[23403] = 3;
    handCardsMap[23404] = 2;
    handCardsMap[23410] = 2;
    handCardsMap[23411] = 1;
    handCardsMap[23412] = 3;
    handCardsMap[23413] = 2;
    handCardsMap[23414] = 1;
    handCardsMap[23420] = 1;
    handCardsMap[23421] = 0;
    handCardsMap[23422] = 2;
    handCardsMap[23423] = 1;
    handCardsMap[23430] = 3;
    handCardsMap[23431] = 2;
    handCardsMap[23432] = 1;
    handCardsMap[23440] = 2;
    handCardsMap[23441] = 1;
    handCardsMap[24000] = 3;
    handCardsMap[24001] = 5;
    handCardsMap[24002] = 4;
    handCardsMap[24003] = 3;
    handCardsMap[24004] = 5;
    handCardsMap[24010] = 2;
    handCardsMap[24011] = 4;
    handCardsMap[24012] = 3;
    handCardsMap[24013] = 2;
    handCardsMap[24014] = 4;
    handCardsMap[24020] = 4;
    handCardsMap[24021] = 3;
    handCardsMap[24022] = 5;
    handCardsMap[24023] = 4;
    handCardsMap[24024] = 3;
    handCardsMap[24030] = 3;
    handCardsMap[24031] = 5;
    handCardsMap[24032] = 4;
    handCardsMap[24033] = 3;
    handCardsMap[24034] = 5;
    handCardsMap[24040] = 2;
    handCardsMap[24041] = 4;
    handCardsMap[24042] = 3;
    handCardsMap[24043] = 2;
    handCardsMap[24044] = 4;
    handCardsMap[24100] = 2;
    handCardsMap[24101] = 4;
    handCardsMap[24102] = 3;
    handCardsMap[24103] = 2;
    handCardsMap[24104] = 4;
    handCardsMap[24110] = 1;
    handCardsMap[24111] = 3;
    handCardsMap[24112] = 2;
    handCardsMap[24113] = 1;
    handCardsMap[24114] = 3;
    handCardsMap[24120] = 3;
    handCardsMap[24121] = 2;
    handCardsMap[24122] = 4;
    handCardsMap[24123] = 3;
    handCardsMap[24124] = 2;
    handCardsMap[24130] = 2;
    handCardsMap[24131] = 4;
    handCardsMap[24132] = 3;
    handCardsMap[24133] = 2;
    handCardsMap[24134] = 4;
    handCardsMap[24140] = 1;
    handCardsMap[24141] = 3;
    handCardsMap[24142] = 2;
    handCardsMap[24143] = 1;
    handCardsMap[24200] = 1;
    handCardsMap[24201] = 3;
    handCardsMap[24202] = 2;
    handCardsMap[24203] = 1;
    handCardsMap[24204] = 3;
    handCardsMap[24210] = 3;
    handCardsMap[24211] = 2;
    handCardsMap[24212] = 4;
    handCardsMap[24213] = 3;
    handCardsMap[24214] = 2;
    handCardsMap[24220] = 2;
    handCardsMap[24221] = 1;
    handCardsMap[24222] = 3;
    handCardsMap[24223] = 2;
    handCardsMap[24224] = 1;
    handCardsMap[24230] = 1;
    handCardsMap[24231] = 3;
    handCardsMap[24232] = 2;
    handCardsMap[24233] = 1;
    handCardsMap[24240] = 3;
    handCardsMap[24241] = 2;
    handCardsMap[24242] = 4;
    handCardsMap[24300] = 3;
    handCardsMap[24301] = 2;
    handCardsMap[24302] = 4;
    handCardsMap[24303] = 3;
    handCardsMap[24304] = 2;
    handCardsMap[24310] = 2;
    handCardsMap[24311] = 1;
    handCardsMap[24312] = 3;
    handCardsMap[24313] = 2;
    handCardsMap[24314] = 1;
    handCardsMap[24320] = 1;
    handCardsMap[24321] = 3;
    handCardsMap[24322] = 2;
    handCardsMap[24323] = 1;
    handCardsMap[24330] = 3;
    handCardsMap[24331] = 2;
    handCardsMap[24332] = 1;
    handCardsMap[24340] = 2;
    handCardsMap[24341] = 1;
    handCardsMap[24400] = 2;
    handCardsMap[24401] = 4;
    handCardsMap[24402] = 3;
    handCardsMap[24403] = 2;
    handCardsMap[24404] = 4;
    handCardsMap[24410] = 1;
    handCardsMap[24411] = 3;
    handCardsMap[24412] = 2;
    handCardsMap[24413] = 1;
    handCardsMap[24420] = 0;
    handCardsMap[24421] = 2;
    handCardsMap[24422] = 1;
    handCardsMap[24430] = 2;
    handCardsMap[24431] = 1;
    handCardsMap[24440] = 1;
    handCardsMap[30000] = 0;
    handCardsMap[30001] = 2;
    handCardsMap[30002] = 1;
    handCardsMap[30003] = 0;
    handCardsMap[30004] = 2;
    handCardsMap[30010] = 2;
    handCardsMap[30011] = 1;
    handCardsMap[30012] = 3;
    handCardsMap[30013] = 2;
    handCardsMap[30014] = 1;
    handCardsMap[30020] = 1;
    handCardsMap[30021] = 3;
    handCardsMap[30022] = 2;
    handCardsMap[30023] = 1;
    handCardsMap[30024] = 3;
    handCardsMap[30030] = 0;
    handCardsMap[30031] = 2;
    handCardsMap[30032] = 1;
    handCardsMap[30033] = 0;
    handCardsMap[30034] = 2;
    handCardsMap[30040] = 2;
    handCardsMap[30041] = 1;
    handCardsMap[30042] = 3;
    handCardsMap[30043] = 2;
    handCardsMap[30044] = 1;
    handCardsMap[30100] = 2;
    handCardsMap[30101] = 1;
    handCardsMap[30102] = 3;
    handCardsMap[30103] = 2;
    handCardsMap[30104] = 1;
    handCardsMap[30110] = 1;
    handCardsMap[30111] = 0;
    handCardsMap[30112] = 2;
    handCardsMap[30113] = 1;
    handCardsMap[30114] = 0;
    handCardsMap[30120] = 3;
    handCardsMap[30121] = 2;
    handCardsMap[30122] = 1;
    handCardsMap[30123] = 3;
    handCardsMap[30124] = 2;
    handCardsMap[30130] = 2;
    handCardsMap[30131] = 1;
    handCardsMap[30132] = 3;
    handCardsMap[30133] = 2;
    handCardsMap[30134] = 1;
    handCardsMap[30140] = 1;
    handCardsMap[30141] = 0;
    handCardsMap[30142] = 2;
    handCardsMap[30143] = 1;
    handCardsMap[30144] = 0;
    handCardsMap[30200] = 1;
    handCardsMap[30201] = 3;
    handCardsMap[30202] = 2;
    handCardsMap[30203] = 1;
    handCardsMap[30204] = 3;
    handCardsMap[30210] = 3;
    handCardsMap[30211] = 2;
    handCardsMap[30212] = 1;
    handCardsMap[30213] = 3;
    handCardsMap[30214] = 2;
    handCardsMap[30220] = 2;
    handCardsMap[30221] = 1;
    handCardsMap[30222] = 0;
    handCardsMap[30223] = 2;
    handCardsMap[30224] = 1;
    handCardsMap[30230] = 1;
    handCardsMap[30231] = 3;
    handCardsMap[30232] = 2;
    handCardsMap[30233] = 1;
    handCardsMap[30234] = 3;
    handCardsMap[30240] = 3;
    handCardsMap[30241] = 2;
    handCardsMap[30242] = 1;
    handCardsMap[30243] = 3;
    handCardsMap[30244] = 2;
    handCardsMap[30300] = 0;
    handCardsMap[30301] = 2;
    handCardsMap[30302] = 1;
    handCardsMap[30303] = 0;
    handCardsMap[30304] = 2;
    handCardsMap[30310] = 2;
    handCardsMap[30311] = 1;
    handCardsMap[30312] = 3;
    handCardsMap[30313] = 2;
    handCardsMap[30314] = 1;
    handCardsMap[30320] = 1;
    handCardsMap[30321] = 3;
    handCardsMap[30322] = 2;
    handCardsMap[30323] = 1;
    handCardsMap[30324] = 3;
    handCardsMap[30330] = 0;
    handCardsMap[30331] = 2;
    handCardsMap[30332] = 1;
    handCardsMap[30333] = 0;
    handCardsMap[30334] = 2;
    handCardsMap[30340] = 2;
    handCardsMap[30341] = 1;
    handCardsMap[30342] = 3;
    handCardsMap[30343] = 2;
    handCardsMap[30344] = 1;
    handCardsMap[30400] = 2;
    handCardsMap[30401] = 1;
    handCardsMap[30402] = 3;
    handCardsMap[30403] = 2;
    handCardsMap[30404] = 1;
    handCardsMap[30410] = 1;
    handCardsMap[30411] = 0;
    handCardsMap[30412] = 2;
    handCardsMap[30413] = 1;
    handCardsMap[30414] = 0;
    handCardsMap[30420] = 3;
    handCardsMap[30421] = 2;
    handCardsMap[30422] = 1;
    handCardsMap[30423] = 3;
    handCardsMap[30424] = 2;
    handCardsMap[30430] = 2;
    handCardsMap[30431] = 1;
    handCardsMap[30432] = 3;
    handCardsMap[30433] = 2;
    handCardsMap[30434] = 1;
    handCardsMap[30440] = 1;
    handCardsMap[30441] = 0;
    handCardsMap[30442] = 2;
    handCardsMap[30443] = 1;
    handCardsMap[31000] = 2;
    handCardsMap[31001] = 4;
    handCardsMap[31002] = 3;
    handCardsMap[31003] = 2;
    handCardsMap[31004] = 4;
    handCardsMap[31010] = 1;
    handCardsMap[31011] = 3;
    handCardsMap[31012] = 2;
    handCardsMap[31013] = 1;
    handCardsMap[31014] = 3;
    handCardsMap[31020] = 3;
    handCardsMap[31021] = 2;
    handCardsMap[31022] = 4;
    handCardsMap[31023] = 3;
    handCardsMap[31024] = 2;
    handCardsMap[31030] = 2;
    handCardsMap[31031] = 4;
    handCardsMap[31032] = 3;
    handCardsMap[31033] = 2;
    handCardsMap[31034] = 4;
    handCardsMap[31040] = 1;
    handCardsMap[31041] = 3;
    handCardsMap[31042] = 2;
    handCardsMap[31043] = 1;
    handCardsMap[31044] = 3;
    handCardsMap[31100] = 1;
    handCardsMap[31101] = 3;
    handCardsMap[31102] = 2;
    handCardsMap[31103] = 1;
    handCardsMap[31104] = 3;
    handCardsMap[31110] = 0;
    handCardsMap[31111] = 2;
    handCardsMap[31112] = 1;
    handCardsMap[31113] = 0;
    handCardsMap[31114] = 2;
    handCardsMap[31120] = 2;
    handCardsMap[31121] = 1;
    handCardsMap[31122] = 3;
    handCardsMap[31123] = 2;
    handCardsMap[31124] = 1;
    handCardsMap[31130] = 1;
    handCardsMap[31131] = 3;
    handCardsMap[31132] = 2;
    handCardsMap[31133] = 1;
    handCardsMap[31134] = 3;
    handCardsMap[31140] = 0;
    handCardsMap[31141] = 2;
    handCardsMap[31142] = 1;
    handCardsMap[31143] = 0;
    handCardsMap[31144] = 2;
    handCardsMap[31200] = 3;
    handCardsMap[31201] = 2;
    handCardsMap[31202] = 4;
    handCardsMap[31203] = 3;
    handCardsMap[31204] = 2;
    handCardsMap[31210] = 2;
    handCardsMap[31211] = 1;
    handCardsMap[31212] = 3;
    handCardsMap[31213] = 2;
    handCardsMap[31214] = 1;
    handCardsMap[31220] = 1;
    handCardsMap[31221] = 0;
    handCardsMap[31222] = 2;
    handCardsMap[31223] = 1;
    handCardsMap[31224] = 0;
    handCardsMap[31230] = 3;
    handCardsMap[31231] = 2;
    handCardsMap[31232] = 1;
    handCardsMap[31233] = 3;
    handCardsMap[31234] = 2;
    handCardsMap[31240] = 2;
    handCardsMap[31241] = 1;
    handCardsMap[31242] = 3;
    handCardsMap[31243] = 2;
    handCardsMap[31244] = 1;
    handCardsMap[31300] = 2;
    handCardsMap[31301] = 4;
    handCardsMap[31302] = 3;
    handCardsMap[31303] = 2;
    handCardsMap[31304] = 4;
    handCardsMap[31310] = 1;
    handCardsMap[31311] = 3;
    handCardsMap[31312] = 2;
    handCardsMap[31313] = 1;
    handCardsMap[31314] = 3;
    handCardsMap[31320] = 3;
    handCardsMap[31321] = 2;
    handCardsMap[31322] = 1;
    handCardsMap[31323] = 3;
    handCardsMap[31324] = 2;
    handCardsMap[31330] = 2;
    handCardsMap[31331] = 1;
    handCardsMap[31332] = 0;
    handCardsMap[31333] = 2;
    handCardsMap[31334] = 1;
    handCardsMap[31340] = 1;
    handCardsMap[31341] = 3;
    handCardsMap[31342] = 2;
    handCardsMap[31343] = 1;
    handCardsMap[31400] = 1;
    handCardsMap[31401] = 3;
    handCardsMap[31402] = 2;
    handCardsMap[31403] = 1;
    handCardsMap[31404] = 3;
    handCardsMap[31410] = 0;
    handCardsMap[31411] = 2;
    handCardsMap[31412] = 1;
    handCardsMap[31413] = 0;
    handCardsMap[31414] = 2;
    handCardsMap[31420] = 2;
    handCardsMap[31421] = 1;
    handCardsMap[31422] = 3;
    handCardsMap[31423] = 2;
    handCardsMap[31424] = 1;
    handCardsMap[31430] = 1;
    handCardsMap[31431] = 3;
    handCardsMap[31432] = 2;
    handCardsMap[31433] = 1;
    handCardsMap[31440] = 0;
    handCardsMap[31441] = 2;
    handCardsMap[31442] = 1;
    handCardsMap[32000] = 1;
    handCardsMap[32001] = 3;
    handCardsMap[32002] = 2;
    handCardsMap[32003] = 1;
    handCardsMap[32004] = 3;
    handCardsMap[32010] = 3;
    handCardsMap[32011] = 2;
    handCardsMap[32012] = 4;
    handCardsMap[32013] = 3;
    handCardsMap[32014] = 2;
    handCardsMap[32020] = 2;
    handCardsMap[32021] = 4;
    handCardsMap[32022] = 3;
    handCardsMap[32023] = 2;
    handCardsMap[32024] = 4;
    handCardsMap[32030] = 1;
    handCardsMap[32031] = 3;
    handCardsMap[32032] = 2;
    handCardsMap[32033] = 1;
    handCardsMap[32034] = 3;
    handCardsMap[32040] = 3;
    handCardsMap[32041] = 2;
    handCardsMap[32042] = 4;
    handCardsMap[32043] = 3;
    handCardsMap[32044] = 2;
    handCardsMap[32100] = 3;
    handCardsMap[32101] = 2;
    handCardsMap[32102] = 4;
    handCardsMap[32103] = 3;
    handCardsMap[32104] = 2;
    handCardsMap[32110] = 2;
    handCardsMap[32111] = 1;
    handCardsMap[32112] = 3;
    handCardsMap[32113] = 2;
    handCardsMap[32114] = 1;
    handCardsMap[32120] = 1;
    handCardsMap[32121] = 3;
    handCardsMap[32122] = 2;
    handCardsMap[32123] = 1;
    handCardsMap[32124] = 3;
    handCardsMap[32130] = 3;
    handCardsMap[32131] = 2;
    handCardsMap[32132] = 4;
    handCardsMap[32133] = 3;
    handCardsMap[32134] = 2;
    handCardsMap[32140] = 2;
    handCardsMap[32141] = 1;
    handCardsMap[32142] = 3;
    handCardsMap[32143] = 2;
    handCardsMap[32144] = 1;
    handCardsMap[32200] = 2;
    handCardsMap[32201] = 4;
    handCardsMap[32202] = 3;
    handCardsMap[32203] = 2;
    handCardsMap[32204] = 4;
    handCardsMap[32210] = 1;
    handCardsMap[32211] = 3;
    handCardsMap[32212] = 2;
    handCardsMap[32213] = 1;
    handCardsMap[32214] = 3;
    handCardsMap[32220] = 0;
    handCardsMap[32221] = 2;
    handCardsMap[32222] = 1;
    handCardsMap[32223] = 0;
    handCardsMap[32224] = 2;
    handCardsMap[32230] = 2;
    handCardsMap[32231] = 1;
    handCardsMap[32232] = 3;
    handCardsMap[32233] = 2;
    handCardsMap[32234] = 1;
    handCardsMap[32240] = 1;
    handCardsMap[32241] = 3;
    handCardsMap[32242] = 2;
    handCardsMap[32243] = 1;
    handCardsMap[32300] = 1;
    handCardsMap[32301] = 3;
    handCardsMap[32302] = 2;
    handCardsMap[32303] = 1;
    handCardsMap[32304] = 3;
    handCardsMap[32310] = 3;
    handCardsMap[32311] = 2;
    handCardsMap[32312] = 4;
    handCardsMap[32313] = 3;
    handCardsMap[32314] = 2;
    handCardsMap[32320] = 2;
    handCardsMap[32321] = 1;
    handCardsMap[32322] = 3;
    handCardsMap[32323] = 2;
    handCardsMap[32324] = 1;
    handCardsMap[32330] = 1;
    handCardsMap[32331] = 0;
    handCardsMap[32332] = 2;
    handCardsMap[32333] = 1;
    handCardsMap[32340] = 3;
    handCardsMap[32341] = 2;
    handCardsMap[32342] = 1;
    handCardsMap[32400] = 3;
    handCardsMap[32401] = 2;
    handCardsMap[32402] = 4;
    handCardsMap[32403] = 3;
    handCardsMap[32404] = 2;
    handCardsMap[32410] = 2;
    handCardsMap[32411] = 1;
    handCardsMap[32412] = 3;
    handCardsMap[32413] = 2;
    handCardsMap[32414] = 1;
    handCardsMap[32420] = 1;
    handCardsMap[32421] = 3;
    handCardsMap[32422] = 2;
    handCardsMap[32423] = 1;
    handCardsMap[32430] = 3;
    handCardsMap[32431] = 2;
    handCardsMap[32432] = 1;
    handCardsMap[32440] = 2;
    handCardsMap[32441] = 1;
    handCardsMap[33000] = 0;
    handCardsMap[33001] = 2;
    handCardsMap[33002] = 1;
    handCardsMap[33003] = 0;
    handCardsMap[33004] = 2;
    handCardsMap[33010] = 2;
    handCardsMap[33011] = 1;
    handCardsMap[33012] = 3;
    handCardsMap[33013] = 2;
    handCardsMap[33014] = 1;
    handCardsMap[33020] = 1;
    handCardsMap[33021] = 3;
    handCardsMap[33022] = 2;
    handCardsMap[33023] = 1;
    handCardsMap[33024] = 3;
    handCardsMap[33030] = 0;
    handCardsMap[33031] = 2;
    handCardsMap[33032] = 1;
    handCardsMap[33033] = 0;
    handCardsMap[33034] = 2;
    handCardsMap[33040] = 2;
    handCardsMap[33041] = 1;
    handCardsMap[33042] = 3;
    handCardsMap[33043] = 2;
    handCardsMap[33044] = 1;
    handCardsMap[33100] = 2;
    handCardsMap[33101] = 1;
    handCardsMap[33102] = 3;
    handCardsMap[33103] = 2;
    handCardsMap[33104] = 1;
    handCardsMap[33110] = 1;
    handCardsMap[33111] = 0;
    handCardsMap[33112] = 2;
    handCardsMap[33113] = 1;
    handCardsMap[33114] = 0;
    handCardsMap[33120] = 3;
    handCardsMap[33121] = 2;
    handCardsMap[33122] = 1;
    handCardsMap[33123] = 3;
    handCardsMap[33124] = 2;
    handCardsMap[33130] = 2;
    handCardsMap[33131] = 1;
    handCardsMap[33132] = 3;
    handCardsMap[33133] = 2;
    handCardsMap[33134] = 1;
    handCardsMap[33140] = 1;
    handCardsMap[33141] = 0;
    handCardsMap[33142] = 2;
    handCardsMap[33143] = 1;
    handCardsMap[33200] = 1;
    handCardsMap[33201] = 3;
    handCardsMap[33202] = 2;
    handCardsMap[33203] = 1;
    handCardsMap[33204] = 3;
    handCardsMap[33210] = 3;
    handCardsMap[33211] = 2;
    handCardsMap[33212] = 1;
    handCardsMap[33213] = 3;
    handCardsMap[33214] = 2;
    handCardsMap[33220] = 2;
    handCardsMap[33221] = 1;
    handCardsMap[33222] = 0;
    handCardsMap[33223] = 2;
    handCardsMap[33224] = 1;
    handCardsMap[33230] = 1;
    handCardsMap[33231] = 3;
    handCardsMap[33232] = 2;
    handCardsMap[33233] = 1;
    handCardsMap[33240] = 3;
    handCardsMap[33241] = 2;
    handCardsMap[33242] = 1;
    handCardsMap[33300] = 0;
    handCardsMap[33301] = 2;
    handCardsMap[33302] = 1;
    handCardsMap[33303] = 0;
    handCardsMap[33304] = 2;
    handCardsMap[33310] = 2;
    handCardsMap[33311] = 1;
    handCardsMap[33312] = 3;
    handCardsMap[33313] = 2;
    handCardsMap[33314] = 1;
    handCardsMap[33320] = 1;
    handCardsMap[33321] = 3;
    handCardsMap[33322] = 2;
    handCardsMap[33323] = 1;
    handCardsMap[33330] = 0;
    handCardsMap[33331] = 2;
    handCardsMap[33332] = 1;
    handCardsMap[33340] = 2;
    handCardsMap[33341] = 1;
    handCardsMap[33400] = 2;
    handCardsMap[33401] = 1;
    handCardsMap[33402] = 3;
    handCardsMap[33403] = 2;
    handCardsMap[33404] = 1;
    handCardsMap[33410] = 1;
    handCardsMap[33411] = 0;
    handCardsMap[33412] = 2;
    handCardsMap[33413] = 1;
    handCardsMap[33420] = 3;
    handCardsMap[33421] = 2;
    handCardsMap[33422] = 1;
    handCardsMap[33430] = 2;
    handCardsMap[33431] = 1;
    handCardsMap[33440] = 1;
    handCardsMap[34000] = 2;
    handCardsMap[34001] = 4;
    handCardsMap[34002] = 3;
    handCardsMap[34003] = 2;
    handCardsMap[34004] = 4;
    handCardsMap[34010] = 1;
    handCardsMap[34011] = 3;
    handCardsMap[34012] = 2;
    handCardsMap[34013] = 1;
    handCardsMap[34014] = 3;
    handCardsMap[34020] = 3;
    handCardsMap[34021] = 2;
    handCardsMap[34022] = 4;
    handCardsMap[34023] = 3;
    handCardsMap[34024] = 2;
    handCardsMap[34030] = 2;
    handCardsMap[34031] = 4;
    handCardsMap[34032] = 3;
    handCardsMap[34033] = 2;
    handCardsMap[34034] = 4;
    handCardsMap[34040] = 1;
    handCardsMap[34041] = 3;
    handCardsMap[34042] = 2;
    handCardsMap[34043] = 1;
    handCardsMap[34100] = 1;
    handCardsMap[34101] = 3;
    handCardsMap[34102] = 2;
    handCardsMap[34103] = 1;
    handCardsMap[34104] = 3;
    handCardsMap[34110] = 0;
    handCardsMap[34111] = 2;
    handCardsMap[34112] = 1;
    handCardsMap[34113] = 0;
    handCardsMap[34114] = 2;
    handCardsMap[34120] = 2;
    handCardsMap[34121] = 1;
    handCardsMap[34122] = 3;
    handCardsMap[34123] = 2;
    handCardsMap[34124] = 1;
    handCardsMap[34130] = 1;
    handCardsMap[34131] = 3;
    handCardsMap[34132] = 2;
    handCardsMap[34133] = 1;
    handCardsMap[34140] = 0;
    handCardsMap[34141] = 2;
    handCardsMap[34142] = 1;
    handCardsMap[34200] = 3;
    handCardsMap[34201] = 2;
    handCardsMap[34202] = 4;
    handCardsMap[34203] = 3;
    handCardsMap[34204] = 2;
    handCardsMap[34210] = 2;
    handCardsMap[34211] = 1;
    handCardsMap[34212] = 3;
    handCardsMap[34213] = 2;
    handCardsMap[34214] = 1;
    handCardsMap[34220] = 1;
    handCardsMap[34221] = 0;
    handCardsMap[34222] = 2;
    handCardsMap[34223] = 1;
    handCardsMap[34230] = 3;
    handCardsMap[34231] = 2;
    handCardsMap[34232] = 1;
    handCardsMap[34240] = 2;
    handCardsMap[34241] = 1;
    handCardsMap[34300] = 2;
    handCardsMap[34301] = 4;
    handCardsMap[34302] = 3;
    handCardsMap[34303] = 2;
    handCardsMap[34304] = 4;
    handCardsMap[34310] = 1;
    handCardsMap[34311] = 3;
    handCardsMap[34312] = 2;
    handCardsMap[34313] = 1;
    handCardsMap[34320] = 3;
    handCardsMap[34321] = 2;
    handCardsMap[34322] = 1;
    handCardsMap[34330] = 2;
    handCardsMap[34331] = 1;
    handCardsMap[34340] = 1;
    handCardsMap[34400] = 1;
    handCardsMap[34401] = 3;
    handCardsMap[34402] = 2;
    handCardsMap[34403] = 1;
    handCardsMap[34410] = 0;
    handCardsMap[34411] = 2;
    handCardsMap[34412] = 1;
    handCardsMap[34420] = 2;
    handCardsMap[34421] = 1;
    handCardsMap[34430] = 1;
    handCardsMap[40000] = 2;
    handCardsMap[40001] = 4;
    handCardsMap[40002] = 3;
    handCardsMap[40003] = 2;
    handCardsMap[40004] = 4;
    handCardsMap[40010] = 4;
    handCardsMap[40011] = 3;
    handCardsMap[40012] = 5;
    handCardsMap[40013] = 4;
    handCardsMap[40014] = 3;
    handCardsMap[40020] = 3;
    handCardsMap[40021] = 5;
    handCardsMap[40022] = 4;
    handCardsMap[40023] = 3;
    handCardsMap[40024] = 5;
    handCardsMap[40030] = 2;
    handCardsMap[40031] = 4;
    handCardsMap[40032] = 3;
    handCardsMap[40033] = 2;
    handCardsMap[40034] = 4;
    handCardsMap[40040] = 4;
    handCardsMap[40041] = 3;
    handCardsMap[40042] = 5;
    handCardsMap[40043] = 4;
    handCardsMap[40044] = 3;
    handCardsMap[40100] = 1;
    handCardsMap[40101] = 3;
    handCardsMap[40102] = 2;
    handCardsMap[40103] = 1;
    handCardsMap[40104] = 3;
    handCardsMap[40110] = 3;
    handCardsMap[40111] = 2;
    handCardsMap[40112] = 4;
    handCardsMap[40113] = 3;
    handCardsMap[40114] = 2;
    handCardsMap[40120] = 2;
    handCardsMap[40121] = 4;
    handCardsMap[40122] = 3;
    handCardsMap[40123] = 2;
    handCardsMap[40124] = 4;
    handCardsMap[40130] = 1;
    handCardsMap[40131] = 3;
    handCardsMap[40132] = 2;
    handCardsMap[40133] = 1;
    handCardsMap[40134] = 3;
    handCardsMap[40140] = 3;
    handCardsMap[40141] = 2;
    handCardsMap[40142] = 4;
    handCardsMap[40143] = 3;
    handCardsMap[40144] = 2;
    handCardsMap[40200] = 3;
    handCardsMap[40201] = 2;
    handCardsMap[40202] = 4;
    handCardsMap[40203] = 3;
    handCardsMap[40204] = 2;
    handCardsMap[40210] = 2;
    handCardsMap[40211] = 1;
    handCardsMap[40212] = 3;
    handCardsMap[40213] = 2;
    handCardsMap[40214] = 1;
    handCardsMap[40220] = 4;
    handCardsMap[40221] = 3;
    handCardsMap[40222] = 2;
    handCardsMap[40223] = 4;
    handCardsMap[40224] = 3;
    handCardsMap[40230] = 3;
    handCardsMap[40231] = 2;
    handCardsMap[40232] = 4;
    handCardsMap[40233] = 3;
    handCardsMap[40234] = 2;
    handCardsMap[40240] = 2;
    handCardsMap[40241] = 1;
    handCardsMap[40242] = 3;
    handCardsMap[40243] = 2;
    handCardsMap[40244] = 1;
    handCardsMap[40300] = 2;
    handCardsMap[40301] = 4;
    handCardsMap[40302] = 3;
    handCardsMap[40303] = 2;
    handCardsMap[40304] = 4;
    handCardsMap[40310] = 4;
    handCardsMap[40311] = 3;
    handCardsMap[40312] = 2;
    handCardsMap[40313] = 4;
    handCardsMap[40314] = 3;
    handCardsMap[40320] = 3;
    handCardsMap[40321] = 2;
    handCardsMap[40322] = 1;
    handCardsMap[40323] = 3;
    handCardsMap[40324] = 2;
    handCardsMap[40330] = 2;
    handCardsMap[40331] = 4;
    handCardsMap[40332] = 3;
    handCardsMap[40333] = 2;
    handCardsMap[40334] = 4;
    handCardsMap[40340] = 4;
    handCardsMap[40341] = 3;
    handCardsMap[40342] = 2;
    handCardsMap[40343] = 4;
    handCardsMap[40400] = 1;
    handCardsMap[40401] = 3;
    handCardsMap[40402] = 2;
    handCardsMap[40403] = 1;
    handCardsMap[40404] = 3;
    handCardsMap[40410] = 3;
    handCardsMap[40411] = 2;
    handCardsMap[40412] = 4;
    handCardsMap[40413] = 3;
    handCardsMap[40414] = 2;
    handCardsMap[40420] = 2;
    handCardsMap[40421] = 4;
    handCardsMap[40422] = 3;
    handCardsMap[40423] = 2;
    handCardsMap[40424] = 4;
    handCardsMap[40430] = 1;
    handCardsMap[40431] = 3;
    handCardsMap[40432] = 2;
    handCardsMap[40433] = 1;
    handCardsMap[40440] = 3;
    handCardsMap[40441] = 2;
    handCardsMap[40442] = 4;
    handCardsMap[41000] = 1;
    handCardsMap[41001] = 3;
    handCardsMap[41002] = 2;
    handCardsMap[41003] = 1;
    handCardsMap[41004] = 3;
    handCardsMap[41010] = 3;
    handCardsMap[41011] = 2;
    handCardsMap[41012] = 4;
    handCardsMap[41013] = 3;
    handCardsMap[41014] = 2;
    handCardsMap[41020] = 2;
    handCardsMap[41021] = 4;
    handCardsMap[41022] = 3;
    handCardsMap[41023] = 2;
    handCardsMap[41024] = 4;
    handCardsMap[41030] = 1;
    handCardsMap[41031] = 3;
    handCardsMap[41032] = 2;
    handCardsMap[41033] = 1;
    handCardsMap[41034] = 3;
    handCardsMap[41040] = 3;
    handCardsMap[41041] = 2;
    handCardsMap[41042] = 4;
    handCardsMap[41043] = 3;
    handCardsMap[41044] = 2;
    handCardsMap[41100] = 0;
    handCardsMap[41101] = 2;
    handCardsMap[41102] = 1;
    handCardsMap[41103] = 0;
    handCardsMap[41104] = 2;
    handCardsMap[41110] = 2;
    handCardsMap[41111] = 1;
    handCardsMap[41112] = 3;
    handCardsMap[41113] = 2;
    handCardsMap[41114] = 1;
    handCardsMap[41120] = 1;
    handCardsMap[41121] = 3;
    handCardsMap[41122] = 2;
    handCardsMap[41123] = 1;
    handCardsMap[41124] = 3;
    handCardsMap[41130] = 0;
    handCardsMap[41131] = 2;
    handCardsMap[41132] = 1;
    handCardsMap[41133] = 0;
    handCardsMap[41134] = 2;
    handCardsMap[41140] = 2;
    handCardsMap[41141] = 1;
    handCardsMap[41142] = 3;
    handCardsMap[41143] = 2;
    handCardsMap[41144] = 1;
    handCardsMap[41200] = 2;
    handCardsMap[41201] = 1;
    handCardsMap[41202] = 3;
    handCardsMap[41203] = 2;
    handCardsMap[41204] = 1;
    handCardsMap[41210] = 1;
    handCardsMap[41211] = 0;
    handCardsMap[41212] = 2;
    handCardsMap[41213] = 1;
    handCardsMap[41214] = 0;
    handCardsMap[41220] = 3;
    handCardsMap[41221] = 2;
    handCardsMap[41222] = 1;
    handCardsMap[41223] = 3;
    handCardsMap[41224] = 2;
    handCardsMap[41230] = 2;
    handCardsMap[41231] = 1;
    handCardsMap[41232] = 3;
    handCardsMap[41233] = 2;
    handCardsMap[41234] = 1;
    handCardsMap[41240] = 1;
    handCardsMap[41241] = 0;
    handCardsMap[41242] = 2;
    handCardsMap[41243] = 1;
    handCardsMap[41300] = 1;
    handCardsMap[41301] = 3;
    handCardsMap[41302] = 2;
    handCardsMap[41303] = 1;
    handCardsMap[41304] = 3;
    handCardsMap[41310] = 3;
    handCardsMap[41311] = 2;
    handCardsMap[41312] = 1;
    handCardsMap[41313] = 3;
    handCardsMap[41314] = 2;
    handCardsMap[41320] = 2;
    handCardsMap[41321] = 1;
    handCardsMap[41322] = 0;
    handCardsMap[41323] = 2;
    handCardsMap[41324] = 1;
    handCardsMap[41330] = 1;
    handCardsMap[41331] = 3;
    handCardsMap[41332] = 2;
    handCardsMap[41333] = 1;
    handCardsMap[41340] = 3;
    handCardsMap[41341] = 2;
    handCardsMap[41342] = 1;
    handCardsMap[41400] = 0;
    handCardsMap[41401] = 2;
    handCardsMap[41402] = 1;
    handCardsMap[41403] = 0;
    handCardsMap[41404] = 2;
    handCardsMap[41410] = 2;
    handCardsMap[41411] = 1;
    handCardsMap[41412] = 3;
    handCardsMap[41413] = 2;
    handCardsMap[41414] = 1;
    handCardsMap[41420] = 1;
    handCardsMap[41421] = 3;
    handCardsMap[41422] = 2;
    handCardsMap[41423] = 1;
    handCardsMap[41430] = 0;
    handCardsMap[41431] = 2;
    handCardsMap[41432] = 1;
    handCardsMap[41440] = 2;
    handCardsMap[41441] = 1;
    handCardsMap[42000] = 3;
    handCardsMap[42001] = 5;
    handCardsMap[42002] = 4;
    handCardsMap[42003] = 3;
    handCardsMap[42004] = 5;
    handCardsMap[42010] = 2;
    handCardsMap[42011] = 4;
    handCardsMap[42012] = 3;
    handCardsMap[42013] = 2;
    handCardsMap[42014] = 4;
    handCardsMap[42020] = 4;
    handCardsMap[42021] = 3;
    handCardsMap[42022] = 5;
    handCardsMap[42023] = 4;
    handCardsMap[42024] = 3;
    handCardsMap[42030] = 3;
    handCardsMap[42031] = 5;
    handCardsMap[42032] = 4;
    handCardsMap[42033] = 3;
    handCardsMap[42034] = 5;
    handCardsMap[42040] = 2;
    handCardsMap[42041] = 4;
    handCardsMap[42042] = 3;
    handCardsMap[42043] = 2;
    handCardsMap[42044] = 4;
    handCardsMap[42100] = 2;
    handCardsMap[42101] = 4;
    handCardsMap[42102] = 3;
    handCardsMap[42103] = 2;
    handCardsMap[42104] = 4;
    handCardsMap[42110] = 1;
    handCardsMap[42111] = 3;
    handCardsMap[42112] = 2;
    handCardsMap[42113] = 1;
    handCardsMap[42114] = 3;
    handCardsMap[42120] = 3;
    handCardsMap[42121] = 2;
    handCardsMap[42122] = 4;
    handCardsMap[42123] = 3;
    handCardsMap[42124] = 2;
    handCardsMap[42130] = 2;
    handCardsMap[42131] = 4;
    handCardsMap[42132] = 3;
    handCardsMap[42133] = 2;
    handCardsMap[42134] = 4;
    handCardsMap[42140] = 1;
    handCardsMap[42141] = 3;
    handCardsMap[42142] = 2;
    handCardsMap[42143] = 1;
    handCardsMap[42200] = 1;
    handCardsMap[42201] = 3;
    handCardsMap[42202] = 2;
    handCardsMap[42203] = 1;
    handCardsMap[42204] = 3;
    handCardsMap[42210] = 0;
    handCardsMap[42211] = 2;
    handCardsMap[42212] = 1;
    handCardsMap[42213] = 0;
    handCardsMap[42214] = 2;
    handCardsMap[42220] = 2;
    handCardsMap[42221] = 1;
    handCardsMap[42222] = 3;
    handCardsMap[42223] = 2;
    handCardsMap[42224] = 1;
    handCardsMap[42230] = 1;
    handCardsMap[42231] = 3;
    handCardsMap[42232] = 2;
    handCardsMap[42233] = 1;
    handCardsMap[42240] = 0;
    handCardsMap[42241] = 2;
    handCardsMap[42242] = 1;
    handCardsMap[42300] = 3;
    handCardsMap[42301] = 2;
    handCardsMap[42302] = 4;
    handCardsMap[42303] = 3;
    handCardsMap[42304] = 2;
    handCardsMap[42310] = 2;
    handCardsMap[42311] = 1;
    handCardsMap[42312] = 3;
    handCardsMap[42313] = 2;
    handCardsMap[42314] = 1;
    handCardsMap[42320] = 1;
    handCardsMap[42321] = 0;
    handCardsMap[42322] = 2;
    handCardsMap[42323] = 1;
    handCardsMap[42330] = 3;
    handCardsMap[42331] = 2;
    handCardsMap[42332] = 1;
    handCardsMap[42340] = 2;
    handCardsMap[42341] = 1;
    handCardsMap[42400] = 2;
    handCardsMap[42401] = 4;
    handCardsMap[42402] = 3;
    handCardsMap[42403] = 2;
    handCardsMap[42404] = 4;
    handCardsMap[42410] = 1;
    handCardsMap[42411] = 3;
    handCardsMap[42412] = 2;
    handCardsMap[42413] = 1;
    handCardsMap[42420] = 3;
    handCardsMap[42421] = 2;
    handCardsMap[42422] = 1;
    handCardsMap[42430] = 2;
    handCardsMap[42431] = 1;
    handCardsMap[42440] = 1;
    handCardsMap[43000] = 2;
    handCardsMap[43001] = 4;
    handCardsMap[43002] = 3;
    handCardsMap[43003] = 2;
    handCardsMap[43004] = 4;
    handCardsMap[43010] = 4;
    handCardsMap[43011] = 3;
    handCardsMap[43012] = 5;
    handCardsMap[43013] = 4;
    handCardsMap[43014] = 3;
    handCardsMap[43020] = 3;
    handCardsMap[43021] = 5;
    handCardsMap[43022] = 4;
    handCardsMap[43023] = 3;
    handCardsMap[43024] = 5;
    handCardsMap[43030] = 2;
    handCardsMap[43031] = 4;
    handCardsMap[43032] = 3;
    handCardsMap[43033] = 2;
    handCardsMap[43034] = 4;
    handCardsMap[43040] = 4;
    handCardsMap[43041] = 3;
    handCardsMap[43042] = 5;
    handCardsMap[43043] = 4;
    handCardsMap[43100] = 1;
    handCardsMap[43101] = 3;
    handCardsMap[43102] = 2;
    handCardsMap[43103] = 1;
    handCardsMap[43104] = 3;
    handCardsMap[43110] = 3;
    handCardsMap[43111] = 2;
    handCardsMap[43112] = 4;
    handCardsMap[43113] = 3;
    handCardsMap[43114] = 2;
    handCardsMap[43120] = 2;
    handCardsMap[43121] = 4;
    handCardsMap[43122] = 3;
    handCardsMap[43123] = 2;
    handCardsMap[43124] = 4;
    handCardsMap[43130] = 1;
    handCardsMap[43131] = 3;
    handCardsMap[43132] = 2;
    handCardsMap[43133] = 1;
    handCardsMap[43140] = 3;
    handCardsMap[43141] = 2;
    handCardsMap[43142] = 4;
    handCardsMap[43200] = 3;
    handCardsMap[43201] = 2;
    handCardsMap[43202] = 4;
    handCardsMap[43203] = 3;
    handCardsMap[43204] = 2;
    handCardsMap[43210] = 2;
    handCardsMap[43211] = 1;
    handCardsMap[43212] = 3;
    handCardsMap[43213] = 2;
    handCardsMap[43214] = 1;
    handCardsMap[43220] = 1;
    handCardsMap[43221] = 3;
    handCardsMap[43222] = 2;
    handCardsMap[43223] = 1;
    handCardsMap[43230] = 3;
    handCardsMap[43231] = 2;
    handCardsMap[43232] = 4;
    handCardsMap[43240] = 2;
    handCardsMap[43241] = 1;
    handCardsMap[43300] = 2;
    handCardsMap[43301] = 4;
    handCardsMap[43302] = 3;
    handCardsMap[43303] = 2;
    handCardsMap[43304] = 4;
    handCardsMap[43310] = 1;
    handCardsMap[43311] = 3;
    handCardsMap[43312] = 2;
    handCardsMap[43313] = 1;
    handCardsMap[43320] = 0;
    handCardsMap[43321] = 2;
    handCardsMap[43322] = 1;
    handCardsMap[43330] = 2;
    handCardsMap[43331] = 1;
    handCardsMap[43340] = 1;
    handCardsMap[43400] = 1;
    handCardsMap[43401] = 3;
    handCardsMap[43402] = 2;
    handCardsMap[43403] = 1;
    handCardsMap[43410] = 3;
    handCardsMap[43411] = 2;
    handCardsMap[43412] = 4;
    handCardsMap[43420] = 2;
    handCardsMap[43421] = 1;
    handCardsMap[43430] = 1;
    handCardsMap[44000] = 1;
    handCardsMap[44001] = 3;
    handCardsMap[44002] = 2;
    handCardsMap[44003] = 1;
    handCardsMap[44004] = 3;
    handCardsMap[44010] = 3;
    handCardsMap[44011] = 2;
    handCardsMap[44012] = 4;
    handCardsMap[44013] = 3;
    handCardsMap[44014] = 2;
    handCardsMap[44020] = 2;
    handCardsMap[44021] = 4;
    handCardsMap[44022] = 3;
    handCardsMap[44023] = 2;
    handCardsMap[44024] = 4;
    handCardsMap[44030] = 1;
    handCardsMap[44031] = 3;
    handCardsMap[44032] = 2;
    handCardsMap[44033] = 1;
    handCardsMap[44040] = 3;
    handCardsMap[44041] = 2;
    handCardsMap[44042] = 4;
    handCardsMap[44100] = 0;
    handCardsMap[44101] = 2;
    handCardsMap[44102] = 1;
    handCardsMap[44103] = 0;
    handCardsMap[44104] = 2;
    handCardsMap[44110] = 2;
    handCardsMap[44111] = 1;
    handCardsMap[44112] = 3;
    handCardsMap[44113] = 2;
    handCardsMap[44114] = 1;
    handCardsMap[44120] = 1;
    handCardsMap[44121] = 3;
    handCardsMap[44122] = 2;
    handCardsMap[44123] = 1;
    handCardsMap[44130] = 0;
    handCardsMap[44131] = 2;
    handCardsMap[44132] = 1;
    handCardsMap[44140] = 2;
    handCardsMap[44141] = 1;
    handCardsMap[44200] = 2;
    handCardsMap[44201] = 1;
    handCardsMap[44202] = 3;
    handCardsMap[44203] = 2;
    handCardsMap[44204] = 1;
    handCardsMap[44210] = 1;
    handCardsMap[44211] = 0;
    handCardsMap[44212] = 2;
    handCardsMap[44213] = 1;
    handCardsMap[44220] = 3;
    handCardsMap[44221] = 2;
    handCardsMap[44222] = 1;
    handCardsMap[44230] = 2;
    handCardsMap[44231] = 1;
    handCardsMap[44240] = 1;
    handCardsMap[44300] = 1;
    handCardsMap[44301] = 3;
    handCardsMap[44302] = 2;
    handCardsMap[44303] = 1;
    handCardsMap[44310] = 3;
    handCardsMap[44311] = 2;
    handCardsMap[44312] = 1;
    handCardsMap[44320] = 2;
    handCardsMap[44321] = 1;
    handCardsMap[44330] = 1;
    handCardsMap[44400] = 0;
    handCardsMap[44401] = 2;
    handCardsMap[44402] = 1;
    handCardsMap[44410] = 2;
    handCardsMap[44411] = 1;
    handCardsMap[44420] = 1;

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
    while (iStart < 8)
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

        set<long long>::iterator iterCur = curset.begin();
        while (iterCur != curset.end())
        {
            handCardsMap.insert(std::pair<long long, long long>(*iterCur, GetMinLaiZi(*iterCur)));
            ++iterCur;
        }

        vec.push_back(curset);
        iStart++;
    }

    vShowVec(vec);
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
            printf("handCardsMap[%lld] = %lld;\n", *iterSet, GetMinLaiZi(*iterSet));
            ++iterSet;
        }
        printf("\n\n");
        ++iterVec;
    }
    printf("### vShowVec End\n");
}
