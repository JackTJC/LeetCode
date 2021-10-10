package array;

import java.util.LinkedList;
import java.util.Queue;
/* 
剑指 Offer II 041. 滑动窗口的平均值
给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。
实现 MovingAverage 类：
MovingAverage(int size) 用窗口大小 size 初始化对象。
double next(int val) 成员函数 next 每次调用的时候都会往滑动窗口增加一个整数，请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。
*/

public class MovingAverage {
    
    public Queue<Integer> queue;
    public int size;
    /** Initialize your data structure here. */
    public MovingAverage(int size) {
        this.queue= new LinkedList<Integer>();
        this.size=size;
    }
    
    public double next(int val) {
        if (this.queue.size()==this.size){
            this.queue.poll();
        }
        this.queue.offer(val);
        int sum=0;
        for (Integer i : this.queue) {
            sum+=i;
        }
        return Double.valueOf(sum)/this.queue.size();
    }
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * MovingAverage obj = new MovingAverage(size);
 * double param_1 = obj.next(val);
 */
