struct Point<T, U> {
    x: T,
    y: U,
}
fn main() {
    let p1 = Point { x: 10, y: 0.1 };
    println!("x:{}", p1.x)
}
