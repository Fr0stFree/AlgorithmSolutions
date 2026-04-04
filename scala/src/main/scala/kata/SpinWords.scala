package kata

object SpinWords extends App {
  def spinWords(sentence: String) = {
    sentence.split(' ').map({
      case word if word.length >= 5 => word.reverse
      case word                     => word
    }).mkString(" ")
  }
  println(spinWords("Hey wollef sroirraw"))
}
