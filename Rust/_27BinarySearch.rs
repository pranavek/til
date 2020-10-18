/**
 *
 * The binary search requires a sorted list. 
 *
 * Recursive approach
 * 1. Check if the list is empty, if empty return -1
 * 2. Find the middle element
 * 3. check if the middle element is what we are searhing for
 * 4. If the element which we need to find is less than middle element search in the left half
 * 5. If the element which we need to find is grater than middle element search in the right half
 * 6. Do steps from 3 to 6 until we find the element in the middle
 *
 */


/**
 * Why is this called binary search?
 * What is the complexity of the recursive approach?
 * What will change when I convert this to linear approach?
 * What if there are duplicates in the list? How to find the index of all duplicate elements?
 *
 */

fn main(){

    let input: [i32;11] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11];
    let x: i32 = 10;
    let index = search(input, x, 0, input.len() as i32);
    println!("The index is {}", index);
}

fn search(input: [i32;11], x: i32, min: i32, max: i32) -> i32 {
    if max > min {
        let mid: i32 = (min + max) /2;
        if input[mid as usize] == x {
            return mid;
        }
        if x < input[mid as usize] {
            return search(input, x, min, mid - 1);
        }
        return search(input, x, mid + 1, max); 
        
    }
    return -1;
}
