package kata

object SumOfDigits extends App {
  def digitalRoot(n: Int): Int = {
    n match
      case number if number < 10 => number
      case number => digitalRoot(number.toString.map(_.asDigit).sum)
  }

  println(digitalRoot(942))
}
