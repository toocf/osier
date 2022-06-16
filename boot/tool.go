package boot

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 检测文件所属路径目录是否存在
func ChkDir(fpath string) string {

	_, err := os.Stat(filepath.Dir(fpath))
	if err == nil {
		return fpath
	}

	if os.IsNotExist(err) {
		_ = os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
	}

	return fpath
}

// 首字母大写
func FirstUpper(word string) string {

	if word == "" {
		return ""
	}

	return strings.ToUpper(word[:1]) + word[1:]
}

// 实例化结构体-自带默认值
func NewClass[T any](cls T, dfKey string) T {

	rt := reflect.ValueOf(&cls).Elem()
	vt := rt.Type()

	for i := 0; i < rt.NumField(); i++ {

		f := rt.Field(i)
		tag, ok := vt.Field(i).Tag.Lookup(dfKey)
		if ok {
			switch f.Kind() {

			case reflect.Bool:
				i, _ := strconv.ParseBool(tag)
				f.SetBool(i)

			case reflect.String:
				f.SetString(tag)

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				i, _ := strconv.ParseInt(tag, 10, 64)
				f.SetInt(i)

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				i, _ := strconv.ParseUint(tag, 10, 64)
				f.SetUint(i)

			case reflect.Float32, reflect.Float64:
				i, _ := strconv.ParseFloat(tag, 64)
				f.SetFloat(i)
			}
		}
	}

	return cls
}

// 获取缓存
func Get(key string) (string, error) {

	return Rds.Get(context.Background(), key).Result()
}

// 设置缓存
func Set(key, val string, tm time.Duration) error {

	return Rds.Set(context.Background(), key, val, tm).Err()
}

// 语言转换
func Lang(k string) string {

	val, ok := langText[k]
	if !ok {
		return k
	}
	return val
}
