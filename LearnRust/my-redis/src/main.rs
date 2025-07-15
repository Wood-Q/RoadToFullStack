use mini_redis::{Connection, Frame, blocking_client::connect};
use tokio::net::{TcpListener, TcpStream};

#[tokio::main]
async fn main() {
    let listener = TcpListener::bind("127.0.0.1:6379").await.unwrap();
    loop {
        let (socket, _) = listener.accept().await.unwrap();
        tokio::spawn(async move {
            process(socket).await;
        });
    }
}

async fn process(socket: TcpStream) {
    use mini_redis::Command::{self, Get, Set};
    use std::collections::HashMap;
    //使用hashmap来存数据
    let mut db = HashMap::new();

    //使用返回的connection获取socket数据
    let mut connection = Connection::new(socket);
    //Some是检测右边是否有返回的值
    while let Some(frame) = connection.read_frame().await.unwrap() {
        //模式匹配
        let response = match Command::from_frame(frame).unwrap() {
            //如果是set，就插入数据
            Set(cmd) => {
                db.insert(cmd.key().to_string(), cmd.value().to_vec());
                Frame::Simple("OK".to_string())
            }
            //如果是get，就检查数据
            Get(cmd) => {
                //如果有值，就返回，否则就返回null
                if let Some(value) = db.get(cmd.key()) {
                    // `Frame::Bulk` 期待数据的类型是 `Bytes`， 该类型会在后面章节讲解，
                    // 此时，你只要知道 `&Vec<u8>` 可以使用 `into()` 方法转换成 `Bytes` 类型
                    Frame::Bulk(value.clone().into())
                } else {
                    Frame::Null
                }
            }
            cmd => panic!("unimplemented {:?}", cmd),
        };
        connection.write_frame(&response).await.unwrap();
    }
}
