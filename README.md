CRAWL AMAP
=========
```
go mod download
go mod tidy
go mod vendor
```

Usage
-------
Step 1:
```
Create database table

CREATE TABLE `cities` (
  `id` int(10) unsigned NOT NULL auto_increment COMMENT 'primary key',
  `code` int(10) unsigned NOT NULL COMMENT 'code, level=city 是城市code',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '名称',
  `parent_code` int(10) NOT NULL DEFAULT 0 COMMENT '父code',
  `level` varchar(32) NOT NULL default '' COMMENT '等级,city,province',
  `x_axis` DECIMAL(12, 6) NOT NULL default 0 COMMENT 'X轴坐标',
  `y_axis` DECIMAl(12, 6) NOT NULL default 0 COMMENT 'Y轴坐标',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_mn` (`code`, `name`),
  KEY `idx_pid` (`parent_code`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4 COMMENT='行政区域地州市信息表';
```
Step 2:
```
set conf/conf.go
const (
	DB_USER string = "root"      //数据库连接用户名
	DB_PWD  string = "123456"    //数据库连接密码
	DB_HOST string = "127.0.0.1" //数据库连接地址
	DB_PORT int    = 3306        //数据库连接端口
	DB_NAME string = "test"      //数据库名称

	URL string = "https://restapi.amap.com/v3/config/district?keywords=%s&subdistrict=%d&key=%s"
	Key string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" //高德API KEY,高德开放平台申请
	/*
		一级
		https://restapi.amap.com/v3/config/district?keywords=%E4%B8%AD%E5%9B%BD&subdistrict=1&key=xxxxxxxxxxxxxxxxxxxxxxxxxx
		二级
		https://restapi.amap.com/v3/config/district?keywords=%E6%B2%B3%E5%8D%97%E7%9C%81&subdistrict=1&key=xxxxxxxxxxxxxxxxxxxx
	*/
)
```
Step 3:
```
go run main.go
```
