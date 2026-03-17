package kata

object Dominator extends App {
  def dominator(xs: Seq[Int]) = {
    val half = xs.size / 2
    xs.groupBy(identity)
      .map(pair => (pair._1, pair._2.size))
      .filter(pair => pair._2 > half)
      .keys
      .headOption
      .getOrElse(-1)
  }

  println(dominator(Seq(3,4,3,2,3,1,3,3)))
}