#### 用 == 比较数组

* 相同维数且含有相同个数元素的数组才可以比较
* 每个元素都相同的才相等

#### &^按位置零
右操作数为1，则反转左操作数，右操作数为0，则左操作数保持不变
 1 &^ 0 -- 1
 1 &^ 1 -- 0
 0 &^ 1 -- 0
 0 &^ 0 -- 0