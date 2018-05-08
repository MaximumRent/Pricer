package unique_parsers

import (
	. "../entity"
	"../util"
	"strconv"
)

const _MINIMUM_PRICE_BORDER = 10000

func ExtractKufarPrice(value string, result *ParseResult) {
	var firstPosition int
	var secondPosition int
	for position, char := range value {
		if char == '.' {
			firstPosition = position + 1
		} else if char == '$' {
			secondPosition = position
		}
	}
	extractedPrice := value[firstPosition : secondPosition]
	if util.IsNonTrashString(extractedPrice) {
		price, err := strconv.Atoi(extractedPrice)
		if err == nil && price > _MINIMUM_PRICE_BORDER {
			result.Price = price
		}
	}
}
