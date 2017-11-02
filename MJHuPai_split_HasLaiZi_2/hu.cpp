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

void vGetJiang(int iCardsNum[MAX_CARD_ARRAY_SIZE], Card aucJiang[MAX_HANDCARD_NUM], int & riJiangNum)
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

/*
@return:
true: 返回true代表不需要赖子就可以构成顺子或者刻子。
false: 返回false代表需要赖子，需要的赖子数量是iNeedLaiZiNum。
*/
bool bTongHuaSeCanBePu(int iCardsNum[MAX_CARD_ARRAY_SIZE], int iStartIndex, int iEndIndex, int & iNeedLaiZiNum)
{
	bool bNeedLaiZi = false;
	iNeedLaiZiNum = 0;

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
			}
			else
			{
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
				if (i + 1 <= iEndIndex && iCardsNum[i + 1] > 0) iCardsNum[i + 1] = 0;
				if (i + 2 <= iEndIndex && iCardsNum[i + 2] > 0) iCardsNum[i + 2] = 0;
			}
			else if ((i + 1 <= iEndIndex && iCardsNum[i + 1] == 2 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 0) ||
				(i + 1 <= iEndIndex && iCardsNum[i + 1] == 0 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 2) ||
				(i + 1 == iEndIndex && iCardsNum[i + 1] == 2)
				)
			{
				iNeedLaiZiNum += 3;
				if (i + 1 <= iEndIndex && iCardsNum[i + 1] > 0) iCardsNum[i + 1] = 0;
				if (i + 2 <= iEndIndex && iCardsNum[i + 2] > 0) iCardsNum[i + 2] = 0;
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
                if (i + 1 <= iEndIndex && iCardsNum[i + 1] > 0) iCardsNum[i + 1] = 0;
                if (i + 2 <= iEndIndex && iCardsNum[i + 2] > 0) iCardsNum[i + 2] = 0;
            }
            else if ((i + 1 <= iEndIndex && iCardsNum[i + 1] == 2 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 0) ||
                (i + 1 <= iEndIndex && iCardsNum[i + 1] == 0 && i + 2 <= iEndIndex && iCardsNum[i + 2] == 2) ||
                (i + 1 == iEndIndex && iCardsNum[i + 1] == 2)
                )
            {
                iNeedLaiZiNum += 2;
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

    for (int i = iStartIndex; i < iEndIndex; i++)
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

bool bCanBePu(int iCardsNum[MAX_CARD_ARRAY_SIZE], Card ucLaiZi)
{
	/*计算赖子的数量，同时先把赖子移除。*/
	int iLaiZiNum = iCalcLaiZiNum(iCardsNum, ucLaiZi);

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

	/*找出可能做将的牌。在这里，默认赖子也可以做将牌。*/
	Card aucJiang[MAX_HANDCARD_NUM] = { 0 };
	int iJiangNum = 0;
	vGetJiang(iCardsNum, aucJiang, iJiangNum);

	/*遍历所有将牌的情况。*/
	for (int i = 0; i < iJiangNum; i++)
	{
		Card ucJiang = aucJiang[i];
		int iCardsNumNoJiang[MAX_CARD_ARRAY_SIZE] = { 0 };
		memcpy(iCardsNumNoJiang, iCardsNum, MAX_CARD_ARRAY_SIZE * sizeof(int));
		iCardsNumNoJiang[ucJiang] -= 2;

		/*判断扣除了将牌之后，剩余的牌能不能构成顺子(1万、2万、3万)或者刻子(1万、1万、1万)。*/
		if (bCanBePu(iCardsNumNoJiang, ucLaiZi))
		{
			return true;
		}
	}

	return false;
}
