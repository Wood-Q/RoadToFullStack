fn main() {
    // let s1 = String::from("hello");
    //s1所存储的内容被转移到s2中，原s1不再有效
    //s1无法再被使用，会被编译器阻止，直到生命周期结束被回收
    // let s2 = s1;
    /*s1报错
        let s1 = String::from("hello");
      |         -- move occurs because `s1` has type `String`, which does not implement the `Copy` trait
    3 |     let s2 = s1;
      |              -- value moved here
    4 |
    5 |     println!("s1 = {s1}");
      |                    ^^^^ value borrowed here after move */
    // println!("s1 = {s1}");
    // println!("s2 = {s2}");
    //--------------------------------------------------
    let s1 = String::from("hello");
    let len = calculate_length(&s1);
    let mut s2 = String::from("hello");
    change(&mut s2);
    println!("The length of '{}' is {}.", s1, len);
}
//这里的s是一个引用，不是所有权的转移
fn calculate_length(s: &String) -> usize {
    //这一步借用修改了引用值，没有所有权无法修改，所以报错
    s.push_str("world");
    //这一步借用没有修改，所以可以运行
    s.len()
}
//这里因为加上了mut，所以是可变引用，就可以进行修改
fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
