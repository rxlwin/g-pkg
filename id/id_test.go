package id

import (
	"fmt"
	"testing"
	"time"
)

func TestGetId(t *testing.T) {
	m := make(map[int64]int64)
	for i := 0; i < 100000; i++ {
		newId := GetId()
		fmt.Printf("newID %#b,%d\n", newId, newId)
		m[newId]++
	}
	for id, c := range m {
		if c > 1 {
			fmt.Println(id, c)
			return
		}
	}
	fmt.Println("没有重复")

}

func TestGetBaseTime(t *testing.T) {
	//最高时间差: 1111111111111111111111111111111111111111 1099511627775
	//c:=0b1111111111111111111111111111111111111111
	c := 0b1111111111111111111111111111111111111111
	fmt.Printf("最高时间差: %#b, \n%d\n", c, c)
	//取 2021-09-26 13:14:52 为基准时间 (时间戳: 1632633292000)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	baseTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-09-26 13:14:52", loc)
	fmt.Println(baseTime.UnixMilli())
	//最大时间戳 1633732803627775
	lastTime := 1099511627775 - baseTime.UnixMilli()
	fmt.Println(lastTime)
	//fmt.Printf("%#b,%#b, %d\n",lastTime<<23,lastTime, lastTime)
	//换算年月日 2056-07-30 09:08:39
	//lastTm := time.Unix(lastTime/1000,0)
	//fmt.Println(lastTm.Format("2006-01-02 15:04:05"))
	//就是说这样的位数可以使用到 2056-07-30 09:08:39
}

func TestTime(t *testing.T) {
	//t1 := math.MaxInt
	//fmt.Printf("%#b, %d\n", t1, t1)
	//t2 := 1632633292000
	//fmt.Printf("%#b\n", t2)
	//fmt.Printf("%#b\n", t2<<22)
	//fmt.Println(549755813887 - 1632633292000)

	t3 := 215517050000
	fmt.Printf("%08b, %d", t3, t3)
}

func TestGetId2(t *testing.T) {
	id := GetId()
	idStr := Int64ToStr(id)

	fmt.Println("idInt:", id)
	fmt.Println("idStr:", idStr)
}
