package kata
import scala.collection.mutable.Map

object FibonacciMemo extends App {
  def fib(n: Int): BigInt = {
    val memo = Map[Int, BigInt](0 -> 0, 1 -> 1)

    lazy val calculateFib: Int => BigInt = (number: Int) => 
      memo.getOrElseUpdate(number, calculateFib(number - 1) + calculateFib(number - 2))
    
    calculateFib(n)
  }
}
