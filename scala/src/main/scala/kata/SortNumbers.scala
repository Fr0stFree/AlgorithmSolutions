package kata

object SortNumbers extends App {
  def sol(nums: List[Int]): List[Int] = nums.sortBy(identity)
}
