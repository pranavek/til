class EuclidsAlgorithm{
    public static void main(String[] args) {
        System.out.println(new EuclidsAlgorithm().gcd(5,10));
    }

    /**
     * Use recursion to divide the divisor by the reminder until reminder becomes zero. 
     * Return the the divisor when the reminder is zero
     */
    int gcd(int a, int b){
        if(b ==0) return a;
        int r = a % b;
        return gcd(b,r);
    }
}