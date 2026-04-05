package kata

object FindTheDifference extends App {
  def findTheDifference(s: String, t: String) = {
    t.toList
      .sorted
      .zipAll(s.toList.sorted, ' ', ' ')
      .find((left, right) => left != right)
      .head
      ._1
  }
  
  println(findTheDifference("abcd", "abecd"))
}
