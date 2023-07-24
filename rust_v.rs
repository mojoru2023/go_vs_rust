
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