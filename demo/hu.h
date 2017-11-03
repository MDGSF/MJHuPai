#pragma once

#include <stdio.h>
#include <iostream>
#include <set>
#include <map>
#include <vector>
#include <queue>
#include "public.h"

using namespace std;

struct Sub;

class CLaiZiHu
{
public:
    CLaiZiHu();
    ~CLaiZiHu();

    bool bHu(Card aucHandCards[MAX_HANDCARD_NUM], int iHandCardsLen, Card ucLaiZi);

    //要保证输入的只有01234
    long long GetMinLaiZi(long long iNum);

private:
    void vGetJiang(int iCardsNum[MAX_CARD_ARRAY_SIZE], Card aucJiang[MAX_HANDCARD_NUM], int & riJiangNum);

    bool bCanBePu(int iCardsNum[MAX_CARD_ARRAY_SIZE], Card ucLaiZi, int iLaiZiNum);

    long long ToNum(int iCardsNum[MAX_CARD_ARRAY_SIZE], int iStartIndex, int iEndIndex);

    void vShowVec(vector< set<long long> > & vec);

    bool bIsNumBitMoreThan14(long long iNum);

    void vGetNext(long long iNum, long long iNumBits, long long iInsertNum, set<long long> & curset);

    void vInit();

    int iGetNumBits(long long iNum);

    int iGetNumRightBit(long long iNum);

    int iGetNumLeftBit(long long iNum);

    void vReverseNumArray(long long arr[], int iArrLen);

    Sub SplitNum(long long iNum, int i);

    void CalcMainSub(long long iNum, Sub MainSub, vector<Sub> & subVec);

    void GetNumSub(long long iNum, vector<Sub> & subVec);

    bool bSplitWithTwoBlank(long long iNum, vector<long long> & setNum);

    /*
    @brief: 删除iNum最低位的零
    */
    void vTrimNumZero(long long & iNum);

    bool bTongHuaSeCanBeKeZi(int iCardsNum[MAX_CARD_ARRAY_SIZE], int iStartIndex, int iEndIndex, int & iNeedLaiZiNum);

    void vTest2();

    //<handcards, need LaiZi Num>
    map<long long, long long> handCardsMap;

    vector< set<long long> > vec;
};
