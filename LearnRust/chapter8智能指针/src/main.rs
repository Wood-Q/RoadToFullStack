fn main(){
    let arr=Box::new([0;1000]);
    let arr1=arr;
    println!("{:?}",arr1.len());
}