use std::thread;
use std::time::Duration;
fn main() {
    // let handle = thread::spawn(|| {
    //     for i in 1..5 {
    //         println!("我是子线程数字{}", i);
    //         thread::sleep(Duration::from_millis(1));
    //     }
    // });
    // handle.join().unwrap();
    // for i in 1..5 {
    //     println!("我是主线程数字{}", i);
    //     thread::sleep(Duration::from_millis(1));
    // }
    let v = vec![1, 2, 3];
    let handle = thread::spawn(move || {
        println!("这个数字是{:?}", v);
    });
    handle.join().unwrap();
}
