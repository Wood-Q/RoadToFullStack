use bytes::Bytes;
use mini_redis::{Connection, Frame};
use std::{
    collections::HashMap,
    sync::{Arc, Mutex},
};
use tokio::net::{TcpListener, TcpStream};

type Db = Arc<Mutex<HashMap<String, Bytes>>>;

#[tokio::main]
async fn main() {
    let listener = TcpListener::bind("127.0.0.1:6379").await.unwrap();

    println!("Listening");

    let db = Arc::new(Mutex::new(HashMap::new()));
    loop {
        let (socket, _) = listener.accept().await.unwrap();
        let db = db.clone();
        tokio::spawn(async move {
            process(socket, db).await;
        });
    }
}

async fn process(socket: TcpStream, db: Db) {
    use mini_redis::Command::{self, Get, Set};

    //使用返回的connection获取socket数据
    let mut connection = Connection::new(socket);
    //Some是检测右边是否有返回的值
    while let Some(frame) = connection.read_frame().await.unwrap() {
        //模式匹配
        let response = match Command::from_frame(frame).unwrap() {
            //如果是set，就插入数据
            Set(cmd) => {
                let mut db = db.lock().unwrap();
                db.insert(cmd.key().to_string(), cmd.value().clone());
                Frame::Simple("OK".to_string())
            }
            //如果是get，就检查数据
            Get(cmd) => {
                // //如果有值，就返回，否则就返回null
                // if let Some(value) = db.get(cmd.key()) {
                //     // `Frame::Bulk` 期待数据的类型是 `Bytes`， 该类型会在后面章节讲解，
                //     // 此时，你只要知道 `&Vec<u8>` 可以使用 `into()` 方法转换成 `Bytes` 类型
                //     Frame::Bulk(value.clone().into())
                // } else {
                //     Frame::Null
                // }
                let db = db.lock().unwrap();
                if let Some(value) = db.get(cmd.key()) {
                    Frame::Bulk(value.clone())
                } else {
                    Frame::Null
                }
            }
            cmd => panic!("unimplemented {:?}", cmd),
        };
        connection.write_frame(&response).await.unwrap();
    }
}
