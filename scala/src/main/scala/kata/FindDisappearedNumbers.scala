package kata


object FindDisappearedNumbers extends App {

  def findDisappearedNumbers(nums: Array[Int]): List[Int] = {
    val aSet = nums.toSet
    (1 to nums.length).filter(!aSet.contains(_)).toList
  }
  println(findDisappearedNumbers(Array(4, 3, 2, 7, 8, 2, 3, 1)))
}
