import days/day1
import days/day2
import gleam/int
import gleam/io

pub fn main() {
  io.println("Hello from advent_of_code!")
  io.println("Day 1, Part 1: " <> int.to_string(day1.day1_pt1(False)))
  io.println("Day 1, Part 2: " <> int.to_string(day1.day1_pt2(False)))
  io.println("Day 2, Part 1: " <> int.to_string(day2.day2_pt1(False)))
  io.println("Day 2, Part 2: " <> int.to_string(day2.day2_pt2(False)))
}
