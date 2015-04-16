首先 import opdb这个包

连接数据库
opdb.OpenDb(dbpath) 成功返回true，反之返回false
添加数据
opdb.AddData(key,val) 成功返回true，反之返回false

判断是否存在某个key
opdb.HasData(key) 成功返回true，反之返回false

根据key返回对应的值
opdb.FindData(key) 返回类型 byte

根据key删除数据
opdb.DelData(key) 成功返回true，反之返回false

遍历数据库中的所有数据 
opdb.GetAllData() 返回类型 map

根据相同的key前缀遍历数据
opdb.GetDataByPrefix(pre) 返回类型 map

获取数据库中数据总条目
opdb.GetNum() 返回类型 int64

