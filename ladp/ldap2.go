package main

/**
  ldap 使用教程
 */
import (
	"fmt"
	"github.com/samuel/go-ldap/ldap"
	"strings"
)

var attributes []string
var Dn string

//type entry *ldap.Entry

func main() {
	/**  如果ldap禁止匿名查询，需要使用账户才能开始查询，bind账户需要完整的DN信息。
	   "cn=Manager,dc=maxcrc,dc=com"
	 */
	// 要检查的用户名和密码
	//userName := "王一飞"
	//passWord := "123456"

	//binUername := "cn=Manager,dc=maxcrc,dc=com"
	//bindPassword := "secret"
	cli, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "192.168.0.195", 389))
	if err != nil {
		fmt.Println("err", err)
	}
	defer cli.Close()
	//// 建立StartTls连接，
	//err2 := cli.StartTLS(&tls.Config{InsecureSkipVerify: true})
	//if err2 != nil {
	// fmt.Println("err2", err2)
	//}
	//
	////  使用账号绑定
	//errB := cli.Bind(binUername, bindPassword)
	//if errB != nil {
	// fmt.Println("errB", errB)
	//}

	//绑定完成后就有了查询权限，构造查询请求
	searchRequest := ldap.SearchRequest{
		BaseDN:       "ou=bzks,ou=People,dc=ecust,dc=edu,dc=cn",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		TimeLimit:    10,
		TypesOnly:    false,
		SizeLimit:    999999,
	}
	//  filter Ldap的查询条件
	//filter := fmt.Sprintf("(&(objectClass=inetOrgPerson)(uid=%s))", userName) // 查询指定用户

	// 查询 inetOrgPerson分 支下的所有‘用户

	filter2 := "(&(objectClass=inetOrgPerson))"

	//searchRequest.Filter = filter2 // 设置属性

	fil, err := ldap.ParseFilter(filter2)

	// Attributes  想要获取的属性
	attributes = append(attributes, "mail")
	attributes = append(attributes, "mobile")
	attributes = append(attributes, "cn")


	atrr :="mail,mobile,cn"
	searchRequest.Filter = fil

	attr := strings.Split(atrr, ",")
	searchRequest.Attributes = make(map[string]bool)
	for _, a := range attr {
		searchRequest.Attributes[a] = true
	}

	//searchRequest.Attributes = attributes

	//查询
	//sr, errS := cli.SearchWithPaging(&searchRequest, 100)
	sr, errS := cli.Search(&searchRequest)
	if errS != nil {
		fmt.Println("errS", errS)
	}
	//循环取出属性

	//len := sr.Entries // 获取记录的条数

	//fmt.Println(len(sr.Entries))

	//s := make([]*ldap.Entry, 0)
	s := make([] *ldap.SearchResult, 0)
	for _, Att := range sr {
		s = append(s, Att)
		//fmt.Println(Att)
		//searchRes = append(searchRes,Att)
		if len(s) == 500 {
			//println(s)
			go handleAttributes_v2(s)
			s = nil // 将原先得切片置空

		}

		//fmt.Println(reflect.TypeOf(Att))
		//handleAttributes(*Att)

	}

	//userdn := sr.Entries[1]
	/*
	 fmt.Println(userdn.DN)
	 fmt.Println(userdn.Attributes)

	 for _, value := range userdn.Attributes {
	  fmt.Println("name:", value.Name, "Value:", value.Values[0])
	 }*/

}

//// 处理用户的属性名及信息
//func handleAttributes(entry ldap.Entry) {
// fmt.Println("DN:=", entry.DN) // 详细信息
// for _, Value := range entry.Attributes {
//  fmt.Println("Name:", Value.Name, "Value:", Value.Values)
// }
//}

func handleAttributes_v2(str [] *ldap.SearchResult) {
	//fmt.Println(str)

	for _, e := range str {

		for i, v := range e.Attributes {

			fmt.Println(i, string(v[0]))
		}

	}
}