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

sealed trait FilterCondition {
  def apply(info: LineInfo): Boolean
}
object FilterCondition {
  case class Contains(substring: String) extends FilterCondition {
    def apply(info: LineInfo): Boolean = info.line.toLowerCase.contains(substring.toLowerCase)
  }
  case class NotContains(substring: String) extends FilterCondition {
    def apply(info: LineInfo): Boolean = !info.line.toLowerCase.contains(substring.toLowerCase)
  }
  case class Or(first: FilterCondition, second: FilterCondition) extends FilterCondition {
    def apply(info: LineInfo): Boolean = first(info) || second(info)
  }
  case class And(first: FilterCondition, second: FilterCondition) extends FilterCondition {
    def apply(info: LineInfo): Boolean = first(info) && second(info)
  }
}

object Grep extends App {
  val filePath  = "C:\\Users\\me\\Desktop\\Projects\\AlgorithmSolutions\\scala\\src\\main\\scala\\kata\\file.txt"
  val condition = FilterCondition.And(FilterCondition.Contains("Lorem"), FilterCondition.Contains("of"))
  val processor = StringIndexProcessor(condition)

  Using(Source.fromFile(filePath))(source => processor.process(source.getLines())).fold(
    error => println(s"Failed to open the file due to an error ${error}"),
    _ => println("Finished processing")
  )
}
