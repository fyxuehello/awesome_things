package tools

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

var ShopInfoCsvHeader = []string{"出资商家名称", "出资shop_id", "商家出资比例", "平台出资比例"}

// 生成csv文件
func GenCsvFile() error {
	// 不存在则创建;存在则清空;读写模式;
	file, err := os.Create("../files/shop_info.csv") //文件生成在test_file下，也可以指定路径
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
		return err
	}
	// 延迟关闭
	defer file.Close()

	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(file)

	// 写入表头
	_ = w.Write(ShopInfoCsvHeader)
	w.Flush()
	for i := 0; i < 300; i++ {
		info := []string{
			"测试商家名称" + fmt.Sprint(i),
			fmt.Sprint(rand.Intn(10000000)),
			fmt.Sprint(rand.Intn(100000000)),
			fmt.Sprint(rand.Intn(100000000)),
		}
		_ = w.Write(info)
		// 刷新缓冲
		w.Flush()
	}
	return nil
}

// 解析csv文件
func ParseCsvFile() error {
	file, err := os.Open("../files/shop_info.csv")
	defer file.Close()
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
		return err
	}
	r := csv.NewReader(file)
	csvFile, err := r.ReadAll()
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
		return err
	}
	data := csvFile[1:]
	for i := range data {
		fmt.Println(data[i])
	}
	return nil
}
