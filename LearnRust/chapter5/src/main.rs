// struct User {
//     name: String,
//     email: String,
//     age: u64,
// }
// fn main() {
//     let user1 = User {
//         name: String::from("张三"),
//         email: String::from("woodq@gmail.com"),
//         age: 18,
//     };
//     let user2=User{
//         name:String::from("李四"),
//         ..user1
//     };
//     //这种情况不可行，因为user1的email和age转移给user2了，所有权归user2
//     // println!("user1:{}",user1.email);
//     println!("user2:{}",user2.email);
// }

//-------------------------------------------------
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
    fn ifWidth(&self)->bool{
        self.width>0
    }
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };

    println!(
        "The area of the rectangle is {} square pixels.",
        rect1.area()
    );

    println!("This rectangle width is {} > 0",rect1.ifWidth());
}