package conf

const (
	DB_USER string = "root"      //数据库连接用户名
	DB_PWD  string = "123456"    //数据库连接密码
	DB_HOST string = "127.0.0.1" //数据库连接地址
	DB_PORT int    = 3306        //数据库连接端口
	DB_NAME string = "test"      //数据库名称

	URL string = "https://restapi.amap.com/v3/config/district?keywords=%s&subdistrict=%d&key=%s"
	Key string = "4a84914ef05fc1fd9dc785a513cce58d" //高德API KEY,高德开放平台申请
	/*
		一级
		https://restapi.amap.com/v3/config/district?keywords=%E4%B8%AD%E5%9B%BD&subdistrict=1&key=4a84914ef05fc1fd9dc785a513cce58d
		二级
		https://restapi.amap.com/v3/config/district?keywords=%E6%B2%B3%E5%8D%97%E7%9C%81&subdistrict=1&key=4a84914ef05fc1fd9dc785a513cce58d
	*/
)
