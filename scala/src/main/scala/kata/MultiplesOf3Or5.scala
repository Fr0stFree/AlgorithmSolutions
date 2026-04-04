package kata

object MultiplesOf3Or5 extends App {
  def solution(number: Int): Long = (1L until number).filter(num => num % 3 == 0 || num % 5 == 0).sum
}
