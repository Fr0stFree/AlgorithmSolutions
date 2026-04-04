package kata

object FindTheOddInt extends App {
  def findOdd(xs: Seq[Int]) = {
    xs.groupBy(identity)
      .collectFirst({ case (number, group) if group.size % 2 == 1 => number })
      .get
  }
  println(findOdd(Seq(20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5)))
}
