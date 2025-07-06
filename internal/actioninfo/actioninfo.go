package actioninfo

import "fmt"

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i := 0; i < len(dataset); i++ {
		datastring := dataset[i]
		err := dp.Parse(datastring)
		if err != nil {
			fmt.Println("ошибка парсинга: ", err)
			continue
		}
		action, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("ошибка при ActionInfo(): ", err)
			continue
		}
		fmt.Println(action)
	}

}
