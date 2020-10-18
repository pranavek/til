#include<iostream>
using namespace std;


int search(int array[], int key, int size){

    for(unsigned int i=0;i<size;i++){
        if(array[i] == key){
            return i;
        }
    }
    return -1;

}

int main(){
    int key = 6;
    int array[] = {4, 5, 6, 2};
    int size = sizeof(array)/sizeof(int);
    int index = search(array, key, size);
    (index == -1) ? cout << "Key not found" << endl : cout<< "Key found at index " << index << endl;
}
