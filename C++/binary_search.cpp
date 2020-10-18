#include<iostream>
using namespace std;

int binarySearch(int array[],int key ,int l, int h){
    
    if(h >= l){
         
        int m = (l + h) / 2;

        if(array[m] == key){ 
            return m;
        }else if(key < array[m]){ 
            return binarySearch(array, key, l, m - 1);
        }else{
            return binarySearch(array, key, m + 1 ,h);
        }

    }
    return -1;    
}

int main(){
    int array[] = {4, 45, 49, 84, 91, 105};
    int key = 49;
    int size = sizeof(array)/sizeof(array[0]);

    int index = binarySearch(array, key, 0, size - 1);
    (index == -1) ? cout << "Key not found!" << endl: cout << "Key found at " << index << endl;
}