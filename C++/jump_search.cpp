#include<iostream>
#include<cmath>
using namespace std;

int linearSearch(int arr[], int key, int l, int h){
    for(unsigned int i=l;i<h;i++){
        if(arr[i] == key){
            return i;
        }
    }
    return -1;
}


int jumpSearch(int arr[], int key, int size, int jump){

    for(unsigned int i=0; i < size; i+=jump){
        if(key < arr[i]){
            return linearSearch(arr, key, i - jump, i);
        }
    }
    return -1;
}

int main(){
    cout << "**** Jump search ****" << endl;

    int arr[] = {0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610};
    int key = 55;

    int size = sizeof(arr)/sizeof(arr[0]);
    int jump = sqrt(size);

    int index = jumpSearch(arr, key, size, jump);
    cout << "Index " << index << endl;


}