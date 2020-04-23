public class EuclidsAlgorithm{
    public static void main(String[] args) {
        EuclidsAlgorithm ea = new EuclidsAlgorithm();
        System.out.println(ea.gcd(7,10));
    }

    private int gcd(int p,int q){
        if(q == 0) return p;
        int r = p % q;
        return gcd(q, r);
    }
}