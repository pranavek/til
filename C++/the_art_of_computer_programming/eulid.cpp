#include<iostream>
using namespace std;

/**
m = 5 n= 7

r = m % n = 5

m = 7
n = 5

r = 7 % 5 = 2

m = 5
n = 2

r = 5 % 2 = 1

m = 2
n = 1

r = 2 % 1 = 0

m = 1
n = 0

1
**/

int gcd(int m, int n){
    if(n == 0) return m;
    int r = m % n;
    return gcd(n,r);
}

int main(){
    cout << gcd(2166,6099) << endl;
    return 0;   
}


