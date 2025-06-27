fn main() {
    let s1 = String::from("hello");
    //s1所存储的内容被转移到s2中，原s1不再有效
    let s2 = s1;
    /*s1报错
        let s1 = String::from("hello");
      |         -- move occurs because `s1` has type `String`, which does not implement the `Copy` trait
    3 |     let s2 = s1;
      |              -- value moved here
    4 |
    5 |     println!("s1 = {s1}");
      |                    ^^^^ value borrowed here after move */
    println!("s1 = {s1}");
    println!("s2 = {s2}");
}
