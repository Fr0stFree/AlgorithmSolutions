package kata

import scala.util.chaining._

object LowestAndHighest extends App {
  def highAndLow(numbers: String) = numbers.split("\\s")
      .map(_.toInt)
      .sorted
      .toList
      .pipe(list => s"${list.last} ${list.head}")

  println(highAndLow("1 9 3 4 -5"))
}
