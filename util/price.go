package util

import "strconv"

func PriceDbToExport(dbPrice int32) (resPrice string) {
	//fmt.Println("dbPrice", dbPrice)
	priceInt := dbPrice / 100
	//fmt.Println("priceInt", priceInt)
	priceDecimal1 := dbPrice - (dbPrice/10)*10
	//fmt.Println("priceDecimal1", priceDecimal1)
	priceDecimal2 := dbPrice/10 - (dbPrice/100)*10
	//fmt.Println("priceDecimal2", priceDecimal2)
	return Val2string(priceInt) + "." + Val2string(priceDecimal2) + Val2string(priceDecimal1)
}

func ScoreDbToExport(score int32) string {
	scoreInt := score / 10
	scoreDecimal := score - scoreInt*10
	return Val2string(scoreInt) + "." + Val2string(scoreDecimal)
}

func PriceExportToDb(exportPrice string) int32 {
	price, _ := strconv.ParseFloat(exportPrice, 32)
	return int32(price*100 + 0.5)
}
