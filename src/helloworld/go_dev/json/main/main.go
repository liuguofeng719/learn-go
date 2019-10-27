package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type RepayCard struct {
	BankCode  string   `json:"bank_code,omitempty"`
	BankName  string   `json:"bank_name,omitempty"`
	CardNo    string   `json:"card_no,omitempty"`
	Mobile    string   `json:"mobile,omitempty"`
	Name      string   `json:"name,omitempty"`
	CardType  string   `json:"card_type,omitempty"`
	UnionMark int      `json:"union_mark,omitempty"`
	IsBool    bool     `json:"is_bool,omitempty"`
	Monster   *Monster `json:"monster,omitempty"`
}

type Monster struct {
	Name     string  `json:"name,omitempty"`
	Age      int     `json:"age,omitempty"`
	Birthday string  `json:"birthday,omitempty"`
	Sal      float64 `json:"sal,omitempty"`
	Skill    string  `json:"skill,omitempty"`
}
type Project struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
	Docs string `json:"docs,omitempty"`
}

func main() {
	p1 := Project{
		Name: "hello name",
		Url:  "https://blog.csdn.net/qq_30505673",
	}

	data, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// Docs定义为omitempty所以不会出现Docs的字段
	fmt.Printf("%s\n", data)
	testStruct()
	// unMarshalStruct()
	// testSliceMap()
	// testMarshalMap()
	// testSliceMarshal()
	// testSliceUnMarshal()

	var month = Month{
		MonthNumber: 1,
		YearNumber:  2019,
	}

	month.UnmarshalJSON([]byte("2012/7"))
	fmt.Printf("Year=%v,Month=%v \n", month.YearNumber, month.MonthNumber)

	m, _ := month.MarshalJSON()
	fmt.Println(string(m))
}

type Month struct {
	MonthNumber int
	YearNumber  int
}

func (m Month) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d/%d", m.MonthNumber, m.YearNumber)), nil
}
func (m *Month) UnmarshalJSON(value []byte) error {
	parts := strings.Split(string(value), "/")
	MonthNumber, _ := strconv.ParseInt(parts[0], 10, 32)
	YearNumber, _ := strconv.ParseInt(parts[1], 10, 32)
	m.MonthNumber = int(MonthNumber)
	m.YearNumber = int(YearNumber)
	return nil
}

// json 转换注意事项
func testStruct() {

	repayCard := RepayCard{
		BankCode:  "28283123",
		Mobile:    "13649297908",
		Name:      "白健",
		CardType:  "credit",
		UnionMark: 10,
	}

	data, err := json.Marshal(repayCard)
	if err != nil {
		fmt.Printf("序列化错误%v", err)
	}

	fmt.Printf("repayCard = %s\n", string(data))
	// JSON格式化输出
	data1, err := json.MarshalIndent(&repayCard, "", "\t")
	if err != nil {
		fmt.Println("序列化错误")
	}
	fmt.Printf("MarshalIndent = %v \n", string(data1))

	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}
	data2, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("monster 序列化后=%v\n", string(data2))
}
func testSliceUnMarshal() {
	const jsonStr = `[{"code":"001","name":"中国银行","type":"credit"},{"code":"002","name":"建设银行","type":"debit"},{"code":"0020","name":"建设银行","type":"debit0"},{"code":"0021","name":"建设银行","type":"debit1"},{"code":"0022","name":"建设银行","type":"debit2"},{"code":"0023","name":"建设银行","type":"debit3"},{"code":"0024","name":"建设银行","type":"debit4"}]`

	var maps []map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &maps)

	for k, v := range maps {
		log.Printf("%v,%v \n", k, v)
	}
}

func testSliceMarshal() {
	var sliceCard []map[string]interface{}
	var map1 = make(map[string]interface{})
	map1["code"] = "001"
	map1["type"] = "credit"
	map1["name"] = "中国银行"
	sliceCard = append(sliceCard, map1)
	var map2 = make(map[string]interface{})
	map2["code"] = "002"
	map2["type"] = "debit"
	map2["name"] = "建设银行"
	sliceCard = append(sliceCard, map2)

	for i := 0; i < 5; i++ {
		map3 := make(map[string]interface{})
		map3["code"] = "002" + strconv.Itoa(i)
		map3["type"] = "debit" + fmt.Sprintf("%d", i)
		map3["name"] = "建设银行"
		sliceCard = append(sliceCard, map3)
	}

	jsonStr, err := json.Marshal(sliceCard)
	if err != nil {
		fmt.Println("序列化出错")
	}
	fmt.Printf("testSliceMarshal = %v \n", string(jsonStr))

}
func testMarshalMap() {
	// var keyStrObject map[string]interface{}
	var newMap = make(map[string]interface{})
	newMap["name"] = "zzhangsan"
	newMap["age"] = 20

	for k, v := range newMap {
		fmt.Printf("%v,%v\n", k, v)
	}
	mapStr, err := json.Marshal(newMap)
	if err != nil {
		fmt.Println("序列化错误")
	}
	fmt.Printf("map  = %v \n", string(mapStr))
}

func testSliceMap() {
	const sliceStr = `[{"name":"123"},{"age":"20"},{"address":"xxx"}]`
	var strs []map[string]string
	json.Unmarshal([]byte(sliceStr), &strs)
	for k, v := range strs {
		fmt.Printf("k = %v,v = %v\n", k, v)
	}
}

func unMarshalStruct() {
	const strJSON = `{
		"bankCode": "28283123",
		"BankName": "中国银行",
		"CardNo": "6228442021749717072",
		"Mobile": "13649297908",
		"Name": "白健",
		"CardType": "credit",
		"UnionMark": 0
	}`

	var repay RepayCard
	err := json.Unmarshal([]byte(strJSON), &repay)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("bankCode=%v\n", repay)
}
