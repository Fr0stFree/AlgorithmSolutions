package kata

import scala.collection.mutable

object ClimbStairs extends App {
  private val memo = mutable.Map[Int, Int](0 -> 0, 1 -> 1, 2 -> 2)
  def climbStairs(n: Int): Int = memo.getOrElseUpdate(n, climbStairs(n - 1) + climbStairs(n - 2))

  println(climbStairs(5))
}
