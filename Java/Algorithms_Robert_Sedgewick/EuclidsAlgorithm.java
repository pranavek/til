class EuclidsAlgorithm{
    public static void main(String[] args) {
        System.out.println(new EuclidsAlgorithm().gcd(5,10));
    }

    int gcd(int a, int b){
        if(b ==0) return a;
        int r = a % b;
        return gcd(b,r);
    }
}