package utils

import (
	"github.com/cong5/persimmon/app/info"
	"encoding/base64"
	"crypto/sha256"
	"encoding/hex"
	"crypto/md5"
	"crypto/rand"
	"strconv"
	"strings"
	"regexp"
	"bytes"
	"math"
	"fmt"
	"os"
	"io"
)

// md5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Guid
func NewGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

func HashFile(path string) (hashSum string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	hashVal := h.Sum(nil)
	return string(hashVal), nil
}

func ToInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return val
}

func Urlencode(data map[string]string) string {
	var buf bytes.Buffer
	for k, v := range data {
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(v)
		buf.WriteByte('&')
	}
	s := buf.String()
	return s[0: len(s)-1]
}

func SQLFilters(sql string) bool {
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return re.MatchString(sql)
}

func IntImplode(intArr []int, delimiter string) string {
	newString := ""
	for _, v := range intArr {
		newString += fmt.Sprintf("%d%s", v, delimiter)
	}
	newString = string([]byte(newString)[:len(newString)-1])
	return newString
}

func GetTotalPage(page, count int) int {
	return int(math.Ceil(float64(count) / float64(page)))
}

func IntDefault(inputVal int, defaultVal int) int {
	newVal := inputVal
	if newVal <= 0 {
		newVal = defaultVal
	}
	return newVal
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func Substr(src string, length int) string {
	if length <= 0 {
		length = 125
	}
	newSrc := TrimHtml(src)
	newStr := []rune(newSrc)
	strLen := len(newStr)
	var desc string
	if strLen > length {
		desc = string(newStr[:length])
	} else {
		desc = newSrc
	}
	return desc
}

func SubstrContent(posts []info.Posts) []info.Posts {
	for key, val := range posts {
		posts[key].Content = Substr(val.Content, 130)
	}
	return posts
}

func CommentRelolver(comments []info.Comments) []info.Comments {
	pat := `@[^\s]+\s?`
	re, _ := regexp.Compile(pat)
	reFun := func(s string) string {
		return fmt.Sprintf("<a href='[anchor]' rel='external nofollow'>%s</a>", s)
	}

	for key, val := range comments {
		//email md5
		comments[key].Md5 = Md5(val.Email)
		comments[key].Email = "Privacy"
		//解析@username，暂时只支持@一个人
		newContent := re.ReplaceAllStringFunc(val.Content, reFun)
		anchor := fmt.Sprintf("#%d", val.ParentId)
		comments[key].Content = strings.Replace(newContent, "[anchor]", anchor, 1)
	}
	return comments
}
