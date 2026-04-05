package kata

object FindTheDifference extends App {
  def findTheDifference(s: String, t: String) = {
    t.toList
      .sorted
      .zipAll(s.toList.sorted, ' ', ' ')
      .collectFirst(({ case (left, right ) if left != right => left}))
      .head
  }
  
  println(findTheDifference("abcd", "abecd"))
}
