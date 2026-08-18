package kata

import scala.io.Source
import scala.util._
import scala.util.chaining.scalaUtilChainingOps

case class LineInfo(line: String, index: Int, originalLine: String)

sealed trait Processor {
  def process(stream: Iterator[String]): Unit
}

final class StringIndexProcessor(condition: FilterCondition) extends Processor {
  def process(stream: Iterator[String]): Unit = {
    stream.zipWithIndex
      .map { case (line, idx) => LineInfo(line, idx + 1, line) }
      .filter(condition.apply)
      .foreach(info => println(s"${info.index}: ${info.line}"))
  }
}

final class CountProcessor(condition: FilterCondition) extends Processor {
  def process(stream: Iterator[String]): Unit = {
    stream.zipWithIndex
      .map { case (line, idx) => LineInfo(line, idx + 1, line) }
      .filter(condition.apply)
      .size
      .tap(println)
  }
}

type FilterCondition = LineInfo => Boolean

object FilterCondition {
  def contains(substring: String): FilterCondition =
    info => info.line.toLowerCase.contains(substring.toLowerCase)

  def notContains(substring: String): FilterCondition =
    info => !info.line.toLowerCase.contains(substring.toLowerCase)

  def or(first: FilterCondition, second: FilterCondition): FilterCondition =
    info => first(info) || second(info)

  def and(first: FilterCondition, second: FilterCondition): FilterCondition =
    info => first(info) && second(info)

  implicit class FilterConditionOps(val condition: FilterCondition) extends AnyVal {
    def ||(other: FilterCondition): FilterCondition = FilterCondition.or(condition, other)
    def &&(other: FilterCondition): FilterCondition = FilterCondition.and(condition, other)
  }
}

object Grep extends App {
  import FilterCondition._

  val filePath  = "C:\\Users\\me\\Desktop\\Projects\\AlgorithmSolutions\\scala\\src\\main\\scala\\kata\\file.txt"
  val processor = StringIndexProcessor(FilterCondition.contains("Lorem") && FilterCondition.contains("of"))

  Using(Source.fromFile(filePath))(source => processor.process(source.getLines())).fold(
    error => println(s"Failed to open the file due to an error ${error}"),
    _ => println("Finished processing")
  )
}
