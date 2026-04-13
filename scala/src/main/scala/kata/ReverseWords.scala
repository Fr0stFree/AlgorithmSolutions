package kata

object ReverseWords extends App {
  def reverseWords(s: String): String = {
    s.split(" ").map(_.reverse).mkString(" ")
  }
  println(reverseWords("Let's take LeetCode contest"))
}
