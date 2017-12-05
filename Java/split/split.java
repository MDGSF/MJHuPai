//http://hp.vector.co.jp/authors/VA046927/mjscore/mjalgorism.html
//http://hp.vector.co.jp/authors/VA046927/mjscore/AgariIndex.java
//http://hp.vector.co.jp/authors/VA046927/mjscore/AgariBacktrack.java
import java.util.Arrays;
import java.util.List;
import java.util.ArrayList;

public class AgariBacktrack {
    static final int MAN = 0;
    static final int MAN1 = 0;
    static final int MAN2 = 1;
    static final int MAN3 = 2;
    static final int MAN4 = 3;
    static final int MAN5 = 4;
    static final int MAN6 = 5;
    static final int MAN7 = 6;
    static final int MAN8 = 7;
    static final int MAN9 = 8;
    static final int PIN = 9;
    static final int PIN1 = 9;
    static final int PIN2 = 10;
    static final int PIN3 = 11;
    static final int PIN4 = 12;
    static final int PIN5 = 13;
    static final int PIN6 = 14;
    static final int PIN7 = 15;
    static final int PIN8 = 16;
    static final int PIN9 = 17;
    static final int SOU = 18;
    static final int SOU1 = 18;
    static final int SOU2 = 19;
    static final int SOU3 = 20;
    static final int SOU4 = 21;
    static final int SOU5 = 22;
    static final int SOU6 = 23;
    static final int SOU7 = 24;
    static final int SOU8 = 25;
    static final int SOU9 = 26;
    static final int TON = 27;
    static final int NAN = 28;
    static final int SHA = 29;
    static final int PEI = 30;
    static final int HAK = 31;
    static final int HAT = 32;
    static final int CHU = 33;

    static final int[] n_zero;
    static {
        n_zero = new int[34];
        Arrays.fill(n_zero, 0);
    }

    // 攙偺庬椶偛偲偺屄悢傪媮傔傞
    static int[] analyse(int[] hai) {
        int[] n = n_zero.clone();

        for (int i : hai) {
            n[i]++;
        }
        return n;
    }

    // 僶僢僋僩儔僢僋朄偱悵摢偲柺巕偺慻傒崌傢偣傪媮傔傞
    static List<Integer[][]> agari(int[] n) {
        List<Integer[][]> ret = new ArrayList<Integer[][]>();

        for (int i = 0; i < 34; i++) {
            for (int kotsu_first = 0; kotsu_first < 2; kotsu_first++) {
                Integer[] janto = new Integer[1];
                ArrayList<Integer> kotsu = new ArrayList<Integer>();
                ArrayList<Integer> shuntsu = new ArrayList<Integer>();

                int[] t = n.clone();
                if (t[i] >= 2) {
                    // 悵摢傪偼偠傔偵庢傝弌偡
                    t[i] -= 2;
                    janto[0] = i;

                    if (kotsu_first == 0) {
                        // 崗巕傪愭偵庢傝弌偡
                        for (int j = 0; j < 34; j++) {
                            if (t[j] >= 3) {
                                t[j] -= 3;
                                kotsu.add(j);
                            }
                        }
                        // 弴巕傪屻偵庢傝弌偡
                        for (int a = 0; a < 3; a++) {
                            for (int b = 0; b < 7;) {
                                if (t[9*a+b] >= 1 &&
                                    t[9*a+b+1] >= 1 &&
                                    t[9*a+b+2] >= 1) {
                                    t[9*a+b]--;
                                    t[9*a+b+1]--;
                                    t[9*a+b+2]--;
                                    shuntsu.add(9*a+b);
                                } else {
                                    b++;
                                }
                            }
                        }
                    } else {
                        // 弴巕傪愭偵庢傝弌偡
                        for (int a = 0; a < 3; a++) {
                            for (int b = 0; b < 7;) {
                                if (t[9*a+b] >= 1 &&
                                    t[9*a+b+1] >= 1 &&
                                    t[9*a+b+2] >= 1) {
                                    t[9*a+b]--;
                                    t[9*a+b+1]--;
                                    t[9*a+b+2]--;
                                    shuntsu.add(9*a+b);
                                } else {
                                    b++;
                                }
                            }
                        }
                        // 崗巕傪屻偵庢傝弌偡
                        for (int j = 0; j < 34; j++) {
                            if (t[j] >= 3) {
                                t[j] -= 3;
                                kotsu.add(j);
                            }
                        }
                    }

                    // 榓椆偺宍偐挷傋傞
                    if (Arrays.equals(t, n_zero)) {
                        ret.add(new Integer[][] {janto, kotsu.toArray(new Integer[0]), shuntsu.toArray(new Integer[0])});
                    }
                }
            }
        }
        return ret;
    }

    public static void main(String[] args) {
        int[] hai = {
            MAN1, MAN1, MAN1,
            MAN2, MAN3, MAN4,
            MAN6, MAN7, MAN8,
            TON, TON, TON,
            SHA, SHA};

        int[] n = null;
        List<Integer[][]> ret = null;

        long time = System.currentTimeMillis();
        for (int i = 0; i < 100000; i++) {
            n = analyse(hai);
            ret = agari(n);
        }
        System.out.println(System.currentTimeMillis() - time);

        for (Integer[][] r : ret) {
            System.out.print("悵摢=");
            System.out.println(r[0][0]);
            for (Integer kotsu : r[1]) {
                System.out.print("崗巕=");
                System.out.println(kotsu);
            }
            for (Integer shuntsu : r[2]) {
                System.out.print("弴巕=");
                System.out.println(shuntsu);
            }
        }
    }
}