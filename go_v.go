package main

// https://mp.weixin.qq.com/s/ES8RBASjUdCrsvV3ErceGg
//1. Print Hello World

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"math/bits"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	fmt.Println("Hello World")
}

// 2. Print Hello 10 times

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
}

// or
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
}

//3. Create a procedure

func finish(name string) {
	fmt.Println("My job here is done. Good bye " + name)
}

func main() {
	finish("Tony")
}

//4. Create a function which returns the square of an integer

func square(x int) int {
	return x * x
}

// 5. Create a 2D Point data structure
//声明一个容器类型,有x、y两个浮点数

type Point struct {
	x, y float64
}

func main() {
	p1 := Point{}
	p2 := Point{2.1, 2.2}
	p3 := Point{
		y: 3.1,
		x: 3.2,
	}
	p4 := &Point{
		x: 4.1,
		y: 4.2,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
}

// 6. Iterate over list values

//遍历列表的值

func main() {
	items := []int{11, 22, 33}

	for _, x := range items {
		doSomething(x)
	}
}

func doSomething(i int) {
	fmt.Println(i)
}

// 7. Iterate over list indexes and values
// 遍历列表的索引和值

func main() {
	items := []string{
		"oranges",
		"apples",
		"bananas",
	}

	for i, x := range items {
		fmt.Printf("Item %d = %v \n", i, x)
	}
}

//8. Initialize a new map (associative array)
//创建一个新的map,提供一些键值对 作为初始内容

func main() {
	x := map[string]int{"one": 1, "two": 2}

	fmt.Println(x)
}

//9. Create a Binary Tree data structure
// 创建一个二叉树

type BinTree struct {
	Value int
	Left  *BinTree
	Right *BinTree
}

func inorder(root *BinTree) {
	if root == nil {
		return
	}

	inorder(root.Left)
	fmt.Printf("%d ", root.Value)
	inorder(root.Right)
}

func main() {
	root := &BinTree{1, nil, nil}
	root.Left = &BinTree{2, nil, nil}
	root.Right = &BinTree{3, nil, nil}
	root.Left.Left = &BinTree{4, nil, nil}
	root.Left.Right = &BinTree{5, nil, nil}
	root.Right.Right = &BinTree{6, nil, nil}
	root.Left.Left.Left = &BinTree{7, nil, nil}

	inorder(root)
}

// 10. Shuffle a list

// 随机排序一个list

// import (
//  "fmt"
//  "math/rand"
// )

func main() {
	x := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for i := range x {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}

	fmt.Println(x)
}

// or

func main() {
	x := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	y := make([]string, len(x))
	perm := rand.Perm(len(x))
	for i, v := range perm {
		y[v] = x[i]
	}

	fmt.Println(y)
}

// rand.Perm(x)挺有意思的一个函数，perm应该是permutation的缩写，即置换，排列。

//会输出一个从0-(x-1)随机顺序排列的数组，类似洗牌，总数不变，打乱顺序

// or

func main() {
	x := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})

	fmt.Println(x)
}

// or

func main() {
	x := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for i := len(x) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}

	fmt.Println(x)
}

// 11. Pick a random element from a list
// 从列表中选择一个随机元素

var x = []string{"bleen", "fuligin", "garrow", "grue", "hooloovoo"}

func main() {
	fmt.Println(x[rand.Intn(len(x))])
}

// or

type T string

func pickT(x []T) T {
	return x[rand.Intn(len(x))]
}

func main() {
	var list = []T{"bleen", "fuligin", "garrow", "grue", "hooloovoo"}
	fmt.Println(pickT(list))
}

//12. Check if list contains a value
// 检查列表中是否包含一个值

func Contains(list []T, x T) bool {
	for _, item := range list {
		if item == x {
			return true
		}
	}
	return false
}

type T string

func main() {
	list := []T{"a", "b", "c"}
	fmt.Println(Contains(list, "b"))
	fmt.Println(Contains(list, "z"))
}

// 13. Iterate over map keys and values

//  遍历关联数组中的每一对 k-v， 并打印出它们

func main() {
	mymap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	for k, x := range mymap {
		fmt.Println("Key =", k, ", Value =", x)
	}
}

//14. Pick uniformly a random floating point number in [a..b)
// 选出一个随机的浮点数，大于或等于a，严格小于b，且a< b

func main() {
	x := pick(-2.0, 6.5)
	fmt.Println(x)
}

func pick(a, b float64) float64 {
	return a + (rand.Float64() * (b - a))
}

//15. Pick uniformly a random integer in [a..b]
// 选出一个随机的整数，大于或等于a，小于或等于b，且a< b

func main() {
	x := pick(3, 7)

	// Note that in the Go Playground, time and random don't change very often.
	fmt.Println(x)
}

func pick(a, b int) int {
	return a + rand.Intn(b-a+1)
}

// 17. Create a Tree data structure

// 创建树数据结构, 该结构必须是递归的。一个节点可以有零个或多个子节点,节点可以访问子节点，但不能访问其父节点

type Tree struct {
	Key      keyType
	Deco     valueType
	Children []*Tree
}

type Tree struct {
	Key      key
	Deco     value
	Children []*Tree
}

type key string
type value string

func (t *Tree) String() string {
	str := "("
	str += string(t.Deco)
	if len(t.Children) == 0 {
		return str + ")"
	}
	str += " ("
	for _, child := range t.Children {
		str += child.String()
	}
	str += "))"
	return str
}

func (this *Tree) AddChild(x key, v value) *Tree {
	child := &Tree{Key: x, Deco: v}
	this.Children = append(this.Children, child)
	return child
}

func main() {
	tree := &Tree{Key: "Granpa", Deco: "Abraham"}
	subtree := tree.AddChild("Dad", "Homer")
	subtree.AddChild("Kid 1", "Bart")
	subtree.AddChild("Kid 2", "Lisa")
	subtree.AddChild("Kid 3", "Maggie")

	fmt.Println(tree)
}

// 18. Depth-first traversing of a tree
// 树的深度优先遍历。按照深度优先的前缀顺序，在树的每个节点上调用函数f

func (t *Tree) Dfs(f func(*Tree)) {
	if t == nil {
		return
	}
	f(t)
	for _, child := range t.Children {
		child.Dfs(f)
	}
}

type key string
type value string

type Tree struct {
	Key      key
	Deco     value
	Children []*Tree
}

func (this *Tree) AddChild(x key, v value) {
	child := &Tree{Key: x, Deco: v}
	this.Children = append(this.Children, child)
}

func NodePrint(node *Tree) {
	Printf("%v (%v)\n", node.Deco, node.Key)
}

func main() {
	tree := &Tree{Key: "Granpa", Deco: "Abraham"}
	tree.AddChild("Dad", "Homer")
	tree.Children[0].AddChild("Kid 1", "Bart")
	tree.Children[0].AddChild("Kid 2", "Lisa")
	tree.Children[0].AddChild("Kid 3", "Maggie")

	tree.Dfs(NodePrint)
}

//19. Reverse a list
//反转链表

func main() {

	s := []int{5, 2, 6, 3, 1, 4}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println(s)
}

// 20. Return two values
// 实现在2D矩阵m中寻找元素x，返回匹配单元格的索引 i，j

func search(m [][]int, x int) (bool, int, int) {
	for i := range m {
		for j, v := range m[i] {
			if v == x {
				return true, i, j
			}
		}
	}
	return false, 0, 0
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for x := 1; x <= 11; x += 2 {
		found, i, j := search(matrix, x)
		if found {
			fmt.Printf("matrix[%v][%v] == %v \n", i, j, x)
		} else {
			fmt.Printf("Value %v not found. \n", x)
		}
	}
}

//21. Swap values
//交换变量a和b的值

func main() {
	a := 3
	b := 10
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

// 22. Convert string to integer
//将字符串转换为整型

func main() {
	// create a string
	s := "123"
	fmt.Println(s)
	fmt.Println("type:", reflect.TypeOf(s))

	// convert string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
	fmt.Println("type:", reflect.TypeOf(i))
}

//or

func main() {
	s := "123"
	fmt.Println("s is", reflect.TypeOf(s), s)

	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("i is", reflect.TypeOf(i), i)
}

//23. Convert real number to string with 2 decimal places
// 给定一个实数，小数点后保留两位小数

func main() {
	x := 3.14159
	s := fmt.Sprintf("%.2f", x)
	fmt.Println(s)
}

//24. Assign to string the japanese word ネコ
//声明一个新的字符串s，并用文字值“ネコ”初始化它(在日语中是“cat”的意思)

func main() {
	s := "ネコ"
	fmt.Println(s)
}

//25. Send a value to another thread
//将字符串值“Alan”与现有的正在运行的进程共享，该进程将显示“你好，Alan”

func main() {
	ch := make(chan string)

	go func() {
		v := <-ch
		fmt.Printf("Hello, %v\n", v)
	}()

	ch <- "Alan"

	// Make sure the non-main goroutine had the chance to finish.
	time.Sleep(time.Second)
}

//26. Create a 2-dimensional array

// 创建一个二维数组

// 声明并初始化一个有m行n列的矩阵x，包含实数。
func main() {
	const m, n = 3, 4
	var x [m][n]float64

	x[1][2] = 8
	fmt.Println(x)
}

// or

func main() {
	x := make2D(2, 3)

	x[1][1] = 8
	fmt.Println(x)
}

func make2D(m, n int) [][]float64 {
	buf := make([]float64, m*n)

	x := make([][]float64, m)
	for i := range x {
		x[i] = buf[:n:n]
		buf = buf[n:]
	}
	return x
}

// 27. Create a 3-dimensional array
// 创建一个三维数组

// 声明并初始化一个三维数组x，它有m，n，p维边界，并且包含实数。

func main() {
	const m, n, p = 2, 2, 3
	var x [m][n][p]float64

	x[1][0][2] = 9

	// Value of x
	fmt.Println(x)

	// Type of x
	fmt.Printf("%T", x)
}

//or

func make3D(m, n, p int) [][][]float64 {
	buf := make([]float64, m*n*p)

	x := make([][][]float64, m)
	for i := range x {
		x[i] = make([][]float64, n)
		for j := range x[i] {
			x[i][j] = buf[:p:p]
			buf = buf[p:]
		}
	}
	return x
}

//or

func main() {
	x := make3D(2, 2, 3)

	x[1][0][2] = 9
	fmt.Println(x)
}

func make3D(m, n, p int) [][][]float64 {
	buf := make([]float64, m*n*p)

	x := make([][][]float64, m)
	for i := range x {
		x[i] = make([][]float64, n)
		for j := range x[i] {
			x[i][j] = buf[:p:p]
			buf = buf[p:]
		}
	}
	return x
}

//28. Sort by a property

// 按x->p的升序对类似数组的集合项的元素进行排序，其中p是项中对象的项类型的字段。

type Item struct {
	label string
	p     int
	lang  string
}

type ItemPSorter []Item

func (s ItemPSorter) Len() int           { return len(s) }
func (s ItemPSorter) Less(i, j int) bool { return s[i].p < s[j].p }
func (s ItemPSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func sortItems(items []Item) {
	sorter := ItemPSorter(items)
	sort.Sort(sorter)
}

func main() {
	items := []Item{
		{"twelve", 12, "english"},
		{"six", 6, "english"},
		{"eleven", 11, "english"},
		{"zero", 0, "english"},
		{"two", 2, "english"},
	}
	fmt.Println("Unsorted: ", items)
	sortItems(items)
	fmt.Println("Sorted: ", items)
}

//or

type Item struct {
	label string
	p     int
	lang  string
}

func main() {
	items := []Item{
		{"twelve", 12, "english"},
		{"six", 6, "english"},
		{"eleven", 11, "english"},
		{"zero", 0, "english"},
		{"two", 2, "english"},
	}
	fmt.Println("Unsorted: ", items)

	less := func(i, j int) bool {
		return items[i].p < items[j].p
	}
	sort.Slice(items, less)

	fmt.Println("Sorted: ", items)
}

//29. Remove item from list, by its index

// 从列表项中删除第I项。这将改变原来的列表或返回一个新的列表，这取决于哪个更习惯。请注意，在大多数语言中，I的最小有效值是0。

func main() {
	items := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(items)

	i := 2
	items = append(items[:i], items[i+1:]...)
	fmt.Println(items)
}

//or

// copy(items[i:], items[i+1:])
// items[len(items)-1] = nil
// items = items[:len(items)-1]

//30. Parallelize execution of 1000 independent tasks

// 用参数I从1到1000启动程序f的并发执行。任务是独立的，f(i)不返回值。任务不需要同时运行，所以可以使用pools

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		go func(j int) {
			f(j)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func f(i int) {
	d := rand.Int() % 10000
	time.Sleep(time.Duration(d))
	fmt.Printf("Hello %v\n", i)
}

func main() {
	for i := 1; i <= 1000; i++ {
		go f(i)
	}

	time.Sleep(4 * time.Second)
}

//31. Recursive factorial (simple)
// 创建递归函数f，该函数返回从f(i-1)计算的非负整数I的阶乘

func f(i int) int {
	if i == 0 {
		return 1
	}
	return i * f(i-1)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("f(%d) = %d\n", i, f(i))
	}
}

//32. Integer exponentiation by squaring

// 创建函数exp，计算(快速)x次方n的值。x和n是非负整数。

func exp(x, n int) int {
	switch {
	case n == 0:
		return 1
	case n == 1:
		return x
	case n%2 == 0:
		return exp(x*x, n/2)
	default:
		return x * exp(x*x, (n-1)/2)
	}
}

func main() {
	fmt.Println(exp(3, 5))
}

//33. Atomically read and update variable
//为变量x分配新值f(x)，确保在读和写之间没有其他线程可以修改x。

func main() {
	var lock sync.RWMutex
	x := 3

	lock.Lock()
	x = f(x)
	lock.Unlock()

	fmt.Println(x)
}

func f(i int) int {
	return 2 * i
}

//34. Create a set of objects
// 声明并初始化一个包含t类型对象的集合x。

type T string

func main() {
	// declare a Set (implemented as a map)
	x := make(map[T]bool)

	// add some elements
	x["A"] = true
	x["B"] = true
	x["B"] = true
	x["C"] = true
	x["D"] = true

	// remove an element
	delete(x, "C")

	for t, _ := range x {
		fmt.Printf("x contains element %v \n", t)
	}
}

//or

type T string

func main() {
	// declare a Set (implemented as a map)
	x := make(map[T]struct{})

	// add some elements
	x["A"] = struct{}{}
	x["B"] = struct{}{}
	x["B"] = struct{}{}
	x["C"] = struct{}{}
	x["D"] = struct{}{}

	// remove an element
	delete(x, "C")

	for t, _ := range x {
		fmt.Printf("x contains element %v \n", t)
	}
}

//35. First-class function : compose[2]
// 用参数f (A -> B)和g (B -> C)实现一个函数compose (A -> C)，返回composition函数g ∘ f

func compose(f func(A) B, g func(B) C) func(A) C {
	return func(x A) C {
		return g(f(x))
	}
}

func main() {
	squareFromStr := compose(str2int, square)
	fmt.Println(squareFromStr("12"))
}

type A string
type B int
type C int

func str2int(a A) B {
	b, _ := strconv.ParseInt(string(a), 10, 32)
	return B(b)
}

func square(b B) C {
	return C(b * b)
}

//36. First-class function : generic composition

// 实现一个函数组合，该函数组合为任何恰好有1个参数的函数f和g返回组合函数g ∘ f。

func composeIntFuncs(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return g(f(x))
	}
}

func main() {
	double := func(x int) int {
		return 2 * x
	}
	addTwo := func(x int) int {
		return x + 2
	}
	h := composeIntFuncs(double, addTwo)

	for i := 0; i < 10; i++ {
		fmt.Println(i, h(i), addTwo(double(i)))
	}
}

//37. Currying
//将一个接受多个参数的函数转换为一个预设了某些参数的函数。

type Company string

type Employee struct {
	FirstName string
	LastName  string
}

func (e *Employee) String() string {
	return "<" + e.FirstName + " " + e.LastName + ">"
}

type Payroll struct {
	Company   Company
	Boss      *Employee
	Employee  *Employee
	StartDate time.Time
	EndDate   time.Time
	Amount    int
}

// Creates a blank payroll for a specific employee with specific boss in specific company
type PayFactory func(Company, *Employee, *Employee) Payroll

// Creates a blank payroll for a specific employee
type CustomPayFactory func(*Employee) Payroll

func CurryPayFactory(pf PayFactory, company Company, boss *Employee) CustomPayFactory {
	return func(e *Employee) Payroll {
		return pf(company, boss, e)
	}
}

func NewPay(company Company, boss *Employee, employee *Employee) Payroll {
	return Payroll{
		Company:  company,
		Boss:     boss,
		Employee: employee,
	}
}

func main() {
	me := Employee{"Jack", "Power"}

	// I happen to be head of the HR department of Richissim Inc.
	var myLittlePayFactory CustomPayFactory = CurryPayFactory(NewPay, "Richissim", &me)

	fmt.Println(myLittlePayFactory(&Employee{"Jean", "Dupont"}))
	fmt.Println(myLittlePayFactory(&Employee{"Antoine", "Pol"}))
}

// 38. Extract a substring

// 查找由字符串s的字符I(包括)到j(不包括)组成的子字符串t。除非另有说明，字符索引从0开始。确保正确处理多字节字符。

func main() {
	s := "hello, utf-8 문자들"
	i, j := 7, 15

	t := string([]rune(s)[i:j])

	fmt.Println(t)
}

//39. Check if string contains a word
// 如果字符串单词作为子字符串包含在字符串s中，则将布尔ok设置为true，否则设置为false。

func main() {
	s := "Let's dance the macarena"

	word := "dance"
	ok := strings.Contains(s, word)
	fmt.Println(ok)

	word = "car"
	ok = strings.Contains(s, word)
	fmt.Println(ok)

	word = "duck"
	ok = strings.Contains(s, word)
	fmt.Println(ok)
}

//41. Reverse a string
//反转字符串

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "The quick brown 狐 jumped over the lazy 犬"
	fmt.Println(Reverse(input))
	// Original string unaltered
	fmt.Println(input)
}

// 42. Continue outer loop

// 打印列表a中不包含在列表b中的每个项目v。 为此，编写一个外部循环来迭代a，编写一个内部循环来迭代b。

func printSubtraction(a []int, b []int) {
mainloop:
	for _, v := range a {
		for _, w := range b {
			if v == w {
				continue mainloop
			}
		}
		fmt.Println(v)
	}
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{2, 4, 6, 8}
	printSubtraction(a, b)
}

//43. Break outer loop
//在2D整数矩阵m中寻找一个负值v，打印出来，停止搜索。

var m = [][]int{
	{1, 2, 3},
	{11, 0, 30},
	{5, -20, 55},
	{0, 0, -60},
}

func main() {
mainloop:
	for i, line := range m {
		fmt.Fprintln(os.Stderr, "Searching in line", i)
		for _, v := range line {
			if v < 0 {
				fmt.Println("Found ", v)
				break mainloop
			}
		}
	}

	fmt.Println("Done.")
}

//44. Insert element in list

// 在列表s的位置I插入元素x。其他元素必须向右移动。

func main() {

	s := make([]int, 2)

	s[0] = 0
	s[1] = 2

	fmt.Println(s)
	// insert one at index one
	s = append(s, 0)
	copy(s[2:], s[1:])
	s[1] = 1

	fmt.Println(s)
}

//45. Pause execution for 5 seconds
// 在继续下一个指令之前，在当前线程中休眠5秒钟。

func main() {
	time.Sleep(5 * time.Second)
	fmt.Println("Done.")
}

//46. Extract beginning of string (prefix)
//创建由字符串s的前5个字符组成的字符串t。 确保正确处理多字节字符。

func main() {
	s := "Привет"
	t := string([]rune(s)[:5])

	fmt.Println(t)
}

//47. Extract string suffix

//创建由字符串s的最后5个字符组成的字符串t。 确保正确处理多字节字符

func main() {
	s := "hello, world! 문자"
	t := string([]rune(s)[len([]rune(s))-5:])
	fmt.Println(t)
}

//48. Multi-line string literal
// 给变量s赋值一个由几行文本组成的字符串，包括换行符。

func main() {
	s := `Huey
Dewey
Louie`

	fmt.Println(s)
}

//49. Split a space-separated string
// 拆分用空格分隔的字符串
//构建由输入字符串的子字符串组成的列表块，由一个或多个空格字符分隔。

func main() {
	s := "Un dos tres"
	chunks := strings.Split(s, " ")
	fmt.Println(len(chunks))
	fmt.Println(chunks)

	s = " Un dos tres "
	chunks = strings.Split(s, " ")
	fmt.Println(len(chunks))
	fmt.Println(chunks)

	s = "Un  dos"
	chunks = strings.Split(s, " ")
	fmt.Println(len(chunks))
	fmt.Println(chunks)
}

//or

func main() {
	s := "hello world"
	chunks := strings.Fields(s)

	fmt.Println(chunks)
}

// or

func main() {
	s := "Un dos tres"
	chunks := strings.Fields(s)
	fmt.Println(len(chunks))
	fmt.Println(chunks)

	s = " Un dos tres "
	chunks = strings.Fields(s)
	fmt.Println(len(chunks))
	fmt.Println(chunks)

	s = "Un  dos"
	chunks = strings.Fields(s)
	fmt.Println(len(chunks))
	fmt.Println(chunks)
}

// strings.Fields 就只能干这个事儿

//50. Make an infinite loop

// 写一个无限循环

// for {
//  // Do something
// }

func main() {
	k := 0
	for {
		fmt.Println("Hello, playground")
		k++
		if k == 5 {
			break
		}
	}
}

//51. Check if map contains key
// 检查map是否有某个key

func main() {
	m := map[string]int{
		"uno":  1,
		"dos":  2,
		"tres": 3,
	}

	k := "cinco"
	_, ok := m[k]
	fmt.Printf("m contains key %q: %v\n", k, ok)

	k = "tres"
	_, ok = m[k]
	fmt.Printf("m contains key %q: %v\n", k, ok)
}

//52. Check if map contains value
// 检查map中是否有某个值

func containsValue(m map[K]T, v T) bool {
	for _, x := range m {
		if x == v {
			return true
		}
	}
	return false
}

// Arbitrary types for K, T.
type K string
type T int

func main() {
	m := map[K]T{
		"uno":  1,
		"dos":  2,
		"tres": 3,
	}

	var v T = 5
	ok := containsValue(m, v)
	fmt.Printf("m contains value %d: %v\n", v, ok)

	v = 3
	ok = containsValue(m, v)
	fmt.Printf("m contains value %d: %v\n", v, ok)
}

//53. Join a list of strings
// 字符串连接

func main() {

	x := []string{"xxx", "bbb", "aaa"}
	y := strings.Join(x, "&")

	fmt.Println(y)

}

// 关于 strings.Joins[2]

// 54. Compute sum of integers
// 计算整数之和

func main() {
	x := []int{1, 2, 3}
	s := 0
	for _, v := range x {
		s += v
	}
	fmt.Println(s)
}

// 55. Convert integer to string
// 将整数转换为字符串

func main() {
	var i int = 1234
	s := strconv.Itoa(i)
	fmt.Println(s)
}

// or
func main() {
	var i int64 = 1234
	s := strconv.FormatInt(i, 10)
	fmt.Println(s)
}

// or

func main() {
	var i int = 1234
	s := fmt.Sprintf("%d", i)
	fmt.Println(s)

	var j int = 5678
	s = fmt.Sprintf("%d", j)
	fmt.Println(s)

	var k *big.Int = big.NewInt(90123456)
	s = fmt.Sprintf("%d", k)
	fmt.Println(s)
}

//56. Launch 1000 parallel tasks and wait for completion

// 创建1000个并行任务，并等待其完成

func f(i int) {
	d := rand.Int() % 10000
	time.Sleep(time.Duration(d))
	fmt.Printf("Hello %v\n", i)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		go func(i int) {
			f(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished")
}

//57. Filter list
// 过滤list中的值

type T int

func main() {
	x := []T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := func(t T) bool { return t%2 == 0 }

	y := make([]T, 0, len(x))
	for _, v := range x {
		if p(v) {
			y = append(y, v)
		}
	}

	fmt.Println(y)
}

// or

type T int

func main() {
	x := []T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := func(t T) bool { return t%2 == 0 }

	n := 0
	for _, v := range x {
		if p(v) {
			n++
		}
	}
	y := make([]T, 0, n)
	for _, v := range x {
		if p(v) {
			y = append(y, v)
		}
	}

	fmt.Println(y)
}

// 58. Extract file content to a string
// 提取字符串的文件内容

func main() {
	f := "data.txt"
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	lines := string(b)

	fmt.Println(lines)
}

// Create file in fake FS of the Playground. init is executed before main.
func init() {
	err := ioutil.WriteFile("data.txt", []byte(`Un
Dos
Tres`), 0644)
	if err != nil {
		panic(err)
	}
}

//59. Write to standard error stream
// 写入标准错误流

func main() {
	x := -2
	fmt.Fprintln(os.Stderr, x, "is negative")
}

//60. Read command line argument

//读取命令行参数

// x := os.Args[1]

//61. Get current date
// 获取当前时间

func main() {
	d := time.Now()
	fmt.Println("Now is", d)
	// The Playground has a special sandbox, so you may get a Time value fixed in the past.
}

// 62. Find substring position
// 字符串查找

// 查找子字符串位置

func main() {
	x := "été chaud"

	{
		y := "chaud"
		i := strings.Index(x, y)
		fmt.Println(i)
	}

	{
		y := "froid"
		i := strings.Index(x, y)
		fmt.Println(i)
	}
}

//63. Replace fragment of a string
//替换字符串片段

func main() {
	x := "oink oink oink"
	y := "oink"
	z := "moo"
	x2 := strings.Replace(x, y, z, -1)
	fmt.Println(x2)
}

// 64. Big integer : value 3 power 247
// 超大整数

func main() {
	x := new(big.Int)
	x.Exp(big.NewInt(3), big.NewInt(247), nil)
	fmt.Println(x)
}

// 65. Format decimal number
// 格式化十进制数

func main() {
	x := 0.15625
	s := fmt.Sprintf("%.1f%%", 100.0*x)
	fmt.Println(s)
}

// 66. Big integer exponentiation
// 大整数幂运算

func exp(x *big.Int, n int) *big.Int {
	nb := big.NewInt(int64(n))
	var z big.Int
	z.Exp(x, nb, nil)
	return &z
}

func main() {
	x := big.NewInt(3)
	n := 5
	z := exp(x, n)
	fmt.Println(z)
}

// 67. Binomial coefficient "n choose k"
// Calculate binom(n, k) = n! / (k! * (n-k)!). Use an integer type able to handle huge numbers.

// 二项式系数“n选择k”

func main() {
	z := new(big.Int)

	z.Binomial(4, 2)
	fmt.Println(z)

	z.Binomial(133, 71)
	fmt.Println(z)
}

// 68. Create a bitset
// 创建位集合

func main() {
	var x *big.Int = new(big.Int)

	x.SetBit(x, 42, 1)

	for _, y := range []int{13, 42} {
		fmt.Println("x has bit", y, "set to", x.Bit(y))
	}
}

//or

const n = 1024

func main() {
	x := make([]bool, n)

	x[42] = true

	for _, y := range []int{13, 42} {
		fmt.Println("x has bit", y, "set to", x[y])
	}
}

// or

func main() {
	const n = 1024

	x := NewBitset(n)

	x.SetBit(13)
	x.SetBit(42)
	x.ClearBit(13)

	for _, y := range []int{13, 42} {
		fmt.Println("x has bit", y, "set to", x.GetBit(y))
	}
}

type Bitset []uint64

func NewBitset(n int) Bitset {
	return make(Bitset, (n+63)/64)
}

func (b Bitset) GetBit(index int) bool {
	pos := index / 64
	j := index % 64
	return (b[pos] & (uint64(1) << j)) != 0
}

func (b Bitset) SetBit(index int) {
	pos := index / 64
	j := index % 64
	b[pos] |= (uint64(1) << j)
}

func (b Bitset) ClearBit(index int) {
	pos := index / 64
	j := index % 64
	b[pos] ^= (uint64(1) << j)
}

//69. Seed random generator

// 随机种子生成器

func main() {
	var s int64 = 42
	rand.Seed(s)
	fmt.Println(rand.Int())
}

// or

func main() {
	var s int64 = 42
	r := rand.New(rand.NewSource(s))
	fmt.Println(r.Int())
}

//70. Use clock as random generator seed
//使用时钟作为随机生成器的种子

func main() {
	rand.Seed(time.Now().UnixNano())
	// Well, the playground date is actually fixed in the past, and the
	// output is cached.
	// But if you run this on your workstation, the output will vary.
	fmt.Println(rand.Intn(999))
}

// or

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Well, the playground date is actually fixed in the past, and the
	// output is cached.
	// But if you run this on your workstation, the output will vary.
	fmt.Println(r.Intn(999))
}

// 71. Echo program implementation
// 实现 Echo 程序

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//74. Compute GCD
//计算大整数a和b的最大公约数x。使用能够处理大数的整数类型。

func main() {
	a, b, x := new(big.Int), new(big.Int), new(big.Int)
	a.SetString("6000000000000", 10)
	b.SetString("9000000000000", 10)
	x.GCD(nil, nil, a, b)
	fmt.Println(x)
}

// 75. Compute LCM
// 计算大整数a和b的最小公倍数x。使用能够处理大数的整数类型。

func main() {
	a, b, gcd, x := new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	a.SetString("6000000000000", 10)
	b.SetString("9000000000000", 10)
	gcd.GCD(nil, nil, a, b)
	x.Div(a, gcd).Mul(x, b)
	fmt.Println(x)
}

//76. Binary digits from an integer
//将十进制整数转换为二进制数字

func main() {
	x := int64(13)
	s := strconv.FormatInt(x, 2)

	fmt.Println(s)
}

// or

func main() {
	x := big.NewInt(13)
	s := fmt.Sprintf("%b", x)

	fmt.Println(s)
}

//77. SComplex number

// 复数

func main() {
	x := 3i - 2
	x *= 1i

	fmt.Println(x)
	fmt.Print(reflect.TypeOf(x))
}

// 78. "do while" loop

// 循环执行

func main() {
	for {
		x := rollDice()
		fmt.Println("Got", x)
		if x == 3 {
			break
		}

	}
}

func rollDice() int {
	return 1 + rand.Intn(6)
}

//or

func main() {
	for done := false; !done; {
		x := rollDice()
		fmt.Println("Got", x)
		done = x == 3
	}
}

func rollDice() int {
	return 1 + rand.Intn(6)
}

//79. Convert integer to floating point number

// 整型转浮点型

// 声明浮点数y并用整数x的值初始化它。

func main() {
	x := 5
	y := float64(x)

	fmt.Println(y)
	fmt.Printf("%.2f\n", y)
	fmt.Println(reflect.TypeOf(y))
}

//80. Truncate floating point number to integer
// /浮点型转整型

func main() {
	a := -6.4
	b := 6.4
	c := 6.6
	fmt.Println(int(a))
	fmt.Println(int(b))
	fmt.Println(int(c))
}

//81. Round floating point number to integer
// 按规则取整

func round(x float64) int {
	y := int(math.Floor(x + 0.5))
	return y
}

func main() {
	for _, x := range []float64{-1.1, -0.9, -0.5, -0.1, 0., 0.1, 0.5, 0.9, 1.1} {
		fmt.Printf("%5.1f %5d\n", x, round(x))
	}
}

// 82. Count substring occurrences
// 统计子字符串出现次数

func main() {
	s := "Romaromamam"
	t := "mam"

	x := strings.Count(s, t)

	fmt.Println(x)
}

// 83. Regex with character repetition
// 正则表达式匹配重复字符

func main() {
	r := regexp.MustCompile("htt+p")

	for _, s := range []string{
		"hp",
		"htp",
		"http",
		"htttp",
		"httttp",
		"htttttp",
		"htttttp",
		"word htttp in a sentence",
	} {
		fmt.Println(s, "=>", r.MatchString(s))
	}
}

// 84. Count bits set in integer binary representation

// 计算十进制整型的二进制表示中 1的个数

func PopCountUInt64(i uint64) (c int) {
	// bit population count, see
	// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
	i -= (i >> 1) & 0x5555555555555555
	i = (i>>2)&0x3333333333333333 + i&0x3333333333333333
	i += i >> 4
	i &= 0x0f0f0f0f0f0f0f0f
	i *= 0x0101010101010101
	return int(i >> 56)
}

func PopCountUInt32(i uint32) (n int) {
	// bit population count, see
	// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
	i -= (i >> 1) & 0x55555555
	i = (i>>2)&0x33333333 + i&0x33333333
	i += i >> 4
	i &= 0x0f0f0f0f
	i *= 0x01010101
	return int(i >> 24)
}

func main() {
	for i := uint64(0); i < 16; i++ {
		c := PopCountUInt64(i)
		fmt.Printf("%4d %04[1]b %d\n", i, c)
	}

	for i := uint32(0); i < 16; i++ {
		c := PopCountUInt32(i)
		fmt.Printf("%4d %04[1]b %d\n", i, c)
	}
}

// or

func main() {
	for i := uint(0); i < 16; i++ {
		c := bits.OnesCount(i)
		fmt.Printf("%4d %04[1]b %d\n", i, c)
	}
}

// 85. Check if integer addition will overflow
// 检查两个整型相加是否溢出

func willAddOverflow(a, b int64) bool {
	return a > math.MaxInt64-b
}

func main() {

	fmt.Println(willAddOverflow(11111111111111111, 2))

}

// 86. Check if integer multiplication will overflow
// 检查整型相乘是否溢出

func multiplyWillOverflow(x, y uint64) bool {
	if x <= 1 || y <= 1 {
		return false
	}
	d := x * y
	return d/y != x
}

func main() {
	{
		var x, y uint64 = 2345, 6789
		if multiplyWillOverflow(x, y) {
			fmt.Println(x, "*", y, "overflows")
		} else {
			fmt.Println(x, "*", y, "doesn't overflow")
		}
	}
	{
		var x, y uint64 = 2345678, 9012345
		if multiplyWillOverflow(x, y) {
			fmt.Println(x, "*", y, "overflows")
		} else {
			fmt.Println(x, "*", y, "doesn't overflow")
		}
	}
	{
		var x, y uint64 = 2345678901, 9012345678
		if multiplyWillOverflow(x, y) {
			fmt.Println(x, "*", y, "overflows")
		} else {
			fmt.Println(x, "*", y, "doesn't overflow")
		}
	}
}

// 87. Stop program

// 停止程序,立即退出。

func main() {

	os.Exit(1)

	print(2222)
}

// 88. Allocate 1M bytes
// 分配1M内存

func main() {
	buf := make([]byte, 1000000)

	for i, b := range buf {
		if b != 0 {
			fmt.Println("Found unexpected value", b, "at position", i)
		}
	}
	fmt.Println("Buffer was correctly initialized with zero values.")
}

// 89. Handle invalid argument
// 处理无效参数

func NewSquareMatrix(N int) ([][]float64, error) {
	if N < 0 {
		return nil, fmt.Errorf("Invalid size %d: order cannot be negative", N)
	}
	matrix := make([][]float64, N)
	for i := range matrix {
		matrix[i] = make([]float64, N)
	}
	return matrix, nil
}

func main() {
	N1 := 3
	matrix1, err1 := NewSquareMatrix(N1)
	if err1 == nil {
		fmt.Println(matrix1)
	} else {
		fmt.Println(err1)
	}

	N2 := -2
	matrix2, err2 := NewSquareMatrix(N2)
	if err2 == nil {
		fmt.Println(matrix2)
	} else {
		fmt.Println(err2)
	}
}

// 90. Read-only outside
// 外部只读

type Foo struct {
	x int
}

func (f *Foo) X() int {
	return f.x
}

// 91. Load JSON file into struct
// json转结构体

func readJSONFile() error {
	var x Person

	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buffer, &x)
	if err != nil {
		return err
	}

	fmt.Println(x)
	return nil
}

func main() {
	err := readJSONFile()
	if err != nil {
		panic(err)
	}
}

type Person struct {
	FirstName string
	Age       int
}

const filename = "/tmp/data.json"

func init() {
	err := ioutil.WriteFile(filename, []byte(`
  {
   "FirstName":"Napoléon",
   "Age": 51 
  }`), 0644)
	if err != nil {
		panic(err)
	}
}

// or

func readJSONFile() error {
	var x Person

	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&x)
	if err != nil {
		return err
	}

	fmt.Println(x)
	return nil
}

func main() {
	err := readJSONFile()
	if err != nil {
		panic(err)
	}
}

type Person struct {
	FirstName string
	Age       int
}

const filename = "/tmp/data.json"

func init() {
	err := ioutil.WriteFile(filename, []byte(`
  {
   "FirstName":"Napoléon",
   "Age": 51 
  }`), 0644)
	if err != nil {
		panic(err)
	}
}

// 92. Save object into JSON file
// 将json对象写入文件

func writeJSONFile() error {
	x := Person{
		FirstName: "Napoléon",
		Age:       51,
	}

	buffer, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buffer, 0644)
}

func main() {
	err := writeJSONFile()
	if err != nil {
		panic(err)
	}
	fmt.Println("Done.")
}

type Person struct {
	FirstName string
	Age       int
}

const filename = "/tmp/data.json"

// 93. Pass a runnable procedure as parameter
// 以函数作为参数

func main() {
	control(greet)
}

func control(f func()) {
	fmt.Println("Before f")
	f()
	fmt.Println("After f")
}

func greet() {
	fmt.Println("Hello, developers")
}

// 94. Print type of variable
// 打印变量的类型

func main() {
	var x interface{}

	x = "Hello"
	fmt.Println(reflect.TypeOf(x))

	x = 4
	fmt.Println(reflect.TypeOf(x))

	x = os.NewFile(0777, "foobar.txt")
	fmt.Println(reflect.TypeOf(x))
}

// or

func main() {
	var x interface{}

	x = "Hello"
	fmt.Printf("%T", x)
	fmt.Println()

	x = 4
	fmt.Printf("%T", x)
	fmt.Println()

	x = os.NewFile(0777, "foobar.txt")
	fmt.Printf("%T", x)
	fmt.Println()
}

// 95. Get file size
// 获取文件的大小

func main() {
	err := printSize("file.txt")
	if err != nil {
		panic(err)
	}
}

func printSize(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	x := info.Size()

	fmt.Println(x)
	return nil
}

func init() {
	// The file will only contains the characters "Hello", no newlines.
	buffer := []byte("Hello")
	err := ioutil.WriteFile("file.txt", buffer, 0644)
	if err != nil {
		panic(err)
	}
}

//96. Check string prefix
// 检查两个字符串前缀是否一致

func check(s, prefix string) {

	b := strings.HasPrefix(s, prefix)

	if b {
		fmt.Println(s, "starts with", prefix)
	} else {
		fmt.Println(s, "doesn't start with", prefix)
	}
}

func main() {
	check("bar", "foo")
	check("foobar", "foo")
}

//97. Check string suffix
// 检查字符串后缀

func check(s, suffix string) {

	b := strings.HasSuffix(s, suffix)

	if b {
		fmt.Println(s, "ends with", suffix)
	} else {
		fmt.Println(s, "doesn't end with", suffix)
	}
}

func main() {
	check("foo", "bar")
	check("foobar", "bar")
}

//98. Epoch seconds to date object

//时间戳转日期

func main() {
	ts := int64(1451606400)
	d := time.Unix(ts, 0)

	fmt.Println(d)
}

//99. Format date YYYY-MM-DD
// 时间格式转换

func main() {
	d := time.Now()
	x := d.Format("2006-01-02")
	fmt.Println(x)

	// The output may be "2009-11-10" because the Playground's clock is fixed in the past.
}

//100. Sort by a comparator
// 根据某个字段排序

type Item struct {
	label string
	p     int
	lang  string
}

// c returns true if x is "inferior to" y (in a custom way)
func c(x, y Item) bool {
	return x.p < y.p
}

type ItemCSorter []Item

func (s ItemCSorter) Len() int           { return len(s) }
func (s ItemCSorter) Less(i, j int) bool { return c(s[i], s[j]) }
func (s ItemCSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func sortItems(items []Item) {
	sorter := ItemCSorter(items)
	sort.Sort(sorter)
}

func main() {
	items := []Item{
		{"twelve", 12, "english"},
		{"six", 6, "english"},
		{"eleven", 11, "english"},
		{"zero", 0, "english"},
		{"two", 2, "english"},
	}
	fmt.Println("Unsorted: ", items)
	sortItems(items)
	fmt.Println("Sorted: ", items)
}

// or

type Item struct {
	label string
	p     int
	lang  string
}

type ItemsSorter struct {
	items []Item
	c     func(x, y Item) bool
}

func (s ItemsSorter) Len() int           { return len(s.items) }
func (s ItemsSorter) Less(i, j int) bool { return s.c(s.items[i], s.items[j]) }
func (s ItemsSorter) Swap(i, j int)      { s.items[i], s.items[j] = s.items[j], s.items[i] }

func sortItems(items []Item, c func(x, y Item) bool) {
	sorter := ItemsSorter{
		items,
		c,
	}
	sort.Sort(sorter)
}

func main() {
	items := []Item{
		{"twelve", 12, "english"},
		{"six", 6, "english"},
		{"eleven", 11, "english"},
		{"zero", 0, "english"},
		{"two", 2, "english"},
	}
	fmt.Println("Unsorted: ", items)

	c := func(x, y Item) bool {
		return x.p < y.p
	}
	sortItems(items, c)

	fmt.Println("Sorted: ", items)
}

// or

type Item struct {
	label string
	p     int
	lang  string
}

// c returns true if x is "inferior to" y (in a custom way)
func c(x, y Item) bool {
	return x.p < y.p
}

func main() {
	items := []Item{
		{"twelve", 12, "english"},
		{"six", 6, "english"},
		{"eleven", 11, "english"},
		{"zero", 0, "english"},
		{"two", 2, "english"},
	}
	fmt.Println("Unsorted: ", items)

	sort.Slice(items, func(i, j int) bool {
		return c(items[i], items[j])
	})

	fmt.Println("Sorted: ", items)
}

//101. Load from HTTP GET request into a string
// 发起http请求

func main() {
	u := "http://" + localhost + "/hello?name=Inigo+Montoya"

	res, err := http.Get(u)
	check(err)
	buffer, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	check(err)
	s := string(buffer)

	fmt.Println("GET  response:", res.StatusCode, s)
}

const localhost = "127.0.0.1:3000"

func init() {
	http.HandleFunc("/hello", myHandler)
	startServer()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.FormValue("name"))
}

func startServer() {
	listener, err := net.Listen("tcp", localhost)
	check(err)

	go http.Serve(listener, nil)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// 102. Load from HTTP GET request into a file

//发起http请求

func main() {
	err := saveGetResponse()
	check(err)
	err = readFile()
	check(err)

	fmt.Println("Done.")
}

func saveGetResponse() error {
	u := "http://" + localhost + "/hello?name=Inigo+Montoya"

	fmt.Println("Making GET request")
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Status: %v", resp.Status)
	}

	fmt.Println("Saving data to file")
	out, err := os.Create("/tmp/result.txt")
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func readFile() error {
	fmt.Println("Reading file")
	buffer, err := ioutil.ReadFile("/tmp/result.txt")
	if err != nil {
		return err
	}
	fmt.Printf("Saved data is %q\n", string(buffer))
	return nil
}

const localhost = "127.0.0.1:3000"

func init() {
	http.HandleFunc("/hello", myHandler)
	startServer()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.FormValue("name"))
}

func startServer() {
	listener, err := net.Listen("tcp", localhost)
	check(err)

	go http.Serve(listener, nil)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//105. Current executable name

//当前可执行文件名称

// 将当前正在执行的程序的名称分配给字符串s(但不是它的完整路径)。

func main() {
	var s string

	path := os.Args[0]
	s = filepath.Base(path)

	fmt.Println(s)
}

// 106. Get program working directory

//获取程序的工作路径

func main() {
	dir, err := os.Getwd()
	fmt.Println(dir, err)
}

//107. Get folder containing current program

//获取包含当前程序的文件夹

func main() {
	var dir string

	programPath := os.Args[0]
	absolutePath, err := filepath.Abs(programPath)
	if err != nil {
		panic(err)
	}
	dir = filepath.Dir(absolutePath)

	fmt.Println(dir)
}

//109. Number of bytes of a type
// 获取某个类型的字节数

func main() {
	var t T
	tType := reflect.TypeOf(t)
	n := tType.Size()

	fmt.Println("A", tType, "object is", n, "bytes.")
}

type Person struct {
	FirstName string
	Age       int
}

// T is a type alias, to stick to the idiom statement.
// T has the same memory footprint per value as Person.
type T Person

// 110. Check if string is blank
// 检查字符串是否空白

func main() {
	for _, s := range []string{
		"",
		"a",
		" ",
		"\t \n",
		"_",
	} {
		blank := strings.TrimSpace(s) == ""

		if blank {
			fmt.Printf("%q is blank\n", s)
		} else {
			fmt.Printf("%q is not blank\n", s)
		}
	}
}

// 111. Launch other program
// 运行其他程序

func main() {
	err := exec.Command("x", "a", "b").Run()
	fmt.Println(err)
}

//112. Iterate over map entries, ordered by keys
// 遍历map，按key排序

func main() {
	mymap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	keys := make([]string, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		x := mymap[k]
		fmt.Println("Key =", k, ", Value =", x)
	}
}

//113. Iterate over map entries, ordered by values
// 遍历map，按值排序

type entry struct {
	key   string
	value int
}
type entries []entry

func (list entries) Len() int           { return len(list) }
func (list entries) Less(i, j int) bool { return list[i].value < list[j].value }
func (list entries) Swap(i, j int)      { list[i], list[j] = list[j], list[i] }

func main() {
	mymap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"dos":   2,
		"deux":  2,
	}

	entries := make(entries, 0, len(mymap))
	for k, x := range mymap {
		entries = append(entries, entry{key: k, value: x})
	}
	sort.Sort(entries)

	for _, e := range entries {
		fmt.Println("Key =", e.key, ", Value =", e.value)
	}
}

// or
func main() {
	mymap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"dos":   2,
		"deux":  2,
	}

	type entry struct {
		key   string
		value int
	}

	entries := make([]entry, 0, len(mymap))
	for k, x := range mymap {
		entries = append(entries, entry{key: k, value: x})
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].value < entries[j].value
	})

	for _, e := range entries {
		fmt.Println("Key =", e.key, ", Value =", e.value)
	}
}

//114. Test deep equality
// 深度判等

func main() {
	x := Foo{9, "Hello", []bool{false, true}, map[int]float64{1: 1.0, 2: 2.0}, &Bar{"Babar"}}

	list := []Foo{
		{9, "Bye", []bool{false, true}, map[int]float64{1: 1.0, 2: 2.0}, &Bar{"Babar"}},
		{9, "Hello", []bool{false, false}, map[int]float64{1: 1.0, 2: 2.0}, &Bar{"Babar"}},
		{9, "Hello", []bool{false, true}, map[int]float64{1: 3.0, 2: 2.0}, &Bar{"Babar"}},
		{9, "Hello", []bool{false, true}, map[int]float64{1: 1.0, 5: 2.0}, &Bar{"Babar"}},
		{9, "Hello", []bool{false, true}, map[int]float64{1: 1.0, 2: 2.0}, &Bar{"Batman"}},
		{9, "Hello", []bool{false, true}, map[int]float64{1: 1.0, 2: 2.0}, &Bar{"Babar"}},
	}

	for i, y := range list {
		b := reflect.DeepEqual(x, y)
		if b {
			fmt.Println("x deep equals list[", i, "]")

		} else {
			fmt.Println("x doesn't deep equal list[", i, "]")
		}
	}
}

type Foo struct {
	int
	str     string
	bools   []bool
	mapping map[int]float64
	bar     *Bar
}

type Bar struct {
	name string
}

//115. Compare dates
// 日期比较

func main() {
	d1 := time.Now()
	d2 := time.Date(2020, time.November, 10, 23, 0, 0, 0, time.UTC)

	b := d1.Before(d2)

	fmt.Println(b)
}

// 116. Remove occurrences of word from string
// 去除指定字符串

func main() {
	var s1 = "foobar"
	var w = "foo"
	s2 := strings.Replace(s1, w, "", -1)
	fmt.Println(s2)
}

// 117. Get list size
// 获取list的大小

func main() {
	// x is a slice
	x := []string{"a", "b", "c"}
	n := len(x)
	fmt.Println(n)

	// y is an array
	y := [4]string{"a", "b", "c"}
	n = len(y)
	fmt.Println(n)
}

//118. List to set
// 从list到set

func main() {
	x := []string{"b", "a", "b", "c"}
	fmt.Println("x =", x)

	y := make(map[string]struct{}, len(x))
	for _, v := range x {
		y[v] = struct{}{}
	}

	fmt.Println("y =", y)
}

// 119. Deduplicate list
// list去重

func main() {
	type T string
	x := []T{"b", "a", "b", "c"}
	fmt.Println("x =", x)

	y := make(map[T]struct{}, len(x))
	for _, v := range x {
		y[v] = struct{}{}
	}
	x2 := make([]T, 0, len(y))
	for _, v := range x {
		if _, ok := y[v]; ok {
			x2 = append(x2, v)
			delete(y, v)
		}
	}
	x = x2

	fmt.Println("x =", x)
}

// or
func main() {
	type T string
	x := []T{"b", "a", "b", "b", "c", "b", "a"}
	fmt.Println("x =", x)

	seen := make(map[T]bool)
	j := 0
	for _, v := range x {
		if !seen[v] {
			x[j] = v
			j++
			seen[v] = true
		}
	}
	x = x[:j]

	fmt.Println("x =", x)
}

// or
type T *int64

func main() {
	var a, b, c, d int64 = 11, 22, 33, 11
	x := []T{&b, &a, &b, &b, &c, &b, &a, &d}
	print(x)

	seen := make(map[T]bool)
	j := 0
	for _, v := range x {
		if !seen[v] {
			x[j] = v
			j++
			seen[v] = true
		}
	}
	for i := j; i < len(x); i++ {
		// Avoid memory leak
		x[i] = nil
	}
	x = x[:j]

	// Now x has only distinct pointers (even if some point to int64 values that are the same)
	print(x)
}

func print(a []T) {
	glue := ""
	for _, p := range a {
		fmt.Printf("%s%d", glue, *p)
		glue = ", "
	}
	fmt.Println()
}

//120. Read integer from stdin

// 从标准输入中读取整数

// This string simulates the keyboard entry.
var userInput string = `42 017`

func main() {
	var i int
	_, err := fmt.Scan(&i)
	fmt.Println(i, err)

	// The second value starts with 0, thus is interpreted as octal!
	var j int
	_, err = fmt.Scan(&j)
	fmt.Println(j, err)
}

// The Go Playground doesn't actually read os.Stdin, so this
// workaround writes some data on virtual FS in a file, and then
// sets this file as the new Stdin.
//
// Note that the init func is run before main.
func init() {
	err := ioutil.WriteFile("/tmp/stdin", []byte(userInput), 0644)
	if err != nil {
		panic(err)
	}
	fileIn, err := os.Open("/tmp/stdin")
	if err != nil {
		panic(err)
	}
	os.Stdin = fileIn
}

// or

// This string simulates the keyboard entry.
var userInput string = `42 017`

func main() {
	var i int
	_, err := fmt.Scanf("%d", &i)
	fmt.Println(i, err)

	var j int
	_, err = fmt.Scanf("%d", &j)
	fmt.Println(j, err)
}

// The Go Playground doesn't actually read os.Stdin, so this
// workaround writes some data on virtual FS in a file, and then
// sets this file as the new Stdin.
//
// Note that the init func is run before main.
func init() {
	err := ioutil.WriteFile("/tmp/stdin", []byte(userInput), 0644)
	if err != nil {
		panic(err)
	}
	fileIn, err := os.Open("/tmp/stdin")
	if err != nil {
		panic(err)
	}
	os.Stdin = fileIn
}

//121. UDP listen and read
// 听端口p上的UDP流量，并将1024字节读入缓冲区b。

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", p)
	if err != nil {
		return err
	}
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		return err
	}
	defer ServerConn.Close()
	n, addr, err := ServerConn.ReadFromUDP(b[:1024])
	if err != nil {
		return err
	}
	if n < 1024 {
		return fmt.Errorf("Only %d bytes could be read.", n)
	}
}

//122. Declare enumeration
// 声明枚举值

type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

func main() {
	fmt.Printf("Hearts has type %T and value %d", Hearts, Hearts)
}

//123. Assert condition
//断言条件

func main() {
	salary = 65000
	employees = 120000
	totalPayroll = salary * employees

	if !isConsistent() {
		panic("State consistency violated")
	}
	fmt.Println("Everything fine")
}

var salary int32
var employees int32
var totalPayroll int32

func isConsistent() bool {
	return salary >= 0 &&
		employees >= 0 &&
		totalPayroll >= 0
}

//124. Binary search for a value in sorted array
// 排序数组中值的二分搜索法

// 二分查找

func binarySearch(a []T, x T) int {
	imin, imax := 0, len(a)-1
	for imin <= imax {
		imid := (imin + imax) / 2
		switch {
		case a[imid] == x:
			return imid
		case a[imid] < x:
			imin = imid + 1
		default:
			imax = imid - 1
		}
	}
	return -1
}

type T int

func main() {
	a := []T{-2, -1, 0, 1, 1, 1, 6, 8, 8, 9, 10}
	for x := T(-5); x <= 15; x++ {
		i := binarySearch(a, x)
		if i == -1 {
			fmt.Println("Value", x, "not found")
		} else {
			fmt.Println("Value", x, "found at index", i)
		}
	}
}

// or

func binarySearch(a []int, x int) int {
	i := sort.SearchInts(a, x)
	if i < len(a) && a[i] == x {
		return i
	}
	return -1
}

func main() {
	a := []int{-2, -1, 0, 1, 1, 1, 6, 8, 8, 9, 10}
	for x := -5; x <= 15; x++ {
		i := binarySearch(a, x)
		if i == -1 {
			fmt.Println("Value", x, "not found")
		} else {
			fmt.Println("Value", x, "found at index", i)
		}
	}
}

// or

func binarySearch(a []T, x T) int {
	f := func(i int) bool { return a[i] >= x }
	i := sort.Search(len(a), f)
	if i < len(a) && a[i] == x {
		return i
	}
	return -1
}

type T int

func main() {
	a := []T{-2, -1, 0, 1, 1, 1, 6, 8, 8, 9, 10}
	for x := T(-5); x <= 15; x++ {
		i := binarySearch(a, x)
		if i == -1 {
			fmt.Println("Value", x, "not found")
		} else {
			fmt.Println("Value", x, "found at index", i)
		}
	}
}

//125. Measure function call duration
// 函数调用时间

func main() {
	t1 := time.Now()
	foo()
	t := time.Since(t1)
	ns := int64(t / time.Nanosecond)

	// Note that the clock is fixed in the Playground, so the resulting duration is always zero
	fmt.Printf("%dns\n", ns)
}

func foo() {
	fmt.Println("Hello")
}

// or

func main() {
	t1 := time.Now()
	foo()
	t := time.Since(t1)
	ns := t.Nanoseconds()
	fmt.Printf("%dns\n", ns)
}

func foo() {
	fmt.Println("Hello")
}

//126. Multiple return values
// 多个返回值

func main() {
	s, b := foo()
	fmt.Println(s, b)
}

func foo() (string, bool) {
	return "Too good to be", true
}

//128. Breadth-first traversing of a tree
// 树的广度优先遍历

func (root *Tree) Bfs(f func(*Tree)) {
	if root == nil {
		return
	}
	queue := []*Tree{root}
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		f(t)
		queue = append(queue, t.Children...)
	}
}

type key string
type value string

type Tree struct {
	Key      key
	Deco     value
	Children []*Tree
}

func (this *Tree) AddChild(x key, v value) {
	child := &Tree{Key: x, Deco: v}
	this.Children = append(this.Children, child)
}

func NodePrint(node *Tree) {
	fmt.Printf("%v (%v)\n", node.Key, node.Deco)
}

func main() {
	tree := &Tree{Key: "World", Deco: "Our planet"}
	tree.AddChild("Europe", "A continent")
	tree.Children[0].AddChild("Germany", "A country")
	tree.Children[0].AddChild("Ireland", "A country")
	tree.Children[0].AddChild("Mediterranean Sea", "A sea")
	tree.AddChild("Asia", "A continent")
	tree.Children[0].AddChild("Japan", "A country")
	tree.Children[0].AddChild("Thailand", "A country")

	tree.Bfs(NodePrint)
}

//129. Breadth-first traversing in a graph
// 图的广度优先遍历

func (start *Vertex) Bfs(f func(*Vertex)) {
	queue := []*Vertex{start}
	seen := map[*Vertex]bool{start: true}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		f(v)
		for next, isEdge := range v.Neighbours {
			if isEdge && !seen[next] {
				queue = append(queue, next)
				seen[next] = true
			}
		}
	}
}

type Vertex struct {
	Id         int
	Label      string
	Neighbours map[*Vertex]bool
}

type Graph []*Vertex

func NewVertex(id int, label string) *Vertex {
	return &Vertex{
		Id:         id,
		Label:      label,
		Neighbours: make(map[*Vertex]bool),
	}
}

func (v *Vertex) AddNeighbour(w *Vertex) {
	v.Neighbours[w] = true
}

func VertexPrint(v *Vertex) {
	fmt.Printf("%v (%v)\n", v.Id, v.Label)
}

func main() {
	// Some cities
	london := NewVertex(0, "London")
	ny := NewVertex(1, "New York City")
	berlin := NewVertex(2, "Berlin")
	paris := NewVertex(3, "Paris")
	tokyo := NewVertex(4, "Tokyo")

	g := Graph{
		london,
		ny,
		berlin,
		paris,
		tokyo,
	}
	_ = g

	london.AddNeighbour(paris)
	london.AddNeighbour(ny)
	ny.AddNeighbour(london)
	ny.AddNeighbour(paris)
	ny.AddNeighbour(tokyo)
	tokyo.AddNeighbour(paris)
	paris.AddNeighbour(tokyo)
	paris.AddNeighbour(berlin)

	london.Bfs(VertexPrint)
}

// 130. Depth-first traversing in a graph
// 图的深度优先遍历

func (v *Vertex) Dfs(f func(*Vertex), seen map[*Vertex]bool) {
	seen[v] = true
	f(v)
	for next, isEdge := range v.Neighbours {
		if isEdge && !seen[next] {
			next.Dfs(f, seen)
		}
	}
}

type Vertex struct {
	Id         int
	Label      string
	Neighbours map[*Vertex]bool
}

type Graph []*Vertex

func NewVertex(id int, label string) *Vertex {
	return &Vertex{
		Id:         id,
		Label:      label,
		Neighbours: make(map[*Vertex]bool),
	}
}

func (v *Vertex) AddNeighbour(w *Vertex) {
	v.Neighbours[w] = true
}

func VertexPrint(v *Vertex) {
	fmt.Printf("%v (%v)\n", v.Id, v.Label)
}

func main() {
	// Some cities
	london := NewVertex(0, "London")
	ny := NewVertex(1, "New York City")
	berlin := NewVertex(2, "Berlin")
	paris := NewVertex(3, "Paris")
	tokyo := NewVertex(4, "Tokyo")

	g := Graph{
		london,
		ny,
		berlin,
		paris,
		tokyo,
	}
	_ = g

	london.AddNeighbour(paris)
	london.AddNeighbour(ny)
	ny.AddNeighbour(london)
	ny.AddNeighbour(paris)
	ny.AddNeighbour(tokyo)
	tokyo.AddNeighbour(paris)
	paris.AddNeighbour(tokyo)
	paris.AddNeighbour(berlin)

	alreadySeen := map[*Vertex]bool{}
	london.Dfs(VertexPrint, alreadySeen)
}

//131. Successive conditions

//连续条件判等

func conditional(x string) {
	switch {
	case c1(x):
		f1()
	case c2(x):
		f2()
	case c3(x):
		f3()
	}
}

func main() {
	conditional("dog Snoopy")
	conditional("fruit Raspberry")
}

func f1() {
	fmt.Println("I'm a Human")
}

func f2() {
	fmt.Println("I'm a Dog")
}

func f3() {
	fmt.Println("I'm a Fruit")
}

var c1, c2, c3 = prefixCheck("human"), prefixCheck("dog"), prefixCheck("fruit")

func prefixCheck(prefix string) func(string) bool {
	return func(x string) bool {
		return strings.HasPrefix(x, prefix)
	}
}

//132. Measure duration of procedure execution
// 度量程序执行时间

func clock(f func()) time.Duration {
	t := time.Now()
	f()
	return time.Since(t)
}

func f() {
	re := regexp.MustCompilePOSIX("|A+{300}")
	re.FindAllString(strings.Repeat("A", 299), -1)
}

func main() {
	d := clock(f)

	// The result is always zero in the playground, which has a fixed clock!
	// Try it on your workstation instead.
	fmt.Println(d)
}

//133. Case-insensitive string contains
// 不区分大小写的字符串包含

// Package _strings has no case-insensitive version of _Contains, so
// we have to make our own.
func containsCaseInsensitive(s, word string) bool {
	lowerS, lowerWord := strings.ToLower(s), strings.ToLower(word)
	ok := strings.Contains(lowerS, lowerWord)
	return ok
}

func main() {
	s := "Let's dance the macarena"

	word := "Dance"
	ok := containsCaseInsensitive(s, word)
	fmt.Println(ok)

	word = "dance"
	ok = containsCaseInsensitive(s, word)
	fmt.Println(ok)

	word = "Duck"
	ok = containsCaseInsensitive(s, word)
	fmt.Println(ok)
}

// 134. Create a new list
// 创建一个新list

func main() {
	var a, b, c T = "This", "is", "wonderful"

	items := []T{a, b, c}

	fmt.Println(items)
}

type T string

//135. Remove item from list, by its value
// 移除列表中的值

func main() {
	items := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(items)

	x := "c"
	for i, y := range items {
		if y == x {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
	fmt.Println(items)
}

// or

func main() {
	for i, y := range items {
		if y == x {
			copy(items[i:], items[i+1:])
			items[len(items)-1] = nil
			items = items[:len(items)-1]
			break
		}
	}
}

//136. Remove all occurrences of a value from a list
//从列表中删除所有出现的值

func main() {
	items := []T{"b", "a", "b", "a", "r"}
	fmt.Println(items)

	var x T = "b"
	items2 := make([]T, 0, len(items))
	for _, v := range items {
		if v != x {
			items2 = append(items2, v)
		}
	}

	fmt.Println(items2)
}

type T string

// or

func main() {
	items := []T{"b", "a", "b", "a", "r"}
	fmt.Println(items)

	x := T("b")
	j := 0
	for i, v := range items {
		if v != x {
			items[j] = items[i]
			j++
		}
	}
	items = items[:j]

	fmt.Println(items)
}

type T string

//or

func main() {
	var items []*image
	{
		red := newUniform(rgb{0xFF, 0, 0})
		white := newUniform(rgb{0xFF, 0xFF, 0xFF})
		items = []*image{red, white, red} // Like the flag of Austria
		fmt.Println("items =", items)

		x := red
		j := 0
		for i, v := range items {
			if v != x {
				items[j] = items[i]
				j++
			}
		}
		for k := j; k < len(items); k++ {
			items[k] = nil
		}
		items = items[:j]
	}

	// At this point, red can be garbage collected

	printAllocInfo()

	fmt.Println("items =", items) // Not the original flag anymore...
	fmt.Println("items undelying =", items[:3])
}

type image [1024][1024]rgb
type rgb [3]byte

func newUniform(color rgb) *image {
	im := new(image)
	for x := range im {
		for y := range im[x] {
			im[x][y] = color
		}
	}
	return im
}

func printAllocInfo() {
	var stats runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats)
	fmt.Println("Bytes allocated (total):", stats.TotalAlloc)
	fmt.Println("Bytes still allocated:  ", stats.Alloc)
}

//137. Check if string contains only digits
// 检查字符串是否只包含数字

func main() {
	for _, s := range []string{
		"123",
		"",
		"abc123def",
		"abc",
		"123.456",
		"123 456",
	} {
		b := true
		for _, c := range s {
			if c < '0' || c > '9' {
				b = false
				break
			}
		}
		fmt.Println(s, "=>", b)
	}
}

// or
func main() {
	for _, s := range []string{
		"123",
		"",
		"abc123def",
		"abc",
		"123.456",
		"123 456",
	} {
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		b := strings.IndexFunc(s, isNotDigit) == -1
		fmt.Println(s, "=>", b)
	}
}

//138. Create temp file

// 创建一个新的临时文件

func main() {
	content := []byte("Big bag of misc data")

	log.Println("Opening new temp file")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	tmpfilename := tmpfile.Name()
	defer os.Remove(tmpfilename) // clean up
	log.Println("Opened new file", tmpfilename)

	log.Println("Writing [[", string(content), "]]")
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("Closed", tmpfilename)

	log.Println("Opening", tmpfilename)
	buffer, err := ioutil.ReadFile(tmpfilename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read[[", string(buffer), "]]")
}

//139. Create temp directory
// 创建一个临时目录

func main() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	inspect(dir)

	tmpfn := filepath.Join(dir, "tmpfile")
	err = ioutil.WriteFile(tmpfn, content, 0666)
	if err != nil {
		log.Fatal(err)
	}

	inspect(dir)
}

func inspect(dirpath string) {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirpath, "contains", len(files), "files")
}

// 140. Delete map entry
// 从map中删除某个key

func main() {
	m := map[string]int{
		"uno":  1,
		"dos":  2,
		"tres": 3,
	}

	delete(m, "dos")
	delete(m, "cinco")

	fmt.Println(m)
}

//141. Iterate in sequence over two lists
//依次迭代两个列表 依次迭代列表项1和项2的元素。每次迭代打印元素。

func main() {
	items1 := []string{"a", "b", "c"}
	items2 := []string{"A", "B", "C"}

	for _, v := range items1 {
		fmt.Println(v)
	}
	for _, v := range items2 {
		fmt.Println(v)
	}
}

//142. Hexadecimal digits of an integer
// 将整数x的十六进制表示(16进制)赋给字符串s。

func main() {
	x := int64(999)
	s := strconv.FormatInt(x, 16)

	fmt.Println(s)
}

// or

func main() {
	x := big.NewInt(999)
	s := fmt.Sprintf("%x", x)

	fmt.Println(s)
}

//143. Iterate alternatively over two lists
// 交替迭代两个列表

func main() {
	items1 := []string{"a", "b"}
	items2 := []string{"A", "B", "C"}

	for i := 0; i < len(items1) || i < len(items2); i++ {
		if i < len(items1) {
			fmt.Println(items1[i])
		}
		if i < len(items2) {
			fmt.Println(items2[i])
		}
	}
}

//144. Check if file exists
// 检查文件是否存在

func main() {
	fp := "foo.txt"
	_, err := os.Stat(fp)
	b := !os.IsNotExist(err)
	fmt.Println(fp, "exists:", b)

	fp = "bar.txt"
	_, err = os.Stat(fp)
	b = !os.IsNotExist(err)
	fmt.Println(fp, "exists:", b)
}

func init() {
	ioutil.WriteFile("foo.txt", []byte(`abc`), 0644)
}

//145. Print log line with datetime
//打印带时间的日志

func main() {
	msg := "Hello, playground"
	log.Println(msg)

	// The date is fixed in the past in the Playground, never mind.
}

// See http://www.programming-idioms.org/idiom/145/print-log-line-with-date/1815/go

//146. Convert string to floating point number
// 字符串转换为浮点型

func main() {
	s := "3.1415926535"

	f, err := strconv.ParseFloat(s, 64)
	fmt.Printf("%T, %v, err=%v\n", f, f, err)
}

//
// http://www.programming-idioms.org/idiom/146/convert-string-to-floating-point-number/1819/go
//

//147. Remove all non-ASCII characters
// 移除所有的非ASCII字符

func main() {
	s := "dæmi : пример : příklad : thí dụ"

	re := regexp.MustCompile("[[:^ascii:]]")
	t := re.ReplaceAllLiteralString(s, "")

	fmt.Println(t)
}

//or

func main() {
	s := "5#∑∂ƒ∞645eyfu"
	t := strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return -1
		}
		return r
	}, s)
	fmt.Println(t)
}

//148. Read list of integers from stdin
// 从stdin(标准输入)中读取整数列表

func main() {
	var ints []int
	s := bufio.NewScanner(osStdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err == nil {
			ints = append(ints, i)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ints)
}

// osStdin simulates os.Stdin
// var osStdin = strings.NewReader(`
// 11
// 22
// 33  `)

//150. Remove trailing slash
// 去除末尾的 /

func main() {
	p := "/usr/bin/"

	p = strings.TrimSuffix(p, "/")

	fmt.Println(p)
}

//151. Remove string trailing path separator
// 删除字符串尾部路径分隔符

func main() {
	p := somePath()
	fmt.Println(p)

	sep := fmt.Sprintf("%c", os.PathSeparator)
	p = strings.TrimSuffix(p, sep)

	fmt.Println(p)
}

func somePath() string {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	p := fmt.Sprintf("%s%c%s%c", dir, os.PathSeparator, "foobar", os.PathSeparator)
	return p
}

// or

func main() {
	p := somePath()
	fmt.Println(p)

	sep := fmt.Sprintf("%c", filepath.Separator)
	p = strings.TrimSuffix(p, sep)

	fmt.Println(p)
}

func somePath() string {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	p := fmt.Sprintf("%s%c%s%c", dir, os.PathSeparator, "foobar", os.PathSeparator)
	return p
}

// 152. Turn a character into a string
// 将字符转换成字符串

func main() {
	var c rune = os.PathSeparator
	fmt.Printf("%c \n", c)

	s := fmt.Sprintf("%c", c)
	fmt.Printf("%#v \n", s)
}

//153. Concatenate string with integer
// 连接字符串和整数

func main() {
	s := "Hello"
	i := 123

	t := fmt.Sprintf("%s%d", s, i)

	fmt.Println(t)
}

//154. Halfway between two hex color codes
// 求两个十六进制颜色代码的中间值

// For concision, halfway assume valid inputs.
// Caller must have explicitly checked that c1, c2 are well-formed color codes.
func halfway(c1, c2 string) string {
	r1, _ := strconv.ParseInt(c1[1:3], 16, 0)
	r2, _ := strconv.ParseInt(c2[1:3], 16, 0)
	r := (r1 + r2) / 2

	g1, _ := strconv.ParseInt(c1[3:5], 16, 0)
	g2, _ := strconv.ParseInt(c2[3:5], 16, 0)
	g := (g1 + g2) / 2

	b1, _ := strconv.ParseInt(c1[5:7], 16, 0)
	b2, _ := strconv.ParseInt(c2[5:7], 16, 0)
	b := (b1 + b2) / 2

	c := fmt.Sprintf("#%02X%02X%02X", r, g, b)
	return c
}

func main() {
	c1 := "#15293E"
	c2 := "#012549"

	if err := checkFormat(c1); err != nil {
		panic(fmt.Errorf("Wrong input %q: %v", c1, err))
	}
	if err := checkFormat(c2); err != nil {
		panic(fmt.Errorf("Wrong input %q: %v", c2, err))
	}

	c := halfway(c1, c2)
	fmt.Println("The average of", c1, "and", c2, "is", c)
}

func checkFormat(color string) error {
	if len(color) != 7 {
		return fmt.Errorf("Hex colors have exactly 7 chars")
	}
	if color[0] != '#' {
		return fmt.Errorf("Hex colors start with #")
	}
	isNotDigit := func(c rune) bool { return (c < '0' || c > '9') && (c < 'a' || c > 'f') }
	if strings.IndexFunc(strings.ToLower(color[1:]), isNotDigit) != -1 {
		return fmt.Errorf("Forbidden char")
	}
	return nil
}

// For concision, halfway assume valid inputs.
// Caller must have explicitly checked that c1, c2 are well-formed color codes.
func halfway(c1, c2 string) string {
	var buf [7]byte
	buf[0] = '#'
	for i := 0; i < 3; i++ {
		sub1 := c1[1+2*i : 3+2*i]
		sub2 := c2[1+2*i : 3+2*i]
		v1, _ := strconv.ParseInt(sub1, 16, 0)
		v2, _ := strconv.ParseInt(sub2, 16, 0)
		v := (v1 + v2) / 2
		sub := fmt.Sprintf("%02X", v)
		copy(buf[1+2*i:3+2*i], sub)
	}
	c := string(buf[:])

	return c
}

func main() {
	c1 := "#15293E"
	c2 := "#012549"

	if err := checkFormat(c1); err != nil {
		panic(fmt.Errorf("Wrong input %q: %v", c1, err))
	}
	if err := checkFormat(c2); err != nil {
		panic(fmt.Errorf("Wrong input %q: %v", c2, err))
	}

	c := halfway(c1, c2)
	fmt.Println("The average of", c1, "and", c2, "is", c)
}

func checkFormat(color string) error {
	if len(color) != 7 {
		return fmt.Errorf("Hex colors have exactly 7 chars")
	}
	if color[0] != '#' {
		return fmt.Errorf("Hex colors start with #")
	}
	isNotDigit := func(c rune) bool { return (c < '0' || c > '9') && (c < 'a' || c > 'f') }
	if strings.IndexFunc(strings.ToLower(color[1:]), isNotDigit) != -1 {
		return fmt.Errorf("Forbidden char")
	}
	return nil
}

//155. Delete file
// 删除文件

func main() {
	for _, filepath := range []string{
		"/tmp/foo.txt",
		"/tmp/bar.txt",
		"/tmp/foo.txt",
	} {
		err := os.Remove(filepath)
		if err == nil {
			fmt.Println("Removed", filepath)
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func init() {
	err := ioutil.WriteFile("/tmp/foo.txt", []byte(`abc`), 0644)
	if err != nil {
		panic(err)
	}
}

//156. Format integer with zero-padding
// 用零填充格式化整数

func main() {
	for _, i := range []int{
		0,
		8,
		64,
		256,
		2048,
	} {
		s := fmt.Sprintf("%03d", i)
		fmt.Println(s)
	}
}

//157. Declare constant string

// 声明常量字符串

const planet = "Earth"

func main() {
	fmt.Println("We live on planet", planet)
}

//158. Random sublist
// 随机子列表

func main() {
	type T string

	x := []T{"Alice", "Bob", "Carol", "Dan", "Eve", "Frank", "Grace", "Heidi"}
	k := 4

	y := make([]T, k)
	perm := rand.Perm(len(x))
	for i, v := range perm[:k] {
		y[i] = x[v]
	}

	fmt.Printf("%q", y)
}

// 159. Trie
//前缀树/字典树

type Trie struct {
	c        rune
	children map[rune]*Trie
	isLeaf   bool
	value    V
}

type V int

func main() {
	t := NewTrie(0)
	for s, v := range map[string]V{
		"to":  7,
		"tea": 3,
		"ted": 4,
		"ten": 12,
		"A":   15,
		"i":   11,
		"in":  5,
		"inn": 9,
	} {
		t.insert(s, v)
	}
	fmt.Println(t.startsWith("te", ""))
}

func NewTrie(c rune) *Trie {
	t := new(Trie)
	t.c = c
	t.children = map[rune]*Trie{}
	return t
}

func (t *Trie) insert(s string, value V) {
	if s == "" {
		t.isLeaf = true
		t.value = value
		return
	}
	c, tail := cut(s)
	child, exists := t.children[c]
	if !exists {
		child = NewTrie(c)
		t.children[c] = child
	}
	child.insert(tail, value)
}

func (t *Trie) startsWith(p string, accu string) []string {
	if t == nil {
		return nil
	}
	if p == "" {
		var result []string
		if t.isLeaf {
			result = append(result, accu)
		}
		for c, child := range t.children {
			rec := child.startsWith("", accu+string(c))
			result = append(result, rec...)
		}
		return result
	}
	c, tail := cut(p)
	return t.children[c].startsWith(tail, accu+string(c))
}

func cut(s string) (head rune, tail string) {
	r, size := utf8.DecodeRuneInString(s)
	return r, s[size:]
}

//160. Detect if 32-bit or 64-bit architecture
// 检测是32位还是64位架构

func main() {
	if strconv.IntSize == 32 {
		f32()
	}
	if strconv.IntSize == 64 {
		f64()
	}
}

func f32() {
	fmt.Println("I am 32-bit")
}

func f64() {
	fmt.Println("I am 64-bit")
}

//161. Multiply all the elements of a list
// 将list中的每个元素都乘以一个数

func main() {
	const c = 5.5
	elements := []float64{2, 4, 9, 30}
	fmt.Println(elements)

	for i := range elements {
		elements[i] *= c
	}
	fmt.Println(elements)
}

//162. Execute procedures depending on options
// 根据选项执行程序

func init() {
	// Just for testing in the Playground, let's simulate
	// the user called this program with command line
	// flags -f and -b
	os.Args = []string{"program", "-f", "-b"}
}

var b = flag.Bool("b", false, "Do bat")
var f = flag.Bool("f", false, "Do fox")

func main() {
	flag.Parse()
	if *b {
		bar()
	}
	if *f {
		fox()
	}
	fmt.Println("The end.")
}

func bar() {
	fmt.Println("BAR")
}

func fox() {
	fmt.Println("FOX")
}

//163. Print list elements by group of 2
// 两个一组打印数组元素

func main() {
	list := []string{"a", "b", "c", "d", "e", "f"}

	for i := 0; i+1 < len(list); i += 2 {
		fmt.Println(list[i], list[i+1])
	}
}

//164. Open URL in default browser
// 在默认浏览器中打开链接

// import "github.com/skratchdot/open-golang/open"
// b := open.Start(s) == nil

// or
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

///165. Last element of list
// 列表中的最后一个元素

func main() {
	items := []string{"what", "a", "mess"}

	x := items[len(items)-1]

	fmt.Println(x)
}

// 166. Concatenate two lists
// 连接两个列表

func main() {
	a := []string{"The ", "quick "}
	b := []string{"brown ", "fox "}

	ab := append(a, b...)

	fmt.Printf("%q", ab)
}

//or

func main() {
	type T string

	a := []T{"The ", "quick "}
	b := []T{"brown ", "fox "}

	var ab []T
	ab = append(append(ab, a...), b...)

	fmt.Printf("%q", ab)
}

/// or

func main() {
	type T string

	a := []T{"The ", "quick "}
	b := []T{"brown ", "fox "}

	ab := make([]T, len(a)+len(b))
	copy(ab, a)
	copy(ab[len(a):], b)

	fmt.Printf("%q", ab)
}

//167. Trim prefix

// 移除前缀

func main() {
	s := "café-society"
	p := "café"

	t := strings.TrimPrefix(s, p)

	fmt.Println(t)
}

// 168. Trim suffix
// 移除后缀
func main() {
	s := "café-society"
	w := "society"

	t := strings.TrimSuffix(s, w)

	fmt.Println(t)
}

// 169. String length
// 字符串长度

func main() {
	s := "Hello, 世界"
	n := utf8.RuneCountInString(s)

	fmt.Println(n)
}

//170. Get map size
// 获取map的大小

func main() {
	mymap := map[string]int{"a": 1, "b": 2, "c": 3}
	n := len(mymap)
	fmt.Println(n)
}

//171. Add an element at the end of a list
// 在list尾部添加元素

func main() {
	s := []int{1, 1, 2, 3, 5, 8, 13}
	x := 21

	s = append(s, x)

	fmt.Println(s)
}

//172. Insert entry in map
// 向map中写入元素

func main() {
	m := map[string]int{"one": 1, "two": 2}
	k := "three"
	v := 3

	m[k] = v

	fmt.Println(m)
}

//173. Format a number with grouped thousands
// 按千位格式化数字

// import (
//  "fmt"

//  "golang.org/x/text/language"
//  "golang.org/x/text/message"
// )

// The Playground doesn't work with import of external packages.
// However, you may copy this source and test it on your workstation.

func main() {
	p := message.NewPrinter(language.English)
	s := p.Sprintf("%d\n", 1000)

	fmt.Println(s)
	// Output:
	// 1,000
}

//import (
//  "fmt"
//  "github.com/floscodes/golang-thousands"
//  "strconv"
// )

// The Playground takes more time when importing external packages.
// However, you may want to copy this source and test it on your workstation.

func main() {
	n := strconv.Itoa(23489)
	s := thousands.Separate(n, "en")

	fmt.Println(s)
	// Output:
	// 23,489
}

//174. Make HTTP POST request
// 发起http POST请求

func main() {
	contentType := "text/plain"
	var body io.Reader
	u := "http://" + localhost + "/hello"

	response, err := http.Post(u, contentType, body)
	check(err)
	buffer, err := ioutil.ReadAll(response.Body)
	check(err)
	fmt.Println("POST response:", response.StatusCode, string(buffer))

	response, err = http.Get(u)
	check(err)
	buffer, err = ioutil.ReadAll(response.Body)
	check(err)
	fmt.Println("GET  response:", response.StatusCode, string(buffer))
}

const localhost = "127.0.0.1:3000"

func init() {
	http.HandleFunc("/hello", myHandler)
	startServer()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Refusing request verb %q", r.Method)
		return
	}
	fmt.Fprintf(w, "Hello POST :)")
}

func startServer() {
	listener, err := net.Listen("tcp", localhost)
	check(err)

	go http.Serve(listener, nil)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// or

func main() {
	formValues := url.Values{
		"who": []string{"Alice"},
	}
	u := "http://" + localhost + "/hello"

	response, err := http.PostForm(u, formValues)
	check(err)
	buffer, err := ioutil.ReadAll(response.Body)
	check(err)
	fmt.Println("POST response:", response.StatusCode, string(buffer))

	response, err = http.Get(u)
	check(err)
	buffer, err = ioutil.ReadAll(response.Body)
	check(err)
	fmt.Println("GET  response:", response.StatusCode, string(buffer))
}

const localhost = "127.0.0.1:3000"

func init() {
	http.HandleFunc("/hello", myHandler)
	startServer()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Refusing request verb %q", r.Method)
		return
	}
	fmt.Fprintf(w, "Hello %s (POST)", r.FormValue("who"))
}

func startServer() {
	listener, err := net.Listen("tcp", localhost)
	check(err)

	go http.Serve(listener, nil)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//175. Bytes to hex string
// 字节转十六进制字符串

func main() {
	a := []byte("Hello")

	s := hex.EncodeToString(a)

	fmt.Println(s)
}

// 176. Hex string to byte array
//  十六进制字符串转字节数组

func main() {
	s := "48656c6c6f"

	a, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)
	fmt.Println(string(a))
}

//178. Check if point is inside rectangle
// 检查点是否在矩形内

func main() {
	x1, y1, x2, y2 := 1, 1, 50, 100
	r := image.Rect(x1, y1, x2, y2)

	x, y := 10, 10
	p := image.Pt(x, y)
	b := p.In(r)
	fmt.Println(b)

	x, y = 100, 100
	p = image.Pt(x, y)
	b = p.In(r)
	fmt.Println(b)
}

//179. Get center of a rectangle
// 获取矩形的中心

//import "image"
// c := image.Pt((x1+x2)/2, (y1+y2)/2)

// 180. List files in directory
// 列出目录中的文件

func main() {
	d := "/"

	x, err := ioutil.ReadDir(d)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range x {
		fmt.Println(f.Name())
	}
}

//182. Quine program
// 输出程序的源代码

func main() {
	fmt.Printf("%s%c%s%c\n", s, 0x60, s, 0x60)
}

var s = `package main

import "fmt"

func main() {
 fmt.Printf("%s%c%s%c\n", s, 0x60, s, 0x60)
}

var s = `

func main() {
	fmt.Printf("%s%c%s%c\n", s, 0x60, s, 0x60)
}

var s = `package main

import "fmt"

func main() {
 fmt.Printf("%s%c%s%c\n", s, 0x60, s, 0x60)
}

var s = `

//另一种写法：

//go:embed 入门[2]

// Quine 是一种可以输出自身源码的程序。利用 go:embed 我们可以轻松实现 quine 程序：

// //
// import (
//     _ "embed"
//     "fmt"
// )

// //go:embed quine.go
// var src string

// func main() {
//     fmt.Print(src)
// }

//184. Tomorrow

//明天的日期

func main() {
	t := time.Now().Add(24 * time.Hour).Format("2006-01-02")
}

//185. Execute function in 30 seconds
// 30秒内执行功能

func main() {
	timer := time.AfterFunc(
		30*time.Second,
		func() {
			f(42)
		})
}

//or

func main() {
	fmt.Println("Scheduling f(42)")

	go func() {
		time.Sleep(3 * time.Second)
		f(42)
	}()

	// Poor man's waiting of completion of f.
	// Don't do this in prod, use proper synchronization instead.
	time.Sleep(4 * time.Second)
}

func f(i int) {
	fmt.Println("Received", i)
}

//186. Exit program cleanly
// 干净地退出程序

func main() {
	fmt.Println("A")
	os.Exit(0)
	fmt.Println("B")
}

// or

func main() {
	process1()
	process2()
	process3()
}

func process1() {
	fmt.Println("process 1")
}

func process2() {
	fmt.Println("process 2")
	defer fmt.Println("A")
	defer os.Exit(0)
	defer fmt.Println("B")
	fmt.Println("C")
}

func process3() {
	fmt.Println("process 3")
}

//189. Filter and transform list

// 过滤和转换列表

func P(e int) bool {
	// Predicate "is even"
	return e%2 == 0
}

type Result = int

func T(e int) Result {
	// Transformation "square"
	return e * e
}

func main() {
	x := []int{4, 5, 6, 7, 8, 9, 10}

	var y []Result
	for _, e := range x {
		if P(e) {
			y = append(y, T(e))
		}
	}

	fmt.Println(y)
}

//190. Call an external C function

// 调用外部C函数

// void foo(double *a, int n);
// double a[] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
// import "C"

// C.foo(C.a, 10)

//191. Check if any value in a list is larger than a limit
//检查列表中是否有任何值大于限制

func f() {
	fmt.Println("Larger found")
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	x := 4
	for _, v := range a {
		if v > x {
			f()
			break
		}
	}
}

//192. Declare a real variable with at least 20 digits

// 声明一个至少有20位数字的实变量

func main() {
	a, _, err := big.ParseFloat("123456789.123456789123465789", 10, 200, big.ToZero)
	if err != nil {
		panic(err)
	}

	fmt.Println(a)
}

//197. Get a list of lines from a file
// 从文件中获取行列表.将文件路径中的内容检索到字符串行列表中，其中每个元素都是文件的一行。

func readLines(path string) ([]string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	return lines, nil
}

func main() {
	lines, err := readLines("/tmp/file1")
	if err != nil {
		log.Fatalln(err)
	}

	for i, line := range lines {
		fmt.Printf("line %d: %s\n", i, line)
	}
}

func init() {
	data := []byte(`foo
bar
baz`)
	err := ioutil.WriteFile("/tmp/file1", data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

//198. Abort program execution with error condition
// 出现错误情况时中止程序执行

func main() {
	x := 1
	os.Exit(x)
}

//200. Return hypotenuse
// 返回三角形的斜边h，其中与直角相邻的边的长度为x和y。

func main() {
	x := 1.0
	y := 1.0

	h := math.Hypot(x, y)
	fmt.Println(h)
}

//202. Sum of squares

// 计算平方和

func main() {
	data := []float64{0.06, 0.82, -0.01, -0.34, -0.55}
	var s float64
	for _, d := range data {
		s += math.Pow(d, 2)
	}
	println(s)
}

//205. Get an environment variable
// 获取环境变量

func main() {
	foo, ok := os.LookupEnv("FOO")
	if !ok {
		foo = "none"
	}

	fmt.Println(foo)
}

//or

func main() {
	foo := os.Getenv("FOO")
	if foo == "" {
		foo = "none"
	}

	fmt.Println(foo)
}

//206. Switch statement with strings

// switch语句

func main() {
	str := "baz"

	switch str {
	case "foo":
		foo()
	case "bar":
		bar()
	case "baz":
		baz()
	case "barfl":
		barfl()
	}
}

func foo() {
	fmt.Println("Called foo")
}

func bar() {
	fmt.Println("Called bar")
}

func baz() {
	fmt.Println("Called baz")
}

func barfl() {
	fmt.Println("Called barfl")
}

//207. Allocate a list that is automatically deallocated

// 分配一个自动解除分配的列表

type T byte

func main() {
	n := 10_000_000
	a := make([]T, n)
	fmt.Println(len(a))
}

//208. Formula with arrays
// 对数组元素进行运算

func applyFormula(a, b, c, d []float64, e float64) {
	for i, v := range a {
		a[i] = e * (v + b[i] + c[i] + math.Cos(d[i]))
	}
}

func main() {
	a := []float64{1, 2, 3, 4}
	b := []float64{5.5, 6.6, 7.7, 8.8}
	c := []float64{9, 10, 11, 12}
	d := []float64{13, 14, 15, 16}
	e := 42.0

	fmt.Println("a is    ", a)
	applyFormula(a, b, c, d, e)
	fmt.Println("a is now", a)
}

//209. Type with automatic deep deallocation
// 自动深度解除分配的类型

func f() {
	type t struct {
		s string
		n []int
	}

	v := t{
		s: "Hello, world!",
		n: []int{1, 4, 9, 16, 25},
	}

	// Pretend to use v (otherwise this is a compile error)
	_ = v

	// When f exits, v and all its fields are garbage-collected, recursively
}

//211. Create folder
// 创建文件夹

func main() {
	path := "foo"
	_, err := os.Stat(path)
	b := !os.IsNotExist(err)
	fmt.Println(path, "exists:", b)

	err = os.Mkdir(path, os.ModeDir)
	if err != nil {
		panic(err)
	}

	info, err2 := os.Stat(path)
	b = !os.IsNotExist(err2)
	fmt.Println(path, "exists:", b)
	fmt.Println(path, "is a directory:", info.IsDir())
}

// or

func main() {
	path := "foo/bar"
	_, err := os.Stat(path)
	b := !os.IsNotExist(err)
	fmt.Println(path, "exists:", b)

	err = os.Mkdir(path, os.ModeDir)
	if err != nil {
		fmt.Println("Could not create", path, "with os.Mkdir")
	}

	info, err2 := os.Stat(path)
	b = !os.IsNotExist(err2)
	fmt.Println(path, "exists:", b)

	err = os.MkdirAll(path, os.ModeDir)
	if err != nil {
		fmt.Println("Could not create", path, "with os.MkdirAll")
	}

	info, err2 = os.Stat(path)
	b = !os.IsNotExist(err2)
	fmt.Println(path, "exists:", b)
	fmt.Println(path, "is a directory:", info.IsDir())
}

//212. Check if folder exists
// 检查文件夹是否存在

func main() {
	path := "foo"
	info, err := os.Stat(path)
	b := !os.IsNotExist(err) && info.IsDir()
	fmt.Println(path, "is a directory:", b)

	err = os.Mkdir(path, os.ModeDir)
	if err != nil {
		panic(err)
	}

	info, err = os.Stat(path)
	b = !os.IsNotExist(err) && info.IsDir()
	fmt.Println(path, "is a directory:", b)
}

//215. Pad string on the left
// 左侧补齐字符串

func main() {
	m := 3
	c := "-"
	for _, s := range []string{
		"",
		"a",
		"ab",
		"abc",
		"abcd",
		"é",
	} {
		if n := utf8.RuneCountInString(s); n < m {
			s = strings.Repeat(c, m-n) + s
		}
		fmt.Println(s)
	}
}

//217. Create a Zip archive
// 创建压缩文件

func main() {
	list := []string{
		"readme.txt",
		"gopher.txt",
		"todo.txt",
	}
	name := "archive.zip"

	err := makeZip(list, name)
	if err != nil {
		log.Fatal(err)
	}
}

func makeZip(list []string, name string) error {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	for _, filename := range list {
		// Open file for reading
		input, err := os.Open(filename)
		if err != nil {
			return err
		}
		// Create ZIP entry for writing
		output, err := w.Create(filename)
		if err != nil {
			return err
		}

		_, err = io.Copy(output, input)
		if err != nil {
			return err
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		return err
	}

	N := buf.Len()
	err = ioutil.WriteFile(name, buf.Bytes(), 0777)
	if err != nil {
		return err
	}
	log.Println("Written a ZIP file of", N, "bytes")

	return nil
}

func init() {
	// Create some files in the filesystem.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		err := ioutil.WriteFile(file.Name, []byte(file.Body), 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//218. List intersection
// 列表的交集

type T int

func main() {
	a := []T{4, 5, 6, 7, 8, 9, 10}
	b := []T{1, 3, 9, 5, 7, 9, 7, 7}

	// Convert to sets
	seta := make(map[T]bool, len(a))
	for _, x := range a {
		seta[x] = true
	}
	setb := make(map[T]bool, len(a))
	for _, y := range b {
		setb[y] = true
	}

	// Iterate in one pass
	var c []T
	for x := range seta {
		if setb[x] {
			c = append(c, x)
		}
	}

	fmt.Println(c)
}

//219. Replace multiple spaces with single space
// 用单个空格替换多个空格

// regexp created only once, and then reused
var whitespaces = regexp.MustCompile(`\s+`)

func main() {
	s := `
 one   two
    three
 `

	t := whitespaces.ReplaceAllString(s, " ")

	fmt.Printf("t=%q", t)
}

//220. Create a tuple value

// 创建元组值a

func main() {
	t := []interface{}{
		2.5,
		"hello",
		make(chan int),
	}
}

// 221. Remove all non-digits characters
// 删除所有非数字字符

func main() {
	s := `height="168px"`

	re := regexp.MustCompile("[^\\d]")
	t := re.ReplaceAllLiteralString(s, "")

	fmt.Println(t)
}

//222. Find first index of an element in list
// 在列表中查找元素的第一个索引

func main() {
	items := []string{"huey", "dewey", "louie"}
	x := "dewey"

	i := -1
	for j, e := range items {
		if e == x {
			i = j
			break
		}
	}

	fmt.Printf("Found %q at position %d in %q", x, i, items)
}

//223. for else loop
// for else循环

func main() {
	items := []string{"foo", "bar", "baz", "qux"}

	for _, item := range items {
		if item == "baz" {
			fmt.Println("found it")
			goto forelse
		}
	}
	{
		fmt.Println("never found it")
	}
forelse:
}

//224. Add element to the beginning of the list

// 将元素添加到列表的开头

type T int

func main() {
	items := []T{42, 1337}
	var x T = 7

	items = append([]T{x}, items...)

	fmt.Println(items)
}

// or

type T int

func main() {
	items := []T{42, 1337}
	var x T = 7

	items = append(items, x)
	copy(items[1:], items)
	items[0] = x

	fmt.Println(items)
}

//225. Declare and use an optional argument

// 声明并使用可选参数

func f(x ...int) {
	if len(x) > 0 {
		println("Present", x[0])
	} else {
		println("Not present")
	}
}

func main() {
	f()
	f(1)
}

//226. Delete last element from list
// 从列表中删除最后一个元素

func main() {
	items := []string{"banana", "apple", "kiwi"}
	fmt.Println(items)

	items = items[:len(items)-1]
	fmt.Println(items)
}

//227. Copy list

// 复制列表

func main() {
	type T string
	x := []T{"Never", "gonna", "shower"}

	y := make([]T, len(x))
	copy(y, x)

	y[2] = "give"
	y = append(y, "you", "up")

	fmt.Println(x)
	fmt.Println(y)
}

//228. Copy a file

// 复制文件

func main() {
	src, dst := "/tmp/file1", "/tmp/file2"

	err := copy(dst, src)
	if err != nil {
		log.Fatalln(err)
	}

	stat, err := os.Stat(dst)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dst, "exists, it has size", stat.Size(), "and mode", stat.Mode())
}

func copy(dst, src string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, data, stat.Mode())
}

func init() {
	data := []byte("Hello")
	err := ioutil.WriteFile("/tmp/file1", data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

// or

func main() {
	src, dst := "/tmp/file1", "/tmp/file2"

	err := copy(dst, src)
	if err != nil {
		log.Fatalln(err)
	}

	stat, err := os.Stat(dst)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dst, "exists, it has size", stat.Size(), "and mode", stat.Mode())
}

func copy(dst, src string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, data, stat.Mode())
	if err != nil {
		return err
	}
	return os.Chmod(dst, stat.Mode())
}

func init() {
	data := []byte("Hello")
	err := ioutil.WriteFile("/tmp/file1", data, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.Chmod("/tmp/file1", 0777)
	if err != nil {
		log.Fatalln(err)
	}
}

// or

func main() {
	src, dst := "/tmp/file1", "/tmp/file2"

	err := copy(dst, src)
	if err != nil {
		log.Fatalln(err)
	}

	stat, err := os.Stat(dst)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dst, "exists, it has size", stat.Size(), "and mode", stat.Mode())
}

func copy(dst, src string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		return err
	}
	g, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, stat.Mode())
	if err != nil {
		return err
	}
	defer g.Close()
	_, err = io.Copy(g, f)
	if err != nil {
		return err
	}
	return os.Chmod(dst, stat.Mode())
}

func init() {
	data := []byte("Hello")
	err := ioutil.WriteFile("/tmp/file1", data, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.Chmod("/tmp/file1", 0777)
	if err != nil {
		log.Fatalln(err)
	}
}

//231. Test if bytes are a valid UTF-8 string
// 测试字节是否是有效的UTF-8字符串

func main() {
	{
		s := []byte("Hello, 世界")
		b := utf8.Valid(s)
		fmt.Println(b)
	}
	{
		s := []byte{0xff, 0xfe, 0xfd}
		b := utf8.Valid(s)
		fmt.Println(b)
	}
}

///234. Encode bytes to base64

//将字节编码为base64

func main() {
	data := []byte("Hello world")
	s := base64.StdEncoding.EncodeToString(data)
	fmt.Println(s)
}

//235. Decode base64
// 解码base64

func main() {
	str := "SGVsbG8gd29ybGQ="

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%q\n", data)
}

//237. Xor integers
// 异或运算

func main() {
	a, b := 230, 42
	c := a ^ b

	fmt.Printf("a is %12b\n", a)
	fmt.Printf("b is %12b\n", b)
	fmt.Printf("c is %12b\n", c)
	fmt.Println("c ==", c)
}

// or

func main() {
	a, b := big.NewInt(230), big.NewInt(42)
	c := new(big.Int)
	c.Xor(a, b)

	fmt.Printf("a is %12b\n", a)
	fmt.Printf("b is %12b\n", b)
	fmt.Printf("c is %12b\n", c)
	fmt.Println("c ==", c)
}

//238. Xor byte arrays

//异或字节数组

func main() {
	a, b := []byte("Hello"), []byte("world")

	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}

	fmt.Printf("a is %08b\n", a)
	fmt.Printf("b is %08b\n", b)
	fmt.Printf("c is %08b\n", c)
	fmt.Println("c ==", c)
	fmt.Printf("c as string would be %q\n", string(c))
}

// or

type T [5]byte

func main() {
	var a, b T
	copy(a[:], "Hello")
	copy(b[:], "world")

	var c T
	for i := range a {
		c[i] = a[i] ^ b[i]
	}

	fmt.Printf("a is %08b\n", a)
	fmt.Printf("b is %08b\n", b)
	fmt.Printf("c is %08b\n", c)
	fmt.Println("c ==", c)
	fmt.Printf("c as string would be %q\n", string(c[:]))
}

// 239. Find first regular expression match

// 查找第一个正则表达式匹配项

func main() {
	re := regexp.MustCompile(`\b\d\d\d\b`)
	for _, s := range []string{
		"",
		"12",
		"123",
		"1234",
		"I have 12 goats, 3988 otters, 224 shrimps and 456 giraffes",
		"See p.456, for word boundaries",
	} {
		x := re.FindString(s)
		fmt.Printf("%q -> %q\n", s, x)
	}
}

// 240. Sort 2 lists together
// 将两个列表排序在一起.列表a和b的长度相同。对a和b应用相同的排列，根据a的值对它们进行排序。
type K int
type T string

type sorter struct {
	k []K
	t []T
}

func (s *sorter) Len() int {
	return len(s.k)
}

func (s *sorter) Swap(i, j int) {
	// Swap affects 2 slices at once.
	s.k[i], s.k[j] = s.k[j], s.k[i]
	s.t[i], s.t[j] = s.t[j], s.t[i]
}

func (s *sorter) Less(i, j int) bool {
	return s.k[i] < s.k[j]
}

func main() {
	a := []K{9, 3, 4, 8}
	b := []T{"nine", "three", "four", "eight"}

	sort.Sort(&sorter{
		k: a,
		t: b,
	})

	fmt.Println(a)
	fmt.Println(b)
}

//241. Yield priority to other threads
// 将优先权让给其他线程

func main() {
	go fmt.Println("aaa")
	go fmt.Println("bbb")
	go fmt.Println("ccc")
	go fmt.Println("ddd")
	go fmt.Println("eee")

	runtime.Gosched()
	busywork()

	time.Sleep(100 * time.Millisecond)
}

func busywork() {
	fmt.Println("main")
}

//242. Iterate over a set
// 迭代一个集合

type T string

func main() {
	// declare a Set (implemented as a map)
	x := make(map[T]bool)

	// add some elements
	x["A"] = true
	x["B"] = true
	x["B"] = true
	x["C"] = true
	x["D"] = true

	// remove an element
	delete(x, "C")

	for e := range x {
		f(e)
	}
}

func f(e T) {
	fmt.Printf("contains element %v \n", e)
}

//243. Print list
// 打印 list

func main() {
	{
		a := []int{11, 22, 33}
		fmt.Println(a)
	}

	{
		a := []string{"aa", "bb"}
		fmt.Println(a)
	}

	{
		type Person struct {
			First string
			Last  string
		}
		x := Person{
			First: "Jane",
			Last:  "Doe",
		}
		y := Person{
			First: "John",
			Last:  "Doe",
		}
		a := []Person{x, y}
		fmt.Println(a)
	}

	{
		x, y := 11, 22
		a := []*int{&x, &y}
		fmt.Println(a)
	}
}

//244. Print map
// 打印 map

func main() {
	{
		m := map[string]int{
			"eleven":     11,
			"twenty-two": 22,
		}
		fmt.Println(m)
	}

	{
		x, y := 7, 8
		m := map[string]*int{
			"seven": &x,
			"eight": &y,
		}
		fmt.Println(m)
	}
}

//245. Print value of custom type

// 打印自定义类型的值

// T represents a tank. It doesn't implement fmt.Stringer.
type T struct {
	name      string
	weight    int
	firePower int
}

// Person implement fmt.Stringer.
type Person struct {
	FirstName   string
	LastName    string
	YearOfBirth int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, born %d", p.FirstName, p.LastName, p.YearOfBirth)
}

func main() {
	{
		x := T{
			name:      "myTank",
			weight:    100,
			firePower: 90,
		}

		fmt.Println(x)
	}
	{
		x := Person{
			FirstName:   "John",
			LastName:    "Doe",
			YearOfBirth: 1958,
		}

		fmt.Println(x)
	}
}

//246. Count distinct elements
// 计算不同的元素的数量

func main() {
	type T string
	items := []T{"a", "b", "b", "aaa", "c", "a", "d"}
	fmt.Println("items has", len(items), "elements")

	distinct := make(map[T]bool)
	for _, v := range items {
		distinct[v] = true
	}
	c := len(distinct)

	fmt.Println("items has", c, "distinct elements")
}

//247. Filter list in-place
// 就地筛选列表

type T int

func main() {
	x := []T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := func(t T) bool { return t%2 == 0 }

	j := 0
	for i, v := range x {
		if p(v) {
			x[j] = x[i]
			j++
		}
	}
	x = x[:j]

	fmt.Println(x)
}

// or

type T int

func main() {
	var x []*T
	for _, v := range []T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		t := new(T)
		*t = v
		x = append(x, t)
	}
	p := func(t *T) bool { return *t%2 == 0 }

	j := 0
	for i, v := range x {
		if p(v) {
			x[j] = x[i]
			j++
		}
	}
	for k := j; k < len(x); k++ {
		x[k] = nil
	}
	x = x[:j]

	for _, pt := range x {
		fmt.Print(*pt, " ")
	}
}

// 249. Declare and assign multiple variables
// 声明并分配多个变量

func main() {
	// a, b and c don't need to have the same type.

	a, b, c := 42, "hello", 5.0

	fmt.Println(a, b, c)
	fmt.Printf("%T %T %T \n", a, b, c)
}

//251. Parse binary digits
// 解析二进制数字

func main() {
	s := "1101"
	fmt.Println("s is", reflect.TypeOf(s), s)

	i, err := strconv.ParseInt(s, 2, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("i is", reflect.TypeOf(i), i)
}

//252. Conditional assignment
// 条件赋值

func main() {
	var x string
	if condition() {
		x = "a"
	} else {
		x = "b"
	}

	fmt.Println(x)
}

func condition() bool {
	return "Socrates" == "dog"
}

//258. Convert list of strings to list of integers
// 将字符串列表转换为整数列表

func main() {
	a := []string{"11", "22", "33"}

	b := make([]int, len(a))
	var err error
	for i, s := range a {
		b[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(b)
}

//259. Split on several separators
// 在几个分隔符上拆分

func main() {
	s := "2021-03-11,linux_amd64"

	re := regexp.MustCompile("[,\\-_]")
	parts := re.Split(s, -1)

	fmt.Printf("%q", parts)
}

//266. Repeating string
//重复字符串

func main() {
	v := "abc"
	n := 5

	s := strings.Repeat(v, n)

	fmt.Println(s)
}
