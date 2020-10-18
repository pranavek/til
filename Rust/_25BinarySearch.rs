fn main(){

    let arr: [i32;7] = [39, 2, 34, 79, 22, 34, 19];
    let x: i32 = 2;
    let index = search(arr, 0, arr.len() as i32, x); 
    println!("Index is {}", index);
}

fn search(arr: [i32;7], l: i32, r: i32,  x: i32) -> i32 {
    if r >= l {
        let mid = l + (r - 1)/2;
        if arr[mid as usize] == x {
            return mid;
        }
        if arr[mid as usize] > x {
            return search(arr, l, mid - 1, x);
        }
        return search(arr, mid + 1, r, x);
    }
    -1
} 


