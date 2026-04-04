package kata

object ArrayDiff extends App {
  def arrayDiff(a: Seq[Int], b: Seq[Int]): Seq[Int] = {
    val bSet = b.toSet
    a.collect({ case number if !bSet.contains(number) => number })
  }
  println(arrayDiff(Seq(1, 2, 2), Seq(1)))
}
