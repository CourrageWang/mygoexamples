package usehttp

//http 提供了http客户端和服务器的实现，支持get head post

type HttpClient struct {
	UserAgent      string //告知服务器客户端使用的OS 以及浏览器版本名称
	Accept         string //浏览器端可接收的媒体类型
	AcceptCharset  string //浏览器生生明自己接收的字符集
	AcceptEncoding string //浏览器编码方式
	AcceptLanguage string //浏览器声明自己接收的语言
	Connection     string
	CacheControl   string // 缓存机制
	ContentType    string //内容类型
	CustomHeader   map[string]string
}

const (
	httpcontenttypeXml   = "application/x-www-form-urlencoded; param=value"
	httpdefaultuserAgent = "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; CIBA; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C; InfoPath.2; .NET4.0E)"
)

//创建一个HttpClient实例
func NewHttpClient(ruls string) *HttpClient {
	return &HttpClient{
		ContentType:    "",
		UserAgent:      httpdefaultuserAgent,
		Accept:         "*/*", // 浏览器可以处理任何媒体类型
		AcceptCharset:  "GBK,utf-8;q=0.7,*;q=0.3",
		AcceptEncoding: "identity",                            //浏览器编码格式
		AcceptLanguage: "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3", //浏览器接收语言
		Connection:     "Close",                               //链接后释放
		CacheControl:   "max-age=0",                           //向server发送http请求确认，该资源是否有修改，有的话返回200没有返回304
		CustomHeader:   make(map[string]string)}
}
