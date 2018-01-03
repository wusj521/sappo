package saprfc

import (
	"fmt"

	saprfc "github.com/sap/gorfc/gorfc"
	//saprfc "simonwaldherr.de/go/saprfc"
)

var SAPconnection *saprfc.Connection

func abapSystem() saprfc.ConnectionParameter {
	return saprfc.ConnectionParameter{
		Dest:   "aa",
		Client: "800",
		User:   "xxxx",
		Passwd: "xxxx",
		Lang:   "ZH",
		Ashost: "192.168.0.0",
		Sysnr:  "00",
		//		Saprouter: "/H/123.125.21.51/H/",
	}
}

func Connect() error {
	var err error
	fmt.Println(abapSystem()) //wusj
	SAPconnection, err = saprfc.ConnectionFromParams(abapSystem())

	return err
}

func Close() {
	SAPconnection.Close()
}

//
/*
func request() {
	params := map[string]interface{}{
		"WERKS":      "YK01",
		"MANTR_MARK": "X",
		//		"DELIMITER":   ";",
		//		"NO_DATA":     "",
		//		"ROWSKIPS":    0,
		//		"ROWCOUNT":    5,
	}
	r, err := SAPconnection.Call("ZRFC_PPBOM_PLC", params)
	if err != nil {
		fmt.Println(err)
		//return []string{}
	}

	//	var ret []string
	/*
		echoStruct := r["DATA"].([]interface{})
		//	echoStruct := r["FIELDS"].([]interface{})
		for _, value := range echoStruct {
			values := value.(map[string]interface{})
			for _, val := range values {
				valstr := strings.Split(fmt.Sprint("%s", val), ";")
				ret = append(ret, strings.TrimSpace(valstr[1]))
			}
		}
		return ret
*/
/*echoStruct := r["IMAKT"].([]interface{})
for _, value := range echoStruct {
	values := value.(map[string]interface{})
	//		fmt.Println(len(values)) //打印行数
	//		fmt.Println(values["MATNR"])//打印某个字段的值
	//	delete(values, "MAKTX")
	//		fmt.Println(values["MATNR"])
	//		fmt.Println(values["MATNR"])
	fmt.Println(values["MAKTX"])*/

/*
			for _, val := range values {
				delete(values, "MAKTX")
				//			fmt.Println(key)
				//			fmt.Println(val)

			}

	}

}

/*
func main() {
	Connect()

	//	user := request()
	request()
	//	for _, usr := range user {
	//		fmt.Println(usr)
	//	}

	Close()
}
*/
