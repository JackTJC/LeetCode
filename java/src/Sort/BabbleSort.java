package Sort;
public class BabbleSort {
    public static void main(String[] args) {
        int [] arr ={1,2,3,4,5,6,7};
        arr=BabbleSort.sort(arr);
        System.out.println("done");
    }
    static int [] sort(int [] arr)
    {
        int len=arr.length;
        for(int i=1;i<len;i++)
            for (int j=0;j<len-i;j++)
            {
                if(arr[j]<arr[j+1])
                {
                    int temp=arr[j+1];
                    arr[j+1]=arr[j];
                    arr[j]=temp;
                }
            }
        return arr;
    }
}
