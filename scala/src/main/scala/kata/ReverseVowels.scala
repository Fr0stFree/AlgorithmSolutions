package kata

object ReverseVowels extends App {
  private val allVowels =
    Set[Char]('a', 'a'.toUpper, 'e', 'e'.toUpper, 'i', 'i'.toUpper, 'o', 'o'.toUpper, 'u', 'u'.toUpper)

  def reverseVowels(s: String): String = {
    val vowels = s.filter(allVowels.contains(_)).reverse.iterator
    s.map(letter => if (allVowels.contains(letter)) then vowels.next() else letter).mkString
  }

  println(reverseVowels("IceCreAm")) // , "AceCreIm"
}
