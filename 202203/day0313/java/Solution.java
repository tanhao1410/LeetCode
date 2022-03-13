class Solution {
    //393. UTF-8 编码验证
    public boolean validUtf8(int[] data) {
        int i = 0;
        while(i < data.length){
            //先取第一个数
            int first = data[i];
            int count = getUniSize(first);
            //System.out.println(count);
            if (count == 5) return false;
            //判断后面的
            for(;count >= 2;count --){
                if (i == data.length - 1 || !isNextByte(data[++i])){
                    return false;
                }
            }
            i++;
        }
        return i == data.length;
    }

    //根据首字节判断该Unicode应该有多少
    private int getUniSize(int first){
        //以0开头
        if ((first & 0x00000080) == 0) return 1;
        int[] mask = new int[]{0x000000c0,0x000000e0,0x000000f0};
        int[] mask2 = new int[]{0x00000020,0x00000010,0x00000008};
        //判断前面有几个1
        for(int i = 2;i <= 4;i ++){
            //保证前面有i个1，保证i + 1位是0
            if (((mask[i - 2] & first) == mask[i - 2]) && ((mask2[i - 2] & first) == 0)) {
                return i;
            }
        }
        return 5;
    }
    //判断是否是10开头。
    private boolean isNextByte(int next){
        return (next & 0x000000c0 )==0x00000080;
    }
}