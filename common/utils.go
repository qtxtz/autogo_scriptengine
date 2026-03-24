package common

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"html"
	"net/url"
	"reflect"
	"time"
)

// AnyPtr 将任意类型的值转换为指针
// 如果值已经是指针，则直接返回
// 如果值是 nil，则返回 nil
// 否则创建一个新的指针指向该值
func AnyPtr(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		return v
	}

	ptr := reflect.New(rv.Type())
	ptr.Elem().Set(rv)
	return ptr.Interface()
}

// Md5 计算字符串的 MD5 哈希值
func Md5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Sha1 计算字符串的 SHA1 哈希值
func Sha1(s string) string {
	hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Sha256 计算字符串的 SHA256 哈希值
func Sha256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Base64Encode 将字符串进行 Base64 编码
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode 将 Base64 编码的字符串解码
func Base64Decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// UrlEncode 对字符串进行 URL 编码
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}

// UrlDecode 对 URL 编码的字符串进行解码
func UrlDecode(s string) (string, error) {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return "", err
	}
	return decoded, nil
}

// HtmlEncode 对字符串进行 HTML 实体编码
func HtmlEncode(s string) string {
	return html.EscapeString(s)
}

// HtmlDecode 对 HTML 实体编码的字符串进行解码
func HtmlDecode(s string) string {
	return html.UnescapeString(s)
}

// JsonEncode 将对象转换为 JSON 字符串
func JsonEncode(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// JsonDecode 将 JSON 字符串解析为对象
func JsonDecode(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// Timestamp 获取当前时间戳（秒）
func Timestamp() int64 {
	return time.Now().Unix()
}
