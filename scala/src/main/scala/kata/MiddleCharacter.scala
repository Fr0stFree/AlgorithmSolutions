package kata

object MiddleCharacter extends App {
  def middle(s: String): String = s.length match {
    case len if len % 2 == 0 => s.substring(len / 2 - 1, len / 2 + 1)
    case len => s(len / 2).toString
  }

  println(middle("middle"))
}