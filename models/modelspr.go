package models

import (
	"fmt"
	//"os"
	//"path"
	"sappo/saprfc"
	//"sappo/utils"
	"strconv"
	//"time"

	//"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/lib/pq"
	//saprfc "github.com/sap/gorfc/gorfc"
)

//字体首写字母大写
type Sappr struct {
	Id    int64
	Banfn string `orm:"index"` //	采购申请号
	Bnfpo string //	采购申请的项目编号
	/////FRGKZ  string  //	批准标识
	Frgst    string  `orm:"index"` //	采购请求中的批准策略
	Ekgrp    string  //	采购组
	Ernam    string  //	对象创建人姓名
	Erdat    string  `orm:"index"` //	最近一次更改的日期
	Txz01    string  //	短文本
	Matnr    string  `orm:"index"` //	物料号
	Werks    string  `orm:"index"` //	工厂
	Menge    float64 //	采购申请数量
	Meins    string  //	采购申请计量单位
	Lfdat    string  //	项目交货日期
	Username string  //已批准人
	Prgcr	 string  `orm:"index"`//PR审批人代码
	Flag     string  `orm:"index"` //审批标识
	Uppo     string  //上传SAP的PO审批订单
	//Uptime    time.Time //上传时间
	//Sptime time.Time //审批时间
}


func GetAllSapprs(prgpr string) ([]*Sappr, error) {
	o := orm.NewOrm()

	cates := make([]*Sappr, 0)

	qs := o.QueryTable("sappr").Filter("prgcr", prgpr)

	//	_, err := qs.All(&cates)
	_, err := qs.OrderBy("flag", "id", "banfn", "bnfpo").All(&cates) //("-flag"）倒序排列

	return cates, err
}

//Topic
func GetDaisppr(tid string) (*Sappr, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	Sappr := new(Sappr)

	qs := o.QueryTable("Sappr")
	err = qs.Filter("id", tidNum).One(Sappr)
	if err != nil {
		return nil, err
	}

	//	topic.Views++
	//	_, err = o.Update(Categories)
	return Sappr, nil
}

//保存审批后的PO并加入标示X
func ModifyDaisppr(tid, banfn, flag string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	fmt.Println(tidNum)

	o := orm.NewOrm()
	//Sappo := &Sappo{Id: tidNum}

	//Sappo := &Sappo{Ebeln: ebeln}
	//依据当前查询条件，进行批量更新操作
	//	num, err := o.QueryTable("user").Filter("name", "slene").Update(orm.Params{
	//		"name": "astaxie",
	//	})
	////更新DB表Sappr
	_, err = o.QueryTable("Sappr").Filter("banfn", banfn).Filter("id", tidNum).Update(orm.Params{
		"flag": flag})
	if err != nil {
		return err
	}
	/*if o.Read(Sappo) != nil {
		fmt.Println("ok", ebeln)
		Sappo.Flag = flag
		//		topic.Content = content
		//Sappo.Sptime = time.Now()
		o.Update(Sappo)
	}*/
	return nil
}

func Checkpr(banfn, prgpr string) (ok bool) {
	//审批前检查sappo本地表是否还有同PO号未审批的PO行项目，因为sappo表是以行项目存取的。
	o := orm.NewOrm()
	var sapprs []Sappr
	//	sappo := &Sappo{
	//		Ebeln: ebeln,
	//		Flag:  " ",
	//		Prgco: prg,
	//	}
	//err := o.Read(&sappo, "Ebeln", "Flag", "Prgco")
	err := o.QueryTable("sappr").Filter("banfn", banfn).Filter("flag", "").Filter("prgcr", prgpr).One(&sapprs)
	if err == orm.ErrNoRows {
		//// 没有找到记录
		ok = true
		fmt.Println(ok)
		return ok

	} else {
		ok = false
		fmt.Println(ok)
		return ok
	}

}


//调取SAP采购订单审批FUNCTION
func PostSapPr(tid, banfn, prgpr, uppo string) error {

	//在之前需要判断PO在本地表中是否还有未审批的，如有放弃审批。
	//原因是sappo表中依照PO的行项目存取

	//连接SAP RFC
	saprfc.Connect()

	//接口开始调用
	params := map[string]interface{}{
		"NUMBER":  banfn,
		"REL_CODE":    prgpr,
		
	}
	r, err := saprfc.SAPconnection.Call("BAPI_REQUISITION_RELEASE_GEN", params)
	if err != nil {
		fmt.Println(err)
		return err
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
	//var b []byte
	//var s []string
	echoStruct := r["RETURN"].([]interface{})
	for _, value := range echoStruct {
		values := value.(map[string]interface{})
		//		fmt.Println(len(values)) //打印行数
		//		fmt.Println(values["MATNR"])//打印某个字段的值
		//	delete(values, "MAKTX")
		//		fmt.Println(values["MATNR"])
		//		fmt.Println(values["MATNR"])
		//TYPE消息类型: S 成功,E 错误,W 警告,I 信息,A 中断
		//CODE消息代码
		//MESSAGE消息文本
		//审批成功消息：ID:ME Type:E Number:103
		if values["CODE"] == "103" {
			//SAVE DB
			fmt.Println(tid)
			fmt.Println(prgpr)
			fmt.Println(uppo)
			//fmt.Println(ebeln)
			fmt.Println(values["TYPE"])
			fmt.Println(values["MESSAGE"])
		} else {

			//返回消息并标记 DB table
			fmt.Println(tid)
			fmt.Println(prgpr)
			fmt.Println(uppo)
			//fmt.Println(ebeln)
			fmt.Println(values["TYPE"])
			fmt.Println(values["MESSAGE"])

		}

		/*for key, val := range values {
			//valstr := strings.Split(fmt.Sprint("%s", val), ";")
			if key == "USERNAME" {
				fmt.Println(key)
				fmt.Println(val)
				s = append(s, val.(string))
			}
		}*/

	}
	//SAP RFC连接关闭
	saprfc.Close()

	return nil
}

//读取审批和未审批行数
func GetPrgcocountpr(prgpr string) (flagxpr, flagnotpr int64) {
	//依据当前的查询条件，返回结果行数
	//cnt, err := o.QueryTable("user").Count() // SELECT COUNT(*) FROM USER
	//fmt.Printf("Count Num: %s, %s", cnt, err)
	//prgcocount := make(map[string]int)
	//查旬SAPPO表中审批的个数
	o := orm.NewOrm()
	qs := o.QueryTable("Sappr")
	flagxpr, _ = qs.Filter("flag", "X").Filter("prgcr", prgpr).Count() //and 条件
	flagnotpr, _ = qs.Filter("flag", "").Filter("prgcr", prgpr).Count()
	//o.QueryTable("post").Filter("Status", 1).All(&posts, "Id", "Title")
	//var Sappolists []Sappo
	
/*	Sapprlists := make([]Sappr, 0, 0) //slice

	//读取表部分字段内容
	o.QueryTable("Sappr").Filter("prgcr", prgpr).Filter("flag", "").All(&Sapprlists, "id", "Menge")

	banfncont = 0
	for _, value := range Sapprlists {
		//ebelncont2 := int64(value.Menge1)
		//类型转换把“11.0001”的string转到浮点类型,最后强转为int64
		//ebelncont2, _ := strconv.ParseFloat(value.Menge1, 64)
		banfncont += int64(value.Menge)
		////fmt.Println(ebelncont)

	}*/

	return flagxpr, flagnotpr
}
