
//1. Print Hello World



fn main() {
    println!("Hello World");
}



// 2. Print Hello 10 times
fn main() {
    for _ in 0..10 {
        println!("Hello");
    }
}

//or 

fn main() {
    print!("{}", "Hello\n".repeat(10));
}

//3. Create a procedure

fn main(){
    finish("Buddy")
}

fn finish(name : &str) {
    println!("My job here is done. Goodbye {}", name);
}


//4. Create a function which returns the square of an integer


fn square(x: u32) -> u32 {
    x * x
}

fn main() {
    let sq = square(9);

    println!("{}", sq);
}

// 5. Create a 2D Point data structure
//声明一个容器类型,有x、y两个浮点数


use std::fmt;

struct Point {
    x: f64,
    y: f64,
}

impl fmt::Display for Point {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}

fn main() {
    let p = Point { x: 2.0, y: -3.5 };

    println!("{}", p);
}


// or 

use std::fmt;

struct Point(f64, f64);

impl fmt::Display for Point {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({}, {})", self.0, self.1)
    }
}

fn main() {
    let p = Point(2.0, -3.5);

    println!("{}", p);
}

// 6. Iterate over list values

//遍历列表的值



fn main() {
    let items = vec![11, 22, 33];

    for x in items {
        do_something(x);
    }
}

fn do_something(n: i64) {
    println!("Number {}", n)
}


// or 

fn main() {
    let items = vec![11, 22, 33];

    items.into_iter().for_each(|x| do_something(x));
}

fn do_something(n: i64) {
    println!("Number {}", n)
}

// 7. Iterate over list indexes and values
// 遍历列表的索引和值


fn main() {
    let items = ["a", "b", "c"];
    for (i, x) in items.iter().enumerate() {
        println!("Item {} = {}", i, x);
    }
}


// or 


fn main() {
    let items = ["a", "b", "c"];
    items.iter().enumerate().for_each(|(i, x)| {
        println!("Item {} = {}", i, x);
    });
}


//8. Initialize a new map (associative array)
//创建一个新的map,提供一些键值对 作为初始内容

use std::collections::BTreeMap;

fn main() {
    let mut x = BTreeMap::new();
    x.insert("one", 1);
    x.insert("two", 2);
    
    println!("{:?}", x);
}


// or 

use std::collections::HashMap;

fn main() {
    let x: HashMap<&str, i32> = [
        ("one", 1),
        ("two", 2),
    ].iter().cloned().collect();
    
    println!("{:?}", x);
}

//9. Create a Binary Tree data structure
// 创建一个二叉树

struct BinTree<T> {
    value: T,
    left: Option<Box<BinTree<T>>>,
    right: Option<Box<BinTree<T>>>,
}

// 10. Shuffle a list

// 随机排序一个list


extern crate rand;
use rand::{Rng, StdRng};

let mut rng = StdRng::new().unwrap();
rng.shuffle(&mut x);


// or 


use rand::seq::SliceRandom;
use rand::thread_rng;

fn main() {
    let mut x = [1, 2, 3, 4, 5];
    println!("Unshuffled: {:?}", x);

    let mut rng = thread_rng();
    x.shuffle(&mut rng);

    println!("Shuffled:   {:?}", x);
}


// 11. Pick a random element from a list
// 从列表中选择一个随机元素

use rand::{self, Rng};

fn main() {
    let x = vec![11, 22, 33];

    let choice = x[rand::thread_rng().gen_range(0..x.len())];

    println!("I picked {}!", choice);
}


//or 


use rand::seq::SliceRandom;
 
fn main() {
    let x = vec![11, 22, 33];

    let mut rng = rand::thread_rng();
    let choice = x.choose(&mut rng).unwrap();
    
    println!("I picked {}!", choice);
}

//12. Check if list contains a value
// 检查列表中是否包含一个值


fn main() {
    let list = [10, 40, 30];

    {
        let num = 30;

        if list.contains(&num) {
            println!("{:?} contains {}", list, num);
        } else {
            println!("{:?} doesn't contain {}", list, num);
        }
    }

    {
        let num = 42;

        if list.contains(&num) {
            println!("{:?} contains {}", list, num);
        } else {
            println!("{:?} doesn't contain {}", list, num);
        }
    }
}


// or 


fn main() {
    let list = [10, 40, 30];
    let x = 30;

    if list.iter().any(|v| v == &x) {
        println!("{:?} contains {}", list, x);
    } else {
        println!("{:?} doesn't contain {}", list, x);
    }
}


// or 

fn main() {
    let list = [10, 40, 30];
    let x = 30;

    if (&list).into_iter().any(|v| v == &x) {
        println!("{:?} contains {}", list, x);
    } else {
        println!("{:?} doesn't contain {}", list, x);
    }
}

// 13. Iterate over map keys and values

//  遍历关联数组中的每一对 k-v， 并打印出它们

use std::collections::BTreeMap;

fn main() {
    let mut mymap = BTreeMap::new();
    mymap.insert("one", 1);
    mymap.insert("two", 2);
    mymap.insert("three", 3);
    mymap.insert("four", 4);

    for (k, x) in &mymap {
        println!("Key={key}, Value={val}", key = k, val = x);
    }
}


//14. Pick uniformly a random floating point number in [a..b)
// 选出一个随机的浮点数，大于或等于a，严格小于b，且a< b


extern crate rand;
use rand::{thread_rng, Rng};

fn main() {
    let (a, b) = (1.0, 3.0);
    let c = thread_rng().gen_range(a..b);
    println!("{}", c);
}

//15. Pick uniformly a random integer in [a..b]
// 选出一个随机的整数，大于或等于a，小于或等于b，且a< b

fn pick(a: i32, b: i32) -> i32 {
    let between = Range::new(a, b);
    let mut rng = rand::thread_rng();
    between.ind_sample(&mut rng)
}


// or 


use rand::distributions::Distribution;
use rand::distributions::Uniform;

fn main() {
    let (a, b) = (3, 5);

    let x = Uniform::new_inclusive(a, b).sample(&mut rand::thread_rng());

    println!("{}", x);
}
// 17. Create a Tree data structure

// 创建树数据结构, 该结构必须是递归的。一个节点可以有零个或多个子节点,节点可以访问子节点，但不能访问其父节点



use std::vec;

struct Node<T> {
    value: T,
    children: Vec<Node<T>>,
}

impl<T> Node<T> {
    pub fn dfs<F: Fn(&T)>(&self, f: F) {
       self.dfs_helper(&f);
    }

    fn dfs_helper<F: Fn(&T)>(&self, f: &F) {
        (f)(&self.value);
        for child in &self.children {
            child.dfs_helper(f);
        }
    }
}


fn main() {
    let t: Node<i32> = Node {
        children: vec![
            Node {
                children: vec![
                    Node {
                        children: vec![],
                        value: 14
                    }
                ],
                value: 28
            },
            Node {
                children: vec![],
                value: 80
            }
        ],
        value: 50
    };

    t.dfs(|node| { println!("{}", node); });
}

// 18. Depth-first traversing of a tree
// 树的深度优先遍历。按照深度优先的前缀顺序，在树的每个节点上调用函数f


use std::vec;

struct Tree<T> {
    children: Vec<Tree<T>>,
    value: T
}

impl<T> Tree<T> {
    pub fn new(value: T) -> Self{
        Tree{
            children: vec![],
            value
        }
    }

    pub fn dfs<F: Fn(&T)>(&self, f: F) {
       self.dfs_helper(&f);
    }

    fn dfs_helper<F: Fn(&T)>(&self, f: &F) {
        (f)(&self.value);
        for child in &self.children {
            child.dfs_helper(f);
        }
    }
}


fn main() {
    let t: Tree<i32> = Tree {
        children: vec![
            Tree {
                children: vec![
                    Tree {
                        children: vec![],
                        value: 14
                    }
                ],
                value: 28
            },
            Tree {
                children: vec![],
                value: 80
            }
        ],
        value: 50
    };

    t.dfs(|node| { println!("{}", node); });
}


//19. Reverse a list
//反转链表

fn main() {
    let x = vec!["Hello", "World"];
    let y: Vec<_> = x.iter().rev().collect();
    println!("{:?}", y);
}


// or


fn main() {
    let mut x = vec![1,2,3];
    x.reverse();
    println!("{:?}", x);
}

// 20. Return two values
// 实现在2D矩阵m中寻找元素x，返回匹配单元格的索引 i，j


fn search<T: Eq>(m: &Vec<Vec<T>>, x: &T) -> Option<(usize, usize)> {
    for (i, row) in m.iter().enumerate() {
        for (j, column) in row.iter().enumerate() {
            if *column == *x {
                return Some((i, j));
            }
        }
    }
    
    None
}

fn main() {
    let a = vec![
        vec![0, 11],
        vec![22, 33],
        vec![44, 55],
    ];
    
    let hit = search(&a, &33);
    
    println!("{:?}", hit);
}


//21. Swap values
//交换变量a和b的值

fn main() {
    let a = 3;
    let b = 10;

    let (a, b) = (b, a);

    println!("a: {a}, b: {b}", a=a, b=b);
}

//or 


fn main() {
    let (a, b) = (12, 42);
    
    println!("a = {}, b = {}", a, b);
    
    let (a, b) = (b, a);
    
    println!("a = {}, b = {}", a, b);
}

// 22. Convert string to integer
//将字符串转换为整型


fn main() {
    // This prints 123
    let mut s = "123";
    let mut i = s.parse::<i32>().unwrap();
    println!("{:?}", i);

    // This panics
    s = "12u3";
    i = s.parse::<i32>().unwrap();
    println!("{:?}", i);
}

//or 

fn main() {
    let mut s = "123";
    let mut i: i32 = s.parse().unwrap_or(0);
    println!("{:?}", i);

    s = "12u3";
    i = s.parse().unwrap_or(0);
    println!("{:?}", i);
}

//or 

fn main() {
    let mut s = "123";
    let mut i = match s.parse::<i32>() {
        Ok(i) => i,
        Err(_e) => -1,
    };
    println!("{:?}", i);

    s = "12u3";
    i = match s.parse::<i32>() {
        Ok(i) => i,
        Err(_e) => -1,
    };
    println!("{:?}", i);
}

//23. Convert real number to string with 2 decimal places
// 给定一个实数，小数点后保留两位小数


fn main() {
    let x = 42.1337;
    let s = format!("{:.2}", x);
    
    println!("{}", s);
}

//24. Assign to string the japanese word ネコ
//声明一个新的字符串s，并用文字值“ネコ”初始化它(在日语中是“cat”的意思)
fn main() {
    let s = "ネコ";
    println!("{}", s);
}

//25. Send a value to another thread
//将字符串值“Alan”与现有的正在运行的进程共享，该进程将显示“你好，Alan”

use std::thread;
use std::sync::mpsc::channel;

fn main() {
    let (send, recv) = channel();

    let handle = thread::spawn(move || loop {
        let msg = recv.recv().unwrap();
        println!("Hello, {:?}", msg);
    });

    send.send("Alan").unwrap();
    
    handle.join().unwrap();
}


//26. Create a 2-dimensional array

// 创建一个二维数组

// 声明并初始化一个有m行n列的矩阵x，包含实数。


fn main() {
    const M: usize = 4;
    const N: usize = 6;

    let x = vec![vec![0.0f64; N]; M];
    
    println!("{:#?}", x);
}

//or 

fn main() {
  const M: usize = 3;
  const N: usize = 4;

  let mut x = [[0.0; N] ; M];

  x[1][3] = 5.0;
  println!("{:#?}", x);
}

// 27. Create a 3-dimensional array
// 创建一个三维数组

// 声明并初始化一个三维数组x，它有m，n，p维边界，并且包含实数。

fn main() {
    let m = 4;
    let n = 6;
    let p = 2;

    let x = vec![vec![vec![0.0f64; p]; n]; m];
    
    println!("{:#?}", x);
}

//or 

fn main() {
    const M: usize = 4;
    const N: usize = 6;
    const P: usize = 2;

    let x = [[[0.0f64; P]; N]; M];

    println!("{:#?}", x);
}


//28. Sort by a property

// 按x->p的升序对类似数组的集合项的元素进行排序，其中p是项中对象的项类型的字段。


#[derive(Debug)]
struct Foo {
    p: i32,
}

fn main() {
    let mut items = vec![Foo { p: 3 }, Foo { p: 1 }, Foo { p: 2 }, Foo { p: 4 }];

    items.sort_by(|a, b| a.p.cmp(&b.p));

    println!("{:?}", items);
}

//or 

#[derive(Debug)]
struct Foo {
    p: i32,
}

fn main() {
    let mut items = vec![Foo { p: 3 }, Foo { p: 1 }, Foo { p: 2 }, Foo { p: 4 }];

    items.sort_by_key(|x| x.p);

    println!("{:?}", items);
}


//29. Remove item from list, by its index

// 从列表项中删除第I项。这将改变原来的列表或返回一个新的列表，这取决于哪个更习惯。请注意，在大多数语言中，I的最小有效值是0。

fn main() {
    let mut v = vec![1, 2, 3];
    assert_eq!(v.remove(1), 2);
    assert_eq!(v, [1, 3]);
    
}

//30. Parallelize execution of 1000 independent tasks

// 用参数I从1到1000启动程序f的并发执行。任务是独立的，f(i)不返回值。任务不需要同时运行，所以可以使用pools



use std::thread;

fn main() {
    let threads: Vec<_> = (0..1000).map(|i| thread::spawn(move || f(i))).collect();

    for thread in threads {
        thread.join();
    }
}

fn f(i: i32) {
    println!("{}", i);
}


// or 

extern crate rayon;

use rayon::prelude::*;

fn main() {
    (0..1000).into_par_iter().for_each(f);
}

fn f(i: i32) {
    println!("{}", i);
}


//31. Recursive factorial (simple)
// 创建递归函数f，该函数返回从f(i-1)计算的非负整数I的阶乘

fn f(n: u32) -> u32 {
    if n < 2 {
        1
    } else {
        n * f(n - 1)
    }
}

fn main() {
    println!("{}", f(4 as u32));
}



//or 

fn factorial(num: u64) -> u64 {
    match num {
        0 | 1=> 1,
        _ => factorial(num - 1) * num,
    }
}

fn main (){
    println!("{}", factorial(0));
    println!("{}", factorial(1));
    println!("{}", factorial(2));
    println!("{}", factorial(3));
    println!("{}", factorial(4));
    println!("{}", factorial(5));
}


//32. Integer exponentiation by squaring

// 创建函数exp，计算(快速)x次方n的值。x和n是非负整数。


fn exp(x: u64, n: u64) -> u64 {
    match n {
        0 => 1,
        1 => x,
        i if i % 2 == 0 => exp(x * x, n / 2),
        _ => x * exp(x * x, (n - 1) / 2),
    }
}

fn main() {
    let x = 16;
    let n = 4;
    
    println!("{}", exp(x, n));
}

//33. Atomically read and update variable
//为变量x分配新值f(x)，确保在读和写之间没有其他线程可以修改x。

use std::sync::Mutex;

fn f(x: i32) -> i32 {
    x + 1
}

fn main() {
    let x = Mutex::new(0);
    let mut x = x.lock().unwrap();
    *x = f(*x);
    
    println!("{:?}", *x);
}

//34. Create a set of objects
// 声明并初始化一个包含t类型对象的集合x。



use std::collections::HashSet;

fn main() {
    let mut m = HashSet::new();
    m.insert("a");
    m.insert("b");

    println!("{:?}", m);
}





//35. First-class function : compose[2]
// 用参数f (A -> B)和g (B -> C)实现一个函数compose (A -> C)，返回composition函数g ∘ f
fn compose<'a, A, B, C, G, F>(f: F, g: G) -> Box<Fn(A) -> C + 'a>
 where F: 'a + Fn(A) -> B, G: 'a + Fn(B) -> C
{
 Box::new(move |x| g(f(x)))
}


// or 


fn compose<A, B, C>(f: impl Fn(A) -> B, g: impl Fn(B) -> C) -> impl Fn(A) -> C {
 move |x| g(f(x))
}

fn main() {
    let f = |x: u32| (x * 2) as i32;
    let g = |x: i32| (x + 1) as f32;
    let c = compose(f, g);
    
    println!("{}", c(2));
}

//36. First-class function : generic composition

// 实现一个函数组合，该函数组合为任何恰好有1个参数的函数f和g返回组合函数g ∘ f。


fn compose<'a, A, B, C, G, F>(f: F, g: G) -> Box<Fn(A) -> C + 'a>
 where F: 'a + Fn(A) -> B, G: 'a + Fn(B) -> C
{
 Box::new(move |x| g(f(x)))
}

// or 
fn compose<A, B, C>(f: impl Fn(A) -> B, g: impl Fn(B) -> C) -> impl Fn(A) -> C {
 move |x| g(f(x))
}

fn main() {
    let f = |x: u32| (x * 2) as i32;
    let g = |x: i32| (x + 1) as f32;
    let c = compose(f, g);
    
    println!("{}", c(2));
}

//37. Currying
//将一个接受多个参数的函数转换为一个预设了某些参数的函数。



fn add(a: u32, b: u32) -> u32 {
    a + b
}

fn main() {
    let add5 = move |x| add(5, x);

    let y = add5(12);
    println!("{}", y);
}

// 38. Extract a substring

// 查找由字符串s的字符I(包括)到j(不包括)组成的子字符串t。除非另有说明，字符索引从0开始。确保正确处理多字节字符。

extern crate unicode_segmentation;
use unicode_segmentation::UnicodeSegmentation;

fn main() {
    let s = "Lorem Ipsüm Dolor";
    let (i, j) = (6, 11);
    let t = s.graphemes(true).skip(i).take(j - i).collect::<String>();
    println!("{}", t);
}

// or 

use substring::Substring;
let t = s.substring(i, j);


//39. Check if string contains a word
// 如果字符串单词作为子字符串包含在字符串s中，则将布尔ok设置为true，否则设置为false。

fn main() {
    let s = "Let's dance the macarena";

    let word = "dance";
    let ok = s.contains(word);
    println!("{}", ok);

    let word = "car";
    let ok = s.contains(word);
    println!("{}", ok);

    let word = "duck";
    let ok = s.contains(word);
    println!("{}", ok);
}

//41. Reverse a string
//反转字符串

let t = s.chars().rev().collect::<String>();


// or 
fn main() {
    let s = "lorém ipsüm dolör sit amor ❤ ";
    let t: String = s.chars().rev().collect();
    println!("{}", t);
}

// 42. Continue outer loop

// 打印列表a中不包含在列表b中的每个项目v。 为此，编写一个外部循环来迭代a，编写一个内部循环来迭代b。


fn main() {
    let a = ['a', 'b', 'c', 'd', 'e'];
    let b = [     'b',      'd'     ];
    
    'outer: for va in &a {
        for vb in &b {
            if va == vb {
                continue 'outer;
            }
        }
        println!("{}", va);
    }
}

//43. Break outer loop
//在2D整数矩阵m中寻找一个负值v，打印出来，停止搜索。

fn main() {
    let m = vec![
        vec![1, 2, 3],
        vec![11, 0, 30],
        vec![5, -20, 55],
        vec![0, 0, -60],
    ];
    
    'outer: for v in m {
        'inner: for i in v {
            if i < 0 {
                println!("Found {}", i);
                break 'outer;
            }
        }
    }
}

//44. Insert element in list

// 在列表s的位置I插入元素x。其他元素必须向右移动。

fn main() {
    let mut vec = vec![1, 2, 3];
    vec.insert(1, 4);
    assert_eq!(vec, [1, 4, 2, 3]);
    vec.insert(4, 5);
    assert_eq!(vec, [1, 4, 2, 3, 5]);
    
}

//45. Pause execution for 5 seconds
// 在继续下一个指令之前，在当前线程中休眠5秒钟。


use std::{thread, time};
thread::sleep(time::Duration::from_secs(5));

//46. Extract beginning of string (prefix)
//创建由字符串s的前5个字符组成的字符串t。 确保正确处理多字节字符。


fn main() {
    let s = "été 😁 torride";
    
    let t = s.char_indices().nth(5).map_or(s, |(i, _)| &s[..i]);

    println!("{}", t);
}


//47. Extract string suffix
//创建由字符串s的最后5个字符组成的字符串t。 确保正确处理多字节字符

fn main() {
    let s = "tükörfúrógép";
    let last5ch = s.chars().count() - 5;
    let s2: String = s.chars().skip(last5ch).collect();
    println!("{}", s2);
}


//48. Multi-line string literal
// 给变量s赋值一个由几行文本组成的字符串，包括换行符。

fn main() {
    let s = "line 1
line 2
line 3";
    
    print!("{}", &s);
}

// or 

fn main() {
    let s = r#"Huey
Dewey
Louie"#;
    
    print!("{}", &s);
}



//49. Split a space-separated string
// 拆分用空格分隔的字符串
//构建由输入字符串的子字符串组成的列表块，由一个或多个空格字符分隔。

fn main() {
    let s = "What a  mess";

    let chunks: Vec<_> = s.split_whitespace().collect();

    println!("{:?}", chunks);
}


//or 


fn main() {
    let s = "What a  mess";

    let chunks: Vec<_> = s.split_ascii_whitespace().collect();

    println!("{:?}", chunks);
}

//or 
fn main() {
    let s = "What a  mess";

    let chunks: Vec<_> = s.split(' ').collect();

    println!("{:?}", chunks);
}

//50. Make an infinite loop

// 写一个无限循环

// loop {
//  // Do something
// }
//51. Check if map contains key
// 检查map是否有某个key


use std::collections::HashMap;

fn main() {
    let mut m = HashMap::new();
    m.insert(1, "a");
    m.insert(2, "b");

    let k = 2;

    let hit = m.contains_key(&k);

    println!("{:?}", hit);
}

//52. Check if map contains value
// 检查map中是否有某个值

use std::collections::BTreeMap;

fn main() {
    let mut m = BTreeMap::new();
    m.insert(11, "one");
    m.insert(22, "twenty-two");

    {
        let v = "eight";
        let does_contain = m.values().any(|&val| *val == *v);
        println!("{:?}", does_contain);
    }

    {
        let v = "twenty-two";
        let does_contain = m.values().any(|&val| *val == *v);
        println!("{:?}", does_contain);
    }
}


//52. Check if map contains value
// 检查map中是否有某个值

fn main() {
    let x = vec!["Lorem", "ipsum", "dolor", "sit", "amet"];
    let y = x.join(", ");
    println!("{}", y);
}


//53. Join a list of strings
// 字符串连接

fn main() {
    let x = vec!["Lorem", "ipsum", "dolor", "sit", "amet"];
    let y = x.join(", ");
    println!("{}", y);
}

// 54. Compute sum of integers
// 计算整数之和

fn main() {
    let x: Vec<usize> = (0..=10_000).collect();
    
    eprintln!("Sum of 0-10,000 = {}", x.iter().sum::<usize>())
}

// 55. Convert integer to string
// 将整数转换为字符串


fn main() {
    let i = 123;
    let s = i.to_string();

    println!("{}", s);
}

// or 

fn main() {
    let i = 123;
    let s = format!("{}", i);

    println!("{}", s);
}



//56. Launch 1000 parallel tasks and wait for completion

// 创建1000个并行任务，并等待其完成


use std::thread;

fn f(i: i32) {
    i + 1;
}

fn main() {
    let threads: Vec<_> = (0..10).map(|i| thread::spawn(move || f(i))).collect();

    for t in threads {
     t.join();
    }
}


//57. Filter list
// 过滤list中的值


fn main() {
    let x = vec![1, 2, 3, 4, 5, 6];

    let y: Vec<_> = x.iter()
        .filter(|&x| x % 2 == 0)
        .collect();

    println!("{:?}", y);
}


// 58. Extract file content to a string
// 提取字符串的文件内容


use std::fs::File;
use std::io::prelude::*;

fn main() -> Result<(), ()> {
    let f = "Cargo.toml";

    let mut file = File::open(f).expect("Can't open file.");
    let mut lines = String::new();
    file.read_to_string(&mut lines)
        .expect("Can't read file contents.");

    println!("{}", lines);

    Ok(())
}


// or 


use std::fs;

fn main() {
    let f = "Cargo.toml";

    let lines = fs::read_to_string(f).expect("Can't read file.");

    println!("{}", lines);
}


//59. Write to standard error stream
// 写入标准错误流


fn main() {
    let x = -3;
    eprintln!("{} is negative", x);
}

//60. Read command line argument

//读取命令行参数

use std::env;

fn main() {
    let first_arg = env::args().skip(1).next();

    let fallback = "".to_owned();
    let x = first_arg.unwrap_or(fallback);
    
    println!("{:?}", x);
}

//61. Get current date
// 获取当前时间



extern crate time;
let d = time::now();


// or 

use std::time::SystemTime;

fn main() {
    let d = SystemTime::now();
    println!("{:?}", d);
}

// 62. Find substring position
// 字符串查找

// 查找子字符串位置

fn main() {
    let x = "été chaud";
    
    let y = "chaud";
    let i = x.find(y);
    println!("{:?}", i);
    
    let y = "froid";
    let i = x.find(y);
    println!("{:?}", i);
}


//63. Replace fragment of a string
//替换字符串片段

fn main() {
    let x = "lorem ipsum dolor lorem ipsum";
    let y = "lorem";
    let z = "LOREM";

    let x2 = x.replace(&y, &z);
    
    println!("{}", x2);
}

// 64. Big integer : value 3 power 247
// 超大整数

extern crate num;
use num::bigint::ToBigInt;

fn main() {
    let a = 3.to_bigint().unwrap();
    let x = num::pow(a, 247);
    println!("{}", x)
}



// 65. Format decimal number
// 格式化十进制数

fn main() {
    let x = 0.15625f64;
    let s = format!("{:.1}%", 100.0 * x);
    
    println!("{}", s);
}

// 66. Big integer exponentiation
// 大整数幂运算

extern crate num;

use num::bigint::BigInt;

fn main() {
    let x = BigInt::parse_bytes(b"600000000000", 10).unwrap();
    let n = 42%
}


// 67. Binomial coefficient "n choose k"
// Calculate binom(n, k) = n! / (k! * (n-k)!). Use an integer type able to handle huge numbers.

// 二项式系数“n选择k”



extern crate num;

use num::bigint::BigInt;
use num::bigint::ToBigInt;
use num::traits::One;

fn binom(n: u64, k: u64) -> BigInt {
    let mut res = BigInt::one();
    for i in 0..k {
        res = (res * (n - i).to_bigint().unwrap()) /
              (i + 1).to_bigint().unwrap();
    }
    res
}

fn main() {
    let n = 133;
    let k = 71;

    println!("{}", binom(n, k));
}



// 68. Create a bitset
// 创建位集合



fn main() {
    let n = 20;

    let mut x = vec![false; n];

    x[3] = true;
    println!("{:?}", x);
}

//69. Seed random generator

// 随机种子生成器


use rand::{Rng, SeedableRng, rngs::StdRng};

fn main() {
    let s = 32;
    let mut rng = StdRng::seed_from_u64(s);
    
    println!("{:?}", rng.gen::<f32>());
}

//70. Use clock as random generator seed
//使用时钟作为随机生成器的种子


use rand::{Rng, SeedableRng, rngs::StdRng};
use std::time::SystemTime;

fn main() {
    let d = SystemTime::now()
        .duration_since(SystemTime::UNIX_EPOCH)
        .expect("Duration since UNIX_EPOCH failed");
    let mut rng = StdRng::seed_from_u64(d.as_secs());
    
    println!("{:?}", rng.gen::<f32>());
}


// 71. Echo program implementation
// 实现 Echo 程序


use std::env;

fn main() {
    println!("{}", env::args().skip(1).collect::<Vec<_>>().join(" "));
}


// or 


use itertools::Itertools;
println!("{}", std::env::args().skip(1).format(" "));


//74. Compute GCD
//计算大整数a和b的最大公约数x。使用能够处理大数的整数类型。



extern crate num;

use num::Integer;
use num::bigint::BigInt;

fn main() {
    let a = BigInt::parse_bytes(b"6000000000000", 10).unwrap();
    let b = BigInt::parse_bytes(b"9000000000000", 10).unwrap();
    
    let x = a.gcd(&b);
 
    println!("{}", x);
}


// 75. Compute LCM
// 计算大整数a和b的最小公倍数x。使用能够处理大数的整数类型。


extern crate num;

use num::bigint::BigInt;
use num::Integer;

fn main() {
    let a = BigInt::parse_bytes(b"6000000000000", 10).unwrap();
    let b = BigInt::parse_bytes(b"9000000000000", 10).unwrap();
    let x = a.lcm(&b);
    println!("x = {}", x);
}


//76. Binary digits from an integer
//将十进制整数转换为二进制数字


fn main() {
    let x = 13;
    let s = format!("{:b}", x);
    
    println!("{}", s);
}


//77. SComplex number

// 复数




extern crate num;
use num::Complex;

fn main() {
    let mut x = Complex::new(-2, 3);
    x *= Complex::i();
    println!("{}", x);
}



// 78. "do while" loop

// 循环执行



loop {
    doStuff();
    if !c { break; }
}


//79. Convert integer to floating point number

// 整型转浮点型

// 声明浮点数y并用整数x的值初始化它。

fn main() {
    let i = 5;
    let f = i as f64;
    
    println!("int {:?}, float {:?}", i, f);
}

//80. Truncate floating point number to integer
// /浮点型转整型

fn main() {
    let x = 41.59999999f64;
    let y = x as i32;
    println!("{}", y);
}



//81. Round floating point number to integer
// 按规则取整

fn main() {
    let x : f64 = 2.71828;
    let y = x.round() as i64;
    
    println!("{} {}", x, y);
}

// 82. Count substring occurrences
// 统计子字符串出现次数



fn main() {
    let s = "lorem ipsum lorem ipsum lorem ipsum lorem ipsum";
    let t = "ipsum";

    let c = s.matches(t).count();

    println!("{} occurrences", c);
}


// 83. Regex with character repetition
// 正则表达式匹配重复字符

extern crate regex;
use regex::Regex;

fn main() {
    let r = Regex::new(r"htt+p").unwrap();
    
    assert!(r.is_match("http"));
    assert!(r.is_match("htttp"));
    assert!(r.is_match("httttp"));
}

// 84. Count bits set in integer binary representation

// 计算十进制整型的二进制表示中 1的个数



fn main() {
    println!("{}", 6usize.count_ones())
}


// 85. Check if integer addition will overflow
// 检查两个整型相加是否溢出


fn adding_will_overflow(x: usize, y: usize) -> bool {
    x.checked_add(y).is_none()
}

fn main() {
    {
        let (x, y) = (2345678, 9012345);

        let overflow = adding_will_overflow(x, y);

        println!(
            "{} + {} {}",
            x,
            y,
            if overflow {
                "overflows"
            } else {
                "doesn't overflow"
            }
        );
    }
    {
        let (x, y) = (2345678901, 9012345678);

        let overflow = adding_will_overflow(x, y);

        println!(
            "{} + {} {}",
            x,
            y,
            if overflow {
                "overflows"
            } else {
                "doesn't overflow"
            }
        );
    }
    {
        let (x, y) = (2345678901234, 9012345678901);

        let overflow = adding_will_overflow(x, y);

        println!(
            "{} + {} {}",
            x,
            y,
            if overflow {
                "overflows"
            } else {
                "doesn't overflow"
            }
        );
    }
    {
        let (x, y) = (23456789012345678, 90123456789012345);

        let overflow = adding_will_overflow(x, y);

        println!(
            "{} + {} {}",
            x,
            y,
            if overflow {
                "overflows"
            } else {
                "doesn't overflow"
            }
        );
    }
    {
        let (x, y) = (12345678901234567890, 9012345678901234567);

        let overflow = adding_will_overflow(x, y);

        println!(
            "{} + {} {}",
            x,
            y,
            if overflow {
                "overflows"
            } else {
                "doesn't overflow"
            }
        );
    }
}



// 86. Check if integer multiplication will overflow
// 检查整型相乘是否溢出



fn main() {
    {
        let (x, y) = (2345, 6789);

        let overflow = multiply_will_overflow(x, y);

        println!(
            "{} * {} {}",
            x,
            y,
            if overflow {
                "overflows"
            } else {
                "doesn't overflow"
            }
        );
    }
    {
        let (x, y) = (2345678, 9012345);

        let overflow = multiply_will_overflow(x, y);

        println!(
            "{} * {} {}",
            x,
            y,
            if overflow {
                "overflows"
            } else {
                "doesn't overflow"
            }
        );
    }
    {
        let (x, y) = (2345678901, 9012345678);

        let overflow = multiply_will_overflow(x, y);

        println!(
            "{} * {} {}",
            x,
            y,
            if overflow {
                "overflows"
            } else {
                "doesn't overflow"
            }
        );
    }
}

fn multiply_will_overflow(x: i64, y: i64) -> bool {
    x.checked_mul(y).is_none()
}

// 87. Stop program

// 停止程序,立即退出。

fn main() {
    std::process::exit(1);
    println!("42");
}


// 88. Allocate 1M bytes
// 分配1M内存

fn main() {
    let buf: Vec<u8> = Vec::with_capacity(1024 * 1024);
    println!("{:?}", buf.capacity());
}


// 89. Handle invalid argument
// 处理无效参数

#[derive(Debug, PartialEq, Eq)]
enum CustomError { InvalidAnswer }

fn do_stuff(x: i32) -> Result<i32, CustomError> {
    if x != 42 {
%2




// 90. Read-only outside
// 外部只读

struct Foo {
    x: usize
}

impl Foo {
    pub fn new(x: usize) -> Self {
        Foo { x }
    }

    pub fn x<'a>(&'a self) -> &'a usize {
        &self.x
    }
}



// 91. Load JSON file into struct
// json转结构体


#[macro_use] extern crate serde_derive;
extern crate serde_json;
use std::fs::File;
let x = ::serde_json::from_reader(File::open("data.json")?)?;





// 92. Save object into JSON file
// 将json对象写入文件



extern crate serde_json;
#[macro_use] extern crate serde_derive;

use std::fs::File;
::serde_json::to_writer(&File::create("data.json")?, &x)?


// 93. Pass a runnable procedure as parameter
// 以函数作为参数

fn control(f: impl Fn()) {
    f();
}

fn hello() {
    println!("Hello,");
}

fn main() {
    control(hello);
    control(|| { println!("Is there anybody in there?"); });
}


// 94. Print type of variable
// 打印变量的类型

#![feature(core_intrinsics)]

fn type_of<T>(_: &T) -> String {
    format!("{}", std::intrinsics::type_name::<T>())
}

fn main() {
    let x: i32 = 1;
    println!("{}", type_of(&x));
}


// 95. Get file size
// 获取文件的大小



use std::fs;

fn filesize(path: &str) -> Result<u64, std::io::Error> {
    let x = fs::metadata(path)?.len();
    Ok(x)
}

fn main() {
    let path = "/etc/hosts";
    let x = filesize(path);
    println!("{}: {:?} bytes", path, x.unwrap());
}


// or 


use std::path::Path;

fn filesize(path: &std::path::Path) -> Result<u64, std::io::Error> {
    let x = path.metadata()?.len();
    Ok(x)
}

fn main() {
    let path = Path::new("/etc/hosts");
    let x = filesize(path);
    println!("{:?}: {:?} bytes", path, x.unwrap());
}


//96. Check string prefix
// 检查两个字符串前缀是否一致

fn main() {
    let s = "bananas";
    let prefix = "bana";

    let b = s.starts_with(prefix);

    println!("{:?}", b);
}



//97. Check string suffix
// 检查字符串后缀

fn main() {
    let s = "bananas";
    let suffix = "nas";

    let b = s.ends_with(suffix);

    println!("{:?}", b);
}

//98. Epoch seconds to date object

//时间戳转日期

extern crate chrono;
use chrono::prelude::*;

fn main() {
    let ts = 1451606400;
    let d = NaiveDateTime::from_timestamp(ts, 0);
    println!("{}", d);
}

//99. Format date YYYY-MM-DD
// 时间格式转换


extern crate chrono;

use chrono::prelude::*;

fn main() {
    println!("{}", Utc::today().format("%Y-%m-%d"))
}


//100. Sort by a comparator
// 根据某个字段排序
fn main() {
    let mut items = [1, 7, 5, 2, 3];
    items.sort_by(i32::cmp);
    println!("{:?}", items);
}

//101. Load from HTTP GET request into a string
// 发起http请求


extern crate reqwest;
use reqwest::Client;
let client = Client::new();
let s = client.get(u).send().and_then(|res| res.text())?;


// or 

[dependencies]
ureq = "1.0"
let s = ureq::get(u).call().into_string()?;


//or 
[dependencies]
error-chain = "0.12.4"
reqwest = { version = "0.11.2", features = ["blocking"] }

use error_chain::error_chain;
use std::io::Read;
let mut response = reqwest::blocking::get(u)?;
let mut s = String::new();
response.read_to_string(&mut s)?;


// 102. Load from HTTP GET request into a file

//发起http请求


extern crate reqwest;
use reqwest::Client;
use std::fs::File;
let client = Client::new();
match client.get(&u).send() {
    Ok(res) => {
        let file = File::create("result.txt")?;
        ::std::io::copy(res, file)?;
    },
    Err(e) => eprintln!("failed to send request: {}", e),
};

//105. Current executable name

//当前可执行文件名称

fn get_exec_name() -> Option<String> {
    std::env::current_exe()
        .ok()
        .and_then(|pb| pb.file_name().map(|s| s.to_os_string()))
        .and_then(|s| s.into_string().ok())
}

fn main() -> () {
    let s = get_exec_name().unwrap();
    println!("{}", s);
}


// or 

fn main() {
    let s = std::env::current_exe()
        .expect("Can't get the exec path")
        .file_name()
        .expect("Can't get the exec name")
        .to_string_lossy()
        .into_owned();
    
    println!("{}", s);
}


 // 106. Get program working directory

 //获取程序的工作路径

 use std::env;

fn main() {
    let dir = env::current_dir().unwrap();

    println!("{:?}", dir);
}


//107. Get folder containing current program

//获取包含当前程序的文件夹


let dir = std::env::current_exe()?
    .canonicalize()
    .expect("the current exe should exist")
    .parent()
    .expect("the current exe should be a file")
    .to_string_lossy()
    .to_owned();


//109. Number of bytes of a type
// 获取某个类型的字节数


// T has (8 + 4) == 12 bytes of data
struct T(f64, i32);

fn main() {
    let n = ::std::mem::size_of::<T>();

    println!("{} bytes", n);
    // T has size 16, which is "the offset in bytes between successive elements in an array with item type T"
}


// 110. Check if string is blank
// 检查字符串是否空白

fn main() {
    let list = vec!["", " ", "  ", "\t", "\n", "a", " b "];
    for s in list {
        let blank = s.trim().is_empty();

        if blank {
            println!("{:?}\tis blank", s)
        } else {
            println!("{:?}\tis not blank", s)
        }
    }
}


// 111. Launch other program
// 运行其他程序

use std::process::Command;

fn main() {
    let child = Command::new("ls")
        .args(&["/etc"])
        .spawn()
        .expect("failed to execute process");

    let output = child.wait_with_output().expect("Failed to wait on child");
    let output = String::from_utf8(output.stdout).unwrap();

    println!("{}", output);
}


// or 

use std::process::Command;

fn main() {
    let output = Command::new("ls")
        .args(&["/etc"])
        .output()
        .expect("failed to execute process");

    let output = String::from_utf8(output.stdout).unwrap();

    println!("{}", output);
}

// or 

use std::process::Command;

fn main() {
    let status = Command::new("ls")
        .args(&["/etc"])
        .status()
        .expect("failed to execute process");

    // exit code is outputted after _ls_ runs
    println!("{}", status);
}

//112. Iterate over map entries, ordered by keys
// 遍历map，按key排序

use std::collections::BTreeMap;

fn main() {
    let mut mymap = BTreeMap::new();
    mymap.insert("one", 1);
    mymap.insert("two", 2);
    mymap.insert("three", 3);
    mymap.insert("four", 4);
    mymap.insert("five", 5);
    mymap.insert("six", 6);

    // Iterate over map entries, ordered by keys, which is NOT the numerical order
    for (k, x) in mymap {
        println!("({}, {})", k, x);
    }
}

//113. Iterate over map entries, ordered by values
// 遍历map，按值排序



use itertools::Itertools;
use std::collections::HashMap;

fn main() {
    let mut mymap = HashMap::new();
    mymap.insert(1, 3);
    mymap.insert(2, 6);
    mymap.insert(3, 4);
    mymap.insert(4, 1);
    
    for (k, x) in mymap.iter().sorted_by_key(|x| x.1) {
        println!("[{},{}]", k, x);
    }
}

// or 

use std::collections::HashMap;

fn main() {
    let mut mymap = HashMap::new();
    mymap.insert(1, 3);
    mymap.insert(2, 6);
    mymap.insert(3, 4);
    mymap.insert(4, 1);
    
    let mut items: Vec<_> = mymap.iter().collect();
    items.sort_by_key(|item| item.1);
    for (k, x) in items {
        println!("[{},{}]", k, x);
    }
}

//114. Test deep equality
// 深度判等


let b = x == y;

//115. Compare dates
// 日期比较



extern crate chrono;
use chrono::prelude::*;
let b = d1 < d2;

// 116. Remove occurrences of word from string
// 去除指定字符串


fn main() {
    let s1 = "foobar";
    let w = "foo";
    let s2 = s1.replace(w, "");
    println!("{}", s2);
}

// or 

fn main() {
    let s1 = "foobar";
    let w = "foo";

    let s2 = str::replace(s1, w, "");

    println!("{}", s2);
}

// 117. Get list size
// 获取list的大小

fn main() {
    let x = vec![11, 22, 33];

    let n = x.len();

    println!("x has {} elements", n);
}

//118. List to set
// 从list到set
use std::collections::HashSet;

fn main() {
    let x: Vec<i32> = vec![1, 7, 3, 1];
    println!("x: {:?}", x);

    let y: HashSet<_> = x.into_iter().collect();
    println!("y: {:?}", y);
}

// 119. Deduplicate list
// list去重

fn main() {
    let mut x = vec![1, 2, 3, 4, 3, 2, 2, 2, 2, 2, 2];
    x.sort();
    x.dedup();

    println!("{:?}", x);
}

//or 

use itertools::Itertools;

fn main() {
    let x = vec![1, 2, 3, 4, 3, 2, 2, 2, 2, 2, 2];
    let dedup: Vec<_> = x.iter().unique().collect();

    println!("{:?}", dedup);
}



//120. Read integer from stdin

// 从标准输入中读取整数



fn get_input() -> String {
    let mut buffer = String::new();
    std::io::stdin().read_line(&mut buffer).expect("Failed");
    buffer
}

let n = get_input().trim().parse::<i64>().unwrap();


//or 

use std::io;
let mut input = String::new();
io::stdin().read_line(&mut input).unwrap();
let n: i32 = input.trim().parse().unwrap();s

// or 

use std::io::BufRead;
let n: i32 = std::io::stdin()
    .lock()
    .lines()
    .next()
    .expect("stdin should be available")
    .expect("couldn't read from stdin")
    .trim()
    .parse()
    .expect("input was not an integer");



//121. UDP listen and read
// 听端口p上的UDP流量，并将1024字节读入缓冲区b。


use std::net::UdpSocket;
let mut b = [0 as u8; 1024];
let sock = UdpSocket::bind(("localhost", p)).unwrap();
sock.recv_from(&mut b).unwrap();


//122. Declare enumeration
// 声明枚举值



enum Suit {
    Spades,
    Hearts,
    Diamonds,
    Clubs,
}

fn main() {
    let _x = Suit::Diamonds;
}
//123. Assert condition
//断言条件

fn main() {
    // i is odd
    let i = 23687;
    let ii = i * i;
    let is_consistent = ii % 2 == 1;

    // i*i must be odd
    assert!(is_consistent);

    println!("Cool.")
}


//125. Measure function call duration
// 函数调用时间


use std::time::{Duration, Instant};
let start = Instant::now();
foo();
let duration = start.elapsed();
println!("{}", duration);



//126. Multiple return values
// 多个返回值



fn foo() -> (String, bool) {
    (String::from("bar"), true)
}

fn main() {
    println!("{:?}", foo());
}

// //128. Breadth-first traversing of a tree
// 树的广度优先遍历


use std::collections::VecDeque;

struct Tree<V> {
    children: Vec<Tree<V>>,
    value: V
}

impl<V> Tree<V> {
    fn bfs(&self, f: impl Fn(&V)) {
        let mut q = VecDeque::new();
        q.push_back(self);

        while let Some(t) = q.pop_front() {
            (f)(&t.value);
            for child in &t.children {
                q.push_back(child);
            }
        }
    }
}

fn main() {
    let t = Tree {
        children: vec![
            Tree {
                children: vec![
                    Tree { children: vec![], value: 5 },
                    Tree { children: vec![], value: 6 }
                ],
                value: 2
            },
            Tree { children: vec![], value: 3 },
            Tree { children: vec![], value: 4 },
        ],
        value: 1
    };
    t.bfs(|v| println!("{}", v));
}


//129. Breadth-first traversing in a graph
// 图的广度优先遍历

use std::rc::{Rc, Weak};
use std::cell::RefCell;

struct Vertex<V> {
    value: V,
    neighbours: Vec<Weak<RefCell<Vertex<V>>>>,
}

type RcVertex<V> = Rc<RefCell<Vertex<V>>>;

struct Graph<V> {
    vertices: Vec<RcVertex<V>>,
}

impl<V> Graph<V> {
    fn new() -> Self {
        Graph { vertices: vec![] }
    }
    
    fn new_vertex(&mut self, value: V) -> RcVertex<V> {
        self.add_vertex(Vertex { value, neighbours: Vec::new() })
    }
    
    fn add_vertex(&mut self, v: Vertex<V>) -> RcVertex<V> {
        let v = Rc::new(RefCell::new(v));
        self.vertices.push(Rc::clone(&v));
        v
    }
    
    fn add_edge(&mut self, v1: &RcVertex<V>, v2: &RcVertex<V>) {
        v1.borrow_mut().neighbours.push(Rc::downgrade(&v2));
        v2.borrow_mut().neighbours.push(Rc::downgrade(&v1));
    }

    fn bft(start: RcVertex<V>, f: impl Fn(&V)) {
        let mut q = vec![start];
        let mut i = 0;
        while i < q.len() {
            let v = Rc::clone(&q[i]);
            i += 1;
            (f)(&v.borrow().value);
            for n in &v.borrow().neighbours {
                let n = n.upgrade().expect("Invalid neighbour");
                if q.iter().all(|v| v.as_ptr() != n.as_ptr()) {
                    q.push(n);
                }
            }
        }
    }
}

fn main() {
    let mut g = Graph::new();
    
    let v1 = g.new_vertex(1);
    let v2 = g.new_vertex(2);
    let v3 = g.new_vertex(3);
    let v4 = g.new_vertex(4);
    let v5 = g.new_vertex(5);
    
    g.add_edge(&v1, &v2);
    g.add_edge(&v1, &v3);
    g.add_edge(&v1, &v4);
    g.add_edge(&v2, &v5);
    g.add_edge(&v3, &v4);
    g.add_edge(&v4, &v5);
    
    Graph::bft(v1, |v| println!("{}", v));
}


// 130. Depth-first traversing in a graph
// 图的深度优先遍历



use std::rc::{Rc, Weak};
use std::cell::RefCell;

struct Vertex<V> {
    value: V,
    neighbours: Vec<Weak<RefCell<Vertex<V>>>>,
}

type RcVertex<V> = Rc<RefCell<Vertex<V>>>;

struct Graph<V> {
    vertices: Vec<RcVertex<V>>,
}

impl<V> Graph<V> {
    fn new() -> Self {
        Graph { vertices: vec![] }
    }
    
    fn new_vertex(&mut self, value: V) -> RcVertex<V> {
        self.add_vertex(Vertex { value, neighbours: Vec::new() })
    }
    
    fn add_vertex(&mut self, v: Vertex<V>) -> RcVertex<V> {
        let v = Rc::new(RefCell::new(v));
        self.vertices.push(Rc::clone(&v));
        v
    }
    
    fn add_edge(&mut self, v1: &RcVertex<V>, v2: &RcVertex<V>) {
        v1.borrow_mut().neighbours.push(Rc::downgrade(&v2));
        v2.borrow_mut().neighbours.push(Rc::downgrade(&v1));
    }
    
    fn dft(start: RcVertex<V>, f: impl Fn(&V)) {
        let mut s = vec![];
        Self::dft_helper(start, &f, &mut s);
    }
    
    fn dft_helper(start: RcVertex<V>, f: &impl Fn(&V), s: &mut Vec<*const Vertex<V>>) {
        s.push(start.as_ptr());
        (f)(&start.borrow().value);
        for n in &start.borrow().neighbours {
            let n = n.upgrade().expect("Invalid neighbor");
            if s.iter().all(|&p| p != n.as_ptr()) {
                Self::dft_helper(n, f, s);
            }
        }
    }
}

fn main() {
    let mut g = Graph::new();
    
    let v1 = g.new_vertex(1);
    let v2 = g.new_vertex(2);
    let v3 = g.new_vertex(3);
    let v4 = g.new_vertex(4);
    let v5 = g.new_vertex(5);
    
    g.add_edge(&v1, &v2);
    g.add_edge(&v1, &v4);
    g.add_edge(&v1, &v5);
    g.add_edge(&v2, &v3);
    g.add_edge(&v3, &v4);
    g.add_edge(&v4, &v5);
    
    Graph::dft(v1, |v| println!("{}", v));
}

//131. Successive conditions

//连续条件判等

if c1 { f1() } else if c2 { f2() } else if c3 { f3() }


// or 

match true {
    _ if c1 => f1(),
    _ if c2 => f2(),
    _ if c3 => f3(),
    _ => (),
}

//132. Measure duration of procedure execution
// 度量程序执行时间


use std::time::Instant;
let start = Instant::now();
f();
let duration = start.elapsed();

//133. Case-insensitive string contains
// 不区分大小写的字符串包含
extern crate regex;
use regex::Regex;

fn main() {
    let s = "Let's dance the macarena";

    {
        let word = "Dance";
        let re = Regex::new(&format!("(?i){}", regex::escape(word))).unwrap();
        let ok = re.is_match(&s);

        println!("{}", ok);
    }
    
    {
        let word = "dance";
        let re = Regex::new(&format!("(?i){}", regex::escape(word))).unwrap();
        let ok = re.is_match(&s);

        println!("{}", ok);
    }
    
    {
        let word = "Duck";
        let re = Regex::new(&format!("(?i){}", regex::escape(word))).unwrap();
        let ok = re.is_match(&s);

        println!("{}", ok);
    }
}

// or 

use regex::RegexBuilder;

fn main() {
    let s = "FooBar";
    let word = "foo";
    
    let re = RegexBuilder::new(&regex::escape(word))
        .case_insensitive(true)
        .build()
        .unwrap();

    let ok = re.is_match(s);
    
    println!("{:?}", ok);
}

//or 

fn main() {
    let s = "Let's dance the macarena";

    {
        let word = "Dance";
        let ok = s.to_ascii_lowercase().contains(&word.to_ascii_lowercase());
        println!("{}", ok);
    }

    {
        let word = "dance";
        let ok = s.to_ascii_lowercase().contains(&word.to_ascii_lowercase());
        println!("{}", ok);
    }

    {
        let word = "Duck";
        let ok = s.to_ascii_lowercase().contains(&word.to_ascii_lowercase());
        println!("{}", ok);
    }
}

// 134. Create a new list
// 创建一个新list

fn main() {
    let (a, b, c) = (11, 22, 33);
    
    let items = vec![a, b, c];
    
    println!("{:?}", items);
}

//135. Remove item from list, by its value
// 移除列表中的值

if let Some(i) = items.first(&x) {
    items.remove(i);
}


//136. Remove all occurrences of a value from a list
//从列表中删除所有出现的值




fn main() {
    let x = 1;
    let mut items = vec![1, 2, 3, 1, 2, 3];
    
    items = items.into_iter().filter(|&item| item != x).collect();
    
    println!("{:?}", items);
}



//137. Check if string contains only digits
// 检查字符串是否只包含数字


fn main() {
    let s = "1023";
    let chars_are_numeric: Vec<bool> = s.chars().map(|c|c.is_numeric()).collect();
    let b = !chars_are_numeric.contains(&false);
    println!("{}", b);
}


//or 

fn main() {
    let b = "0129".chars().all(char::is_numeric);
    println!("{}", b);
}

//138. Create temp file

// 创建一个新的临时文件

use tempdir::TempDir;
use std::fs::File;
let temp_dir = TempDir::new("prefix")?;
let temp_file = File::open(temp_dir.path().join("file_name"))?;



//139. Create temp directory
// 创建一个临时目录


extern crate tempdir;
use tempdir::TempDir;
let tmp = TempDir::new("prefix")?;

// 140. Delete map entry
// 从map中删除某个key
fn main() {
    use std::collections::HashMap;

    let mut m = HashMap::new();
    m.insert(5, "a");
    m.insert(17, "b");
    println!("{:?}", m);

    m.remove(&5);
    println!("{:?}", m);
}



//141. Iterate in sequence over two lists
//依次迭代两个列表 依次迭代列表项1和项2的元素。每次迭代打印元素。



fn main() {
    let item1 = vec!["1", "2", "3"];
    let item2 = vec!["a", "b", "c"];
    for i in item1.iter().chain(item2.iter()) {
        print!("{} ", i);
    }
}

//142. Hexadecimal digits of an integer
// 将整数x的十六进制表示(16进制)赋给字符串s。

fn main() {
    let x = 999;

    let s = format!("{:X}", x);
    println!("{}", s);

    let s = format!("{:x}", x);
    println!("{}", s);
}


//143. Iterate alternatively over two lists
// 交替迭代两个列表




extern crate itertools;
use itertools::izip;

fn main() {
    let items1 = [5, 15, 25];
    let items2 = [10, 20, 30];

    for pair in izip!(&items1, &items2) {
        println!("{}", pair.0);
        println!("{}", pair.1);
    }
}

//144. Check if file exists
// 检查文件是否存在



fn main() {
    let fp = "/etc/hosts";
    let b = std::path::Path::new(fp).exists();
    println!("{}: {}", fp, b);

    let fp = "/etc/kittens";
    let b = std::path::Path::new(fp).exists();
    println!("{}: {}", fp, b);
}


//145. Print log line with datetime
//打印带时间的日志

fn main() {
    let msg = "Hello";
    eprintln!("[{}] {}", humantime::format_rfc3339_seconds(std::time::SystemTime::now()), msg);
}


//146. Convert string to floating point number
// 字符串转换为浮点型



fn main() {
    let s = "3.14159265359";
    let f = s.parse::<f32>().unwrap();
    
    println!("{}² = {}" , f, f * f);
}


// or 

fn main() {
    let s = "3.14159265359";
    let f: f32 = s.parse().unwrap();

    println!("{}² = {}", f, f * f);
}


//147. Remove all non-ASCII characters
// 移除所有的非ASCII字符



fn main() {
    println!("{}", "do👍ot".replace(|c: char| !c.is_ascii(), ""))
}

// or 

fn main() {
    println!("{}", "do👍ot".replace(|c: char| !c.is_ascii(), ""))
}

//148. Read list of integers from stdin
// 从stdin(标准输入)中读取整数列表



use std::{
    io::{self, Read},
    str::FromStr,
};

// dummy io::stdin
fn io_stdin() -> impl Read {
    "123
456
789"
    .as_bytes()
}

fn main() -> io::Result<()> {
    let mut string = String::new();
    io_stdin().read_to_string(&mut string)?;
    let result = string
        .lines()
        .map(i32::from_str)
        .collect::<Result<Vec<_>, _>>();
    println!("{:#?}", result);
    Ok(())
}


//150. Remove trailing slash
// 去除末尾的 /


fn main() {
    let mut p = String::from("Dddd/");
    if p.ends_with('/') {
        p.pop();
    }
    println!("{}", p);
}


//151. Remove string trailing path separator
// 删除字符串尾部路径分隔符



fn main() {
    {
        let p = "/tmp/";

        let p = if ::std::path::is_separator(p.chars().last().unwrap()) {
            &p[0..p.len() - 1]
        } else {
            p
        };

        println!("{}", p);
    }
    
    {
        let p = "/tmp";

        let p = if ::std::path::is_separator(p.chars().last().unwrap()) {
            &p[0..p.len() - 1]
        } else {
            p
        };

        println!("{}", p);
    }
}


// or 


fn main() {
    {
        let mut p = "/tmp/";

        p = p.strip_suffix(std::path::is_separator).unwrap_or(p);

        println!("{}", p);
    }
    
    {
        let mut p = "/tmp";

        p = p.strip_suffix(std::path::is_separator).unwrap_or(p);

        println!("{}", p);
    }
}


// 152. Turn a character into a string
// 将字符转换成字符串



fn main() {
    let c = 'a';
    
    let s = c.to_string();
    
    println!("{}", s);
}

//153. Concatenate string with integer
// 连接字符串和整数


fn main() {
    let s = "Foo";
    let i = 1;
    let t = format!("{}{}", s, i);
    
    println!("{}" , t);
}



//154. Halfway between two hex color codes
// 求两个十六进制颜色代码的中间值



use std::str::FromStr;
use std::fmt;

#[derive(Debug)]
struct Colour {
    r: u8,
    g: u8,
    b: u8
}

#[derive(Debug)]
enum ColourError {
    MissingHash,
    InvalidRed,
    InvalidGreen,
    InvalidBlue
}

impl fmt::Display for Colour {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "#{:02x}{:02x}{:02x}", self.r, self.g, self.b)
    }
}

impl FromStr for Colour {
    type Err = ColourError;
    
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        if !s.starts_with('#') {
            Err(ColourError::MissingHash)
        } else {
            Ok(Colour {
                r: u8::from_str_radix(&s[1..3], 16).map_err(|_| ColourError::InvalidRed)?,
                g: u8::from_str_radix(&s[3..5], 16).map_err(|_| ColourError::InvalidGreen)?,
                b: u8::from_str_radix(&s[5..7], 16).map_err(|_| ColourError::InvalidBlue)?
            })
        }
    }
}

fn mid_colour(c1: &str, c2: &str) -> Result<String, ColourError> {
    let c1 = c1.parse::<Colour>()?;
    let c2 = c2.parse::<Colour>()?;
    let c = Colour {
        r: (((c1.r as u16) + (c2.r as u16))/2) as u8,
        g: (((c1.g as u16) + (c2.g as u16))/2) as u8,
        b: (((c1.b as u16) + (c2.b as u16))/2) as u8
    };
    Ok(format!("{}", c))
}

fn main() {
    println!("{}", mid_colour("#15293E", "#012549").unwrap())
}


//155. Delete file
// 删除文件

use std::fs;

fn main() {
    let filepath = "/tmp/abc";

    println!("Creating {}", filepath);
    let _file = fs::File::create(filepath);

    let b = std::path::Path::new(filepath).exists();
    println!("{} exists: {}", filepath, b);

    println!("Deleting {}", filepath);
    let r = fs::remove_file(filepath);
    println!("{:?}", r);

    let b = std::path::Path::new(filepath).exists();
    println!("{} exists: {}", filepath, b);
}

//156. Format integer with zero-padding
// 用零填充格式化整数

fn main() {
    let i = 1;
    let s = format!("{:03}", i);
    
    println!("{}", s);
    
    
    let i = 1000;
    let s = format!("{:03}", i);
    
    println!("{}", s);
}

//157. Declare constant string

// 声明常量字符串



fn main() {
    const PLANET: &str = "Earth";
    
    println!("{}", PLANET);
}


//158. Random sublist
// 随机子列表

use rand::prelude::*;
let mut rng = &mut rand::thread_rng();
let y = x.choose_multiple(&mut rng, k).cloned().collect::<Vec<_>>();


// 159. Trie
//前缀树/字典树



struct Trie {
    val: String,
    nodes: Vec<Trie>
}
//160. Detect if 32-bit or 64-bit architecture
// 检测是32位还是64位架构

fn main() {
    match std::mem::size_of::<&char>() {
        4 => f32(),
        8 => f64(),
        _ => {}
    }
}

fn f32() {
    println!("I am 32-bit");
}

fn f64() {
    println!("I am 64-bit");
}



//161. Multiply all the elements of a list
// 将list中的每个元素都乘以一个数



fn main() {
    let elements: Vec<f32> = vec![2.0, 3.5, 4.0];
    let c = 2.0;

    let elements = elements.into_iter().map(|x| c * x).collect::<Vec<_>>();

    println!("{:?}", elements);
}


//162. Execute procedures depending on options
// 根据选项执行程序



if let Some(arg) = ::std::env::args().nth(1) {
    if &arg == "f" {
        fox();
    } else if &arg = "b" {
        bat();
    } else {
 eprintln!("invalid argument: {}", arg),
    }
} else {
    eprintln!("missing argument");
}

// or 

if let Some(arg) = ::std::env::args().nth(1) {
    match arg.as_str() {
        "f" => fox(),
        "b" => box(),
        _ => eprintln!("invalid argument: {}", arg),
    };
} else {
    eprintln!("missing argument");
}
//163. Print list elements by group of 2
// 两个一组打印数组元素


fn main() {
    let list = [1,2,3,4,5,6];
    for pair in list.chunks(2) {
        println!("({}, {})", pair[0], pair[1]);
    }
}
//164. Open URL in default browser
// 在默认浏览器中打开链接


use webbrowser;
webbrowser::open(s).expect("failed to open URL");


///165. Last element of list
// 列表中的最后一个元素



fn main() {
    let items = vec![5, 6, 8, -20, 9, 42];
    let x = items[items.len()-1];
    println!("{:?}", x);
}


//or 

fn main() {
    let items = [5, 6, 8, -20, 9, 42];
    let x = items.last().unwrap();
    println!("{:?}", x);
}

// 166. Concatenate two lists
// 连接两个列表


fn main() {
    let a = vec![1, 2];
    let b = vec![3, 4];
    let ab = [a, b].concat();
    println!("{:?}", ab);
}

//167. Trim prefix

// 移除前缀


fn main() {
    {
        let s = "pre_thing";
        let p = "pre_";
        let t = s.trim_start_matches(p);
        println!("{}", t);
    }
    {
        // Warning: trim_start_matches removes several leading occurrences of p, if present.
        let s = "pre_pre_thing";
        let p = "pre_";
        let t = s.trim_start_matches(p);
        println!("{}", t);
    }
}


// or 


fn main() {
    let s = "pre_pre_thing";
    let p = "pre_";

    let t = if s.starts_with(p) { &s[p.len()..] } else { s };
    println!("{}", t);
}

// or 

fn main() {
    {
        let s = "pre_thing";
        let p = "pre_";
        let t = s.strip_prefix(p).unwrap_or_else(|| s);
        println!("{}", t);
    }
    {
        // If prefix p is repeated in s, it is removed only once by strip_prefix
        let s = "pre_pre_thing";
        let p = "pre_";
        let t = s.strip_prefix(p).unwrap_or_else(|| s);
        println!("{}", t);
    }
}

//168. Trim suffix
// 移除后缀

fn main() {
    let s = "thing_suf";
    let w = "_suf";
    let t = s.trim_end_matches(w);
    println!("{}", t);

    let s = "thing";
    let w = "_suf";
    let t = s.trim_end_matches(w); // s does not end with w, it is left intact
    println!("{}", t);

    let s = "thing_suf_suf";
    let w = "_suf";
    let t = s.trim_end_matches(w); // removes several occurrences of w
    println!("{}", t);
}

//or 


fn main() {
    let s = "thing_suf";
    let w = "_suf";
    let t = s.strip_suffix(w).unwrap_or(s);
    println!("{}", t);

    let s = "thing";
    let w = "_suf";
    let t = s.strip_suffix(w).unwrap_or(s); // s does not end with w, it is left intact
    println!("{}", t);

    let s = "thing_suf_suf";
    let w = "_suf";
    let t = s.strip_suffix(w).unwrap_or(s); // only 1 occurrence of w is removed
    println!("{}", t);
}

// 169. String length
// 字符串长度



fn main() {
    let s = "世界";

    let n = s.chars().count();

    println!("{} characters", n);
}


//170. Get map size
// 获取map的大小

use std::collections::HashMap;

fn main() {
    let mut mymap: HashMap<&str, i32> = [("one", 1), ("two", 2)].iter().cloned().collect();
    mymap.insert("three", 3);

    let n = mymap.len();

    println!("mymap has {:?} entries", n);
}


//171. Add an element at the end of a list
// 在list尾部添加元素


fn main() {
    let mut s = vec![1, 2, 3];
    let x = 99;

    s.push(x);

    println!("{:?}", s);
}
//172. Insert entry in map
// 向map中写入元素



use std::collections::HashMap;

fn main() {
    let mut m: HashMap<&str, i32> = [("one", 1), ("two", 2)].iter().cloned().collect();

    let (k, v) = ("three", 3);

    m.insert(k, v);

    println!("{:?}", m);
}

//173. Format a number with grouped thousands
// 按千位格式化数字


use separator::Separatable;
println!("{}", 1000.separated_string());


//174. Make HTTP POST request
// 发起http POST请求



[dependencies]
error-chain = "0.12.4"
reqwest = { version = "0.11.2", features = ["blocking"] }

use error_chain::error_chain;
use std::io::Read;
let client = reqwest::blocking::Client::new();
let mut response = client.post(u).body("abc").send()?;



//175. Bytes to hex string
// 字节转十六进制字符串


use core::fmt::Write;

fn main() -> core::fmt::Result {
    let a = vec![22, 4, 127, 193];
    let n = a.len();
    
    let mut s = String::with_capacity(2 * n);
    for byte in a {
        write!(s, "{:02X}", byte)?;
    }
    
    dbg!(s);
    Ok(())
}

// 176. Hex string to byte array
//  十六进制字符串转字节数组


use hex::FromHex
let a: Vec<u8> = Vec::from_hex(s).expect("Invalid Hex String");

//178. Check if point is inside rectangle
// 检查点是否在矩形内

struct Rect {
    x1: i32,
    x2: i32,
    y1: i32,
    y2: i32,
}

impl Rect {
    fn contains(&self, x: i32, y: i32) -> bool {
        return self.x1 < x && x < self.x2 && self.y1 < y && y < self.y2;
    }
}
//179. Get center of a rectangle
// 获取矩形的中心


struct Rectangle {
    x1: f64,
    y1: f64,
    x2: f64,
    y2: f64,
}

impl Rectangle {
    pub fn center(&self) -> (f64, f64) {
     ((self.x1 + self.x2) / 2.0, (self.y1 + self.y2) / 2.0)
    }
}

fn main() {
    let r = Rectangle {
        x1: 5.,
        y1: 5.,
        x2: 10.,
        y2: 10.,
    };
    
    println!("{:?}", r.center());
}

// 180. List files in directory
// 列出目录中的文件


use std::fs;

fn main() {
    let d = "/etc";

    let x = fs::read_dir(d).unwrap();

    for entry in x {
        let entry = entry.unwrap();
        println!("{:?}", entry.path());
    }
}



// or 


fn main() {
    let d = "/etc";

    let x = std::fs::read_dir(d)
        .unwrap()
        .collect::<Result<Vec<_>, _>>()
        .unwrap();

    for entry in x {
        println!("{:?}", entry.path());
    }
}


//182. Quine program
// 输出程序的源代码



fn main() {
    let x = "fn main() {\n    let x = ";
    let y = "print!(\"{}{:?};\n    let y = {:?};\n    {}\", x, x, y, y)\n}\n";
    print!("{}{:?};
    let y = {:?};
    {}", x, x, y, y)
}

// or 


fn main(){print!("{},{0:?})}}","fn main(){print!(\"{},{0:?})}}\"")}


//184. Tomorrow

//明天的日期

fn main() {
    let t = chrono::Utc::now().date().succ().to_string();
    println!("{}", t);
}

//185. Execute function in 30 seconds
// 30秒内执行功能
use std::time::Duration;
use std::thread::sleep;
sleep(Duration::new(30, 0));
f(42);


//186. Exit program cleanly
// 干净地退出程序



use std::process::exit;

fn main() {
    println!("A");
    exit(0);
    println!("B");
}

//189. Filter and transform list

// 过滤和转换列表




let y = x.iter()
 .filter(P)
        .map(T)
 .collect::<Vec<_>>();

 
//190. Call an external C function

// 调用外部C函数


extern "C" {
    /// # Safety
    ///
    /// `a` must point to an array of at least size 10
    fn foo(a: *mut libc::c_double, n: libc::c_int);
}

let mut a = [0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0];
let n = 10;
unsafe {
    foo(a.as_mut_ptr(), n);
}

//191. Check if any value in a list is larger than a limit
//检查列表中是否有任何值大于限制




fn main() {
    let a = [5, 6, 8, -20, 9, 42];

    let x = 35;
    if a.iter().any(|&elem| elem > x) {
        f()
    }

    let x = 50;
    if a.iter().any(|&elem| elem > x) {
        g()
    }
}

fn f() {
    println!("F")
}

fn g() {
    println!("G")
}


//192. Declare a real variable with at least 20 digits

// 声明一个至少有20位数字的实变量


use rust_decimal::Decimal;
use std::str::FromStr;
let a = Decimal::from_str("1234567890.123456789012345").unwrap();



//197. Get a list of lines from a file
// 从文件中获取行列表.将文件路径中的内容检索到字符串行列表中，其中每个元素都是文件的一行。

use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

fn main() {
    let path = "/etc/hosts";

    let lines = BufReader::new(File::open(path).unwrap())
        .lines()
        .collect::<Vec<_>>();

    println!("{:?}", lines);
}


//198. Abort program execution with error condition
// 出现错误情况时中止程序执行



use std::process;
process::exit(x);

//200. Return hypotenuse
// 返回三角形的斜边h，其中与直角相邻的边的长度为x和y。


fn main() {
    let (x, y) = (1.0, 1.0);
    let h = hypot(x, y);
    println!("{}", h);
}

fn hypot(x: f64, y: f64) -> f64 {
    let num = x.powi(2) + y.powi(2);
    num.powf(0.5)
}




//202. Sum of squares

// 计算平方和



fn main() {
    let data: Vec<f32> = vec![2.0, 3.5, 4.0];

    let s = data.iter().map(|x| x.powi(2)).sum::<f32>();

    println!("{}", s);
}

//205. Get an environment variable
// 获取环境变量



use std::env;

fn main() {
    let foo;
    match env::var("FOO") {
        Ok(val) => foo = val,
        Err(_e) => foo = "none".to_string(),
    }
    println!("{}", foo);
    
    let user;
    match env::var("USER") {
        Ok(val) => user = val,
        Err(_e) => user = "none".to_string(),
    }
    println!("{}", user);
}


//or 

use std::env;

fn main() {
    let foo = env::var("FOO").unwrap_or("none".to_string());
    println!("{}", foo);

    let user = env::var("USER").unwrap_or("none".to_string());
    println!("{}", user);
}

// or 

use std::env;

fn main() {
    let foo = match env::var("FOO") {
        Ok(val) => val,
        Err(_e) => "none".to_string(),
    };
    println!("{}", foo);

    let user = match env::var("USER") {
        Ok(val) => val,
        Err(_e) => "none".to_string(),
    };
    println!("{}", user);
}



// or 

use std::env;
if let Ok(tnt_root) = env::var("TNT_ROOT") {
     //
}

//206. Switch statement with strings

// switch语句


fn main() {
    fn foo() {}
    fn bar() {}
    fn baz() {}
    fn barfl() {}
    let str_ = "x";
    match str_ {
        "foo" => foo(),
        "bar" => bar(),
        "baz" => baz(),
        "barfl" => barfl(),
        _ => {}
    }
}



//207. Allocate a list that is automatically deallocated

// 分配一个自动解除分配的列表

let a = vec![0; n];

//208. Formula with arrays
// 对数组元素进行运算




fn main() {
    let mut a: [f32; 5] = [5., 2., 8., 9., 0.5]; // we want it to be mutable
    let b: [f32; 5] = [7., 9., 8., 0.965, 0.98]; 
    let c: [f32; 5] = [0., 0.8, 789456., 123456., 0.0003]; 
    let d: [f32; 5] = [332., 0.1, 8., 9874., 0.3]; 

    const e: f32 = 85.;

    for i in 0..a.len() {
        a[i] = e * (a[i] + b[i] * c[i] + d[i].cos());
    }

    println!("{:?}", a); //Don't have any idea about the output
}


//209. Type with automatic deep deallocation
// 自动深度解除分配的类型



struct T {
 s: String,
 n: Vec<usize>,
}

fn main() {
 let v = T {
  s: "Hello, world!".into(),
  n: vec![1,4,9,16,25]
 };
}

//211. Create folder
// 创建文件夹



use std::fs;
use std::path::Path;

fn main() {
    let path = "/tmp/goofy";
    let b: bool = Path::new(path).is_dir();
    println!("{} exists: {}", path, b);

    let r = fs::create_dir(path);
    match r {
        Err(e) => {
            println!("error creating {}: {}", path, e);
            std::process::exit(1);
        }
        Ok(_) => println!("created {}: OK", path),
    }

    let b: bool = Path::new(path).is_dir();
    println!("{} exists: {}", path, b);
}


//or 

use std::fs;
use std::path::Path;

fn main() {
    let path = "/tmp/friends/goofy";
    let b: bool = Path::new(path).is_dir();
    println!("{} exists: {}", path, b);

    // fs::create_dir can't create parent folders
    let r = fs::create_dir(path);
    match r {
        Err(e) => println!("error creating {}: {}", path, e),
        Ok(_) => println!("created {}: OK", path),
    }

    let b: bool = Path::new(path).is_dir();
    println!("{} exists: {}", path, b);

    // fs::create_dir_all does create parent folders
    let r = fs::create_dir_all(path);
    match r {
        Err(e) => println!("error creating {}: {}", path, e),
        Ok(_) => println!("created {}: OK", path),
    }

    let b: bool = Path::new(path).is_dir();
    println!("{} exists: {}", path, b);
}


//212. Check if folder exists
// 检查文件夹是否存在


use std::path::Path;

fn main() {
    let path = "/etc";
    let b: bool = Path::new(path).is_dir();
    println!("{}: {}", path, b);

    let path = "/goofy";
    let b: bool = Path::new(path).is_dir();
    println!("{}: {}", path, b);
}


//215. Pad string on the left
// 左侧补齐字符串
use unicode_width::{UnicodeWidthChar, UnicodeWidthStr};
if let Some(columns_short) = m.checked_sub(s.width()) {
    let padding_width = c
        .width()
        .filter(|n| *n > 0)
        .expect("padding character should be visible");
    // Saturate the columns_short
    let padding_needed = columns_short + padding_width - 1 / padding_width;
    let mut t = String::with_capacity(s.len() + padding_needed);
    t.extend((0..padding_needed).map(|_| c)
    t.push_str(&s);
    s = t;
}


//217. Create a Zip archive
// 创建压缩文件


use zip::write::FileOptions;
let path = std::path::Path::new(_name);
let file = std::fs::File::create(&path).unwrap();
let mut zip = zip::ZipWriter::new(file); zip.start_file("readme.txt", FileOptions::default())?;                                                          
zip.write_all(b"Hello, World!\n")?;
zip.finish()?;

// or 

use zip::write::FileOptions;
fn zip(_name: &str, _list: Vec<&str>) -> zip::result::ZipResult<()>
{
    let path = std::path::Path::new(_name);
    let file = std::fs::File::create(&path).unwrap();
    let mut zip = zip::ZipWriter::new(file);
    for i in _list.iter() {
        zip.start_file(i as &str, FileOptions::default())?;
    }
    zip.finish()?;
    Ok(())
}

//218. List intersection
// 列表的交集

use std::collections::HashSet;

fn main() {
    let a = vec![1, 2, 3, 4];
    let b = vec![2, 4, 6, 8, 2, 2];

    let unique_a = a.iter().collect::<HashSet<_>>();
    let unique_b = b.iter().collect::<HashSet<_>>();
    let c = unique_a.intersection(&unique_b).collect::<Vec<_>>();

    println!("c: {:?}", c);
}

//or 

use std::collections::HashSet;

fn main() {
    let a = vec![1, 2, 3, 4];
    let b = vec![2, 4, 6, 8, 2, 2];

    let set_a: HashSet<_> = a.into_iter().collect();
    let set_b: HashSet<_> = b.into_iter().collect();
    let c = set_a.intersection(&set_b);

    println!("c: {:?}", c);
}


//219. Replace multiple spaces with single space
// 用单个空格替换多个空格


use regex::Regex;

fn main() {
    let s = "
 one   two
    three
 ";
    let re = Regex::new(r"\s+").unwrap();
    let t = re.replace_all(s, " ");
    println!("{}", t);
}

//220. Create a tuple value

// 创建元组值a

fn main() {
    let mut t = (2.5, "hello", -1);
    
    t.2 -= 4;
    println!("{:?}", t);
}



// 221. Remove all non-digits characters
// 删除所有非数字字符


fn main() {
    let t: String = "Today is the 14th of July"
        .chars()
        .filter(|c| c.is_digit(10))
        .collect();

    dbg!(t);
}


//222. Find first index of an element in list
// 在列表中查找元素的第一个索引



fn main() {
    let items = ['A', '🎂', '㍗'];
    let x = '💩';

    match items.iter().position(|y| *y == x) {
        Some(i) => println!("Found {} at position {}.", x, i),
        None => println!("There is no {} in the list.", x),
    }
}


// or 


fn main() {
    let items = [42, -3, 12];

    {
        let x = 12;

        let i = items.iter().position(|y| *y == x).map_or(-1, |n| n as i32);

        println!("{} => {}", x, i)
    }

    {
        let x = 13;

        let i = items.iter().position(|y| *y == x).map_or(-1, |n| n as i32);

        println!("{} => {}", x, i)
    }
}


//223. for else loop
// for else循环


fn main() {
    let items: &[&str] = &["foo", "bar", "baz", "qux"];

    items
        .iter()
        .find(|&&item| item == "rockstar programmer")
        .or_else(|| {
            println!("NotFound");
            Some(&"rockstar programmer")
        });
}

//224. Add element to the beginning of the list

// 将元素添加到列表的开头




use std::collections::VecDeque;

fn main() {
    let mut items = VecDeque::new();
    items.push_back(22);
    items.push_back(33);
    let x = 11;

    items.push_front(x);

    println!("{:?}", items);
}


//225. Declare and use an optional argument

// 声明并使用可选参数


fn f(x: Option<()>) {
    match x {
        Some(x) => println!("Present {}", x),
        None => println!("Not present"),
    }
}

//226. Delete last element from list
// 从列表中删除最后一个元素



fn main() {
    let mut items = vec![11, 22, 33];

    items.pop();

    println!("{:?}", items);
}



//227. Copy list

// 复制列表



fn main() {
    let mut x = vec![4, 3, 2];

    let y = x.clone();

    x[0] = 99;
    println!("x is {:?}", x);
    println!("y is {:?}", y);
}


//228. Copy a file

// 复制文件



use std::fs;

fn main() {
    let src = "/etc/fstabZ";
    let dst = "fstab.bck";

    let r = fs::copy(src, dst);

    match r {
        Ok(v) => println!("Copied {:?} bytes", v),
        Err(e) => println!("error copying {:?} to {:?}: {:?}", src, dst, e),
    }
}



//231. Test if bytes are a valid UTF-8 string
// 测试字节是否是有效的UTF-8字符串



fn main() {
    {
        let bytes = [0xc3, 0x81, 0x72, 0x76, 0xc3, 0xad, 0x7a];

        let b = std::str::from_utf8(&bytes).is_ok();
        println!("{}", b);
    }

    {
        let bytes = [0xc3, 0x81, 0x81, 0x76, 0xc3, 0xad, 0x7a];

        let b = std::str::from_utf8(&bytes).is_ok();
        println!("{}", b);
    }
}
///234. Encode bytes to base64

//将字节编码为base64

//use base64;

fn main() {
    let d = "Hello, World!";

    let b64txt = base64::encode(d);
    println!("{}", b64txt);
}

//235. Decode base64
// 解码base64


//use base64;

fn main() {
    let d = "SGVsbG8sIFdvcmxkIQ==";

    let bytes = base64::decode(d).unwrap();
    println!("Hex: {:x?}", bytes);
    println!("Txt: {}", std::str::from_utf8(&bytes).unwrap());
}


//237. Xor integers
// 异或运算

fn main() {
    let a = 230;
    let b = 42;
    let c = a ^ b;

    println!("{}", c);
}

//238. Xor byte arrays

//异或字节数组

fn main() {
    let a: &[u8] = "Hello".as_bytes();
    let b: &[u8] = "world".as_bytes();

    let c: Vec<_> = a.iter().zip(b).map(|(x, y)| x ^ y).collect();

    println!("{:?}", c);
}

// 239. Find first regular expression match

// 查找第一个正则表达式匹配项



use regex::Regex;

fn main() {
    let sentences = vec![
        "",
        "12",
        "123",
        "1234",
        "I have 12 goats, 3988 otters, 224 shrimps and 456 giraffes",
        "See p.456, for word boundaries",
    ];
    for s in sentences {
        let re = Regex::new(r"\b\d\d\d\b").expect("failed to compile regex");
        let x = re.find(s).map(|x| x.as_str()).unwrap_or("");
        println!("[{}] -> [{}]", &s, &x);
    }
}


//240. Sort 2 lists together
// 将两个列表排序在一起.列表a和b的长度相同。对a和b应用相同的排列，根据a的值对它们进行排序。

fn main() {
    let a = vec![30, 20, 40, 10];
    let b = vec![101, 102, 103, 104];

    let mut tmp: Vec<_> = a.iter().zip(b).collect();
    tmp.as_mut_slice().sort_by_key(|(&x, _y)| x);
    let (aa, bb): (Vec<i32>, Vec<i32>) = tmp.into_iter().unzip();

    println!("{:?}, {:?}", aa, bb);
}


//241. Yield priority to other threads
// 将优先权让给其他线程


::std::thread::yield_now();
busywork();


//242. Iterate over a set
// 迭代一个集合


use std::collections::HashSet;

fn main() {
    let mut x = HashSet::new();
    x.insert("a");
    x.insert("b");

    for item in &x {
        f(item);
    }
}

fn f(s: &&str) {
    println!("Element {}", s);
}

// x is a HashSet

// Element a
// Element b

//243. Print list
// 打印 list


fn main() {
    let a = [11, 22, 33];

    println!("{:?}", a);
}

//244. Print map
// 打印 map

use std::collections::HashMap;

fn main() {
    let mut m = HashMap::new();
    m.insert("Áron".to_string(), 23);
    m.insert("Béla".to_string(), 35);
    println!("{:?}", m);
}

//245. Print value of custom type

// 打印自定义类型的值


#[derive(Debug)]

// T represents a tank
struct T<'a> {
    name: &'a str,
    weight: &'a i32,
    fire_power: &'a i32,
}

fn main() {
    let x = T {
        name: "mytank",
        weight: &100,
        fire_power: &90,
    };

    println!("{:?}", x);
}

//246. Count distinct elements
// 计算不同的元素的数量

use itertools::Itertools;

fn main() {
    let items = vec!["víz", "árvíz", "víz", "víz", "ár", "árvíz"];
    let c = items.iter().unique().count();
    println!("{}", c);
}

//247. Filter list in-place
// 就地筛选列表


fn p(t: i32) -> bool {
    t % 2 == 0
}

fn main() {
    let mut x = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    let mut j = 0;
    for i in 0..x.len() {
        if p(x[i]) {
            x[j] = x[i];
            j += 1;
        }
    }
    x.truncate(j);
    println!("{:?}", x);
}



// or 

fn p(t: &i64) -> bool {
    t % 2 == 0
}

fn main() {
    let mut x: Vec<i64> = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

    x.retain(p);

    println!("{:?}", x);
}



// 249. Declare and assign multiple variables
// 声明并分配多个变量

fn main() {
    // a, b and c don't need to have the same type.

    let (a, b, c) = (42, "hello", 5.0);

    println!("{} {} {}", a, b, c);
}
//251. Parse binary digits
// 解析二进制数字

fn main() {
    let s = "1101"; // binary digits
    
    let i = i32::from_str_radix(s, 2).expect("Not a binary number!");
    
    println!("{}", i);
}
//252. Conditional assignment
 // 条件赋值


 x = if condition() { "a" } else { "b" };

 //258. Convert list of strings to list of integers
// 将字符串列表转换为整数列表

fn main() {
    let a: Vec<&str> = vec!["11", "-22", "33"];

    let b: Vec<i64> = a.iter().map(|x| x.parse::<i64>().unwrap()).collect();

    println!("{:?}", b);
}


//259. Split on several separators
// 在几个分隔符上拆分

fn main() {
    let s = "2021-03-11,linux_amd64";

    let parts: Vec<_> = s.split(&[',', '-', '_'][..]).collect();

    println!("{:?}", parts);
}



//266. Repeating string
//重复字符串


fn main() {
    let v = "abc";
    let n = 5;

    let s = v.repeat(n);
    println!("{}", s);
}
