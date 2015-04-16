/*
 * @functional 操作goleveldb的相关方法
 * @author GannicusLiu
 */
package opdb

import (
	. "../util"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"runtime"
)

var (
	db *leveldb.DB
)

/*
 * @functional 打开数据库
 * @return bool
 */
func OpenDb(dbpath string) bool {
	var err error
	db, err = leveldb.OpenFile(dbpath, nil)
	if err != nil {
		return false
	}
	return true
}

/*
 * @functional 添加数据
 * @param string key 数据条目的key
 * @param string val 数据条目的value
 * @return bool
 */
func AddData(key, val string) bool {
	err := db.Put([]byte(key), []byte(val), nil)
	if err != nil {
		return false
	}
	return true
}

/*
 * @functional 查询某个key是否存在
 * @return bool
 */
func HasData(key string) bool {
	ret, err := db.Has([]byte(key), nil)
	if err != nil {
		return false
	}
	return ret
}

/*
 * @functional 返回单条数据
 * @param byte key 数据条目的key
 * @return byte
 */
func FindData(key string) []byte {
	data, _ := db.Get([]byte(key), nil)
	return data
}

/*
 * @functional 删除key对应的记录
 * @param string key 数据条目的key
 * @return bool
 */
func DelData(key string) bool {
	err := db.Delete([]byte(key), nil)
	if err != nil {
		return false
	}
	return true
}

/*
 * @functional 获取所有的记录
 * @return map
 */
func GetAllData() map[string]string {
	//根据key遍历所有记录
	data := make(map[string]string)
	iter := db.NewIterator(nil, nil)
	for ok := iter.Seek([]byte("")); ok; ok = iter.Next() {
		data[string(iter.Key())] = string(iter.Value())
	}
	return data
}

/*
 * @functional 根据前缀获取数据
 * @return map
 */
func GetDataByPrefix(key string) map[string]string {
	data := make(map[string]string)
	iter := db.NewIterator(util.BytesPrefix([]byte(key)), nil)
	for iter.Next() {
		data[string(iter.Key())] = string(iter.Value())
	}
	iter.Release()
	return data
}

/*
 * 获取key的数量
 */
func GetNum() int64 {
	iter := db.NewIterator(nil, nil)
	var num int64
	for iter.Next() {
		num++
	}
	iter.Release()
	return num
}

/*
 * @functional 检测错误并输出
 * @param error e
 */
func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Fatalf("Bad Happened: %s, %s", file, line)
	}
}
