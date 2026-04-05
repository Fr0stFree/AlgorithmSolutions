package kata
import scala.collection.mutable.Map

object LuckCheck extends App {
  def luckCheck(s: String): Boolean = {
    if (!s.forall(_.isDigit)) throw new NumberFormatException(s"Invalid input: $s")
    s.map(_.asDigit) match {
      case array if array.size % 2 == 0 => array.take(s.size / 2).sum == array.drop(s.size / 2).sum
      case array => array.take(s.size / 2).sum == array.drop(s.size / 2 + 1).sum
    }
  }
  println(luckCheck("683179"))
}
