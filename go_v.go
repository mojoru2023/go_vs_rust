package main

// https://mp.weixin.qq.com/s/ES8RBASjUdCrsvV3ErceGg
//1. Print Hello World

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
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
