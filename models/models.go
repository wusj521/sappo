package models

import (
	"fmt"
	"os"
	"path"
	"sappo/saprfc"
	"sappo/utils"
	"strconv"
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/lib/pq"
	//saprfc "github.com/sap/gorfc/gorfc"
)

const (
	_DB_NAME        = "data/sappo.db"
	_SQLITE3_DRIVER = "sqlite3"
)

/*
//字体首写字母大写
*/
type Sappo struct {
	Id        int64
	Username  string  //已批准人
	Bedat     string  `orm:"index"` //采购日期
	Matnr     string  `orm:"index"` //物料号
	Maktx     string  `orm:"index"` //物料名称
	Lifnr     string  //供应商代码
	Name1     string  `orm:"index"` //供应商描述
	Menge1    float64 //本次量
	Meins     string  //订单单位
	Netpr     string  //采购凭证中的净价格(以凭证货币计)
	Priceunit string  //价格单位 /1kg or /1000kg
	Netpr1    float64 //本次单价
	Zgys05k   string  //价格类型
	Werks     string  //工厂代码
	Zgys0502  string  //到站
	Name      string  //送货地点
	Eindt     string  //到货日期
	Vtext     string  //付款方式
	Zgys081   string  //上次单价
	Zgys0901  string  //上次时间
	Zgys06    string  //安全库存
	Zgys071   string  //库存量
	Zgys091   string  //日耗量
	//Zgztl1    float64  //在途量
	Zgys06701 string //掌握量
	Zgys06901 int64  //掌握天数
	Zgys03    string //是否招标
	Zgys01t   string //发票方式
	Zgys02t   string //运费方式
	Ekkotext  string //抬头文本
	Ebeln     string `orm:"index"` //采购单号
	Ebelp     string `orm:"index"` //行项目号
	Namelast  string //经办人
	Udate     string //审批日期及时间
	Prgco     string `orm:"index"` //PO审批代码
	Frget     string //状态批准尚未完全生效
	Frggr     string //审批组+批准策略为索引
	Frgsx     string //批准策略
	Flag      string `orm:"index"` //审批标识
	Uppo      string //上传SAP的PO审批订单
	//Uptime    time.Time //上传时间
	//Sptime time.Time //审批时间
}

//工厂物料发货每日合计清单
type Gilist struct {
	Id    int64  `orm:"index"`
	Werks string `orm:"index"` //工厂
	Mcomp string `orm:"index"` //BOM组件
	Sptag string `orm:"index"` //分析日期
	Enmng string //发货数量
}

//工厂物料采购价格清单表
type Matnrcgjiage struct {
	Id     int64   `orm:"index"`
	Werks  string  `orm:"index"` //工厂
	Matnr  string  `orm:"index"` //物料号
	Maktx  string  `orm:"index"` //物料描述
	Bedat  string  `orm:"index"` //分析日期
	Netpr1 float64 //本次采购价格
}

// 分类
type Category struct {
	Id              int64
	Title           string    `orm:"size(18)"`
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
	Flag            string `orm:"index"`
}

// 文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

//用户
type User struct {
	Id       int64
	Uname    string `orm "unique" "index"` //用户名 两条记录不能重复
	Unamemd5 string `orm "unique" "index"` //用户名 两条记录不能重复
	Pwd      string
	Tel      string `orm:"index"`
	Prgco    string `orm "unique"` //PO审批代码 两条记录不能重复
	Prgcr    string `orm "unique"` //PR审批代码 两条记录不能重复

	Frggr string //批准尚未完全生效+审批组+批准策略为索引
	Frgsx string //批准策略

}

/*
//未审批的物料可用天数
type Getmatnr struct {
	Id    int64  `orm:"index"`
	Werks string `orm:"index"` //工厂
	Mcomp string `orm:"index"` //BOM组件
	////Sptag string `orm:"index"` //分析日期
	////Enmng string //发货数量
	Zgys06901 string //掌握天数
}*/

//全局变量
//var Prgco_x string
//var Prgco_not string

func RegisterDB() {
	// 检查数据库文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	// 注册模型
	orm.RegisterModel(new(Sappr), new(Topic), new(User), new(Sappo), new(Gilist), new(Matnrcgjiage))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	//	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

//以上是注册DB和自动建表

//
func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}
	//检查数据
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	//插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")

	//	_, err := qs.All(&cates)
	_, err := qs.OrderBy("flag", "ebeln").All(&cates) //("-flag"）倒序排列

	return cates, err
}

func GetAllSappos(prg string) ([]*Sappo, error) {
	o := orm.NewOrm()

	cates := make([]*Sappo, 0)

	qs := o.QueryTable("sappo").Filter("prgco", prg)

	//	_, err := qs.All(&cates)
	_, err := qs.OrderBy("flag", "id", "ebeln", "ebelp").All(&cates) //("-flag"）倒序排列

	return cates, err
}

//Topic
func GetDaisp(tid string) (*Sappo, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	Sappo := new(Sappo)

	qs := o.QueryTable("Sappo")
	err = qs.Filter("id", tidNum).One(Sappo)
	if err != nil {
		return nil, err
	}

	//	topic.Views++
	//	_, err = o.Update(Categories)
	return Sappo, nil
}

//保存审批后的PO并加入标示X
func ModifyDaisp(tid, ebeln, flag string) error {
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
	_, err = o.QueryTable("Sappo").Filter("ebeln", ebeln).Filter("id", tidNum).Update(orm.Params{
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

func Checkpo(ebeln, prg string) (ok bool) {
	//审批前检查sappo本地表是否还有同PO号未审批的PO行项目，因为sappo表是以行项目存取的。
	o := orm.NewOrm()
	var sappos []Sappo
	//	sappo := &Sappo{
	//		Ebeln: ebeln,
	//		Flag:  " ",
	//		Prgco: prg,
	//	}
	//err := o.Read(&sappo, "Ebeln", "Flag", "Prgco")
	err := o.QueryTable("sappo").Filter("ebeln", ebeln).Filter("flag", "").Filter("prgco", prg).One(&sappos)
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

func GetUser(uname, pwdmd5 string) (*User, error) {
	//tidNum, err := strconv.ParseInt(tid, 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	var err error
	o := orm.NewOrm()

	User := new(User)
	qs := o.QueryTable("User")
	err = qs.Filter("uname", uname).Filter("pwd", pwdmd5).One(User)
	if err != nil {
		return nil, err
	}

	return User, err
}
func GetUsermd5(unamemd5, pwdmd5 string) (*User, error) {
	//tidNum, err := strconv.ParseInt(tid, 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	var err error
	o := orm.NewOrm()

	User := new(User)
	qs := o.QueryTable("User")
	err = qs.Filter("unamemd5", unamemd5).Filter("pwd", pwdmd5).One(User)
	if err != nil {
		return nil, err
	}

	return User, err
}

func InsertUser(uname, unamemd5, pwd, prg, prgpr, tel string) error {
	//tidNum, err := strconv.ParseInt(tid, 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	o := orm.NewOrm()
	pwdmd5 := utils.Md5(pwd)
	cate := &User{
		Uname:    uname,
		Unamemd5: unamemd5,
		Pwd:      pwdmd5,
		Tel:      tel,
		Prgco:    prg,
		Prgcr:    prgpr,
		//	TopicTime: time.Now(),
	}
	//检查数据
	var err error
	qs := o.QueryTable("user")
	err = qs.Filter("uname", uname).One(cate)
	if err == nil {
		return err
	}
	//插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}
func UpdatetUser(uname, pwd, prg, prgpr, tel string) error {
	//tidNum, err := strconv.ParseInt(tid, 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	o := orm.NewOrm()
	pwdmd5 := utils.Md5(pwd)

	//更新数据
	_, err := o.QueryTable("User").Filter("uname", uname).Update(orm.Params{
		"Pwd":   pwdmd5,
		"Tel":   tel,
		"Prgco": prg,
		"Prgcr": prgpr,
	})
	if err != nil {
		return err
	}

	//_, err := o.Update(cate, "Pwd", "Tel", "Prgco") //指定字段锁定
	//_, err = o.Update(cate)
	//if err != nil {
	//	return err
	//}

	return nil
}

//一、从RFC中读取未审批的采购订单
func GetSappo(prg, prgpr string) error {
	//连接SAP系统
	saprfc.Connect()
	////调取采购订单和申请单 PO PR RFC："ZMM_PO_RELEASE"
	params := map[string]interface{}{
		"I_FRGCO": prg,   //PO审批代码
		"I_FRGCR": prgpr, //PR审批代码
		/////"I_FRGGR": ,//审批组
		//		"GLTRP":      "20170829",
		//"MANTR_MARK": "X",
		//		"DELIMITER":   ";",
		//		"NO_DATA":     "",
		//		"ROWSKIPS":    0,
		//		"ROWCOUNT":    5,
	}
	r, err := saprfc.SAPconnection.Call("ZMM_PO_RELEASE", params)
	if err != nil {
		//fmt.Println(err)
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
	//采购订单信息PO
	echoStruct := r["T_ZMM_EKPO"].([]interface{})
	for _, value := range echoStruct {
		values := value.(map[string]interface{})
		//		fmt.Println(len(values)) //打印行数
		//		fmt.Println(values["MATNR"])//打印某个字段的值
		//	delete(values, "MAKTX")
		//		fmt.Println(values["MATNR"])
		//		fmt.Println(values["MATNR"])

		ebeln := values["EBELN"]
		ebelp := values["EBELP"]

		/*efor key, val := range values {
			//valstr := strings.Split(fmt.Sprint("%s", val), ";")
			if key == "USERNAME" {
				fmt.Println(key)
				fmt.Println(val)
				s = append(s, val.(string))
			}
		}
		//wausr := values["USERNAME"].(string)
		b := []byte(s[0])
		usr, _ := text.GbkToUtf8(b)*/

		fmt.Println(values["NAME1"])
		fmt.Println(values["ZGYS0901"])
		//SAP数据类型转换所有SAP接口过来的数据初步查看，类型都是string，需要转换为合适的类型，以免调取RFC时报类型错误。
		Menge1String := values["MENGE"].(string) //本次量
		Menge1Float64, _ := strconv.ParseFloat(Menge1String, 64)
		Netpr1String := values["NETPR"].(string) //本次单价
		Netpr1Float64, _ := strconv.ParseFloat(Netpr1String, 64)
		Zgys06901String := values["ZGYS0691"].(string) //掌握天数
		Zgys06901Int64, _ := strconv.ParseInt(Zgys06901String, 10, 0)
		//SAP数据类型转换
		////jiages, _ := strconv.ParseFloat(jiage, 64)
		//jiages, _ := strconv.ParseInt(jiage, 10, 0)
		//类型转换

		o := orm.NewOrm()
		cate := &Sappo{
			Username: values["USERNAME"].(string), //已批准人
			//Username:  string(usr),
			Bedat:  values["BEDAT"].(time.Time).Format("2006-01-02"), //采购日期
			Matnr:  values["MATNR"].(string),                         //物料号
			Maktx:  values["MAKTX"].(string),                         //物料名称
			Lifnr:  values["LIFNR"].(string),                         //供应商代码
			Name1:  values["NAME1"].(string),                         //供应商描述
			Menge1: Menge1Float64,                                    //本次量
			Netpr1: Netpr1Float64,                                    //本次单价
			Meins:  values["MEINS"].(string),                         //订单单位
			//Netpr: values["NETPR"].(float64), //采购凭证中的净价格(以凭证货币计)
			Priceunit: values["PRICEUNIT"].(string), //价格单位 /1kg or /1000kg

			//Zgys05k:  values["ZGYS05K"].(string),                       //价格类型
			//Zgys0502: values["ZGYS0502"].(string),                      //到站
			Name:  values["NAME"].(string),  //送货地点 /工厂名
			Werks: values["WERKS"].(string), //工厂代码

			Eindt:    values["EINDT"].(time.Time).Format("2006-01-02"),    //time.Now(),//到货日期time.Now().Format("2006-01-02 15:04:05")
			Vtext:    values["VTEXT"].(string),                            //付款方式
			Zgys081:  values["ZGYS08"].(string),                           //上次单价
			Zgys0901: values["ZGYS0901"].(time.Time).Format("2006-01-02"), //上次时间
			Zgys071:  values["ZGYS07"].(string),                           //当前库存量
			Zgys091:  values["ZGYS09"].(string),                           //日耗量
			//Zgztl1:    values["ZGZTL1"].(float64),   //在途量
			Zgys06:    values["ZGYS06"].(string),   //安全库存量
			Zgys06701: values["ZGYS0671"].(string), //掌握量
			Zgys06901: Zgys06901Int64,              //掌握天数
			Zgys03:    values["ZGYS03T"].(string),  //是否招标
			Zgys01t:   values["ZGYS01T"].(string),  //发票方式
			Zgys02t:   values["ZGYS02T"].(string),  //运费方式
			Ebeln:     ebeln.(string),              //采购订单 .(string)是类型转换
			Ebelp:     ebelp.(string),              //行项目号
			//Namelast:  values["NAMELAST"].(string), //经办人
			Ekkotext: values["EKKOTEXT"].(string), //抬头文本

			Udate: values["UDATE"].(string), //审批日期及时间
			Prgco: prg,                      //审批代码
			//Frget: values["FRGET"].(string), //状态批准尚未完全生效
			//Frggr: values["FRGGR"].(string), //审批组+批准策略为索引
			//Frgsx: values["FRGSX"].(string), //批准策略

			//currentTime := time.Now().Local()
			//timeStr := currentTime.Format("2006-01-02 15:04:05.000")

		}

		//检查数据
		//qs := o.QueryTable("category")
		//fmt.Println(cate)
		err := o.QueryTable("sappo").Filter("ebeln", ebeln).Filter("ebelp", ebelp).Filter("Prgco", prg).One(cate)
		if err == orm.ErrNoRows { //没找到相同数据 insert DB
			//插入数据
			_, err = o.Insert(cate)
			//_, err = o.Update(cate)
			if err != nil {
				return err
			}
		}

		//类型转换
		////jiages, _ := strconv.ParseFloat(jiage, 64)
		//jiages, _ := strconv.ParseInt(jiage, 10, 0)
		Netpr1String1 := values["NETPR"].(string) //本次单价
		Netpr1Float641, _ := strconv.ParseFloat(Netpr1String1, 64)

		//类型转换
		//定义slice工厂物料采购价格清单表
		cgjiage := &Matnrcgjiage{
			Werks:  values["WERKS"].(string),
			Matnr:  values["MATNR"].(string),                         //物料号
			Maktx:  values["MAKTX"].(string),                         //物料名称
			Bedat:  values["BEDAT"].(time.Time).Format("2006-01-02"), //采购日期
			Netpr1: Netpr1Float641,                                   //本次单价

		}
		//存储DB表中
		werks := values["WERKS"]
		matnr := values["MATNR"]
		bedat := values["BEDAT"].(time.Time).Format("2006-01-02")
		err = o.QueryTable("matnrcgjiage").Filter("werks", werks).Filter("matnr", matnr).Filter("bedat", bedat).One(cgjiage)
		if err == orm.ErrNoRows { //没找到相同数据 insert DB
			//插入数据
			_, err = o.Insert(cgjiage)
			//_, err = o.Update(cate)
			if err != nil {
				return err
			}
		}
		//采购申请信息PR
		echoStructpr := r["T_ZMM_EBAN"].([]interface{})
		for _, value := range echoStructpr {
			values := value.(map[string]interface{})
			//		fmt.Println(len(values)) //打印行数
			//		fmt.Println(values["MATNR"])//打印某个字段的值
			//	delete(values, "MAKTX")
			//		fmt.Println(values["MATNR"])
			//		fmt.Println(values["MATNR"])

			banfn := values["BANFN"]
			bnfpo := values["BNFPO"]
			//类型转换
			////jiages, _ := strconv.ParseFloat(jiage, 64)
			//jiages, _ := strconv.ParseInt(jiage, 10, 0)
			MengeString2 := values["MENGE"].(string) //	采购申请数量
			MengeFloat642, _ := strconv.ParseFloat(MengeString2, 64)
			//类型转换
			catepr := &Sappr{
				Banfn:    banfn.(string),                                   //	采购申请号
				Bnfpo:    bnfpo.(string),                                   //	采购申请的项目编号
				Frgst:    values["FRGST"].(string),                         //	采购请求中的批准策略
				Ekgrp:    values["EKGRP"].(string),                         //	采购组
				Ernam:    values["ERNAM"].(string),                         //	对象创建人姓名
				Erdat:    values["ERDAT"].(time.Time).Format("2006-01-02"), //	最近一次更改的日期
				Txz01:    values["TXZ01"].(string),                         //	短文本
				Matnr:    values["MATNR"].(string),                         //	物料号
				Werks:    values["WERKS"].(string),                         //	工厂
				Menge:    MengeFloat642,                                    //	采购申请数量
				Meins:    values["MEINS"].(string),                         //	采购申请计量单位
				Lfdat:    values["LFDAT"].(time.Time).Format("2006-01-02"), //	项目交货日期
				Username: values["USERNAME"].(string),                      //已批准人
				Prgcr:    prgpr,                                            //PR审批代码
			}
			err := o.QueryTable("sappr").Filter("banfn", banfn).Filter("bnfpo", bnfpo).Filter("Prgcr", prgpr).One(catepr)
			if err == orm.ErrNoRows { //没找到相同数据 insert DB
				//插入数据
				_, err = o.Insert(catepr)
				//_, err = o.Update(cate)
				if err != nil {
					return err
				}
			}

		}

	}
	/*
		//二、从RFC中读取发货明细
		var fenxdate string
		fenxdate = time.Now().Format("20060102")
		fmt.Println(fenxdate)
		params1 := map[string]interface{}{
			//"SPTAG": "20171101",
			"SPTAG": fenxdate, //RFC接口如果是日期格式，改为char格式
			//		"GLTRP":      "20170829",
			//"MANTR_MARK": "X",
			//		"DELIMITER":   ";",
			//		"NO_DATA":     "",
			//		"ROWSKIPS":    0,
			//		"ROWCOUNT":    5,
		}
		r1, err1 := saprfc.SAPconnection.Call("ZMM_GI_LIST", params1)
		if err1 != nil {
			fmt.Println(err1)
			return err1
		}

		echoStruct1 := r1["GILIST"].([]interface{})
		for _, value := range echoStruct1 {
			values := value.(map[string]interface{})

			werks := values["WERKS"]
			mcomp := values["MCOMP"]
			sptag := values["SPTAG"].(time.Time).Format("2006-01-02")

			o := orm.NewOrm()
			list := &Gilist{
				Werks: values["WERKS"].(string),
				Mcomp: values["MCOMP"].(string),
				Sptag: values["SPTAG"].(time.Time).Format("2006-01-02"),
				Enmng: values["ENMNG"].(string),
			}
			//检查数据  .Filter("SPTAG", sptag)
			err1 := o.QueryTable("gilist").Filter("werks", werks).Filter("MCOMP", mcomp).Filter("SPTAG", sptag).One(list)
			if err1 == orm.ErrNoRows { //没找到相同数据 insert DB
				//插入数据
				_, err = o.Insert(list)
				if err != nil {
					return err
				}
			} else {
				_, err = o.Update(list)
				//_, err = o.Update(cate)
				if err != nil {
					return err
				}
			}
		}
	*/

	saprfc.Close()
	//SAPconnection.Close()
	return err
}

//调取SAP采购订单审批FUNCTION
func PostSapPo(tid, ebeln, prg, uppo string) error {

	//在之前需要判断PO在本地表中是否还有未审批的，如有放弃审批。
	//原因是sappo表中依照PO的行项目存取

	//连接SAP RFC
	saprfc.Connect()

	//接口开始调用
	params := map[string]interface{}{
		"PURCHASEORDER":  ebeln,
		"PO_REL_CODE":    prg,
		"USE_EXCEPTIONS": "X",
	}
	r, err := saprfc.SAPconnection.Call("BAPI_PO_RELEASE", params)
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
			fmt.Println(prg)
			fmt.Println(uppo)
			//fmt.Println(ebeln)
			fmt.Println(values["TYPE"])
			fmt.Println(values["MESSAGE"])
		} else {

			//返回消息并标记 DB table
			fmt.Println(tid)
			fmt.Println(prg)
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
func GetPrgcocount(prg string) (flagx, flagnot, ebelncont int64) {
	//依据当前的查询条件，返回结果行数
	//cnt, err := o.QueryTable("user").Count() // SELECT COUNT(*) FROM USER
	//fmt.Printf("Count Num: %s, %s", cnt, err)
	//prgcocount := make(map[string]int)
	//查旬SAPPO表中审批的个数
	o := orm.NewOrm()
	qs := o.QueryTable("Sappo")
	flagx, _ = qs.Filter("flag", "X").Filter("prgco", prg).Count() //and 条件
	flagnot, _ = qs.Filter("flag", "").Filter("prgco", prg).Count()
	//o.QueryTable("post").Filter("Status", 1).All(&posts, "Id", "Title")
	//var Sappolists []Sappo
	Sappolists := make([]Sappo, 0, 0) //slice
	//读取表部分字段内容
	o.QueryTable("Sappo").Filter("prgco", prg).Filter("flag", "").All(&Sappolists, "id", "menge1")

	ebelncont = 0
	for _, value := range Sappolists {
		ebelncont2 := int64(value.Menge1)
		//类型转换把“11.0001”的string转到浮点类型,最后强转为int64
		//ebelncont2, _ := strconv.ParseInt(value.Menge1, 10, 0)
		//ebelncont += int64(value.Menge1)
		ebelncont += ebelncont2
		////fmt.Println(ebelncont)

	}

	return flagx, flagnot, ebelncont
}

func GetPricelist(sptag, werks, mcomp string) ([]*Gilist, error) {

	saprfc.Connect()

	params := map[string]interface{}{
		"SPTAG": sptag,
		"WERKS": werks,
		"MCOMP": mcomp,
		//		"GLTRP":      "20170829",

	}
	fmt.Println(sptag, werks, mcomp)

	r, err := saprfc.SAPconnection.Call("ZMM_GI_LIST", params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	echoStruct := r["GILIST"].([]interface{})
	for _, value := range echoStruct {
		values := value.(map[string]interface{})

		werks := values["WERKS"]
		mcomp := values["MCOMP"]
		sptag := values["SPTAG"].(time.Time).Format("2006-01-02")

		o := orm.NewOrm()
		list := &Gilist{
			Werks: values["WERKS"].(string),
			Mcomp: values["MCOMP"].(string),
			Sptag: values["SPTAG"].(time.Time).Format("2006-01-02"),
			Enmng: values["ENMNG"].(string),
		}
		//存DB table检查数据  .Filter("SPTAG", sptag)
		err1 := o.QueryTable("gilist").Filter("werks", werks).Filter("MCOMP", mcomp).Filter("SPTAG", sptag).One(list)
		if err1 == orm.ErrNoRows { //没找到相同数据 insert DB
			//插入数据
			_, err = o.Insert(list)
			if err != nil {
				return nil, err
			}
		} else {
			_, err = o.Update(list)
			//_, err = o.Update(cate)
			if err != nil {
				return nil, err
			}

		}
	}
	//SAP RFC连接关闭
	saprfc.Close()
	//取数据
	o := orm.NewOrm()

	cates := make([]*Gilist, 0)
	//werks, sptag, mcomp
	qs := o.QueryTable("gilist").Filter("werks", werks).Filter("sptag", sptag).Filter("mcomp", mcomp)

	//	_, err := qs.All(&cates)
	_, err = qs.OrderBy("sptag", "sptag").All(&cates) //("-flag"）倒序排列

	return cates, err
}

func GetPricelistOut(matnr string) ([]string, []float64, error) {
	//取数据
	o := orm.NewOrm()

	cates := make([]*Matnrcgjiage, 0)
	//werks, sptag, mcomp
	//qs.Filter("profile__age__lte", 18)
	// WHERE profile.age <= 18
	qs := o.QueryTable("matnrcgjiage").Filter("matnr", matnr).Limit(10)
	_, err := qs.OrderBy("bedat").All(&cates) //("-flag"）倒序排列

	var netpr1sl []float64
	var bedatsl []string
	for _, values := range cates {
		//以下是切片操作方法
		datetime := values.Werks + "-" + values.Bedat
		jiage := values.Netpr1
		//类型转换
		////jiages, _ := strconv.ParseFloat(jiage, 64)
		//jiages, _ := strconv.ParseInt(jiage, 10, 0)

		//netpr1sl = append(netpr1sl, jiages)  //采购价格
		netpr1sl = append(netpr1sl, jiage)  //采购价格
		bedatsl = append(bedatsl, datetime) //工作代码+日期
	}

	return bedatsl, netpr1sl, err
}

func GetMatnrkday(flag, prgco string) ([]string, []int64, error) {
	//取物料可用天数数据
	o := orm.NewOrm()
	cates := make([]*Sappo, 0)
	qs := o.QueryTable("sappo").Filter("flag", flag).Filter("prgco", prgco)
	_, err := qs.OrderBy("Zgys06901").All(&cates, "Maktx", "Name", "Zgys06901") //("-flag"）倒序排列

	////Zgys06901s := make([]int64, cap(cates))//加入cap(cates)会出现在原有空间后append
	////Maktxs := make([]string, cap(cates))
	var Zgys06901s []int64
	var Maktxs []string
	for _, values := range cates {
		//fmt.Println(values.Maktx)//输入切片 的某字段值
		//以下是切片操作方法
		text := values.Name + "-" + values.Maktx
		kyday := values.Zgys06901
		//类型转换
		//kydays, _ := strconv.ParseInt(kyday, 10, 0)

		Maktxs = append(Maktxs, text)          //物料加工厂描述
		Zgys06901s = append(Zgys06901s, kyday) //库存可用天数
		////fmt.Println(text)
	}

	return Maktxs, Zgys06901s, err

}
