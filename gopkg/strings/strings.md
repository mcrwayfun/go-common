# strings
src/strings/strings.go

1. Contains(s, substr string) bool
   判断s是否包含substr，若包含则返回true，否则返回false

2. IndexByte(s string, c byte) int
   返回c在s中的第一次出现的位置，不存在则返回-1

3. ContainsAny(s, chars string) bool
   判断s是否包含chars中的任意字符

   ```go
fmt.Println(strings.ContainsAny("abc","ab"))
   fmt.Println(strings.ContainsAny("abc","z"))
   true
   false
   ```
   
4. IndexAny(s, chars string) int
   返回chars中任意字符在s中第一次出现的位置，不存在则返回-1

   ```go
   fmt.Println(strings.IndexAny("abc","ab"))
   fmt.Println(strings.IndexAny("abc","z"))
   0
   -1
   ```

5. ContainsRune(s string, r rune) bool

   判断s是否包含r的utf8，若包含则返回true，否则返回false

   ```go
   fmt.Println(strings.ContainsRune("geeksforgeeks", 97))
   fmt.Println(strings.ContainsRune("a", 97))
   false
   true
   ```

6. Count(s, substr string) int

   判断s中有几个不重复的substr的子串

7. Title(s string) string

   返回s中每个单词的首字母都改为标题格式的字符串拷贝。

   BUG: Title用于划分单词的规则不能很好的处理Unicode标点符号。

   ```go
   var TitleTests = []struct {
   	in, out string
   }{
   	{"", ""},
   	{"a", "A"},
   	{" aaa aaa aaa ", " Aaa Aaa Aaa "},
   	{" Aaa Aaa Aaa ", " Aaa Aaa Aaa "},
   	{"123a456", "123a456"},
   	{"double-blind", "Double-Blind"},
   	{"ÿøû", "Ÿøû"},
   	{"with_underscore", "With_underscore"},
   	{"unicode \xe2\x80\xa8 line separator", "Unicode \xe2\x80\xa8 Line Separator"},
   }
   ```

8. 