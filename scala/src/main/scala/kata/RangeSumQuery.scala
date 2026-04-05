package kata

object RangeSumQuery extends App {
  class NumArray(_nums: Array[Int]) {
      def sumRange(left: Int, right: Int): Int = _nums.drop(left).take(right-left+1).sum
  }
  val anArray = NumArray(Array(-2, 0, 3, -5, 2, -1))
  
  println(anArray.sumRange(0, 2)) // 1
  println(anArray.sumRange(2, 5)) // -1
  println(anArray.sumRange(0, 5)) // -3
}
