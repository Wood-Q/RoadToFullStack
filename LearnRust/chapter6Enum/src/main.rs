// enum Coin {
//     Penny,
//     Nickel,
//     Dime,
//     Quarter,
// }

// fn value_in_cents(coin: Coin) -> u8 {
//     match coin {
//         Coin::Penny => {
//             println!("Lucky penny!");
//             1
//         }
//         Coin::Nickel => 5,
//         Coin::Dime => 10,
//         Coin::Quarter => 25,
//     }
// }

fn if_a_larger_than_b(a: i32, b: i32) -> Result<String, String> {
    if a > b {
        Ok("a is bigger".to_string())
    } else {
        Err("b is bigger".to_string())
    }
}
fn main() {
    // let coin = Coin::Penny;
    // println!("{}", value_in_cents(coin));
    let a = 3;
    let b = 2;
    let result = if_a_larger_than_b(a, b);
    // if let Ok(res) = result {
    //     println!("{}", res);
    // } else {
    //     println!("{}", result.unwrap_err());
    // }
    match result {
        Ok(res) => println!("{}", res),
        Err(err) => println!("{}", err),
    }
}
