#pragma once

#include "public.h"

#ifdef __cplusplus
extern "C" {
#endif

	/*
	@brief: 判断手牌能不能胡牌，符合3+3+3+3+2则胡牌，包含赖子(鬼牌)。
	@param aucHandCards[in]: 14张手牌，uint8_t类型在数组。
	@return bool: 
		true: 可以胡牌。
		false: 不可以胡牌。
	*/
	bool bHu(Card aucHandCards[MAX_HANDCARD_NUM], Card ucLaiZi);


#ifdef __cplusplus
}
#endif

