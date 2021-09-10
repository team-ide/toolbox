package base

import (
	"errors"
	"reflect"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
)

var (
	//全局的JSON转换对象
	JSON = jsoniter.ConfigCompatibleWithStandardLibrary
)

func GetJSON(data interface{}) string {
	if data != nil {
		bs, _ := JSON.Marshal(data)
		if bs != nil {
			return string(bs)
		}
	}
	return ""
}

//获取当前时间戳
func GetNowTime() int64 {
	return time.Now().UnixNano() / 1e6
}

//string是否在某个string数组中
func StringInArray(target string, str_array []string) bool {
	if str_array == nil {
		return false
	}
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

//生成UUID
func GenerateUUID() string {
	uuid := uuid.NewV4().String()
	slice := strings.Split(uuid, "-")
	var uuidNew string
	for _, str := range slice {
		uuidNew += str
	}
	return uuidNew
}

//复制属性
func CopyProperties(dst interface{}, src interface{}) (err error) {

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针，.Elem()类似于*ptr的操作返回指针指向的地址反射类型
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// 属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}

	return nil
}
