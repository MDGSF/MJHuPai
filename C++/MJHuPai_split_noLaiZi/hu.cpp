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

bool bTongHuaSeCanBePu(int iCardsNum[MAX_CARD_ARRAY_SIZE], int iStartIndex, int iEndIndex)
{
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
				return false;
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
			else
			{
				return false;
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
				return false;
			}
		}
		else
		{
			//错误检查
		}
	}

	return true;
}

bool bCanBePu(int iCardsNum[MAX_CARD_ARRAY_SIZE])
{
	if (!bTongHuaSeCanBePu(iCardsNum, CardYiWan, CardJiuWan)) return false; //判断万牌能不能构成顺子或者刻子
	if (!bTongHuaSeCanBePu(iCardsNum, CardYaoJi, CardJiuTiao)) return false; //判断条牌能不能构成顺子或者刻子
	if (!bTongHuaSeCanBePu(iCardsNum, CardYiTong, CardJiuTong)) return false; //判断筒牌能不能构成顺子或者刻子
	if (!bTongHuaSeCanBePu(iCardsNum, CardDONG, CardBEI)) return false; //判断风牌能不能构成顺子或者刻子
	if (!bTongHuaSeCanBePu(iCardsNum, CardZHONG, CardBAI)) return false; //判断箭牌能不能构成顺子或者刻子
	return true;
}

bool bHu(Card aucHandCards[MAX_HANDCARD_NUM])
{
	/*计算出每张牌的张数。iCardsNum的下标代表每张牌，value就是这张牌的数量。*/
	int iCardsNum[MAX_CARD_ARRAY_SIZE] = { 0 };
	vCalcCardsNum(aucHandCards, iCardsNum);

	/*找出可能做将的牌。*/
	Card aucJiang[MAX_HANDCARD_NUM] = { 0 };
	int iJiangNum = 0;
	vGetJiang(iCardsNum, aucJiang, iJiangNum);

	/*遍历所有将牌的情况。*/
	for (int i = 0; i < iJiangNum; i++)
	{
		Card ucJiang = aucJiang[i];
		int iCardsNumNoJiang[MAX_CARD_ARRAY_SIZE] = { 0 };
		memcpy(iCardsNumNoJiang, iCardsNum, MAX_CARD_ARRAY_SIZE);
		iCardsNumNoJiang[ucJiang] -= 2;

		/*判断扣除了将牌之后，剩余的牌能不能构成顺子(1万、2万、3万)或者刻子(1万、1万、1万)。*/
		if (bCanBePu(iCardsNumNoJiang))
		{
			return true;
		}
	}

	return false;
}
