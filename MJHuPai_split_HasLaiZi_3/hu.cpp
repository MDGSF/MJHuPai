#include "hu.h"
#include <stdlib.h>
#include <string.h>


void vCalcCardsNum(Card aucHandCards[MAX_HANDCARD_NUM], int iCardsNum[MAX_CARD_ARRAY_SIZE])
{
	for (int i = 0; i < MAX_HANDCARD_NUM; i++)
	{
		iCardsNum[aucHandCards[i]]++;
	}
}

/*
@return:
true: 返回true代表不需要赖子就可以构成顺子或者刻子。
false: 返回false代表需要赖子，需要的赖子数量是iNeedLaiZiNum。
*/
bool bTongHuaSeCanBePu(int iOCardsNum[MAX_CARD_ARRAY_SIZE], int iStartIndex, int iEndIndex, int & iNeedLaiZiNum)
{
	bool bNeedLaiZi = false;
	iNeedLaiZiNum = 0;

    int iCardsNum[MAX_CARD_ARRAY_SIZE] = { 0 };
    memcpy(iCardsNum, iOCardsNum, MAX_CARD_ARRAY_SIZE * sizeof(int));

    //先把3张和4张的牌变成0张和1张。
    /*
    0 --> 0
    1 --> 1
    2 --> 2
    3 --> 0
    4 --> 1
    */
    for (int i = iStartIndex; i <= iEndIndex; i++)
    {
        int iNum = iCardsNum[i];
        if (iNum == 3)
        {
            iCardsNum[i] -= 3;
        }
        else if (iNum == 4)
        {
            iCardsNum[i] -= 3;
        }
    }

    //先处理1张的牌，处理离2张牌最远的1张牌。
    int iDis[MAX_CARD_ARRAY_SIZE] = { 0 };
    for (int i = iStartIndex; i <= iEndIndex; i++)
    {
        int iNum = iCardsNum[i];
        if (iNum != 2)
        {
            continue;
        }

        int iLeftIndex = i - 1;
        int iRightIndex = i + 1;
        int iLeftCount = 0;
        int iRightCount = 0;
        bool bLeftEnd = false;
        bool bRightEnd = false;
        while (!bLeftEnd || !bRightEnd)
        {
            if (!bLeftEnd && iLeftIndex >= iStartIndex && iCardsNum[iLeftIndex]<2)
            {
                if (iCardsNum[iLeftIndex] == 1)
                {
                    if (iDis[iLeftIndex] == 0)
                    {
                        iDis[iLeftIndex] = ++iLeftCount;
                    }
                    else if (iLeftCount + 1 < iDis[iLeftIndex])
                    {
                        iDis[iLeftIndex] = ++iLeftCount;
                    }
                    else
                    {
                        bLeftEnd = true;
                    }
                }
                else
                {
                    ++iLeftCount;
                }
            }
            else
            {
                bLeftEnd = true;
            }


            if (!bRightEnd && iRightIndex <= iEndIndex && iCardsNum[iRightIndex] < 2)
            {
                if (iCardsNum[iRightIndex] == 1)
                {
                    if (iDis[iRightIndex] == 0)
                    {
                        iDis[iRightIndex] = ++iRightCount;
                    }
                    else if (iRightCount + 1 < iDis[iRightIndex])
                    {
                        iDis[iRightIndex] = ++iRightCount;
                    }
                    else
                    {
                        bRightEnd = true;
                    }
                }
                else
                {
                    ++iRightCount;
                }
            }
            else
            {
                bRightEnd = true;
            }

            iLeftIndex--;
            iRightIndex++;
        }
    }

    //printf("11111111111111\n");

    int iNumCount = 9;
    while (iNumCount--)
    {
        bool bHasMoreThan3Dis = false;
        int iMaxDis = 0;
        int iMaxIndex = 0;

        for (int i = iStartIndex; i <= iEndIndex; i++)
        {
            if (iDis[i] >= 2)
            {
                bHasMoreThan3Dis = true;
                if (iDis[i] > iMaxDis)
                {
                    iMaxDis = iDis[i];
                    iMaxIndex = i;
                }
            }
        }

        if (!bHasMoreThan3Dis)
        {
            break;
        }

        if (iMaxIndex - 1 >= iStartIndex && iCardsNum[iMaxIndex - 1] >= 1 && iMaxIndex + 1 <= iEndIndex && iCardsNum[iMaxIndex + 1] >= 1)
        {
            iCardsNum[iMaxIndex - 1]--;
            iCardsNum[iMaxIndex]--;
            iCardsNum[iMaxIndex + 1]--;

            iDis[iMaxIndex - 1] = 0;
            iDis[iMaxIndex] = 0;
            iDis[iMaxIndex + 1] = 0;
        }
        else if (iMaxIndex - 1 >= iStartIndex && iCardsNum[iMaxIndex - 1] >= 1 && iMaxIndex - 2 >= iStartIndex && iCardsNum[iMaxIndex - 2] >= 1)
        {
            iCardsNum[iMaxIndex - 2]--;
            iCardsNum[iMaxIndex - 1]--;
            iCardsNum[iMaxIndex]--;

            iDis[iMaxIndex - 2] = 0;
            iDis[iMaxIndex - 1] = 0;
            iDis[iMaxIndex] = 0;
        }
        else if (iMaxIndex + 1 <= iEndIndex && iCardsNum[iMaxIndex + 1] >= 1 && iMaxIndex +2 <= iEndIndex && iCardsNum[iMaxIndex + 2] >= 1)
        {
            iCardsNum[iMaxIndex + 2]--;
            iCardsNum[iMaxIndex + 1]--;
            iCardsNum[iMaxIndex]--;

            iDis[iMaxIndex + 2] = 0;
            iDis[iMaxIndex + 1] = 0;
            iDis[iMaxIndex] = 0;
        }

        if(iMaxIndex - 2 >= iStartIndex) iDis[iMaxIndex - 2] = 0;
        if (iMaxIndex - 1 >= iStartIndex) iDis[iMaxIndex - 1] = 0;
        iDis[iMaxIndex] = 0;
        if (iMaxIndex + 1 <= iEndIndex) iDis[iMaxIndex + 1] = 0;
        if (iMaxIndex + 2 >= iEndIndex) iDis[iMaxIndex + 2] = 0;
    }


    //11111
	for (int i = iStartIndex; i <= iEndIndex; i++)
	{
		int iNum = iCardsNum[i];
		if (iNum == 1)
		{
			if (i + 1 <= iEndIndex && i + 2 <= iEndIndex && iCardsNum[i + 1] > 0 && iCardsNum[i + 2] > 0)
			{
				iCardsNum[i]--;
				iCardsNum[i + 1]--;
				iCardsNum[i + 2]--;
			}
			else
			{
                i += 2;
				bNeedLaiZi = true;
			}
		}
		else if (iNum == 2)
		{
			if (i + 1 <= iEndIndex && i + 2 <= iEndIndex && iCardsNum[i + 1] > 1 && iCardsNum[i + 2] > 1)
			{
				iCardsNum[i] -= 2;;
				iCardsNum[i + 1] -= 2;
				iCardsNum[i + 2] -= 2;
			}
			else if(i + 1 <= iEndIndex && i + 2 <= iEndIndex && iCardsNum[i + 1] > 0 && iCardsNum[i + 2] > 0)
			{
				iCardsNum[i]--;
				iCardsNum[i + 1]--;
				iCardsNum[i + 2]--;

                i += 2;
                bNeedLaiZi = true;
			}
			else
			{
                i += 2;
				bNeedLaiZi = true;
			}
		}
		else if (iNum == 3)
		{
			iCardsNum[i] -= 3;
		}
		else if (iNum == 4)
		{
			iCardsNum[i] -= 3;
			if (i + 1 <= iEndIndex && i + 2 <= iEndIndex && iCardsNum[i + 1] > 0 && iCardsNum[i + 2] > 0)
			{
				iCardsNum[i]--;
				iCardsNum[i + 1]--;
				iCardsNum[i + 2]--;
			}
			else
			{
                i += 2;
				bNeedLaiZi = true;
			}
		}
		else
		{
			//错误检查
		}
	}
	if (!bNeedLaiZi)
	{
		return true;
	}

    //22222
    for (int i = iStartIndex; i <= iEndIndex; i++)
    {
        int iNum = iCardsNum[i];
        if (iNum == 1)
        {
            if (i + 1 <= iEndIndex && i + 2 <= iEndIndex && iCardsNum[i + 1] > 0 && iCardsNum[i + 2] > 0)
            {
                iCardsNum[i]--;
                iCardsNum[i + 1]--;
                iCardsNum[i + 2]--;
            }
        }
        else if (iNum == 2)
        {
            if (i + 1 <= iEndIndex && i + 2 <= iEndIndex && iCardsNum[i + 1] > 1 && iCardsNum[i + 2] > 1)
            {
                iCardsNum[i] -= 2;;
                iCardsNum[i + 1] -= 2;
                iCardsNum[i + 2] -= 2;
            }
            else if (i + 1 <= iEndIndex && i + 2 <= iEndIndex && iCardsNum[i + 1] > 0 && iCardsNum[i + 2] > 0)
            {
                iCardsNum[i]--;
                iCardsNum[i + 1]--;
                iCardsNum[i + 2]--;
            }
        }
        else if (iNum == 3)
        {
            iCardsNum[i] -= 3;
        }
        else if (iNum == 4)
        {
            iCardsNum[i] -= 3;
            if (i + 1 <= iEndIndex && i + 2 <= iEndIndex && iCardsNum[i + 1] > 0 && iCardsNum[i + 2] > 0)
            {
                iCardsNum[i]--;
                iCardsNum[i + 1]--;
                iCardsNum[i + 2]--;
            }
        }
        else
        {
            //错误检查
        }
    }

    //33333
    for (int i = iStartIndex; i <= iEndIndex; i++)
    {
        int iNum = iCardsNum[i];
        if (iNum == 1)
        {
            if ((i + 1 <= iEndIndex && iCardsNum[i + 1] == 1 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 0) ||
                (i + 1 <= iEndIndex && iCardsNum[i + 1] == 0 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 1) ||
                (i + 1 == iEndIndex && iCardsNum[i + 1] == 1)
                )
            {
                iNeedLaiZiNum += 1;
                iCardsNum[i] = 0;
                if (i + 1 <= iEndIndex && iCardsNum[i + 1] > 0) iCardsNum[i + 1] = 0;
                if (i + 2 <= iEndIndex && iCardsNum[i + 2] > 0) iCardsNum[i + 2] = 0;
            }
        }
    }


	for (int i = iStartIndex; i <= iEndIndex; i++)
	{
		int iNum = iCardsNum[i];
		if (iNum == 1)
		{
			if ((i + 1 <= iEndIndex && iCardsNum[i + 1] == 0 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 0) ||
				(i + 1 == iEndIndex && iCardsNum[i + 1] == 0) ||
				(i == iEndIndex)
				)
			{
				iNeedLaiZiNum += 2;
			}
			else if ((i + 1 <= iEndIndex && iCardsNum[i + 1] == 1 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 0) ||
				(i + 1 <= iEndIndex && iCardsNum[i + 1] == 0 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 1) ||
				(i + 1 == iEndIndex && iCardsNum[i + 1] == 1)
				)
			{
				iNeedLaiZiNum += 1;
                iCardsNum[i] = 0;
				if (i + 1 <= iEndIndex && iCardsNum[i + 1] > 0) iCardsNum[i + 1] = 0;
				if (i + 2 <= iEndIndex && iCardsNum[i + 2] > 0) iCardsNum[i + 2] = 0;
			}
			else if ((i + 1 <= iEndIndex && iCardsNum[i + 1] == 2 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 0) ||
				(i + 1 <= iEndIndex && iCardsNum[i + 1] == 0 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 2) ||
				(i + 1 == iEndIndex && iCardsNum[i + 1] == 2)
				)
			{
				iNeedLaiZiNum += 1;
                iCardsNum[i] = 0;
				if (i + 1 <= iEndIndex && iCardsNum[i + 1] > 0) iCardsNum[i + 1]--;
				if (i + 2 <= iEndIndex && iCardsNum[i + 2] > 0) iCardsNum[i + 2]--;
			}
		}
		else if (iNum == 2)
		{
			if ((i + 1 <= iEndIndex && iCardsNum[i + 1] == 0 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 0) ||
				(i + 1 == iEndIndex && iCardsNum[i + 1] == 0) ||
				(i == iEndIndex)
				)
			{
				iNeedLaiZiNum += 1;
			}
            else if ((i + 1 <= iEndIndex && iCardsNum[i + 1] == 1 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 0) ||
                (i + 1 <= iEndIndex && iCardsNum[i + 1] == 0 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 1) ||
                (i + 1 == iEndIndex && iCardsNum[i + 1] == 1)
                )
            {
                iNeedLaiZiNum += 3;
                iCardsNum[i] = 0;
                if (i + 1 <= iEndIndex && iCardsNum[i + 1] > 0) iCardsNum[i + 1] = 0;
                if (i + 2 <= iEndIndex && iCardsNum[i + 2] > 0) iCardsNum[i + 2] = 0;
            }
            else if ((i + 1 <= iEndIndex && iCardsNum[i + 1] == 2 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 0) ||
                (i + 1 <= iEndIndex && iCardsNum[i + 1] == 0 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 2) ||
                (i + 1 == iEndIndex && iCardsNum[i + 1] == 2)
                )
            {
                iNeedLaiZiNum += 2;
                iCardsNum[i] = 0;
                if (i + 1 <= iEndIndex && iCardsNum[i + 1] > 0) iCardsNum[i + 1] = 0;
                if (i + 2 <= iEndIndex && iCardsNum[i + 2] > 0) iCardsNum[i + 2] = 0;
            }
		}
		else
		{
			//错误检查
		}
	}

	return false;
}

/*
@brief: 能否构成刻子。
*/
bool bTongHuaSeCanBeKeZi(int iCardsNum[MAX_CARD_ARRAY_SIZE], int iStartIndex, int iEndIndex, int & iNeedLaiZiNum)
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

int iCalcLaiZiNum(int iCardsNum[MAX_CARD_ARRAY_SIZE], Card ucLaiZi)
{
	int iNum = iCardsNum[ucLaiZi];
	iCardsNum[ucLaiZi] = 0;
	return iNum;
}

bool bCanBePu(int iCardsNum[MAX_CARD_ARRAY_SIZE], Card ucLaiZi, int iLaiZiNum)
{
	//判断万牌能不能构成顺子或者刻子
	int iNeedLaiZiNum = 0;
	if (!bTongHuaSeCanBePu(iCardsNum, CardYiWan, CardJiuWan, iNeedLaiZiNum))
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
		
	//判断条牌能不能构成顺子或者刻子
	iNeedLaiZiNum = 0;
	if (!bTongHuaSeCanBePu(iCardsNum, CardYaoJi, CardJiuTiao, iNeedLaiZiNum))
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

	//判断筒牌能不能构成顺子或者刻子
	if (!bTongHuaSeCanBePu(iCardsNum, CardYiTong, CardJiuTong, iNeedLaiZiNum))
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

bool bHu(Card aucHandCards[MAX_HANDCARD_NUM], Card ucLaiZi)
{
	/*计算出每张牌的张数。iCardsNum的下标代表每张牌，value就是这张牌的数量。*/
	int iCardsNum[MAX_CARD_ARRAY_SIZE] = { 0 };
	vCalcCardsNum(aucHandCards, iCardsNum);

    /*计算赖子的数量，同时先把赖子移除。*/
    int iLaiZiNum = iCalcLaiZiNum(iCardsNum, ucLaiZi);

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
